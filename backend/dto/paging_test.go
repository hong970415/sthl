package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// ****Test_PagingValidate
type pagingValidateTestCase struct {
	name  string
	input *Paging
	exec  func(error)
}

func Test_PagingValidate(t *testing.T) {
	assert := assert.New(t)

	testCases := []pagingValidateTestCase{
		{
			name:  "invalid page",
			input: NewPaging(0, 0, ""),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "invalid limit",
			input: NewPaging(1, 0, ""),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "valid paging",
			input: NewPaging(1, 1, ""),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}

// ****Test_PagingEnsureValid
type pagingEnsureValidTestCase struct {
	name  string
	input *Paging
	exec  func(Paging)
}

func Test_PagingEnsureValid(t *testing.T) {
	assert := assert.New(t)
	testCases := []pagingEnsureValidTestCase{
		{
			name:  "invalid param, return defult value",
			input: NewPaging(0, 0, ""),
			exec: func(result Paging) {
				assert.NotEmpty(result)
				assert.NotEmpty(result.Page)
				assert.NotEmpty(result.Limit)
				assert.Empty(result.Query)

				assert.Equal(1, result.Page)
				assert.Equal(20, result.Limit)
			},
		},
		{
			name:  "valid param, return new copy",
			input: NewPaging(2, 10, ""),
			exec: func(result Paging) {
				assert.NotEmpty(result)
				assert.NotEmpty(result.Page)
				assert.NotEmpty(result.Limit)
				assert.Empty(result.Query)

				assert.Equal(2, result.Page)
				assert.Equal(10, result.Limit)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(*test.input.Ensure())
		})
	}
}
