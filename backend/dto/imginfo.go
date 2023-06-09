package dto

import "sthl/ent"

// ****CreateImgDto
type CreateImgDto struct {
	ImgName    *string `json:"imgName"`
	ImgURL     *string `json:"imgUrl"`
	ImgSize    *int64  `json:"imgSize"`
	ImgS3IdKey *string `json:"imgS3IdKey"`
}

func NewCreateImgDto(imgName *string, imgURL *string, imgSize *int64, imgS3IdKey *string) *CreateImgDto {
	return &CreateImgDto{
		ImgName:    imgName,
		ImgURL:     imgURL,
		ImgSize:    imgSize,
		ImgS3IdKey: imgS3IdKey,
	}
}

// ****QueryImgsInfoDto
type QueryImgsInfoDto struct {
	Paging
}

func NewQueryImgsInfoDto(paging Paging) *QueryImgsInfoDto {
	return &QueryImgsInfoDto{
		Paging: paging,
	}
}

type QueryImgsInfoResponseDto struct {
	Data           []*ent.Imageinfo `json:"imgs"`
	PagingResponse `json:""`
}

func NewQueryImgsInfoResponseDto(data []*ent.Imageinfo, paging PagingResponse) *QueryImgsInfoResponseDto {
	return &QueryImgsInfoResponseDto{
		Data:           data,
		PagingResponse: paging,
	}
}

// ****UpdateImgInfoDto
type UpdateImgInfoDto struct {
	// ImgName *string `json:"imgName"`
	ImgSize *int64 `json:"imgSize"`
}

func NewUpdateImgInfoDto(imgSize *int64) *UpdateImgInfoDto {
	return &UpdateImgInfoDto{
		// ImgName: imgName,
		ImgSize: imgSize,
	}
}

// func (d UpdateImgInfoDto) Validate() error {
// 	return validation.ValidateStruct(&d,
// 		validation.Field(&d.Name, ProductNameRule...),
// 		validation.Field(&d.Price, ProductPriceRule...),
// 		validation.Field(&d.Quantity, ProductQuantityRule...),
// 		validation.Field(&d.Description, ProductDescriptionRule...),
// 		validation.Field(&d.Status, ProductStatusRule...),
// 	)
// }
