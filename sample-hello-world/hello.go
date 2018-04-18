// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/usagisan/", handleUsagisan)
	http.HandleFunc("/nukosama/", handleNukosama)
	http.HandleFunc("/kumasan/", handleKumasan)
	appengine.Main()
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
