// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/faysalahmed-dev/wherehouse-order-picklist/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldName, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldAmount, v))
}

// UnitType applies equality check predicate on the "unit_type" field. It's identical to UnitTypeEQ.
func UnitType(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUnitType, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldName, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldAmount, v))
}

// AmountContains applies the Contains predicate on the "amount" field.
func AmountContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldAmount, v))
}

// AmountHasPrefix applies the HasPrefix predicate on the "amount" field.
func AmountHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldAmount, v))
}

// AmountHasSuffix applies the HasSuffix predicate on the "amount" field.
func AmountHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldAmount, v))
}

// AmountEqualFold applies the EqualFold predicate on the "amount" field.
func AmountEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldAmount, v))
}

// AmountContainsFold applies the ContainsFold predicate on the "amount" field.
func AmountContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldAmount, v))
}

// UnitTypeEQ applies the EQ predicate on the "unit_type" field.
func UnitTypeEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUnitType, v))
}

// UnitTypeNEQ applies the NEQ predicate on the "unit_type" field.
func UnitTypeNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldUnitType, v))
}

// UnitTypeIn applies the In predicate on the "unit_type" field.
func UnitTypeIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldUnitType, vs...))
}

// UnitTypeNotIn applies the NotIn predicate on the "unit_type" field.
func UnitTypeNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldUnitType, vs...))
}

// UnitTypeGT applies the GT predicate on the "unit_type" field.
func UnitTypeGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldUnitType, v))
}

// UnitTypeGTE applies the GTE predicate on the "unit_type" field.
func UnitTypeGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldUnitType, v))
}

// UnitTypeLT applies the LT predicate on the "unit_type" field.
func UnitTypeLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldUnitType, v))
}

// UnitTypeLTE applies the LTE predicate on the "unit_type" field.
func UnitTypeLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldUnitType, v))
}

// UnitTypeContains applies the Contains predicate on the "unit_type" field.
func UnitTypeContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldUnitType, v))
}

// UnitTypeHasPrefix applies the HasPrefix predicate on the "unit_type" field.
func UnitTypeHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldUnitType, v))
}

// UnitTypeHasSuffix applies the HasSuffix predicate on the "unit_type" field.
func UnitTypeHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldUnitType, v))
}

// UnitTypeEqualFold applies the EqualFold predicate on the "unit_type" field.
func UnitTypeEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldUnitType, v))
}

// UnitTypeContainsFold applies the ContainsFold predicate on the "unit_type" field.
func UnitTypeContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldUnitType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasSubCategories applies the HasEdge predicate on the "sub_categories" edge.
func HasSubCategories() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SubCategoriesTable, SubCategoriesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubCategoriesWith applies the HasEdge predicate on the "sub_categories" edge with a given conditions (other predicates).
func HasSubCategoriesWith(preds ...predicate.SubCategory) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newSubCategoriesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}
