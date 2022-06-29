package handler

import (
	"Ozon_fintech/pkg/model"
	"Ozon_fintech/pkg/service"
	"Ozon_fintech/pkg/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Summary Get full link by short form
// @Tags API
// @Description Get full link
// @ID get-full-link
// @Accept plain
// @Produce plain
// @Param link query string false "full link"
// @Success 200 "OK"
// @Failure 400 "Bad request"
// @Failure 404 "Not found"
// @Failure 500 "Internal server error"
// @Router /api/get-full-link [get]
func (h *Handler) getFullLink(w http.ResponseWriter, r *http.Request) {
	shortLink := r.URL.Query().Get("link")

	if err := h.service.ValidateLink(shortLink); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s: link must contains only symbols [a-z, A-Z, 0-9, _] and length = %d",
			err.Error(), service.LengthLink)
		return
	}

	fullLink, err := h.service.GetFullLink(shortLink)
	if err != nil {
		if err.Error() == storage.KeyError {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "%s: short link is not defined", err.Error())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err.Error())
		}
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", fullLink)
	}
}

// @Summary Get short link
// @Tags API
// @Description Post link
// @ID post-link
// @Accept json
// @Produce plain
// @Param linkInfo body model.LinkInfo true "full link"
// @Success 201 "Created"
// @Failure 400 "Bad request"
// @Failure 405 "Method not allowed"
// @Failure 500 "Internal server error"
// @Router /api/post-link [post]
func (h *Handler) postLink(w http.ResponseWriter, r *http.Request) {
	var linkInfo model.LinkInfo
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&linkInfo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "incorrect body format")
		return
	}

	shortLink, err := h.service.SetShortLink(linkInfo.FullLink)
	if err != nil && err.Error() == storage.KeyExistError {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "short link is already in system. short link: %s", shortLink)
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
