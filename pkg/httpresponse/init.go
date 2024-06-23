package httpresponse

import "github.com/sirupsen/logrus"

func NewHttpResponse(log *logrus.Logger) *HttpResponse {
	return &HttpResponse{log}
}
