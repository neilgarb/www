package controllers

import (
	"encoding/json"
	"net/http"
)

func writeJsonResult(w http.ResponseWriter, result string, code int, data interface{}) error {
	payload := &struct {
		Result string      `json:"result"`
		Code   int         `json:"code"`
		Data   interface{} `json:"data"`
	}{result, code, data}
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "text/json")
	w.Write(b)
	return nil
}
