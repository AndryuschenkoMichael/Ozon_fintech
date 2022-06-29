package handler

import (
	"Ozon_fintech/pkg/service"
	"Ozon_fintech/pkg/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) getFullLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("link")

	if err := h.service.ValidateLink(shortLink); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s: link must contains only symbols [a-z, A-Z, _] and length = %d",
			err.Error(), service.LengthLink)
		return
	}

	fullLink, err := h.service.GetFullLink(shortLink)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s: %s", err.Error(), "short link is not defined")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", fullLink)
	}
}

func (h *Handler) postLink(w http.ResponseWriter, r *http.Request) {
	var linkInfo struct {
		FullLink string `json:"full_link"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&linkInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "incorrect body format")
		return
	}

	shortLink, err := h.service.SetShortLink(linkInfo.FullLink)
	if err != nil && err.Error() == storage.KeyExistError {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s %s", "short link is already in system. short link:", shortLink)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", shortLink)
	}
}
