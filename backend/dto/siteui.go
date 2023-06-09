package dto

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpsertSiteUiDto struct {
	Sitename          *string `json:"sitename"`
	HomepageImgUrl    *string `json:"homepageImgUrl"`
	HomepageText      *string `json:"homepageText"`
	HomepageTextColor *string `json:"homepageTextColor"`
}

func NewUpsertSiteUiDto(
	sitename *string,
	homepageImgSrc *string,
	homepageText *string,
	homepageTextColor *string,
) *UpsertSiteUiDto {
	return &UpsertSiteUiDto{
		Sitename:          sitename,
		HomepageImgUrl:    homepageImgSrc,
		HomepageText:      homepageText,
		HomepageTextColor: homepageTextColor,
	}
}

func (d UpsertSiteUiDto) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Sitename, SiteNameRule...),
		validation.Field(&d.HomepageImgUrl, HomepageImgUrlRule...),
		validation.Field(&d.HomepageText, HomepageTextRule...),
		validation.Field(&d.HomepageTextColor, HomepageTextColorRule...),
	)
}
