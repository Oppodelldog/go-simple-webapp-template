package page

import (
	"io"
	"net/http"
	"time"
)

func RenderIndexPage(w io.Writer, r *http.Request) error {
	t, err := NewPageTemplate("index.html")
	if err != nil {
		return err
	}

	return t.Execute(w,
		struct {
			TheTime time.Time
		}{
			TheTime: time.Now(),
		})
}
