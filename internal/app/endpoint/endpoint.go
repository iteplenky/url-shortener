package endpoint

import (
	"fmt"
	"github.com/iteplenky/url-shortener/internal/app/generator"
	"github.com/iteplenky/url-shortener/internal/app/store"
	"io"
	"net/http"
)

const RandomStringLength = 8

type Endpoint struct {
	g generator.Generator
	s *store.Store
}

func New(g generator.Generator, s *store.Store) *Endpoint {
	return &Endpoint{
		g: g,
		s: s,
	}
}

func (ep *Endpoint) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id := ep.g.RandomString(RandomStringLength)

	body, _ := io.ReadAll(r.Body)
	ep.s.Set(id, string(body))

	shortenURL := CorrectURL(r) + id

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(shortenURL))
}

func (ep *Endpoint) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")
	if len(id) == 0 {
		http.Error(w, "expected id in url", http.StatusBadRequest)
		return
	}

	val, ok := ep.s.Get(id)
	if !ok {
		http.Error(w, "not found provided id", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, val, http.StatusTemporaryRedirect)
}

func CorrectURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)
}
