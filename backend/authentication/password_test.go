package authentication

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

// ****Test_HashPassword
type hashPasswordTestCase struct {
	name  string
	input string
	exec  func(string, error)
}

func Test_HashPassword(t *testing.T) {
	assert := assert.New(t)

	testCases := []hashPasswordTestCase{
		{
			name:  "valid case",
			input: gofakeit.Password(true, true, true, true, false, 6),
			exec: func(result string, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:  "invalid empty",
			input: "",
			exec: func(result string, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(HashPassword(test.input))
		})
	}
}

// ****Test_CompareHashPassword
type compareHashPasswordTestCase struct {
	name        string
	inputHashed string
	input       string
	exec        func(error)
}

func Test_CompareHashPassword(t *testing.T) {
	assert := assert.New(t)

	password := gofakeit.Password(true, true, true, true, false, 6)
	hashedPassword, err := HashPassword(password)
	assert.NotEmpty(hashedPassword)
	assert.NoError(err)

	testCases := []compareHashPasswordTestCase{
		{
			name:        "valid case",
			inputHashed: hashedPassword,
			input:       password,
			exec: func(e error) {
				assert.NoError(e)
			},
		},
		{
			name:        "invalid case wrong pw",
			inputHashed: hashedPassword,
			input:       "wrongpassword",
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:        "invalid case empty hashpw",
			inputHashed: "",
			input:       password,
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:        "invalid case empty pw",
			inputHashed: hashedPassword,
			input:       "",
			exec: func(e error) {
				assert.Error(e)
			},
		},
		{
			name:        "invalid case empty",
			inputHashed: "",
			input:       "",
			exec: func(e error) {
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(CompareHashPassword(test.inputHashed, test.input))
		})
	}
}
