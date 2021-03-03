package model

import (
	"fmt"
	"net/http"
)

type ErrAPI struct {
	Code    int
	Message string
}

func (e ErrAPI) Error() string {
	return fmt.Sprintf("code:%v error:%v", e.Code, e.Message)
}

var ErrAPINotfound = ErrAPI{
	Code:    http.StatusNotFound,
	Message: "not found",
}

var ErrAPIBadRequest = ErrAPI{
	Code:    http.StatusBadRequest,
	Message: "bad request",
}
