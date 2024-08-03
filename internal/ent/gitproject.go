// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/duc-cnzj/mars/api/v4/mars"
	"github.com/duc-cnzj/mars/v4/internal/ent/gitproject"
)

// GitProject is the model entity for the GitProject schema.
type GitProject struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// DefaultBranch holds the value of the "default_branch" field.
	DefaultBranch string `json:"default_branch,omitempty"`
	// GitProjectID holds the value of the "git_project_id" field.
	GitProjectID int `json:"git_project_id,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
	// GlobalEnabled holds the value of the "global_enabled" field.
	GlobalEnabled bool `json:"global_enabled,omitempty"`
	// 全局配置
	GlobalConfig *mars.Config `json:"global_config,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GitProjectQuery when eager-loading is set.
	Edges        GitProjectEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GitProjectEdges holds the relations/edges for other nodes in the graph.
type GitProjectEdges struct {
	// Changelogs holds the value of the changelogs edge.
	Changelogs []*Changelog `json:"changelogs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ChangelogsOrErr returns the Changelogs value or an error if the edge
// was not loaded in eager-loading.
func (e GitProjectEdges) ChangelogsOrErr() ([]*Changelog, error) {
	if e.loadedTypes[0] {
		return e.Changelogs, nil
	}
	return nil, &NotLoadedError{edge: "changelogs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GitProject) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gitproject.FieldGlobalConfig:
			values[i] = new([]byte)
		case gitproject.FieldEnabled, gitproject.FieldGlobalEnabled:
			values[i] = new(sql.NullBool)
		case gitproject.FieldID, gitproject.FieldGitProjectID:
			values[i] = new(sql.NullInt64)
		case gitproject.FieldName, gitproject.FieldDefaultBranch:
			values[i] = new(sql.NullString)
		case gitproject.FieldCreatedAt, gitproject.FieldUpdatedAt, gitproject.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GitProject fields.
func (gp *GitProject) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gitproject.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gp.ID = int(value.Int64)
		case gitproject.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gp.CreatedAt = value.Time
			}
		case gitproject.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gp.UpdatedAt = value.Time
			}
		case gitproject.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gp.DeletedAt = new(time.Time)
				*gp.DeletedAt = value.Time
			}
		case gitproject.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gp.Name = value.String
			}
		case gitproject.FieldDefaultBranch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_branch", values[i])
			} else if value.Valid {
				gp.DefaultBranch = value.String
			}
		case gitproject.FieldGitProjectID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field git_project_id", values[i])
			} else if value.Valid {
				gp.GitProjectID = int(value.Int64)
			}
		case gitproject.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				gp.Enabled = value.Bool
			}
		case gitproject.FieldGlobalEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field global_enabled", values[i])
			} else if value.Valid {
				gp.GlobalEnabled = value.Bool
			}
		case gitproject.FieldGlobalConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field global_config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gp.GlobalConfig); err != nil {
					return fmt.Errorf("unmarshal field global_config: %w", err)
				}
			}
		default:
			gp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GitProject.
// This includes values selected through modifiers, order, etc.
func (gp *GitProject) Value(name string) (ent.Value, error) {
	return gp.selectValues.Get(name)
}

// QueryChangelogs queries the "changelogs" edge of the GitProject entity.
func (gp *GitProject) QueryChangelogs() *ChangelogQuery {
	return NewGitProjectClient(gp.config).QueryChangelogs(gp)
}

// Update returns a builder for updating this GitProject.
// Note that you need to call GitProject.Unwrap() before calling this method if this GitProject
// was returned from a transaction, and the transaction was committed or rolled back.
func (gp *GitProject) Update() *GitProjectUpdateOne {
	return NewGitProjectClient(gp.config).UpdateOne(gp)
}

// Unwrap unwraps the GitProject entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gp *GitProject) Unwrap() *GitProject {
	_tx, ok := gp.config.driver.(*txDriver)
	if !ok {
		panic("ent: GitProject is not a transactional entity")
	}
	gp.config.driver = _tx.drv
	return gp
}

// String implements the fmt.Stringer.
func (gp *GitProject) String() string {
	var builder strings.Builder
	builder.WriteString("GitProject(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(gp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(gp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := gp.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gp.Name)
	builder.WriteString(", ")
	builder.WriteString("default_branch=")
	builder.WriteString(gp.DefaultBranch)
	builder.WriteString(", ")
	builder.WriteString("git_project_id=")
	builder.WriteString(fmt.Sprintf("%v", gp.GitProjectID))
	builder.WriteString(", ")
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", gp.Enabled))
	builder.WriteString(", ")
	builder.WriteString("global_enabled=")
	builder.WriteString(fmt.Sprintf("%v", gp.GlobalEnabled))
	builder.WriteString(", ")
	builder.WriteString("global_config=")
	builder.WriteString(fmt.Sprintf("%v", gp.GlobalConfig))
	builder.WriteByte(')')
	return builder.String()
}

// GitProjects is a parsable slice of GitProject.
type GitProjects []*GitProject
