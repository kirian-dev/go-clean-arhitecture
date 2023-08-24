package main

import (
	"fmt"
	"go-clean-architecture/pkg/storage/postgres"
	deliveryHttp "go-clean-architecture/services/contact/internal/delivery/http"
	repoStorage "go-clean-architecture/services/contact/internal/repository/storage/postgres"
	useCaseContact "go-clean-architecture/services/contact/internal/useCase/contact"
	useCaseGroup "go-clean-architecture/services/contact/internal/useCase/group"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	fmt.Println(conn.Pool.Stat())
	fmt.Println("Connection to postgres")

	var (
		repoStorage  = repoStorage.New(conn.Pool, repoStorage.Options{})
		contactUC    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		groupUC      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		httpListener = deliveryHttp.New(contactUC, groupUC, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on port %v\n", viper.GetUint("HTTP_PORT"))
		if err = httpListener.Run(); err != nil {
			fmt.Printf("service failed to start")
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
