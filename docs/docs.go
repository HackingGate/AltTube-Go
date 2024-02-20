// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/like/video": {
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Add like to video",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "Add like to video",
                "parameters": [
                    {
                        "description": "Like Video Request",
                        "name": "likeVideoRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LikeVideoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Video liked successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "{\"message\": \"pong\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/piped/opensearch/suggestions": {
            "get": {
                "description": "Suggestions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "piped"
                ],
                "summary": "Suggestions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Suggestions",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/piped/search": {
            "get": {
                "description": "Search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "piped"
                ],
                "summary": "Search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search results",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/piped/streams/{videoID}": {
            "get": {
                "description": "Get video streams by video ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "piped"
                ],
                "summary": "Get video streams",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Video ID",
                        "name": "videoID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Video streams",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/": {
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete user",
                "responses": {
                    "200": {
                        "description": "{\"message\": \"User deleted successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/devices": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Get devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DeviceList"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Delete devices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete devices",
                "parameters": [
                    {
                        "description": "Devices to be deleted",
                        "name": "devices",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "integer"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Devices deleted successfully\", \"deleted\": [1, 2, 3]}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/email": {
            "patch": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Edit email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Edit email",
                "parameters": [
                    {
                        "description": "EditEmail",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.EditEmail"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Email updated successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"access_token\": \"access_token\", \"refresh_token\": \"refresh_token\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/logout": {
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "Logout user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Logged out successfully\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/refresh_token": {
            "post": {
                "security": [
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "Refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Refresh token",
                "responses": {
                    "200": {
                        "description": "{\"access_token\": \"access_token\", \"refresh_token\": \"refresh_token\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Signup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Signup",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Signup"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"message\": \"Registration successful\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Device": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "ip_address": {
                    "type": "string"
                },
                "last_active": {
                    "type": "string"
                },
                "user_agent": {
                    "type": "string"
                }
            }
        },
        "models.DeviceList": {
            "type": "object",
            "properties": {
                "current_device_id": {
                    "type": "integer"
                },
                "devices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Device"
                    }
                }
            }
        },
        "models.EditEmail": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "models.LikeVideoRequest": {
            "type": "object",
            "required": [
                "video_id"
            ],
            "properties": {
                "video_id": {
                    "type": "string"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.Signup": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "RefreshToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "AltTube API",
	Description:      "This is the API documentation for the AltTube application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
