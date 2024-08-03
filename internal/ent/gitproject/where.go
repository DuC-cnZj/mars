// Code generated by ent, DO NOT EDIT.

package gitproject

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/duc-cnzj/mars/v4/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldDeletedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldName, v))
}

// DefaultBranch applies equality check predicate on the "default_branch" field. It's identical to DefaultBranchEQ.
func DefaultBranch(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldDefaultBranch, v))
}

// GitProjectID applies equality check predicate on the "git_project_id" field. It's identical to GitProjectIDEQ.
func GitProjectID(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldGitProjectID, v))
}

// Enabled applies equality check predicate on the "enabled" field. It's identical to EnabledEQ.
func Enabled(v bool) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldEnabled, v))
}

// GlobalEnabled applies equality check predicate on the "global_enabled" field. It's identical to GlobalEnabledEQ.
func GlobalEnabled(v bool) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldGlobalEnabled, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GitProject {
	return predicate.GitProject(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GitProject {
	return predicate.GitProject(sql.FieldNotNull(FieldDeletedAt))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldContainsFold(FieldName, v))
}

// DefaultBranchEQ applies the EQ predicate on the "default_branch" field.
func DefaultBranchEQ(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldDefaultBranch, v))
}

// DefaultBranchNEQ applies the NEQ predicate on the "default_branch" field.
func DefaultBranchNEQ(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldDefaultBranch, v))
}

// DefaultBranchIn applies the In predicate on the "default_branch" field.
func DefaultBranchIn(vs ...string) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldDefaultBranch, vs...))
}

// DefaultBranchNotIn applies the NotIn predicate on the "default_branch" field.
func DefaultBranchNotIn(vs ...string) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldDefaultBranch, vs...))
}

// DefaultBranchGT applies the GT predicate on the "default_branch" field.
func DefaultBranchGT(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldDefaultBranch, v))
}

// DefaultBranchGTE applies the GTE predicate on the "default_branch" field.
func DefaultBranchGTE(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldDefaultBranch, v))
}

// DefaultBranchLT applies the LT predicate on the "default_branch" field.
func DefaultBranchLT(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldDefaultBranch, v))
}

// DefaultBranchLTE applies the LTE predicate on the "default_branch" field.
func DefaultBranchLTE(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldDefaultBranch, v))
}

// DefaultBranchContains applies the Contains predicate on the "default_branch" field.
func DefaultBranchContains(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldContains(FieldDefaultBranch, v))
}

// DefaultBranchHasPrefix applies the HasPrefix predicate on the "default_branch" field.
func DefaultBranchHasPrefix(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldHasPrefix(FieldDefaultBranch, v))
}

// DefaultBranchHasSuffix applies the HasSuffix predicate on the "default_branch" field.
func DefaultBranchHasSuffix(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldHasSuffix(FieldDefaultBranch, v))
}

// DefaultBranchEqualFold applies the EqualFold predicate on the "default_branch" field.
func DefaultBranchEqualFold(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldEqualFold(FieldDefaultBranch, v))
}

// DefaultBranchContainsFold applies the ContainsFold predicate on the "default_branch" field.
func DefaultBranchContainsFold(v string) predicate.GitProject {
	return predicate.GitProject(sql.FieldContainsFold(FieldDefaultBranch, v))
}

// GitProjectIDEQ applies the EQ predicate on the "git_project_id" field.
func GitProjectIDEQ(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldGitProjectID, v))
}

// GitProjectIDNEQ applies the NEQ predicate on the "git_project_id" field.
func GitProjectIDNEQ(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldGitProjectID, v))
}

// GitProjectIDIn applies the In predicate on the "git_project_id" field.
func GitProjectIDIn(vs ...int) predicate.GitProject {
	return predicate.GitProject(sql.FieldIn(FieldGitProjectID, vs...))
}

// GitProjectIDNotIn applies the NotIn predicate on the "git_project_id" field.
func GitProjectIDNotIn(vs ...int) predicate.GitProject {
	return predicate.GitProject(sql.FieldNotIn(FieldGitProjectID, vs...))
}

// GitProjectIDGT applies the GT predicate on the "git_project_id" field.
func GitProjectIDGT(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldGT(FieldGitProjectID, v))
}

// GitProjectIDGTE applies the GTE predicate on the "git_project_id" field.
func GitProjectIDGTE(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldGTE(FieldGitProjectID, v))
}

// GitProjectIDLT applies the LT predicate on the "git_project_id" field.
func GitProjectIDLT(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldLT(FieldGitProjectID, v))
}

// GitProjectIDLTE applies the LTE predicate on the "git_project_id" field.
func GitProjectIDLTE(v int) predicate.GitProject {
	return predicate.GitProject(sql.FieldLTE(FieldGitProjectID, v))
}

// EnabledEQ applies the EQ predicate on the "enabled" field.
func EnabledEQ(v bool) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldEnabled, v))
}

// EnabledNEQ applies the NEQ predicate on the "enabled" field.
func EnabledNEQ(v bool) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldEnabled, v))
}

// GlobalEnabledEQ applies the EQ predicate on the "global_enabled" field.
func GlobalEnabledEQ(v bool) predicate.GitProject {
	return predicate.GitProject(sql.FieldEQ(FieldGlobalEnabled, v))
}

// GlobalEnabledNEQ applies the NEQ predicate on the "global_enabled" field.
func GlobalEnabledNEQ(v bool) predicate.GitProject {
	return predicate.GitProject(sql.FieldNEQ(FieldGlobalEnabled, v))
}

// GlobalConfigIsNil applies the IsNil predicate on the "global_config" field.
func GlobalConfigIsNil() predicate.GitProject {
	return predicate.GitProject(sql.FieldIsNull(FieldGlobalConfig))
}

// GlobalConfigNotNil applies the NotNil predicate on the "global_config" field.
func GlobalConfigNotNil() predicate.GitProject {
	return predicate.GitProject(sql.FieldNotNull(FieldGlobalConfig))
}

// HasChangelogs applies the HasEdge predicate on the "changelogs" edge.
func HasChangelogs() predicate.GitProject {
	return predicate.GitProject(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChangelogsTable, ChangelogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChangelogsWith applies the HasEdge predicate on the "changelogs" edge with a given conditions (other predicates).
func HasChangelogsWith(preds ...predicate.Changelog) predicate.GitProject {
	return predicate.GitProject(func(s *sql.Selector) {
		step := newChangelogsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GitProject) predicate.GitProject {
	return predicate.GitProject(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GitProject) predicate.GitProject {
	return predicate.GitProject(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GitProject) predicate.GitProject {
	return predicate.GitProject(sql.NotPredicates(p))
}
