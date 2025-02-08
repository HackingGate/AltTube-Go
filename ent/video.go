// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AltTube-Go/ent/video"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Video is the model entity for the Video schema.
type Video struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// UploadDate holds the value of the "upload_date" field.
	UploadDate time.Time `json:"upload_date,omitempty"`
	// Uploader holds the value of the "uploader" field.
	Uploader string `json:"uploader,omitempty"`
	// UploaderURL holds the value of the "uploader_url" field.
	UploaderURL string `json:"uploader_url,omitempty"`
	// ThumbnailURL holds the value of the "thumbnail_url" field.
	ThumbnailURL string `json:"thumbnail_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoQuery when eager-loading is set.
	Edges        VideoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VideoEdges holds the relations/edges for other nodes in the graph.
type VideoEdges struct {
	// LikeVideos holds the value of the like_videos edge.
	LikeVideos []*LikeVideo `json:"like_videos,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LikeVideosOrErr returns the LikeVideos value or an error if the edge
// was not loaded in eager-loading.
func (e VideoEdges) LikeVideosOrErr() ([]*LikeVideo, error) {
	if e.loadedTypes[0] {
		return e.LikeVideos, nil
	}
	return nil, &NotLoadedError{edge: "like_videos"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Video) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case video.FieldID, video.FieldTitle, video.FieldDescription, video.FieldUploader, video.FieldUploaderURL, video.FieldThumbnailURL:
			values[i] = new(sql.NullString)
		case video.FieldCreateTime, video.FieldUpdateTime, video.FieldUploadDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video fields.
func (v *Video) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case video.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				v.ID = value.String
			}
		case video.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				v.CreateTime = value.Time
			}
		case video.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				v.UpdateTime = value.Time
			}
		case video.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case video.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				v.Description = value.String
			}
		case video.FieldUploadDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field upload_date", values[i])
			} else if value.Valid {
				v.UploadDate = value.Time
			}
		case video.FieldUploader:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uploader", values[i])
			} else if value.Valid {
				v.Uploader = value.String
			}
		case video.FieldUploaderURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uploader_url", values[i])
			} else if value.Valid {
				v.UploaderURL = value.String
			}
		case video.FieldThumbnailURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field thumbnail_url", values[i])
			} else if value.Valid {
				v.ThumbnailURL = value.String
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Video.
// This includes values selected through modifiers, order, etc.
func (v *Video) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryLikeVideos queries the "like_videos" edge of the Video entity.
func (v *Video) QueryLikeVideos() *LikeVideoQuery {
	return NewVideoClient(v.config).QueryLikeVideos(v)
}

// Update returns a builder for updating this Video.
// Note that you need to call Video.Unwrap() before calling this method if this Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Video) Update() *VideoUpdateOne {
	return NewVideoClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Video entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Video) Unwrap() *Video {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Video is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Video) String() string {
	var builder strings.Builder
	builder.WriteString("Video(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("create_time=")
	builder.WriteString(v.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(v.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(v.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(v.Description)
	builder.WriteString(", ")
	builder.WriteString("upload_date=")
	builder.WriteString(v.UploadDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("uploader=")
	builder.WriteString(v.Uploader)
	builder.WriteString(", ")
	builder.WriteString("uploader_url=")
	builder.WriteString(v.UploaderURL)
	builder.WriteString(", ")
	builder.WriteString("thumbnail_url=")
	builder.WriteString(v.ThumbnailURL)
	builder.WriteByte(')')
	return builder.String()
}

// Videos is a parsable slice of Video.
type Videos []*Video
