// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/haikara-dev/haikara/ent/article"
	"github.com/haikara-dev/haikara/ent/ogpimage"
)

// OGPImage is the model entity for the OGPImage schema.
type OGPImage struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
	// FilePath holds the value of the "file_path" field.
	FilePath string `json:"file_path,omitempty"`
	// OriginURL holds the value of the "origin_url" field.
	OriginURL string `json:"origin_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OGPImageQuery when eager-loading is set.
	Edges             OGPImageEdges `json:"edges"`
	article_ogp_image *int
}

// OGPImageEdges holds the relations/edges for other nodes in the graph.
type OGPImageEdges struct {
	// Article holds the value of the article edge.
	Article *Article `json:"article,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ArticleOrErr returns the Article value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OGPImageEdges) ArticleOrErr() (*Article, error) {
	if e.loadedTypes[0] {
		if e.Article == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: article.Label}
		}
		return e.Article, nil
	}
	return nil, &NotLoadedError{edge: "article"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OGPImage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ogpimage.FieldID:
			values[i] = new(sql.NullInt64)
		case ogpimage.FieldFileName, ogpimage.FieldFilePath, ogpimage.FieldOriginURL:
			values[i] = new(sql.NullString)
		case ogpimage.FieldCreatedAt, ogpimage.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case ogpimage.ForeignKeys[0]: // article_ogp_image
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OGPImage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OGPImage fields.
func (oi *OGPImage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ogpimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int(value.Int64)
		case ogpimage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oi.CreatedAt = value.Time
			}
		case ogpimage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oi.UpdatedAt = value.Time
			}
		case ogpimage.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				oi.FileName = value.String
			}
		case ogpimage.FieldFilePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_path", values[i])
			} else if value.Valid {
				oi.FilePath = value.String
			}
		case ogpimage.FieldOriginURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field origin_url", values[i])
			} else if value.Valid {
				oi.OriginURL = value.String
			}
		case ogpimage.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field article_ogp_image", value)
			} else if value.Valid {
				oi.article_ogp_image = new(int)
				*oi.article_ogp_image = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryArticle queries the "article" edge of the OGPImage entity.
func (oi *OGPImage) QueryArticle() *ArticleQuery {
	return (&OGPImageClient{config: oi.config}).QueryArticle(oi)
}

// Update returns a builder for updating this OGPImage.
// Note that you need to call OGPImage.Unwrap() before calling this method if this OGPImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OGPImage) Update() *OGPImageUpdateOne {
	return (&OGPImageClient{config: oi.config}).UpdateOne(oi)
}

// Unwrap unwraps the OGPImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OGPImage) Unwrap() *OGPImage {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OGPImage is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OGPImage) String() string {
	var builder strings.Builder
	builder.WriteString("OGPImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("created_at=")
	builder.WriteString(oi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("file_name=")
	builder.WriteString(oi.FileName)
	builder.WriteString(", ")
	builder.WriteString("file_path=")
	builder.WriteString(oi.FilePath)
	builder.WriteString(", ")
	builder.WriteString("origin_url=")
	builder.WriteString(oi.OriginURL)
	builder.WriteByte(')')
	return builder.String()
}

// OGPImages is a parsable slice of OGPImage.
type OGPImages []*OGPImage

func (oi OGPImages) config(cfg config) {
	for _i := range oi {
		oi[_i].config = cfg
	}
}