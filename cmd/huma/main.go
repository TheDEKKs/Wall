package main

import (
	"fmt"
	"os"
	"thedekk/WWT/internal/env"
	"thedekk/WWT/internal/transport"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

func main() {
	var config env.Config

	config.Load()

	fmt.Println("Start app")
	ctx := context.Background()

	conn, err := pgxpool.New(ctx, config.PostgresURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	
	if err := transport.NewService(conn); err != nil {
		fmt.Println(err)
	}

}	