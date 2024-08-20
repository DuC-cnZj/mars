// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/duc-cnzj/mars/v4/internal/ent/favorite"
	"github.com/duc-cnzj/mars/v4/internal/ent/namespace"
	"github.com/duc-cnzj/mars/v4/internal/ent/predicate"
	"github.com/duc-cnzj/mars/v4/internal/ent/project"
)

// NamespaceUpdate is the builder for updating Namespace entities.
type NamespaceUpdate struct {
	config
	hooks    []Hook
	mutation *NamespaceMutation
}

// Where appends a list predicates to the NamespaceUpdate builder.
func (nu *NamespaceUpdate) Where(ps ...predicate.Namespace) *NamespaceUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NamespaceUpdate) SetUpdatedAt(t time.Time) *NamespaceUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetDeletedAt sets the "deleted_at" field.
func (nu *NamespaceUpdate) SetDeletedAt(t time.Time) *NamespaceUpdate {
	nu.mutation.SetDeletedAt(t)
	return nu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nu *NamespaceUpdate) SetNillableDeletedAt(t *time.Time) *NamespaceUpdate {
	if t != nil {
		nu.SetDeletedAt(*t)
	}
	return nu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nu *NamespaceUpdate) ClearDeletedAt() *NamespaceUpdate {
	nu.mutation.ClearDeletedAt()
	return nu
}

// SetName sets the "name" field.
func (nu *NamespaceUpdate) SetName(s string) *NamespaceUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nu *NamespaceUpdate) SetNillableName(s *string) *NamespaceUpdate {
	if s != nil {
		nu.SetName(*s)
	}
	return nu
}

// SetImagePullSecrets sets the "image_pull_secrets" field.
func (nu *NamespaceUpdate) SetImagePullSecrets(s []string) *NamespaceUpdate {
	nu.mutation.SetImagePullSecrets(s)
	return nu
}

// AppendImagePullSecrets appends s to the "image_pull_secrets" field.
func (nu *NamespaceUpdate) AppendImagePullSecrets(s []string) *NamespaceUpdate {
	nu.mutation.AppendImagePullSecrets(s)
	return nu
}

// SetDescription sets the "description" field.
func (nu *NamespaceUpdate) SetDescription(s string) *NamespaceUpdate {
	nu.mutation.SetDescription(s)
	return nu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nu *NamespaceUpdate) SetNillableDescription(s *string) *NamespaceUpdate {
	if s != nil {
		nu.SetDescription(*s)
	}
	return nu
}

// ClearDescription clears the value of the "description" field.
func (nu *NamespaceUpdate) ClearDescription() *NamespaceUpdate {
	nu.mutation.ClearDescription()
	return nu
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (nu *NamespaceUpdate) AddProjectIDs(ids ...int) *NamespaceUpdate {
	nu.mutation.AddProjectIDs(ids...)
	return nu
}

// AddProjects adds the "projects" edges to the Project entity.
func (nu *NamespaceUpdate) AddProjects(p ...*Project) *NamespaceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nu.AddProjectIDs(ids...)
}

// AddFavoriteIDs adds the "favorites" edge to the Favorite entity by IDs.
func (nu *NamespaceUpdate) AddFavoriteIDs(ids ...int) *NamespaceUpdate {
	nu.mutation.AddFavoriteIDs(ids...)
	return nu
}

// AddFavorites adds the "favorites" edges to the Favorite entity.
func (nu *NamespaceUpdate) AddFavorites(f ...*Favorite) *NamespaceUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return nu.AddFavoriteIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nu *NamespaceUpdate) Mutation() *NamespaceMutation {
	return nu.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (nu *NamespaceUpdate) ClearProjects() *NamespaceUpdate {
	nu.mutation.ClearProjects()
	return nu
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (nu *NamespaceUpdate) RemoveProjectIDs(ids ...int) *NamespaceUpdate {
	nu.mutation.RemoveProjectIDs(ids...)
	return nu
}

// RemoveProjects removes "projects" edges to Project entities.
func (nu *NamespaceUpdate) RemoveProjects(p ...*Project) *NamespaceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nu.RemoveProjectIDs(ids...)
}

// ClearFavorites clears all "favorites" edges to the Favorite entity.
func (nu *NamespaceUpdate) ClearFavorites() *NamespaceUpdate {
	nu.mutation.ClearFavorites()
	return nu
}

// RemoveFavoriteIDs removes the "favorites" edge to Favorite entities by IDs.
func (nu *NamespaceUpdate) RemoveFavoriteIDs(ids ...int) *NamespaceUpdate {
	nu.mutation.RemoveFavoriteIDs(ids...)
	return nu
}

// RemoveFavorites removes "favorites" edges to Favorite entities.
func (nu *NamespaceUpdate) RemoveFavorites(f ...*Favorite) *NamespaceUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return nu.RemoveFavoriteIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NamespaceUpdate) Save(ctx context.Context) (int, error) {
	if err := nu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NamespaceUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NamespaceUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NamespaceUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NamespaceUpdate) defaults() error {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		if namespace.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized namespace.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := namespace.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (nu *NamespaceUpdate) check() error {
	if v, ok := nu.mutation.Name(); ok {
		if err := namespace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Namespace.name": %w`, err)}
		}
	}
	return nil
}

