package handlers

import (
	"encoding/json"
	"net/http"
)

func replyWithError(err error, w http.ResponseWriter, httpStatus int) {
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})

}
