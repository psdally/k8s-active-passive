// [START kubernetes_engine_hello_app]
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
        "time"
        "math/rand"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/health", health)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

        // seed the random number generator
        rand.Seed(time.Now().UnixNano())

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

// In a real implementation, this would contain the logic to determine if
// the application was healthy, instead of just using a random 10% probability
func health(w http.ResponseWriter, r *http.Request) {

    max := 5
    x := rand.Intn(max) 

    if (x == 1) {
       w.WriteHeader(http.StatusInternalServerError)
       fmt.Fprintf(w, "500 - Something bad happened\n")
    } else {
       fmt.Fprintf(w, "200 - PASS.  Random number: %d\n",x)
    }

}

// hello responds to the request with a plain-text "Hello, world" message.
func hello(w http.ResponseWriter, r *http.Request) {

	log.Printf("Serving request: %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}
// [END kubernetes_engine_hello_app]