package metrics

import "time"

type HTTPBase struct {
	Handler string
	Method  string
}

type HTTPError struct {
	HTTPBase
	Status int
}

type HTTPDuration struct {
	HTTPBase
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   float64
}

func NewHTTPBase(handler string, method string) *HTTPBase {
	return &HTTPBase{
		Handler: handler,
		Method:  method,
	}
}

func NewHTTPError(handler string, method string, status int) *HTTPError {
	http := NewHTTPBase(handler, method)
	return &HTTPError{
		HTTPBase: *http,
		Status:   status,
	}
}
func NewHTTPDuration(handler string, method string) *HTTPDuration {
	http := NewHTTPBase(handler, method)
	return &HTTPDuration{
		HTTPBase: *http,
	}
}

//Started start monitoring the app
func (h *HTTPDuration) Started() {
	h.StartedAt = time.Now()
}

// Finished app finished
func (h *HTTPDuration) Finished() {
	h.FinishedAt = time.Now()
	h.Duration = time.Since(h.StartedAt).Seconds()
}
