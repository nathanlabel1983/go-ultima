package api

import (
	"encoding/json"
	"net/http"

	"github.com/nathanlabel1983/go-ultima/pkg/services/tcpserver"
)

type API struct {
	server    *tcpserver.TCPServer // The TCP server
	webserver *http.Server         // The web server
	running   bool                 // Is the service running?
}

func NewAPI(server *tcpserver.TCPServer) *API {
	return &API{
		server: server,
		webserver: &http.Server{
			Addr:    ":8080",
			Handler: nil,
		},
	}
}

// GetServerStatus returns the server status
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

// StartAPI starts the API Web Service
func (api *API) Start() error {
	getServerStatus := func(w http.ResponseWriter, r *http.Request) {
		api.GetServerStatus(api.server, w, r)
	}

	api.webserver.HandleFunc("/GetServerStatus", getServerStatus)
	go http.ListenAndServe(":8080", nil)
	api.running = true
	return nil
}

func (api *API) Stop() error {
	api.running = false
	return nil
}
