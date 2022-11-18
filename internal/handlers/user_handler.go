package handlers

import (
	"go-template/internal/dto"
	"net/http"
)

type UserHandler struct {
	Handler
}

func NewUserHandler(
	appHandler Handler,
) *UserHandler {
	return &UserHandler{
		Handler: appHandler,
	}
}

func (h *UserHandler) getUsers(w http.ResponseWriter, r *http.Request) {
	var response ResponseData

	err := r.ParseForm()
	if err != nil {
		response.ClientErrors = err
		h.httpResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	ulDto := dto.UserListDTO{}
	ulDto.ParseRequest(r.Form)

	userLogs, count, err := h.userService.GetList(ulDto)
	if err != nil {
		response.Code = http.StatusNotFound
		response.ClientErrors = err
		h.httpResponseJSON(w, response.Code, response)
		return
	}

	response.Data = userLogs
	response.Meta = struct {
		Total       interface{} `json:"total"`
		CurrentPage int         `json:"current_page"`
	}{
		Total:       count,
		CurrentPage: ulDto.Page,
	}

	h.httpResponseJSON(w, http.StatusOK, response)
}
