package http

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lucat1/me/templates"
)

type PageDataData struct {
	Title string
	Body  template.HTML
}

type PageData struct {
	Data      PageDataData
	RequestID uuid.UUID
	Timestamp time.Time
}

const ROOT_TEMPLATE = "root"

func RenderPage[T any](w http.ResponseWriter, r *http.Request, name, title string, data T) (err error) {
	body, err := RenderBlock[T](name, data)
	if err != nil {
		err = fmt.Errorf("Error while rendering page (non root): %v", err)
		return
	}

	templates := templates.Get()
	pageData := PageData{
		Data: PageDataData{
			Title: title,
			Body:  template.HTML(body),
		},
		RequestID: GetRequestId(r),
		Timestamp: time.Now(),
	}
	err = templates.ExecuteTemplate(w, ROOT_TEMPLATE, pageData)
	if err != nil {
		err = fmt.Errorf("Error while rendering page root: %v", err)
	}
	return
}

func RenderBlock[T any](name string, data T) (block []byte, err error) {
	templates := templates.Get()
	buffer := bytes.NewBuffer([]byte{})
	err = templates.ExecuteTemplate(buffer, name, data)
	if err != nil {
		err = fmt.Errorf("Error while rendering block: %v", err)
	}
	block = buffer.Bytes()
	return
}