func (nu *NamespaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(namespace.Table, namespace.Columns, sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(namespace.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.DeletedAt(); ok {
		_spec.SetField(namespace.FieldDeletedAt, field.TypeTime, value)
	}
	if nu.mutation.DeletedAtCleared() {
		_spec.ClearField(namespace.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(namespace.FieldName, field.TypeString, value)
	}
	if value, ok := nu.mutation.ImagePullSecrets(); ok {
		_spec.SetField(namespace.FieldImagePullSecrets, field.TypeJSON, value)
	}
	if value, ok := nu.mutation.AppendedImagePullSecrets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, namespace.FieldImagePullSecrets, value)
		})
	}
	if value, ok := nu.mutation.Description(); ok {
		_spec.SetField(namespace.FieldDescription, field.TypeString, value)
	}
	if nu.mutation.DescriptionCleared() {
		_spec.ClearField(namespace.FieldDescription, field.TypeString)
	}
	if nu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ProjectsTable,
			Columns: []string{namespace.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !nu.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ProjectsTable,
			Columns: []string{namespace.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ProjectsTable,
			Columns: []string{namespace.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.FavoritesTable,
			Columns: []string{namespace.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedFavoritesIDs(); len(nodes) > 0 && !nu.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.FavoritesTable,
			Columns: []string{namespace.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.FavoritesTable,
			Columns: []string{namespace.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namespace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NamespaceUpdateOne is the builder for updating a single Namespace entity.
type NamespaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NamespaceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NamespaceUpdateOne) SetUpdatedAt(t time.Time) *NamespaceUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetDeletedAt sets the "deleted_at" field.
func (nuo *NamespaceUpdateOne) SetDeletedAt(t time.Time) *NamespaceUpdateOne {
	nuo.mutation.SetDeletedAt(t)
	return nuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nuo *NamespaceUpdateOne) SetNillableDeletedAt(t *time.Time) *NamespaceUpdateOne {
	if t != nil {
		nuo.SetDeletedAt(*t)
	}
	return nuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nuo *NamespaceUpdateOne) ClearDeletedAt() *NamespaceUpdateOne {
	nuo.mutation.ClearDeletedAt()
	return nuo
}

// SetName sets the "name" field.
func (nuo *NamespaceUpdateOne) SetName(s string) *NamespaceUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nuo *NamespaceUpdateOne) SetNillableName(s *string) *NamespaceUpdateOne {
	if s != nil {
		nuo.SetName(*s)
	}
	return nuo
}

// SetImagePullSecrets sets the "image_pull_secrets" field.
func (nuo *NamespaceUpdateOne) SetImagePullSecrets(s []string) *NamespaceUpdateOne {
	nuo.mutation.SetImagePullSecrets(s)
	return nuo
}

// AppendImagePullSecrets appends s to the "image_pull_secrets" field.
func (nuo *NamespaceUpdateOne) AppendImagePullSecrets(s []string) *NamespaceUpdateOne {
	nuo.mutation.AppendImagePullSecrets(s)
	return nuo
}

// SetDescription sets the "description" field.
func (nuo *NamespaceUpdateOne) SetDescription(s string) *NamespaceUpdateOne {
	nuo.mutation.SetDescription(s)
	return nuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (nuo *NamespaceUpdateOne) SetNillableDescription(s *string) *NamespaceUpdateOne {
	if s != nil {
		nuo.SetDescription(*s)
	}
	return nuo
}

// ClearDescription clears the value of the "description" field.
func (nuo *NamespaceUpdateOne) ClearDescription() *NamespaceUpdateOne {
	nuo.mutation.ClearDescription()
	return nuo
}

// AddProjectIDs adds the "projects" edge to the Project entity by IDs.
func (nuo *NamespaceUpdateOne) AddProjectIDs(ids ...int) *NamespaceUpdateOne {
	nuo.mutation.AddProjectIDs(ids...)
	return nuo
}

// AddProjects adds the "projects" edges to the Project entity.
func (nuo *NamespaceUpdateOne) AddProjects(p ...*Project) *NamespaceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nuo.AddProjectIDs(ids...)
}

// AddFavoriteIDs adds the "favorites" edge to the Favorite entity by IDs.
func (nuo *NamespaceUpdateOne) AddFavoriteIDs(ids ...int) *NamespaceUpdateOne {
	nuo.mutation.AddFavoriteIDs(ids...)
	return nuo
}

// AddFavorites adds the "favorites" edges to the Favorite entity.
func (nuo *NamespaceUpdateOne) AddFavorites(f ...*Favorite) *NamespaceUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return nuo.AddFavoriteIDs(ids...)
}

// Mutation returns the NamespaceMutation object of the builder.
func (nuo *NamespaceUpdateOne) Mutation() *NamespaceMutation {
	return nuo.mutation
}

// ClearProjects clears all "projects" edges to the Project entity.
func (nuo *NamespaceUpdateOne) ClearProjects() *NamespaceUpdateOne {
	nuo.mutation.ClearProjects()
	return nuo
}

// RemoveProjectIDs removes the "projects" edge to Project entities by IDs.
func (nuo *NamespaceUpdateOne) RemoveProjectIDs(ids ...int) *NamespaceUpdateOne {
	nuo.mutation.RemoveProjectIDs(ids...)
	return nuo
}

// RemoveProjects removes "projects" edges to Project entities.
func (nuo *NamespaceUpdateOne) RemoveProjects(p ...*Project) *NamespaceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nuo.RemoveProjectIDs(ids...)
}

// ClearFavorites clears all "favorites" edges to the Favorite entity.
func (nuo *NamespaceUpdateOne) ClearFavorites() *NamespaceUpdateOne {
	nuo.mutation.ClearFavorites()
	return nuo
}

// RemoveFavoriteIDs removes the "favorites" edge to Favorite entities by IDs.
func (nuo *NamespaceUpdateOne) RemoveFavoriteIDs(ids ...int) *NamespaceUpdateOne {
	nuo.mutation.RemoveFavoriteIDs(ids...)
	return nuo
}

// RemoveFavorites removes "favorites" edges to Favorite entities.
func (nuo *NamespaceUpdateOne) RemoveFavorites(f ...*Favorite) *NamespaceUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return nuo.RemoveFavoriteIDs(ids...)
}

// Where appends a list predicates to the NamespaceUpdate builder.
func (nuo *NamespaceUpdateOne) Where(ps ...predicate.Namespace) *NamespaceUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NamespaceUpdateOne) Select(field string, fields ...string) *NamespaceUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Namespace entity.
func (nuo *NamespaceUpdateOne) Save(ctx context.Context) (*Namespace, error) {
	if err := nuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NamespaceUpdateOne) SaveX(ctx context.Context) *Namespace {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NamespaceUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NamespaceUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NamespaceUpdateOne) defaults() error {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		if namespace.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized namespace.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := namespace.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NamespaceUpdateOne) check() error {
	if v, ok := nuo.mutation.Name(); ok {
		if err := namespace.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Namespace.name": %w`, err)}
		}
	}
	return nil
}

