package server

import (
	"log"
	"net/http"
	"time"

	"github.com/akhil-skyflow/networktokenize/pkg/services"
	"github.com/gorilla/mux"
)

func StartHttpServer() {
	router := mux.NewRouter()

	//Routes
	router.HandleFunc("/tokenize", services.TokenizeHandler).Methods("POST")
	router.HandleFunc("/getpaymentbundle", services.GetPaymentBundleHandler).Methods("POST")

	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
