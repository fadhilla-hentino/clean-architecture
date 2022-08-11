package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	errorMessage string
}

func ReadEntity(r *http.Request, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

func WriteOKResponse(w http.ResponseWriter, errorCode int, model interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	if err := json.NewEncoder(w).Encode(&model); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

func WriteOKNoContentResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNoContent)
}

func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	if err := json.NewEncoder(w).Encode(&ErrorResponse{errorMessage: errorMessage}); err != nil {

	}
}
