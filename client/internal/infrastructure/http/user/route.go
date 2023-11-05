package user

import (
	"encoding/json"
	"net/http"

	"github.com/elliot14A/tcorp/assessment/client/internal/domain/repositories"
	"github.com/elliot14A/tcorp/assessment/pkg/logger"
	"github.com/go-chi/chi"
)

type httpHandler struct {
	userRepository repositories.UserRepository
	logger         logger.LoggerPort
}

func (h *httpHandler) routes() chi.Router {
	router := chi.NewRouter()
	router.Get("/{userId}", h.Details)
	router.Get("/", h.List)
	return router
}

func InitRoutes(r *chi.Mux, userRepo repositories.UserRepository, logger logger.LoggerPort) {
	httpHandler := &httpHandler{userRepository: userRepo, logger: logger}
	r.Mount("/users", httpHandler.routes())
}

func writeResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonReponse, _ := json.Marshal(body)
	w.Write(jsonReponse)
}
