package handler

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerUser) DeleteSoftUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	w.Header().Set("Content-Type", "application/json")

	if err := h.GetUserUseCase.DeleteSoftUser(r.Context(), id); err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"deletedId": id})

}
