package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ekhabarov/plz-crosscompile-issue/app/prober"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, prober.Probe())
}

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
