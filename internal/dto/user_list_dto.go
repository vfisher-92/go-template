package dto

import (
	"net/url"
	"strconv"
)

type UserListDTO struct {
	ServiceDTO
	Sort  string
	Page  int
	Limit int
}

func (dto *UserListDTO) ParseRequest(values url.Values) {
	dto.Sort = values.Get("sort")

	page, err := strconv.Atoi(values.Get("page"))
	if err == nil {
		dto.Page = page
	}

	limit, err := strconv.Atoi(values.Get("limit"))

	if err == nil {
		dto.Limit = limit
	}
}

func (dto *UserListDTO) Validate() {

}
