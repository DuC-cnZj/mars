// Code generated by ent, DO NOT EDIT.

package file

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldUploadType holds the string denoting the upload_type field in the database.
	FieldUploadType = "upload_type"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldNamespace holds the string denoting the namespace field in the database.
	FieldNamespace = "namespace"
	// FieldPod holds the string denoting the pod field in the database.
	FieldPod = "pod"
	// FieldContainer holds the string denoting the container field in the database.
	FieldContainer = "container"
	// FieldContainerPath holds the string denoting the container_path field in the database.
	FieldContainerPath = "container_path"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// Table holds the table name of the file in the database.
	Table = "files"
	// EventsTable is the table that holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "file_id"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldUploadType,
	FieldPath,
	FieldSize,
	FieldUsername,
	FieldNamespace,
	FieldPod,
	FieldContainer,
	FieldContainerPath,
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
	// DefaultUploadType holds the default value on creation for the "upload_type" field.
	DefaultUploadType schematype.UploadType
	// UploadTypeValidator is a validator for the "upload_type" field. It is called by the builders before save.
	UploadTypeValidator func(string) error
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// DefaultSize holds the default value on creation for the "size" field.
	DefaultSize uint64
	// DefaultUsername holds the default value on creation for the "username" field.
	DefaultUsername string
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultNamespace holds the default value on creation for the "namespace" field.
	DefaultNamespace string
	// NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	NamespaceValidator func(string) error
	// DefaultPod holds the default value on creation for the "pod" field.
	DefaultPod string
	// PodValidator is a validator for the "pod" field. It is called by the builders before save.
	PodValidator func(string) error
	// DefaultContainer holds the default value on creation for the "container" field.
	DefaultContainer string
	// ContainerValidator is a validator for the "container" field. It is called by the builders before save.
	ContainerValidator func(string) error
	// DefaultContainerPath holds the default value on creation for the "container_path" field.
	DefaultContainerPath string
	// ContainerPathValidator is a validator for the "container_path" field. It is called by the builders before save.
	ContainerPathValidator func(string) error
)

// OrderOption defines the ordering options for the File queries.
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

// ByUploadType orders the results by the upload_type field.
func ByUploadType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUploadType, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByNamespace orders the results by the namespace field.
func ByNamespace(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNamespace, opts...).ToFunc()
}

// ByPod orders the results by the pod field.
func ByPod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPod, opts...).ToFunc()
}

// ByContainer orders the results by the container field.
func ByContainer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContainer, opts...).ToFunc()
}

// ByContainerPath orders the results by the container_path field.
func ByContainerPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContainerPath, opts...).ToFunc()
}

// ByEventsCount orders the results by events count.
func ByEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventsStep(), opts...)
	}
}

// ByEvents orders the results by events terms.
func ByEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
	)
}
