package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"

	"github.com/lizzy596/golang-mongo/pkg/controller"
	"github.com/lizzy596/golang-mongo/pkg/options"
)

var (
	server *http.Server
)

func RunServer(ctx context.Context, cfs *options.ServerFlags) {
	mux := http.NewServeMux() // Create Main Handler
	setupControllers(mux)     // Setup all controllers for server

	server = &http.Server{
		Addr:         cfs.Host + ":" + fmt.Sprintf("%d", cfs.Port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if cfs.CertPath != "" && cfs.KeyPath != "" {
		go runTLS(cfs.CertPath, cfs.KeyPath)
	} else {
		go run()
	}
}

func runTLS(certPath, keyPath string) {
	log.Println("Starting TLS Server...")
	log.Fatal(server.ListenAndServeTLS(certPath, keyPath))
}

func run() {
	log.Println("Starting non-TLS Server...")
	log.Fatal(server.ListenAndServe())
}

func setupControllers(mux *http.ServeMux) {
	// controller.SetupUserHandler(mux)
	mux.HandleFunc("/createProfile", controller.createProfile)
	mux.HandleFunc("/404", controller.NotFound)
	// Add more like so:
	// controller.SetupNAMEHandler(mux)
}

func Close() {
	server.Close()
}
