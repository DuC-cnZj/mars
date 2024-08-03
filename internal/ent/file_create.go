// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/duc-cnzj/mars/v4/internal/ent/event"
	"github.com/duc-cnzj/mars/v4/internal/ent/file"
	"github.com/duc-cnzj/mars/v4/internal/ent/schema/schematype"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (fc *FileCreate) SetCreatedAt(t time.Time) *FileCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FileCreate) SetUpdatedAt(t time.Time) *FileCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FileCreate) SetDeletedAt(t time.Time) *FileCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableDeletedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetUploadType sets the "upload_type" field.
func (fc *FileCreate) SetUploadType(st schematype.UploadType) *FileCreate {
	fc.mutation.SetUploadType(st)
	return fc
}

// SetNillableUploadType sets the "upload_type" field if the given value is not nil.
func (fc *FileCreate) SetNillableUploadType(st *schematype.UploadType) *FileCreate {
	if st != nil {
		fc.SetUploadType(*st)
	}
	return fc
}

// SetPath sets the "path" field.
func (fc *FileCreate) SetPath(s string) *FileCreate {
	fc.mutation.SetPath(s)
	return fc
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(u uint64) *FileCreate {
	fc.mutation.SetSize(u)
	return fc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(u *uint64) *FileCreate {
	if u != nil {
		fc.SetSize(*u)
	}
	return fc
}

// SetUsername sets the "username" field.
func (fc *FileCreate) SetUsername(s string) *FileCreate {
	fc.mutation.SetUsername(s)
	return fc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (fc *FileCreate) SetNillableUsername(s *string) *FileCreate {
	if s != nil {
		fc.SetUsername(*s)
	}
	return fc
}

// SetNamespace sets the "namespace" field.
func (fc *FileCreate) SetNamespace(s string) *FileCreate {
	fc.mutation.SetNamespace(s)
	return fc
}

// SetNillableNamespace sets the "namespace" field if the given value is not nil.
func (fc *FileCreate) SetNillableNamespace(s *string) *FileCreate {
	if s != nil {
		fc.SetNamespace(*s)
	}
	return fc
}

// SetPod sets the "pod" field.
func (fc *FileCreate) SetPod(s string) *FileCreate {
	fc.mutation.SetPod(s)
	return fc
}

// SetNillablePod sets the "pod" field if the given value is not nil.
func (fc *FileCreate) SetNillablePod(s *string) *FileCreate {
	if s != nil {
		fc.SetPod(*s)
	}
	return fc
}

// SetContainer sets the "container" field.
func (fc *FileCreate) SetContainer(s string) *FileCreate {
	fc.mutation.SetContainer(s)
	return fc
}

// SetNillableContainer sets the "container" field if the given value is not nil.
func (fc *FileCreate) SetNillableContainer(s *string) *FileCreate {
	if s != nil {
		fc.SetContainer(*s)
	}
	return fc
}

// SetContainerPath sets the "container_path" field.
func (fc *FileCreate) SetContainerPath(s string) *FileCreate {
	fc.mutation.SetContainerPath(s)
	return fc
}

// SetNillableContainerPath sets the "container_path" field if the given value is not nil.
func (fc *FileCreate) SetNillableContainerPath(s *string) *FileCreate {
	if s != nil {
		fc.SetContainerPath(*s)
	}
	return fc
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (fc *FileCreate) AddEventIDs(ids ...int) *FileCreate {
	fc.mutation.AddEventIDs(ids...)
	return fc
}

// AddEvents adds the "events" edges to the Event entity.
func (fc *FileCreate) AddEvents(e ...*Event) *FileCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddEventIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if err := fc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		if file.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized file.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := file.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		if file.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized file.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := file.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.UploadType(); !ok {
		v := file.DefaultUploadType
		fc.mutation.SetUploadType(v)
	}
	if _, ok := fc.mutation.Size(); !ok {
		v := file.DefaultSize
		fc.mutation.SetSize(v)
	}
	if _, ok := fc.mutation.Username(); !ok {
		v := file.DefaultUsername
		fc.mutation.SetUsername(v)
	}
	if _, ok := fc.mutation.Namespace(); !ok {
		v := file.DefaultNamespace
		fc.mutation.SetNamespace(v)
	}
	if _, ok := fc.mutation.Pod(); !ok {
		v := file.DefaultPod
		fc.mutation.SetPod(v)
	}
	if _, ok := fc.mutation.Container(); !ok {
		v := file.DefaultContainer
		fc.mutation.SetContainer(v)
	}
	if _, ok := fc.mutation.ContainerPath(); !ok {
		v := file.DefaultContainerPath
		fc.mutation.SetContainerPath(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "File.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "File.updated_at"`)}
	}
	if _, ok := fc.mutation.UploadType(); !ok {
		return &ValidationError{Name: "upload_type", err: errors.New(`ent: missing required field "File.upload_type"`)}
	}
	if v, ok := fc.mutation.UploadType(); ok {
		if err := file.UploadTypeValidator(string(v)); err != nil {
			return &ValidationError{Name: "upload_type", err: fmt.Errorf(`ent: validator failed for field "File.upload_type": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "File.path"`)}
	}
	if v, ok := fc.mutation.Path(); ok {
		if err := file.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "File.path": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "File.size"`)}
	}
	if _, ok := fc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "File.username"`)}
	}
	if v, ok := fc.mutation.Username(); ok {
		if err := file.UsernameValidator(v); err != nil {
			return &ValidationError{Name: "username", err: fmt.Errorf(`ent: validator failed for field "File.username": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Namespace(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required field "File.namespace"`)}
	}
	if v, ok := fc.mutation.Namespace(); ok {
		if err := file.NamespaceValidator(v); err != nil {
			return &ValidationError{Name: "namespace", err: fmt.Errorf(`ent: validator failed for field "File.namespace": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Pod(); !ok {
		return &ValidationError{Name: "pod", err: errors.New(`ent: missing required field "File.pod"`)}
	}
	if v, ok := fc.mutation.Pod(); ok {
		if err := file.PodValidator(v); err != nil {
			return &ValidationError{Name: "pod", err: fmt.Errorf(`ent: validator failed for field "File.pod": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Container(); !ok {
		return &ValidationError{Name: "container", err: errors.New(`ent: missing required field "File.container"`)}
	}
	if v, ok := fc.mutation.Container(); ok {
		if err := file.ContainerValidator(v); err != nil {
			return &ValidationError{Name: "container", err: fmt.Errorf(`ent: validator failed for field "File.container": %w`, err)}
		}
	}
	if _, ok := fc.mutation.ContainerPath(); !ok {
		return &ValidationError{Name: "container_path", err: errors.New(`ent: missing required field "File.container_path"`)}
	}
	if v, ok := fc.mutation.ContainerPath(); ok {
		if err := file.ContainerPathValidator(v); err != nil {
			return &ValidationError{Name: "container_path", err: fmt.Errorf(`ent: validator failed for field "File.container_path": %w`, err)}
		}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(file.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.SetField(file.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := fc.mutation.UploadType(); ok {
		_spec.SetField(file.FieldUploadType, field.TypeString, value)
		_node.UploadType = value
	}
	if value, ok := fc.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeUint64, value)
		_node.Size = value
	}
	if value, ok := fc.mutation.Username(); ok {
		_spec.SetField(file.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := fc.mutation.Namespace(); ok {
		_spec.SetField(file.FieldNamespace, field.TypeString, value)
		_node.Namespace = value
	}
	if value, ok := fc.mutation.Pod(); ok {
		_spec.SetField(file.FieldPod, field.TypeString, value)
		_node.Pod = value
	}
	if value, ok := fc.mutation.Container(); ok {
		_spec.SetField(file.FieldContainer, field.TypeString, value)
		_node.Container = value
	}
	if value, ok := fc.mutation.ContainerPath(); ok {
		_spec.SetField(file.FieldContainerPath, field.TypeString, value)
		_node.ContainerPath = value
	}
	if nodes := fc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.EventsTable,
			Columns: []string{file.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fc *FileCreate) OnConflict(opts ...sql.ConflictOption) *FileUpsertOne {
	fc.conflict = opts
	return &FileUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FileCreate) OnConflictColumns(columns ...string) *FileUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertOne{
		create: fc,
	}
}

type (
	// FileUpsertOne is the builder for "upsert"-ing
	//  one File node.
	FileUpsertOne struct {
		create *FileCreate
	}

	// FileUpsert is the "OnConflict" setter.
	FileUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsert) SetUpdatedAt(v time.Time) *FileUpsert {
	u.Set(file.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsert) UpdateUpdatedAt() *FileUpsert {
	u.SetExcluded(file.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FileUpsert) SetDeletedAt(v time.Time) *FileUpsert {
	u.Set(file.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FileUpsert) UpdateDeletedAt() *FileUpsert {
	u.SetExcluded(file.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *FileUpsert) ClearDeletedAt() *FileUpsert {
	u.SetNull(file.FieldDeletedAt)
	return u
}

// SetUploadType sets the "upload_type" field.
func (u *FileUpsert) SetUploadType(v schematype.UploadType) *FileUpsert {
	u.Set(file.FieldUploadType, v)
	return u
}

// UpdateUploadType sets the "upload_type" field to the value that was provided on create.
func (u *FileUpsert) UpdateUploadType() *FileUpsert {
	u.SetExcluded(file.FieldUploadType)
	return u
}

// SetPath sets the "path" field.
func (u *FileUpsert) SetPath(v string) *FileUpsert {
	u.Set(file.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *FileUpsert) UpdatePath() *FileUpsert {
	u.SetExcluded(file.FieldPath)
	return u
}

// SetSize sets the "size" field.
func (u *FileUpsert) SetSize(v uint64) *FileUpsert {
	u.Set(file.FieldSize, v)
	return u
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsert) UpdateSize() *FileUpsert {
	u.SetExcluded(file.FieldSize)
	return u
}

// AddSize adds v to the "size" field.
func (u *FileUpsert) AddSize(v uint64) *FileUpsert {
	u.Add(file.FieldSize, v)
	return u
}

// SetUsername sets the "username" field.
func (u *FileUpsert) SetUsername(v string) *FileUpsert {
	u.Set(file.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *FileUpsert) UpdateUsername() *FileUpsert {
	u.SetExcluded(file.FieldUsername)
	return u
}

// SetNamespace sets the "namespace" field.
func (u *FileUpsert) SetNamespace(v string) *FileUpsert {
	u.Set(file.FieldNamespace, v)
	return u
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *FileUpsert) UpdateNamespace() *FileUpsert {
	u.SetExcluded(file.FieldNamespace)
	return u
}

// SetPod sets the "pod" field.
func (u *FileUpsert) SetPod(v string) *FileUpsert {
	u.Set(file.FieldPod, v)
	return u
}

// UpdatePod sets the "pod" field to the value that was provided on create.
func (u *FileUpsert) UpdatePod() *FileUpsert {
	u.SetExcluded(file.FieldPod)
	return u
}

// SetContainer sets the "container" field.
func (u *FileUpsert) SetContainer(v string) *FileUpsert {
	u.Set(file.FieldContainer, v)
	return u
}

// UpdateContainer sets the "container" field to the value that was provided on create.
func (u *FileUpsert) UpdateContainer() *FileUpsert {
	u.SetExcluded(file.FieldContainer)
	return u
}

// SetContainerPath sets the "container_path" field.
func (u *FileUpsert) SetContainerPath(v string) *FileUpsert {
	u.Set(file.FieldContainerPath, v)
	return u
}

// UpdateContainerPath sets the "container_path" field to the value that was provided on create.
func (u *FileUpsert) UpdateContainerPath() *FileUpsert {
	u.SetExcluded(file.FieldContainerPath)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileUpsertOne) UpdateNewValues() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(file.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FileUpsertOne) Ignore() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertOne) DoNothing() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreate.OnConflict
// documentation for more info.
func (u *FileUpsertOne) Update(set func(*FileUpsert)) *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsertOne) SetUpdatedAt(v time.Time) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUpdatedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FileUpsertOne) SetDeletedAt(v time.Time) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateDeletedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *FileUpsertOne) ClearDeletedAt() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUploadType sets the "upload_type" field.
func (u *FileUpsertOne) SetUploadType(v schematype.UploadType) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUploadType(v)
	})
}

// UpdateUploadType sets the "upload_type" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUploadType() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUploadType()
	})
}

// SetPath sets the "path" field.
func (u *FileUpsertOne) SetPath(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *FileUpsertOne) UpdatePath() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdatePath()
	})
}

// SetSize sets the "size" field.
func (u *FileUpsertOne) SetSize(v uint64) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertOne) AddSize(v uint64) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateSize() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// SetUsername sets the "username" field.
func (u *FileUpsertOne) SetUsername(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUsername() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUsername()
	})
}

// SetNamespace sets the "namespace" field.
func (u *FileUpsertOne) SetNamespace(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetNamespace(v)
	})
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateNamespace() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateNamespace()
	})
}

