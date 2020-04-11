package shared

import (
	""
)

func ResponseHandler(w http.ResponseWriter, result interface{}, apiError *APIError) {
	res := []byte("Hello World")
	w.WriteHeader(500)
	w.Write(res)
}