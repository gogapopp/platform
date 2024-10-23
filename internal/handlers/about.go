package handlers

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