// SetPod sets the "pod" field.
func (u *FileUpsertOne) SetPod(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetPod(v)
	})
}

// UpdatePod sets the "pod" field to the value that was provided on create.
func (u *FileUpsertOne) UpdatePod() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdatePod()
	})
}

// SetContainer sets the "container" field.
func (u *FileUpsertOne) SetContainer(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetContainer(v)
	})
}

// UpdateContainer sets the "container" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateContainer() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateContainer()
	})
}

// SetContainerPath sets the "container_path" field.
func (u *FileUpsertOne) SetContainerPath(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetContainerPath(v)
	})
}

// UpdateContainerPath sets the "container_path" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateContainerPath() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateContainerPath()
	})
}

// Exec executes the query.
func (u *FileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	err      error
	builders []*FileCreate
	conflict []sql.ConflictOption
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileUpsertBulk {
	fcb.conflict = opts
	return &FileUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflictColumns(columns ...string) *FileUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertBulk{
		create: fcb,
	}
}

// FileUpsertBulk is the builder for "upsert"-ing
// a bulk of File nodes.
type FileUpsertBulk struct {
	create *FileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileUpsertBulk) UpdateNewValues() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(file.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FileUpsertBulk) Ignore() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertBulk) DoNothing() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreateBulk.OnConflict
