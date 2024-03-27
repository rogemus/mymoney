package web

import (
	"encoding/json"
	"net/http"
)

type ResBody struct {
	Data string
}

func GetHello(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  hello := ResBody{Data: "hello"}
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(hello)
}
