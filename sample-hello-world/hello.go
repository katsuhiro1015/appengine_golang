package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/usagisan/", handleUsagisan)
	http.HandleFunc("/nukosama/", handleNukosama)
	http.HandleFunc("/kumasan/", handleKumasan)
	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func handleUsagisan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, usagisan!")
}
func handleNukosama(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, nukosama!")
}
func handleKumasan(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, kumasan!")
}