// documentation for more info.
func (u *FileUpsertBulk) Update(set func(*FileUpsert)) *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *FileUpsertBulk) SetUpdatedAt(v time.Time) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUpdatedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *FileUpsertBulk) SetDeletedAt(v time.Time) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateDeletedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *FileUpsertBulk) ClearDeletedAt() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUploadType sets the "upload_type" field.
func (u *FileUpsertBulk) SetUploadType(v schematype.UploadType) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUploadType(v)
	})
}

// UpdateUploadType sets the "upload_type" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUploadType() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUploadType()
	})
}

// SetPath sets the "path" field.
func (u *FileUpsertBulk) SetPath(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdatePath() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdatePath()
	})
}

// SetSize sets the "size" field.
func (u *FileUpsertBulk) SetSize(v uint64) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetSize(v)
	})
}

// AddSize adds v to the "size" field.
func (u *FileUpsertBulk) AddSize(v uint64) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.AddSize(v)
	})
}

// UpdateSize sets the "size" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateSize() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateSize()
	})
}

// SetUsername sets the "username" field.
func (u *FileUpsertBulk) SetUsername(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUsername() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUsername()
	})
}

// SetNamespace sets the "namespace" field.
func (u *FileUpsertBulk) SetNamespace(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetNamespace(v)
	})
}

// UpdateNamespace sets the "namespace" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateNamespace() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateNamespace()
	})
}

// SetPod sets the "pod" field.
func (u *FileUpsertBulk) SetPod(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetPod(v)
	})
}

// UpdatePod sets the "pod" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdatePod() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdatePod()
	})
}

// SetContainer sets the "container" field.
func (u *FileUpsertBulk) SetContainer(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetContainer(v)
	})
}

// UpdateContainer sets the "container" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateContainer() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateContainer()
	})
}

// SetContainerPath sets the "container_path" field.
func (u *FileUpsertBulk) SetContainerPath(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetContainerPath(v)
	})
}

// UpdateContainerPath sets the "container_path" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateContainerPath() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateContainerPath()
	})
}

// Exec executes the query.
func (u *FileUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
