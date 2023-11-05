package http

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/lucat1/me/templates"
)

type PageData[T any] struct {
	Data      T
	RequestID uuid.UUID
	Timestamp time.Time
}

func RenderPage[T any](w http.ResponseWriter, r *http.Request, name string, data T) (err error) {
	templates := templates.Get()
	pageData := PageData[T]{
		Data:      data,
		RequestID: GetRequestId(r),
		Timestamp: time.Now(),
	}
	err = templates.ExecuteTemplate(w, name, pageData)
	if err != nil {
		err = fmt.Errorf("Error while rendering page template: %v", err)
	}
	return
}

func RenderBlock[T any](name string, data PageData[T]) (block []byte, err error) {
	templates := templates.Get()
	buffer := bytes.NewBuffer([]byte{})
	err = templates.ExecuteTemplate(buffer, name, data)
	if err != nil {
		err = fmt.Errorf("Error while rendering page template: %v", err)
	}
	block = buffer.Bytes()
	return
}
