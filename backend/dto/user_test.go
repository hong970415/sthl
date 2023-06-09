package dto

import (
	"sthl/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

// ****Test_CreateUserDtoValidate
type createUserDtoValidateTestCase struct {
	name  string
	input *CreateUserDto
	exec  func(error)
}

func Test_CreateUserDtoValidate(t *testing.T) {
	assert := assert.New(t)
	validParam := NewCreateUserDto(utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	invalidEmailParam := NewCreateUserDto(utils.PtrOf(gofakeit.Name()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	emptyParam := NewCreateUserDto(utils.PtrOf(""), utils.PtrOf(""))
	nilParam := NewCreateUserDto(nil, nil)

	testCases := []createUserDtoValidateTestCase{
		{
			name:  "validate with valid param",
			input: validParam,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name:  "validate with invalid param, email",
			input: invalidEmailParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, empty",
			input: emptyParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, nil",
			input: nilParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, pw equal email",
			input: NewCreateUserDto(validParam.Email, validParam.Email),
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}

// ****Test_LoginDtoValidate
type loginDtoValidateTestCase struct {
	name  string
	input *LoginDto
	exec  func(error)
}

func Test_LoginDtoValidate(t *testing.T) {
	assert := assert.New(t)
	validParam := NewLoginDto(utils.PtrOf(gofakeit.Email()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	invalidEmailParam := NewLoginDto(utils.PtrOf(gofakeit.Name()), utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)))
	emptyParam := NewLoginDto(utils.PtrOf(""), utils.PtrOf(""))
	nilParam := NewLoginDto(nil, nil)

	testCases := []loginDtoValidateTestCase{
		{
			name:  "validate with valid param",
			input: validParam,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name:  "validate with invalid param, email",
			input: invalidEmailParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, empty",
			input: emptyParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, nil",
			input: nilParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, pw equal email",
			input: NewLoginDto(validParam.Email, validParam.Email),
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}

// ****Test_UpdateUserPasswordDtoValidate
type updateUserPasswordDtoValidateTestCase struct {
	name  string
	input *UpdateUserPasswordDto
	exec  func(error)
}

func Test_UpdateUserPasswordDtoValidate(t *testing.T) {
	assert := assert.New(t)
	validParam := NewUpdateUserPasswordDto(
		utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)),
		utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)),
	)
	invalidOldPwParam := NewUpdateUserPasswordDto(
		utils.PtrOf(gofakeit.Password(true, true, true, true, false, 5)),
		utils.PtrOf(gofakeit.Password(true, true, true, true, false, 6)),
	)
	emptyParam := NewUpdateUserPasswordDto(utils.PtrOf(""), utils.PtrOf(""))
	nilParam := NewUpdateUserPasswordDto(nil, nil)

	testCases := []updateUserPasswordDtoValidateTestCase{
		{
			name:  "validate with valid param",
			input: validParam,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name:  "validate with invalid param",
			input: invalidOldPwParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param",
			input: emptyParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, nil",
			input: nilParam,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "validate with invalid param, same",
			input: NewUpdateUserPasswordDto(validParam.CurrentPassword, validParam.CurrentPassword),
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}

// ****Test_UpdateUserPasswordDtoValidate
type queryUsersDtoTestCase struct {
	name  string
	input *QueryUsersDto
	exec  func(error)
}

func Test_QueryUsersDtoValidate(t *testing.T) {
	assert := assert.New(t)

	testCases := []queryUsersDtoTestCase{
		{
			name:  "valid param",
			input: NewQueryUsersDto(*NewPaging(1, 1, "")),
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name:  "invalid param, zero page",
			input: NewQueryUsersDto(*NewPaging(0, 1, "")),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "invalid param, zero limit",
			input: NewQueryUsersDto(*NewPaging(1, 0, "")),
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:  "invalid param, zero",
			input: NewQueryUsersDto(*NewPaging(0, 0, "")),
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(test.input.Validate())
		})
	}
}
