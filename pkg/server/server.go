package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Start(ctx context.Context, shutdownChannel chan struct{}) {

	err := LoadDependencies(ctx)
	defer closeDependencies()
	if err != nil {
		log.Printf("failure in loading dependencies, err is: %v", err)
	} else {
		fmt.Println("dependencies loaded")
	}

	log.Printf("Starting REST Server")
	router := loadRoutes()

	log.Printf("REST Server listening on port 9080")

	srv := http.Server{
		Addr:    ":9080",
		Handler: removeTrailingSlash(router),
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server crashed! Error: %s", err.Error())
		}
	}()

	<-shutdownChannel
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shutdown the server gracefully,err is %v", err)
	}

	fmt.Println("Exiting after gracefully closing dependencies")
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
