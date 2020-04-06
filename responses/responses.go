package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, success bool, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(struct {
		Success bool        `json:success`
		Data    interface{} `json:data`
	}{
		Success: success,
		Data:    data,
	})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, false, struct {
			Success bool   `json:success`
			Error   string `json:"error"`
		}{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
}
