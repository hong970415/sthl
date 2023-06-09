package dto

import (
	"sthl/ent"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ****CreateUserDto
type CreateUserDto struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
type CreateUserMappedDto struct {
	Email    *string
	HashedPw *string
}

func NewCreateUserDto(email *string, pw *string) *CreateUserDto {
	return &CreateUserDto{
		Email:    email,
		Password: pw,
	}
}

func (d CreateUserDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Email,
			RulesCombine(UserEmailRule, validation.By(NotEquals(d.Password, "email and pw")))...),
		validation.Field(&d.Password, UserPasswordRule...),
	)
}
func (d *CreateUserDto) MapToSchema(hashedPw string) *CreateUserMappedDto {
	return &CreateUserMappedDto{
		Email:    d.Email,
		HashedPw: &hashedPw,
	}
}

/* ****LoginDto
 */
type LoginDto struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func NewLoginDto(email *string, pw *string) *LoginDto {
	return &LoginDto{
		Email:    email,
		Password: pw,
	}
}
func (d LoginDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Email,
			RulesCombine(UserEmailRule, validation.By(NotEquals(d.Password, "email and pw")))...),
		validation.Field(&d.Password, UserPasswordRule...),
	)
}

/* ****RefreshAccessTokenDto
 */
type RefreshAccessTokenDto struct {
	RefreshToken *string `json:"refreshToken"`
}

func NewRefreshAccessTokenDto(rt *string) *RefreshAccessTokenDto {
	return &RefreshAccessTokenDto{
		RefreshToken: rt,
	}
}
func (d RefreshAccessTokenDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.RefreshToken, validation.NotNil),
	)
}

/* ****UpdateUserPasswordDto
 */
type UpdateUserPasswordDto struct {
	CurrentPassword *string `json:"currentPassword"`
	NewPassword     *string `json:"newPassword"`
}

func NewUpdateUserPasswordDto(oldP *string, newP *string) *UpdateUserPasswordDto {
	return &UpdateUserPasswordDto{
		CurrentPassword: oldP,
		NewPassword:     newP,
	}
}
func (d UpdateUserPasswordDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.CurrentPassword,
			RulesCombine(UserPasswordRule, validation.By(NotEquals(d.NewPassword, "current and new pw")))...),
		validation.Field(&d.NewPassword, UserPasswordRule...),
	)
}

/* ****QueryUsersDto
 */
type QueryUsersDto struct {
	Paging
}

func NewQueryUsersDto(paging Paging) *QueryUsersDto {
	return &QueryUsersDto{
		Paging: paging,
	}
}

type QueryUsersResponseDto struct {
	Data           []*ent.User `json:"users"`
	PagingResponse `json:""`
}

func NewQueryUsersResponseDto(data []*ent.User, paging PagingResponse) *QueryUsersResponseDto {
	return &QueryUsersResponseDto{
		Data:           data,
		PagingResponse: paging,
	}
}
