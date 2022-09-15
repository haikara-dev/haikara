// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/cubdesign/dailyfj/ent/sitecategory"
)

// SiteCategory is the model entity for the SiteCategory schema.
type SiteCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Label holds the value of the "label" field.
	Label string `json:"label,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SiteCategoryQuery when eager-loading is set.
	Edges SiteCategoryEdges `json:"edges"`
}

// SiteCategoryEdges holds the relations/edges for other nodes in the graph.
type SiteCategoryEdges struct {
	// Sites holds the value of the sites edge.
	Sites []*Site `json:"sites,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SitesOrErr returns the Sites value or an error if the edge
// was not loaded in eager-loading.
func (e SiteCategoryEdges) SitesOrErr() ([]*Site, error) {
	if e.loadedTypes[0] {
		return e.Sites, nil
	}
	return nil, &NotLoadedError{edge: "sites"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SiteCategory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case sitecategory.FieldID:
			values[i] = new(sql.NullInt64)
		case sitecategory.FieldLabel:
			values[i] = new(sql.NullString)
		case sitecategory.FieldCreatedAt, sitecategory.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SiteCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SiteCategory fields.
func (sc *SiteCategory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sitecategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = int(value.Int64)
		case sitecategory.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = value.Time
			}
		case sitecategory.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = value.Time
			}
		case sitecategory.FieldLabel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field label", values[i])
			} else if value.Valid {
				sc.Label = value.String
			}
		}
	}
	return nil
}

// QuerySites queries the "sites" edge of the SiteCategory entity.
func (sc *SiteCategory) QuerySites() *SiteQuery {
	return (&SiteCategoryClient{config: sc.config}).QuerySites(sc)
}

// Update returns a builder for updating this SiteCategory.
// Note that you need to call SiteCategory.Unwrap() before calling this method if this SiteCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SiteCategory) Update() *SiteCategoryUpdateOne {
	return (&SiteCategoryClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the SiteCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SiteCategory) Unwrap() *SiteCategory {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SiteCategory is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SiteCategory) String() string {
	var builder strings.Builder
	builder.WriteString("SiteCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("label=")
	builder.WriteString(sc.Label)
	builder.WriteByte(')')
	return builder.String()
}

// SiteCategories is a parsable slice of SiteCategory.
type SiteCategories []*SiteCategory

func (sc SiteCategories) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}