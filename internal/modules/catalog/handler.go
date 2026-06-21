package catalog

import (
	"ecom/internal/shared/httpx"
	"encoding/json"
	"net/http"

	"log/slog"
)

type Handlers struct {
	service CatalogService
	logger  *slog.Logger
}

func NewHandlers(service CatalogService, logger *slog.Logger) *Handlers {
	return &Handlers{service: service, logger: logger}
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var out = &Product{}
	ReadJson(r, out, w)

	created, err := h.service.Create(r.Context(), out)
	if err != nil {
		h.logger.Error("service Create failed", "err", err)
		httpx.SendErrorResponse(w, &httpx.ErrorEnvlop{Code: http.StatusBadRequest, Message: "Faild to create", Details: make(map[string]string, 0)})
		return
	}

	httpx.WriteJson(w, created, http.StatusCreated)
}

func ReadJson(r *http.Request, in *Product, w http.ResponseWriter) error {
	err := json.NewDecoder(r.Body).Decode(&in)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return nil
}
