package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/services/employee/pkg/server"
)

func SetTransactionIDInContext(ctx context.Context, transactionID string) context.Context {
	if transactionID == "" {
		transactionID = uuid.New().String()
	}

	return context.WithValue(ctx, "transactionId", transactionID)
}

func main() {
	fmt.Println("inside main.go")
	shutdownChannel := make(chan struct{})

	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	ctx := context.Background()
	ctx = SetTransactionIDInContext(ctx, "")
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-signals
		fmt.Println("Initiate graceful shutdown of web servers")
		close(shutdownChannel)
	}()
	//call server
	server.Start(ctx, shutdownChannel)
}
