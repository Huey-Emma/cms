package person

import (
	"net/http"
	"regexp"
)

var (
	homePath   = regexp.MustCompile(`^\/api\/?$`)
	homeIdPath = regexp.MustCompile(`^\/api\/(\d+)$`)
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && homePath.Match([]byte(r.URL.Path)):
		h.InsertPerson(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) InsertPerson(w http.ResponseWriter, r *http.Request) {}
