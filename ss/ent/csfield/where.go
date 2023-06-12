// Code generated by ent, DO NOT EDIT.

package csfield

import (
	"CSBackendTmp/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// MasterID applies equality check predicate on the "master_id" field. It's identical to MasterIDEQ.
func MasterID(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMasterID), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// ModeEQ applies the EQ predicate on the "mode" field.
func ModeEQ(v Mode) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMode), v))
	})
}

// ModeNEQ applies the NEQ predicate on the "mode" field.
func ModeNEQ(v Mode) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMode), v))
	})
}

// ModeIn applies the In predicate on the "mode" field.
func ModeIn(vs ...Mode) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMode), v...))
	})
}

// ModeNotIn applies the NotIn predicate on the "mode" field.
func ModeNotIn(vs ...Mode) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMode), v...))
	})
}

// PrivateLevelEQ applies the EQ predicate on the "private_level" field.
func PrivateLevelEQ(v PrivateLevel) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrivateLevel), v))
	})
}

// PrivateLevelNEQ applies the NEQ predicate on the "private_level" field.
func PrivateLevelNEQ(v PrivateLevel) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrivateLevel), v))
	})
}

// PrivateLevelIn applies the In predicate on the "private_level" field.
func PrivateLevelIn(vs ...PrivateLevel) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPrivateLevel), v...))
	})
}

// PrivateLevelNotIn applies the NotIn predicate on the "private_level" field.
func PrivateLevelNotIn(vs ...PrivateLevel) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPrivateLevel), v...))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uint64) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uint64) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// MasterIDEQ applies the EQ predicate on the "master_id" field.
func MasterIDEQ(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMasterID), v))
	})
}

// MasterIDNEQ applies the NEQ predicate on the "master_id" field.
func MasterIDNEQ(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMasterID), v))
	})
}

// MasterIDIn applies the In predicate on the "master_id" field.
func MasterIDIn(vs ...uint64) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMasterID), v...))
	})
}

// MasterIDNotIn applies the NotIn predicate on the "master_id" field.
func MasterIDNotIn(vs ...uint64) predicate.CSField {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMasterID), v...))
	})
}

// MasterIDGT applies the GT predicate on the "master_id" field.
func MasterIDGT(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMasterID), v))
	})
}

// MasterIDGTE applies the GTE predicate on the "master_id" field.
func MasterIDGTE(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMasterID), v))
	})
}

// MasterIDLT applies the LT predicate on the "master_id" field.
func MasterIDLT(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMasterID), v))
	})
}

// MasterIDLTE applies the LTE predicate on the "master_id" field.
func MasterIDLTE(v uint64) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMasterID), v))
	})
}

// MasterIDIsNil applies the IsNil predicate on the "master_id" field.
func MasterIDIsNil() predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMasterID)))
	})
}

// MasterIDNotNil applies the NotNil predicate on the "master_id" field.
func MasterIDNotNil() predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMasterID)))
	})
}

// HasJoinedUser applies the HasEdge predicate on the "joined_user" edge.
func HasJoinedUser() predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JoinedUserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, JoinedUserTable, JoinedUserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJoinedUserWith applies the HasEdge predicate on the "joined_user" edge with a given conditions (other predicates).
func HasJoinedUserWith(preds ...predicate.User) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JoinedUserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, JoinedUserTable, JoinedUserPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasJoins applies the HasEdge predicate on the "joins" edge.
func HasJoins() predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JoinsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, JoinsTable, JoinsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasJoinsWith applies the HasEdge predicate on the "joins" edge with a given conditions (other predicates).
func HasJoinsWith(preds ...predicate.Join) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(JoinsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, JoinsTable, JoinsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CSField) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CSField) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CSField) predicate.CSField {
	return predicate.CSField(func(s *sql.Selector) {
		p(s.Not())
	})
}