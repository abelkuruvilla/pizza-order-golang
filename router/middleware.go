package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pizza-delivery/apputil"
	"pizza-delivery/model"
)

func responseHandler(h func(http.ResponseWriter, *http.Request) (interface{}, int, error)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		data, status, err := h(w, r)
		if err != nil {
			data = err.Error()
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if data != nil {
			// Send JSON response back to the client application
			err = json.NewEncoder(w).Encode(model.Response{Data: data})
			if err != nil {
				apputil.Error(fmt.Sprintf("could not encode response to output: %v", err))
			}
		}
	})
}
