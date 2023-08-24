package ernos

import (
	"net/http"
	"strings"
)

type Ernos struct {
	Status    int    `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	Code      string `json:"code,omitempty"`
	Attribute any    `json:"attribute,omitempty"`
}

func (e Ernos) Error() string {
	return e.Message
}

func NotFound(value ...string) error {
	msg := M.RECORD_NOT_FOUND
	if len(value) > 0 {
		msg = strings.Join(value, "") + " not found"
	}

	return Ernos{
		Status:  http.StatusNotFound,
		Message: msg,
		Code:    C.RECORD_NOT_FOUND,
	}
}

func InternalServerError(value ...string) error {
	msg := M.INTERNAL_SERVER_ERROR
	if len(value) > 0 {
		msg = strings.Join(value, "")
	}

	return Ernos{
		Status:  http.StatusInternalServerError,
		Message: msg,
		Code:    C.INTERNAL_SERVER_ERROR,
	}
}
