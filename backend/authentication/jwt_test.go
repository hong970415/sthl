package authentication

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// ****Test_GenerateJwtToken
type generateJwtTokenTestCase struct {
	name          string
	inputDuration time.Duration
	input         string
	exec          func(string, error)
}

func Test_GenerateJwtToken(t *testing.T) {
	t.Setenv("JWT_SECRET", "welvknmerbginwuenjkvnuer")
	assert := assert.New(t)

	testCases := []generateJwtTokenTestCase{
		{
			name:          "valid token generation",
			inputDuration: time.Hour,
			input:         uuid.NewString(),
			exec: func(result string, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
			},
		},
		{
			name:          "invalid token generation with empty userId",
			inputDuration: time.Hour,
			input:         "",
			exec: func(result string, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:          "invalid token generation with zero duration",
			inputDuration: 0,
			input:         uuid.NewString(),
			exec: func(result string, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:          "invalid token generation with negative duration",
			inputDuration: time.Hour * -10,
			input:         uuid.NewString(),
			exec: func(result string, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(GenerateJwtToken(test.inputDuration, test.input))
		})
	}
}

// ****Test_VerifyJwtToken
type verifyJwtTokenTestCase struct {
	name  string
	input string
	exec  func(*JwtCustomClaims, error)
}

func Test_VerifyJwtToken(t *testing.T) {
	t.Setenv("JWT_SECRET", "welvknmerbginwuenjkvnuer")
	assert := assert.New(t)

	validUserId := "userId"
	validToken, _ := GenerateJwtToken(time.Hour, validUserId)
	expiredToken, _ := GenerateJwtToken(-time.Hour, validUserId)

	testCases := []verifyJwtTokenTestCase{
		{
			name:  "valid token",
			input: validToken,
			exec: func(result *JwtCustomClaims, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
				assert.Equal(validUserId, result.UserId)
			},
		},
		{
			name:  "invalid token",
			input: validToken + "wefwe",
			exec: func(result *JwtCustomClaims, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:  "expired token",
			input: expiredToken,
			exec: func(result *JwtCustomClaims, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
		{
			name:  "empty token",
			input: "",
			exec: func(result *JwtCustomClaims, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(VerifyJwtToken(test.input))
		})
	}
}
