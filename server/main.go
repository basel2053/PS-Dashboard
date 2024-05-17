package main

import (
	"basel2053/ps-board/api"
	"basel2053/ps-board/db"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	dbpool, err := db.NewPG(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	listenAddr := flag.String("listenaddr", os.Getenv("PORT"), "Server port")
	http.HandleFunc("/", api.RootHandler)
	fmt.Printf("Server is up on running on http://localhost%s\n", *listenAddr)
	err = http.ListenAndServe(*listenAddr, nil)
	log.Fatal(err)
}
