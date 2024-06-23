package httpresponse

import (
	"fmt"
	"net/http"

	"github.com/RoseRocket/xerrs"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
)

func (hr *HttpResponse) Nay(w http.ResponseWriter, r *http.Request, status int, err error) {
	message := err.Error()
	reason := make(map[string]string, 0)
	// fetch error code from Xerrs Data
	if errCodefromErr, ok := xerrs.GetData(err, "status_code"); ok {
		status = errCodefromErr.(int)
	}

	if ev, ok := err.(validation.Errors); ok {
		message = "there's some validation issues in request attributes"
		for field, errObj := range ev {
			reason[field] = errObj.Error()
		}
	}

	// For Now log all failed error
	if status >= 400 {
		hr.log.Error(xerrs.Details(err, 5))
	}

	render.Status(r, status)
	_ = render.Render(w, r, &Response{
		RequestId: middleware.GetReqID(r.Context()),
		Status:    status,
		Error: &ErrorWithCode{
			Code:    status,
			Message: message,
			Reasons: reason,
		},
	})
}

func (hr *HttpResponse) Yay(w http.ResponseWriter, r *http.Request, status int, content interface{}) {
	render.Status(r, status)
	_ = render.Render(w, r, &Response{
		RequestId: middleware.GetReqID(r.Context()),
		Content:   content,
		Status:    status,
	})
}

func (hr *HttpResponse) HTMLYay(w http.ResponseWriter, r *http.Request, status int, content string) {
	render.Status(r, status)
	render.HTML(w, r, content)
}

func (hr *HttpResponse) DataYay(w http.ResponseWriter, r *http.Request, filename string, content []byte) {
	contentDisposition := fmt.Sprintf("attachment;filename=%s", filename)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", contentDisposition)
	w.Write(content)
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
