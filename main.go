package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rajeshpudota/wave/handlers"
)

func main() {
	var appConfig map[string]string
	appConfig, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	
	l := log.New(os.Stdout, "payroll-api", log.LstdFlags)

	//create the handlers
	pr := handlers.NewPayrolls(l)
	r := handlers.NewReports(l)

	// creating a new serve mux
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/reports", r.GetReports)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/payrolls", pr.AddPayrolls)


	// create http server

	s := http.Server{
		Addr: ":9090",
		Handler: sm, 
		ErrorLog: l,
		ReadTimeOut: 5 * time.Second,
		WriteTimeOut: 10* time.Second,
		IdleTimeOut: 120 * time.Second,
	}


	// start the server
	go func() {
		l.Prefix("Startinig server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s \n", err)
			os.Exit(1)
		}
	}

	//Gracefully shutdown the server
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Block until a signal is received
	sig := <-c
	log.Println("Got signal:", sig)

	//Wait till other tasks to complete before we shutdown the server
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
