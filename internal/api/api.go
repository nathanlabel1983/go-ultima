package api

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/nathanlabel1983/go-ultima/pkg/tcpserver"
)

//go:embed templates/index.html
var templatesFS embed.FS

func StartAPI(s *tcpserver.TCPServer) {

	serverStatus := func(w http.ResponseWriter, r *http.Request) {
		status := s.GetStatus()
		jsonData, err := json.Marshal(status)
		if err != nil {
			http.Error(w, "Error converting status to JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}

	mainStatus := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(templatesFS, "templates/index.html")
		if err != nil {
			http.Error(w, "Error rendering template1", http.StatusInternalServerError)
			return
		}

		status := s.GetStatus()

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, status)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error rendering template2: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.Write(buf.Bytes())
	}

	http.HandleFunc("/", mainStatus)
	http.HandleFunc("/GetServerStatus", serverStatus)
	http.ListenAndServe(":8080", nil)
}
