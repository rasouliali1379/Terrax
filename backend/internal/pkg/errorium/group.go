package errorium

import (
	"database/sql"
	"errors"
	"strings"
)

type Group interface {
	error
	AddError(err ...string)
	HasError() bool
	EraseErrors()
	Join(g Group)
	ToSqlNullString() sql.NullString
	getErrors() []error
	parse(errs string)
}

type group struct {
	errs []error
}

func NewGroup(initials ...error) Group {
	return &group{errs: initials}
}

func Parse(s string) Group {
	splitted := strings.Split(s, ", ")
	var errs []error
	for _, s := range splitted {
		if len(s) == 0 {
			continue
		}
		errs = append(errs, errors.New(s))
	}
	return &group{errs: errs}
}

func (v *group) AddError(err ...string) {
	for i := range err {
		if !v.errExists(err[i]) {
			v.errs = append(v.errs, errors.New(err[i]))
		}
	}
}

func (v *group) errExists(err string) bool {
	for j := range v.errs {
		if err == v.errs[j].Error() {
			return true
		}
	}

	return false
}

func (v *group) Error() string {
	var msg string

	for i := 0; i < len(v.errs); i++ {
		msg += v.errs[i].Error()
		if i+1 != len(v.errs) {
			msg += ", "
		}
	}

	return msg
}

func (v *group) HasError() bool {
	if len(v.errs) > 0 {
		return true
	}
	return false
}

func (v *group) EraseErrors() {
	v.errs = []error{}
}

func (v *group) Join(g Group) {
	v.errs = append(v.errs, g.getErrors()...)
}

func (v *group) getErrors() []error {
	return v.errs
}

func (v *group) parse(s string) {

}

func (v *group) ToSqlNullString() sql.NullString {
	return sql.NullString{String: v.Error(), Valid: true}
}
