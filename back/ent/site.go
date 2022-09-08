// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cubdesign/dailyfj/ent/site"
)

// Site is the model entity for the Site schema.
type Site struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// FeedURL holds the value of the "feed_url" field.
	FeedURL string `json:"feed_url,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SiteQuery when eager-loading is set.
	Edges SiteEdges `json:"edges"`
}

// SiteEdges holds the relations/edges for other nodes in the graph.
type SiteEdges struct {
	// Articles holds the value of the articles edge.
	Articles []*Article `json:"articles,omitempty"`
	// Feeds holds the value of the feeds edge.
	Feeds []*Feed `json:"feeds,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ArticlesOrErr returns the Articles value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) ArticlesOrErr() ([]*Article, error) {
	if e.loadedTypes[0] {
		return e.Articles, nil
	}
	return nil, &NotLoadedError{edge: "articles"}
}

// FeedsOrErr returns the Feeds value or an error if the edge
// was not loaded in eager-loading.
func (e SiteEdges) FeedsOrErr() ([]*Feed, error) {
	if e.loadedTypes[1] {
		return e.Feeds, nil
	}
	return nil, &NotLoadedError{edge: "feeds"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Site) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case site.FieldActive:
			values[i] = new(sql.NullBool)
		case site.FieldID:
			values[i] = new(sql.NullInt64)
		case site.FieldName, site.FieldURL, site.FieldFeedURL:
			values[i] = new(sql.NullString)
		case site.FieldCreatedAt, site.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Site", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Site fields.
func (s *Site) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case site.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case site.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case site.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case site.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case site.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				s.URL = value.String
			}
		case site.FieldFeedURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field feed_url", values[i])
			} else if value.Valid {
				s.FeedURL = value.String
			}
		case site.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				s.Active = value.Bool
			}
		}
	}
	return nil
}

// QueryArticles queries the "articles" edge of the Site entity.
func (s *Site) QueryArticles() *ArticleQuery {
	return (&SiteClient{config: s.config}).QueryArticles(s)
}

// QueryFeeds queries the "feeds" edge of the Site entity.
func (s *Site) QueryFeeds() *FeedQuery {
	return (&SiteClient{config: s.config}).QueryFeeds(s)
}

// Update returns a builder for updating this Site.
// Note that you need to call Site.Unwrap() before calling this method if this Site
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Site) Update() *SiteUpdateOne {
	return (&SiteClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Site entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Site) Unwrap() *Site {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Site is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Site) String() string {
	var builder strings.Builder
	builder.WriteString("Site(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(s.URL)
	builder.WriteString(", ")
	builder.WriteString("feed_url=")
	builder.WriteString(s.FeedURL)
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", s.Active))
	builder.WriteByte(')')
	return builder.String()
}

// Sites is a parsable slice of Site.
type Sites []*Site

func (s Sites) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
