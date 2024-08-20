// Code generated by ent, DO NOT EDIT.

package namespace

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/duc-cnzj/mars/v4/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldDescription, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Namespace {
	return predicate.Namespace(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Namespace {
	return predicate.Namespace(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Namespace {
	return predicate.Namespace(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Namespace {
	return predicate.Namespace(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Namespace {
	return predicate.Namespace(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Namespace {
	return predicate.Namespace(sql.FieldContainsFold(FieldDescription, v))
}

// HasProjects applies the HasEdge predicate on the "projects" edge.
func HasProjects() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProjectsTable, ProjectsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectsWith applies the HasEdge predicate on the "projects" edge with a given conditions (other predicates).
func HasProjectsWith(preds ...predicate.Project) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := newProjectsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFavorites applies the HasEdge predicate on the "favorites" edge.
func HasFavorites() predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FavoritesTable, FavoritesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFavoritesWith applies the HasEdge predicate on the "favorites" edge with a given conditions (other predicates).
func HasFavoritesWith(preds ...predicate.Favorite) predicate.Namespace {
	return predicate.Namespace(func(s *sql.Selector) {
		step := newFavoritesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Namespace) predicate.Namespace {
	return predicate.Namespace(sql.NotPredicates(p))
}
