package user

import (
	"net/http"

	"github.com/elliot14A/tcorp/assessment/pkg/helper"
)

func (h *httpHandler) List(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.List()

	if err != nil {
		_, message := helper.ErrorStatusConverter(err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	writeResponse(w, users, http.StatusOK)
}
