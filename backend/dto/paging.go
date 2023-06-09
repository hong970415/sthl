package dto

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ****Paging
type Paging struct {
	Page  int
	Limit int
	Query string
}

func ExtractPaging(r *http.Request) *Paging {
	p, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		p = 1
	}

	l, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		l = 20
	}

	q := r.URL.Query().Get("query")

	return NewPaging(p, l, q)
}

func NewPaging(p int, l int, q string) *Paging {
	return &Paging{
		Page:  p,
		Limit: l,
		Query: q,
	}
}

func (d Paging) Validate() error {
	return validation.ValidateStruct(&d,
		validation.Field(&d.Page, validation.Required, validation.Min(1)),
		validation.Field(&d.Limit, validation.Required, validation.Min(1), validation.Max(1000)),
		validation.Field(&d.Query, validation.Length(0, 255)),
	)
}
func (d *Paging) Ensure() *Paging {
	err := d.Validate()
	if err != nil {
		q := d.Query
		if len(q) > 199 {
			q = q[:199]
		}
		return &Paging{
			Page:  1,
			Limit: 20,
			Query: q,
		}
	}
	return &Paging{
		Page:  d.Page,
		Limit: d.Limit,
		Query: d.Query,
	}
}

// ****PagingResponse
type PagingResponse struct {
	Page  int `json:"page"`
	Limit int `json:"limt"`
	Total int `json:"total"`
}

func NewPagingResponse(p int, l int, t int) *PagingResponse {
	return &PagingResponse{
		Page:  p,
		Limit: l,
		Total: t,
	}
}
