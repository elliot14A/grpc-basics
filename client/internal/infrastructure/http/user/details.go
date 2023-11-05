package user

import (
	"net/http"
	"strconv"

	"github.com/elliot14A/tcorp/assessment/pkg/helper"
)

func (h *httpHandler) Details(w http.ResponseWriter, r *http.Request) {
	userId := helper.GetPathParamByName(r, "userId")
	uId, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	user, err := h.userRepository.Details(uId)

	if err != nil {
		h.logger.Error(err.Error())
		code, message := helper.ErrorStatusConverter(err)
		switch code {
		case "NotFound":
			http.Error(w, "User not found", http.StatusNotFound)
			break
		default:
			http.Error(w, message, http.StatusInternalServerError)
			break
		}
		return
	}

	writeResponse(w, user, http.StatusOK)
}
