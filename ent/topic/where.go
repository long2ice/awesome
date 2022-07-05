// Code generated by entc, DO NOT EDIT.

package topic

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/long2ice/awesome/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// SubName applies equality check predicate on the "sub_name" field. It's identical to SubNameEQ.
func SubName(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// GithubURL applies equality check predicate on the "github_url" field. It's identical to GithubURLEQ.
func GithubURL(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubURL), v))
	})
}

// PlatformID applies equality check predicate on the "platform_id" field. It's identical to PlatformIDEQ.
func PlatformID(v int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformID), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// SubNameEQ applies the EQ predicate on the "sub_name" field.
func SubNameEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSubName), v))
	})
}

// SubNameNEQ applies the NEQ predicate on the "sub_name" field.
func SubNameNEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSubName), v))
	})
}

// SubNameIn applies the In predicate on the "sub_name" field.
func SubNameIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSubName), v...))
	})
}

// SubNameNotIn applies the NotIn predicate on the "sub_name" field.
func SubNameNotIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSubName), v...))
	})
}

// SubNameGT applies the GT predicate on the "sub_name" field.
func SubNameGT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSubName), v))
	})
}

// SubNameGTE applies the GTE predicate on the "sub_name" field.
func SubNameGTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSubName), v))
	})
}

// SubNameLT applies the LT predicate on the "sub_name" field.
func SubNameLT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSubName), v))
	})
}

// SubNameLTE applies the LTE predicate on the "sub_name" field.
func SubNameLTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSubName), v))
	})
}

// SubNameContains applies the Contains predicate on the "sub_name" field.
func SubNameContains(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSubName), v))
	})
}

// SubNameHasPrefix applies the HasPrefix predicate on the "sub_name" field.
func SubNameHasPrefix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSubName), v))
	})
}

// SubNameHasSuffix applies the HasSuffix predicate on the "sub_name" field.
func SubNameHasSuffix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSubName), v))
	})
}

// SubNameEqualFold applies the EqualFold predicate on the "sub_name" field.
func SubNameEqualFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSubName), v))
	})
}

// SubNameContainsFold applies the ContainsFold predicate on the "sub_name" field.
func SubNameContainsFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSubName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	})
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	})
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	})
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	})
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	})
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	})
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	})
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	})
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	})
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	})
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	})
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	})
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	})
}

// GithubURLEQ applies the EQ predicate on the "github_url" field.
func GithubURLEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGithubURL), v))
	})
}

// GithubURLNEQ applies the NEQ predicate on the "github_url" field.
func GithubURLNEQ(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGithubURL), v))
	})
}

// GithubURLIn applies the In predicate on the "github_url" field.
func GithubURLIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGithubURL), v...))
	})
}

// GithubURLNotIn applies the NotIn predicate on the "github_url" field.
func GithubURLNotIn(vs ...string) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGithubURL), v...))
	})
}

// GithubURLGT applies the GT predicate on the "github_url" field.
func GithubURLGT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGithubURL), v))
	})
}

// GithubURLGTE applies the GTE predicate on the "github_url" field.
func GithubURLGTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGithubURL), v))
	})
}

// GithubURLLT applies the LT predicate on the "github_url" field.
func GithubURLLT(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGithubURL), v))
	})
}

// GithubURLLTE applies the LTE predicate on the "github_url" field.
func GithubURLLTE(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGithubURL), v))
	})
}

// GithubURLContains applies the Contains predicate on the "github_url" field.
func GithubURLContains(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGithubURL), v))
	})
}

// GithubURLHasPrefix applies the HasPrefix predicate on the "github_url" field.
func GithubURLHasPrefix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGithubURL), v))
	})
}

// GithubURLHasSuffix applies the HasSuffix predicate on the "github_url" field.
func GithubURLHasSuffix(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGithubURL), v))
	})
}

// GithubURLEqualFold applies the EqualFold predicate on the "github_url" field.
func GithubURLEqualFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGithubURL), v))
	})
}

// GithubURLContainsFold applies the ContainsFold predicate on the "github_url" field.
func GithubURLContainsFold(v string) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGithubURL), v))
	})
}

// PlatformIDEQ applies the EQ predicate on the "platform_id" field.
func PlatformIDEQ(v int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlatformID), v))
	})
}

// PlatformIDNEQ applies the NEQ predicate on the "platform_id" field.
func PlatformIDNEQ(v int) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlatformID), v))
	})
}

// PlatformIDIn applies the In predicate on the "platform_id" field.
func PlatformIDIn(vs ...int) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlatformID), v...))
	})
}

// PlatformIDNotIn applies the NotIn predicate on the "platform_id" field.
func PlatformIDNotIn(vs ...int) predicate.Topic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Topic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlatformID), v...))
	})
}

// HasPlatform applies the HasEdge predicate on the "platform" edge.
func HasPlatform() predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlatformTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlatformTable, PlatformColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlatformWith applies the HasEdge predicate on the "platform" edge with a given conditions (other predicates).
func HasPlatformWith(preds ...predicate.Platform) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlatformInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PlatformTable, PlatformColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRepos applies the HasEdge predicate on the "repos" edge.
func HasRepos() predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReposTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReposTable, ReposColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReposWith applies the HasEdge predicate on the "repos" edge with a given conditions (other predicates).
func HasReposWith(preds ...predicate.Repo) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ReposInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReposTable, ReposColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
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
func Not(p predicate.Topic) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		p(s.Not())
	})
}
