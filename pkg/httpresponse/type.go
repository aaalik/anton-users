package httpresponse

import "github.com/sirupsen/logrus"

type (
	Response struct {
		RequestId string         `json:"request_id"`
		Content   interface{}    `json:"content,omitempty"`
		Error     *ErrorWithCode `json:"error,omitempty"`
		Status    int            `json:"status"`
	}
	ErrorWithCode struct {
		Code    int               `json:"code,omitempty"`
		Message string            `json:"message"`
		Reasons map[string]string `json:"reasons,omitempty"`
	}

	HttpResponse struct {
		log *logrus.Logger
	}
)
