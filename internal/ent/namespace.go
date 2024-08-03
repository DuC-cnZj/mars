// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/duc-cnzj/mars/v4/internal/ent/namespace"
)

// Namespace is the model entity for the Namespace schema.
type Namespace struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// 项目空间名
	Name string `json:"name,omitempty"`
	// image pull secrets
	ImagePullSecrets []string `json:"image_pull_secrets,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NamespaceQuery when eager-loading is set.
	Edges        NamespaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NamespaceEdges holds the relations/edges for other nodes in the graph.
type NamespaceEdges struct {
	// Projects holds the value of the projects edge.
	Projects []*Project `json:"projects,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProjectsOrErr returns the Projects value or an error if the edge
// was not loaded in eager-loading.
func (e NamespaceEdges) ProjectsOrErr() ([]*Project, error) {
	if e.loadedTypes[0] {
		return e.Projects, nil
	}
	return nil, &NotLoadedError{edge: "projects"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Namespace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case namespace.FieldImagePullSecrets:
			values[i] = new([]byte)
		case namespace.FieldID:
			values[i] = new(sql.NullInt64)
		case namespace.FieldName:
			values[i] = new(sql.NullString)
		case namespace.FieldCreatedAt, namespace.FieldUpdatedAt, namespace.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Namespace fields.
func (n *Namespace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case namespace.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case namespace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				n.CreatedAt = value.Time
			}
		case namespace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				n.UpdatedAt = value.Time
			}
		case namespace.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				n.DeletedAt = new(time.Time)
				*n.DeletedAt = value.Time
			}
		case namespace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				n.Name = value.String
			}
		case namespace.FieldImagePullSecrets:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field image_pull_secrets", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.ImagePullSecrets); err != nil {
					return fmt.Errorf("unmarshal field image_pull_secrets: %w", err)
				}
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Namespace.
// This includes values selected through modifiers, order, etc.
func (n *Namespace) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// QueryProjects queries the "projects" edge of the Namespace entity.
func (n *Namespace) QueryProjects() *ProjectQuery {
	return NewNamespaceClient(n.config).QueryProjects(n)
}

// Update returns a builder for updating this Namespace.
// Note that you need to call Namespace.Unwrap() before calling this method if this Namespace
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Namespace) Update() *NamespaceUpdateOne {
	return NewNamespaceClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Namespace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Namespace) Unwrap() *Namespace {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Namespace is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Namespace) String() string {
	var builder strings.Builder
	builder.WriteString("Namespace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("created_at=")
	builder.WriteString(n.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(n.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := n.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(n.Name)
	builder.WriteString(", ")
	builder.WriteString("image_pull_secrets=")
	builder.WriteString(fmt.Sprintf("%v", n.ImagePullSecrets))
	builder.WriteByte(')')
	return builder.String()
}

// Namespaces is a parsable slice of Namespace.
type Namespaces []*Namespace
