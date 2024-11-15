package transport

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func (e ErrorResponse) Write(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		log.Fatal(err)
	}
}

func WriteResponseIfErrorExists(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		e := ErrorResponse{
			Message: err.Error(),
		}
		e.Write(w, statusCode)
		return
	}
}

func WriteJsonResponse[T any](w http.ResponseWriter, data T) {
	b, err := json.Marshal(data)
	if err != nil {
		errorResponse := ErrorResponse{
			Message: err.Error(),
		}
		errorResponse.Write(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		errorResponse := ErrorResponse{
			Message: err.Error(),
		}
		errorResponse.Write(w, http.StatusInternalServerError)
	}
}

func ReadRequest[T any](w http.ResponseWriter, r *http.Request, data *T) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		errorResponse := ErrorResponse{
			Message: err.Error(),
		}
		errorResponse.Write(w, http.StatusBadRequest)
		return
	}
}
