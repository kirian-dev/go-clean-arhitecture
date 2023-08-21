package main

import (
	"fmt"
	"go-clean-architecture/pkg/storage/postgres"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	fmt.Println(conn.Pool.Stat())
	fmt.Println("Connection to postgres")
}
