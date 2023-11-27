package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	f, err := os.Open(os.Getenv("DATA_JSON"))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	parsed, err := parseJSON(data)
	if err != nil {
		panic(err)
	}
	handler := generateHandler(parsed, os.Getenv("HTML_TEMPLATE"))
	fmt.Println("Server started on localhost:8080")
	http.ListenAndServe(":8080", handler)
}
