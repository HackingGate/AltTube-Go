package pipedproxy

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func PipedProxy(ctx *gin.Context) {
	// Parse the raw URL string from .env into a URL structure.
	proxyURL, err := url.Parse(os.Getenv("PIPED_PROXY_URL"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse proxy URL"})
		return
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(proxyURL)

	// Modify the director function to handle the request appropriately.
	proxy.Director = func(req *http.Request) {
		originalPath := ctx.Request.URL.Path
		targetPath := strings.TrimPrefix(originalPath, "/pipedproxy")

		// Ensure the request URI is rewritten correctly to the target service.
		req.URL.Scheme = proxyURL.Scheme
		req.URL.Host = proxyURL.Host
		req.URL.Path = targetPath
		req.Host = proxyURL.Host // Update the Host header in the request

		// If the original request includes a query string, append it to the new request
		if ctx.Request.URL.RawQuery != "" {
			req.URL.RawQuery = ctx.Request.URL.RawQuery
		}
	}

	// Modify the response function to replace all "/api/manifest/hls_playlist" with "/pipedproxy/api/manifest/hls_playlist"
	proxy.ModifyResponse = func(res *http.Response) error {
		if strings.HasPrefix(ctx.Request.URL.Path, "/pipedproxy/api/manifest/hls_variant") {
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return err
			}
			err = res.Body.Close()
			if err != nil {
				return err
			}
			body = bytes.ReplaceAll(body, []byte("/api/manifest/hls_playlist"), []byte("/pipedproxy/api/manifest/hls_playlist"))
			res.Body = io.NopCloser(bytes.NewReader(body))
			res.Header["Content-Length"] = []string{string(len(body))}
		}
		return nil
	}

	// Remove the Accept-Encoding header to avoid compressed responses that Gin cannot handle directly
	ctx.Request.Header.Del("Accept-Encoding")

	// ServeHTTP is not directly compatible with Gin, but we can adapt it using a ResponseWriter and Request.
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