func (nuo *NamespaceUpdateOne) sqlSave(ctx context.Context) (_node *Namespace, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(namespace.Table, namespace.Columns, sqlgraph.NewFieldSpec(namespace.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Namespace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, namespace.FieldID)
		for _, f := range fields {
			if !namespace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != namespace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(namespace.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.DeletedAt(); ok {
		_spec.SetField(namespace.FieldDeletedAt, field.TypeTime, value)
	}
	if nuo.mutation.DeletedAtCleared() {
		_spec.ClearField(namespace.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(namespace.FieldName, field.TypeString, value)
	}
	if value, ok := nuo.mutation.ImagePullSecrets(); ok {
		_spec.SetField(namespace.FieldImagePullSecrets, field.TypeJSON, value)
	}
	if value, ok := nuo.mutation.AppendedImagePullSecrets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, namespace.FieldImagePullSecrets, value)
		})
	}
	if value, ok := nuo.mutation.Description(); ok {
		_spec.SetField(namespace.FieldDescription, field.TypeString, value)
	}
	if nuo.mutation.DescriptionCleared() {
		_spec.ClearField(namespace.FieldDescription, field.TypeString)
	}
	if nuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ProjectsTable,
			Columns: []string{namespace.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedProjectsIDs(); len(nodes) > 0 && !nuo.mutation.ProjectsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ProjectsTable,
			Columns: []string{namespace.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.ProjectsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.ProjectsTable,
			Columns: []string{namespace.ProjectsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(project.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.FavoritesTable,
			Columns: []string{namespace.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedFavoritesIDs(); len(nodes) > 0 && !nuo.mutation.FavoritesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.FavoritesTable,
			Columns: []string{namespace.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.FavoritesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   namespace.FavoritesTable,
			Columns: []string{namespace.FavoritesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Namespace{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{namespace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
