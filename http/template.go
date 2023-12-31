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

const (
	ROOT_TEMPLATE     = "root"
	CONTENT_TYPE      = "Content-Type"
	CONTENT_TYPE_HTML = "text/html"
)

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
	w.Header().Add(CONTENT_TYPE, CONTENT_TYPE_HTML)
	err = templates.ExecuteTemplate(w, ROOT_TEMPLATE, pageData)
	if err != nil {
		err = fmt.Errorf("Error while rendering page root: %v", err)
	}
	return
}

func RenderBlockPage[T any](w http.ResponseWriter, r *http.Request, name string, data T) (err error) {
	w.Header().Add(CONTENT_TYPE, CONTENT_TYPE_HTML)
	err = templates.Get().ExecuteTemplate(w, name, data)
	if err != nil {
		err = fmt.Errorf("Error while rendering block page %s: %v", name, err)
	}
	return
}

func RenderBlock[T any](name string, data T) (block []byte, err error) {
	templates := templates.Get()
	buffer := bytes.NewBuffer([]byte{})
	err = templates.ExecuteTemplate(buffer, name, data)
	if err != nil {
		err = fmt.Errorf("Error while rendering block %s: %v", name, err)
	}
	block = buffer.Bytes()
	return
}
