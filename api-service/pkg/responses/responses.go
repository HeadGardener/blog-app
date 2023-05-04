package responses

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func NewErrResponse(w http.ResponseWriter, code int, errorMsg string, logger *zap.Logger) {
	logger.Error(errorMsg)
	NewResponse(w, code, response{
		Message: errorMsg,
	})
}

func NewResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
