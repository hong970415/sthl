package dto

import (
	"errors"
	"sthl/utils"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func RulesCombine(fixed []validation.Rule, rules ...validation.Rule) []validation.Rule {
	return append(fixed, rules...)
}

func NotEquals(target any, fields string) validation.RuleFunc {
	return func(value interface{}) error {
		retreiveErr := errors.New("fail to retreive target type")
		equalErr := errors.New(fields + " cannot be equal")
		unsupportTypeErr := errors.New("unsupported type")

		switch v := value.(type) {
		case string:
			s, ok := target.(string)
			if !ok {
				return retreiveErr
			}
			if v == s {
				return equalErr
			}
			return nil
		case *string:
			s, ok := target.(*string)
			if !ok {
				return retreiveErr
			}
			if *v == *s {
				return equalErr
			}
			return nil
		default:
			return unsupportTypeErr
		}
	}
}

func Equals(target any, fields string) validation.RuleFunc {
	return func(value interface{}) error {
		retreiveErr := errors.New("fail to retreive target type")
		notEqualErr := errors.New(fields + " should be equal")
		unsupportTypeErr := errors.New("unsupported type")

		switch v := value.(type) {
		case string:
			s, ok := target.(string)
			if !ok {
				return retreiveErr
			}
			if v != s {
				return notEqualErr
			}
			return nil
		case *string:
			s, ok := target.(*string)
			if !ok {
				return retreiveErr
			}
			if *v != *s {
				return notEqualErr
			}
			return nil
		default:
			return unsupportTypeErr
		}
	}
}

func InStrings(target []string, field string) validation.RuleFunc {
	return func(value interface{}) error {
		notInStringErr := errors.New(field + " not in target []string")
		unsupportTypeErr := errors.New("unsupported type")

		switch v := value.(type) {
		case string:
			cb := func(i int, j string) bool { return j == v }
			_, foundIndex := utils.Find(target, cb)
			if foundIndex == -1 {
				return notInStringErr
			}
			return nil

		case *string:
			cb := func(i int, j string) bool { return j == *v }
			_, foundIndex := utils.Find(target, cb)
			if foundIndex == -1 {
				return notInStringErr
			}
			return nil

		default:
			return unsupportTypeErr
		}
	}
}
