package main

import (
	"basel2053/ps-board/api"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	listenAddr := flag.String("listenaddr", os.Getenv("PORT"), "Server port")
	http.HandleFunc("/", api.RootHandler)
	fmt.Printf("Server is up on running on http://localhost%s\n", *listenAddr)
	err := http.ListenAndServe(*listenAddr, nil)
	log.Fatal(err)
}
