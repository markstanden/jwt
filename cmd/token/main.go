package main

import (
	"fmt"
	"io"
	"os"

	"github.com/markstanden/token"
)

//var jwtDotIO = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.5YLJzijebmz3VWxPTZugJh1JA-q60WOmsEeSl8Ra55cqAdY7wq0mahPDS3U_912j0LRGL7LEEeO1c57muxgTzg"
var secret = "supersecret"

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	jwt, err := token.Create(secret, "My Server", "User Idenfier", "My Website", "Unique ID for the JWT", "Key ID (to identify secret version used to encrypt signature", int64(10000000))
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Created Token: ", jwt)

	tkn := token.Token{}
	err = token.Decode(jwt, secret, &tkn)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Raw Data: ", tkn)
	return nil
}
