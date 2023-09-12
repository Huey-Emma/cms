package person

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/huey-emma/cms/internal/utils/lib"
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
	case r.Method == http.MethodPost && homePath.Match([]byte(r.URL.Path)):
		h.InsertPerson(w, r)
		return
	case r.Method == http.MethodGet && homeIdPath.Match([]byte(r.URL.Path)):
		h.FindPerson(w, r)
		return
	case r.Method == http.MethodPatch && homeIdPath.Match([]byte(r.URL.Path)):
		h.UpdatePerson(w, r)
		return
	case r.Method == http.MethodDelete && homeIdPath.Match([]byte(r.URL.Path)):
		h.DeletePerson(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) InsertPerson(w http.ResponseWriter, r *http.Request) {
	input := make(map[string]string)

	if err := lib.ReadJSON(w, r, &input); err != nil {
		lib.WriteError(w, 422, err)
		return
	}

	if len(input) == 0 {
		lib.WriteError(w, 422, fmt.Errorf("no content"))
		return
	}

	for k, v := range input {
		if len(strings.TrimSpace(v)) == 0 {
			lib.WriteError(w, 422, fmt.Errorf("%s cannot be blank", k))
			return
		}
	}

	person, err := h.service.InsertPerson(r.Context(), input)
	if err != nil {
		panic(err)
	}

	header := http.Header{
		"Location": []string{"/api/" + fmt.Sprint(person["id"])},
	}

	lib.WriteJSON(w, 201, header, person)
}

func (h *Handler) FindPerson(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(param(r))

	if id == 0 {
		lib.WriteError(w, 404, ErrResourceNotFound)
		return
	}

	person, err := h.service.FindPerson(r.Context(), id)

	if errors.Is(err, ErrResourceNotFound) {
		lib.WriteError(w, 404, err)
		return
	}

	if err != nil {
		panic(err)
	}

	lib.WriteJSON(w, 200, make(http.Header), person)
}

func (h *Handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(param(r))

	if id == 0 {
		lib.WriteError(w, 404, ErrResourceNotFound)
		return
	}

	person, err := h.service.FindPerson(r.Context(), id)

	if errors.Is(err, ErrResourceNotFound) {
		lib.WriteError(w, 404, err)
		return
	}

	if err != nil {
		panic(err)
	}

	input := make(map[string]string)

	if err := lib.ReadJSON(w, r, &input); err != nil {
		lib.WriteError(w, 422, err)
		return
	}

	if len(input) == 0 {
		lib.WriteError(w, 422, fmt.Errorf("no content"))
		return
	}

	for k, v := range input {
		if len(strings.TrimSpace(v)) == 0 {
			lib.WriteError(w, 422, fmt.Errorf("%s cannot be blank", k))
			return
		}

		person[k] = v
	}

	person, err = h.service.UpdatePerson(r.Context(), person)

	if err != nil && !errors.Is(err, ErrResourceNotFound) {
		panic(err)
	}

	lib.WriteJSON(w, 200, make(http.Header), person)
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(param(r))

	if id == 0 {
		lib.WriteError(w, 404, ErrResourceNotFound)
		return
	}

	err := h.service.DeletePerson(r.Context(), id)
	if errors.Is(err, ErrResourceNotFound) {
		lib.WriteError(w, 404, err)
		return
	}

	if err != nil {
		panic(err)
	}

	w.WriteHeader(204)
}

func param(r *http.Request) string {
	match := homeIdPath.FindStringSubmatch(r.URL.Path)

	if len(match) != 2 {
		panic("invalid path parameter")
	}

	return match[1]
}
