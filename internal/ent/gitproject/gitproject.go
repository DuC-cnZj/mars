// Code generated by ent, DO NOT EDIT.

package gitproject

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the gitproject type in the database.
	Label = "git_project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDefaultBranch holds the string denoting the default_branch field in the database.
	FieldDefaultBranch = "default_branch"
	// FieldGitProjectID holds the string denoting the git_project_id field in the database.
	FieldGitProjectID = "git_project_id"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// FieldGlobalEnabled holds the string denoting the global_enabled field in the database.
	FieldGlobalEnabled = "global_enabled"
	// FieldGlobalConfig holds the string denoting the global_config field in the database.
	FieldGlobalConfig = "global_config"
	// EdgeChangelogs holds the string denoting the changelogs edge name in mutations.
	EdgeChangelogs = "changelogs"
	// Table holds the table name of the gitproject in the database.
	Table = "git_projects"
	// ChangelogsTable is the table that holds the changelogs relation/edge.
	ChangelogsTable = "changelogs"
	// ChangelogsInverseTable is the table name for the Changelog entity.
	// It exists in this package in order to avoid circular dependency with the "changelog" package.
	ChangelogsInverseTable = "changelogs"
	// ChangelogsColumn is the table column denoting the changelogs relation/edge.
	ChangelogsColumn = "git_project_id"
)

// Columns holds all SQL columns for gitproject fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldName,
	FieldDefaultBranch,
	FieldGitProjectID,
	FieldEnabled,
	FieldGlobalEnabled,
	FieldGlobalConfig,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/duc-cnzj/mars/v4/internal/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDefaultBranch holds the default value on creation for the "default_branch" field.
	DefaultDefaultBranch string
	// DefaultBranchValidator is a validator for the "default_branch" field. It is called by the builders before save.
	DefaultBranchValidator func(string) error
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
	// DefaultGlobalEnabled holds the default value on creation for the "global_enabled" field.
	DefaultGlobalEnabled bool
)

// OrderOption defines the ordering options for the GitProject queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDefaultBranch orders the results by the default_branch field.
func ByDefaultBranch(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultBranch, opts...).ToFunc()
}

// ByGitProjectID orders the results by the git_project_id field.
func ByGitProjectID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGitProjectID, opts...).ToFunc()
}

// ByEnabled orders the results by the enabled field.
func ByEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnabled, opts...).ToFunc()
}

// ByGlobalEnabled orders the results by the global_enabled field.
func ByGlobalEnabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGlobalEnabled, opts...).ToFunc()
}

// ByChangelogsCount orders the results by changelogs count.
func ByChangelogsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChangelogsStep(), opts...)
	}
}

// ByChangelogs orders the results by changelogs terms.
func ByChangelogs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChangelogsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newChangelogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ChangelogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ChangelogsTable, ChangelogsColumn),
	)
}
