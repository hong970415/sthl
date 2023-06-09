package authentication

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// ****Test_GeneratePassport
type generatePassportTestCase struct {
	name            string
	inputAtDuration time.Duration
	inputRtDuration time.Duration
	input           string
	exec            func(*Passport, error)
}

func Test_GeneratePassport(t *testing.T) {
	t.Setenv("JWT_SECRET", "welvknmerbginwuenjkvnuer")
	assert := assert.New(t)

	validUserId := uuid.NewString()
	testCases := []generatePassportTestCase{
		{
			name:            "valid case",
			inputAtDuration: time.Hour,
			inputRtDuration: 24 * time.Hour,
			input:           validUserId,
			exec: func(result *Passport, e error) {
				assert.NotEmpty(result)
				assert.NoError(e)
				assert.NotEmpty(result.AccessToken)
				assert.NotEmpty(result.RefreshToken)
			},
		},
		{
			name:            "empty userId",
			inputAtDuration: time.Hour,
			inputRtDuration: time.Hour,
			input:           "",
			exec: func(result *Passport, e error) {
				assert.Empty(result)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(GeneratePassport(test.inputAtDuration, test.inputRtDuration, test.input))
		})
	}
}

// ****Test_VerifyPassport
type verifyPassportTestCase struct {
	name  string
	input Passport
	exec  func(*JwtCustomClaims, *JwtCustomClaims, error)
}

func Test_VerifyPassport(t *testing.T) {
	t.Setenv("JWT_SECRET", "welvknmerbginwuenjkvnuer")
	assert := assert.New(t)

	atDuration := time.Hour
	rtDuration := 24 * time.Hour
	validUserId := "userId"
	validPp, err := GeneratePassport(atDuration, rtDuration, validUserId)
	assert.NotEmpty(validPp)
	assert.NoError(err)

	invalidPp := *validPp
	invalidPp.AccessToken += "qwdq"

	testCases := []verifyPassportTestCase{
		{
			name:  "valid case",
			input: *validPp,
			exec: func(resultAt *JwtCustomClaims, resultRt *JwtCustomClaims, e error) {
				assert.NotEmpty(resultAt)
				assert.NotEmpty(resultRt)
				assert.NoError(e)
				assert.NotEmpty(resultAt.UserId)
				assert.NotEmpty(resultRt.UserId)
				assert.True(resultAt.UserId == resultRt.UserId)
				assert.True(resultAt.UserId == validUserId)
				assert.True(time.Now().Before(resultAt.ExpiresAt.Time))
				assert.True(time.Now().Before(resultRt.ExpiresAt.Time))
			},
		},
		{
			name:  "invalid case",
			input: invalidPp,
			exec: func(resultAt *JwtCustomClaims, resultRt *JwtCustomClaims, e error) {
				assert.Empty(resultAt)
				assert.Empty(resultRt)
				assert.Error(e)
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.exec(VerifyPassport(test.input))
		})
	}
}
