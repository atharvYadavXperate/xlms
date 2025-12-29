package cutomeerror

import (
	"encoding/json"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	appErr, ok := err.(*AppError)
	if !ok {
		appErr = ErrInternal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.StatusCode)
	json.NewEncoder(w).Encode(appErr)
}
