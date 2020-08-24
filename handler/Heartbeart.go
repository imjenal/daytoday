package handler

import (
	"encoding/json"
	"net/http"
)

type HeartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func Heartbeat() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		writer.Header().Set("Content-Type", "application/json")
		heartbeatResponse := HeartbeatResponse{
			Status: "OK",
			Code:   http.StatusOK,
		}
		json.NewEncoder(writer).Encode(heartbeatResponse)
	}
}
