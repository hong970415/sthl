package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"sthl/constants"
)

type ResponseMessage[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data *T     `json:"data,omitempty"`
}

func newResponseMessage[T any](code int, msg string, data *T) *ResponseMessage[T] {
	return &ResponseMessage[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (rm *ResponseMessage[T]) Send(rw http.ResponseWriter) {
	rw.WriteHeader(rm.Code)
	_ = json.NewEncoder(rw).Encode(rm)
}

func ResponseSend[T any](rw http.ResponseWriter, code int, msg string, data *T) {
	res := newResponseMessage(code, msg, data)
	switch code {
	case http.StatusOK:
		res.Send(rw)
	case http.StatusCreated:
		res.Send(rw)
	case http.StatusBadRequest:
		res.Msg = "bad request"
		res.Data = nil
		res.Send(rw)
	case http.StatusUnauthorized:
		res.Msg = "unauthorized"
		res.Data = nil
		res.Send(rw)
	case http.StatusForbidden:
		res.Msg = "forbidden"
		res.Data = nil
		res.Send(rw)
	case http.StatusNotFound:
		res.Msg = "not found"
		res.Data = nil
		res.Send(rw)
	default:
		res.Msg = "internal server error"
		res.Data = nil
		res.Send(rw)
	}
}
func HttpErrorResponseSend(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, constants.ErrExisted):
		ResponseSend[any](w, http.StatusBadRequest, "", nil)
	case errors.Is(err, constants.ErrNotFound):
		ResponseSend[any](w, http.StatusNotFound, "", nil)
	case errors.Is(err, constants.ErrBadRequest):
		ResponseSend[any](w, http.StatusBadRequest, "", nil)
	case errors.Is(err, constants.ErrUnauthorized):
		ResponseSend[any](w, http.StatusUnauthorized, "", nil)
	default:
		ResponseSend[any](w, http.StatusInternalServerError, "", nil)
	}
}
