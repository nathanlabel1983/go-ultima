package api

import (
	"encoding/json"
	"net/http"

	"github.com/nathanlabel1983/go-ultima/pkg/tcpserver"
)

func GetServerStatus(s *tcpserver.TCPServer, w http.ResponseWriter, r *http.Request) {
	status := s.GetStatus()
	jsonData, err := json.Marshal(status)
	if err != nil {
		http.Error(w, "Error converting status to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func StartAPI(s *tcpserver.TCPServer) {
	getServerStatus := func(w http.ResponseWriter, r *http.Request) {
		GetServerStatus(s, w, r)
	}

	http.HandleFunc("/GetServerStatus", getServerStatus)
	http.ListenAndServe(":8080", nil)
}
