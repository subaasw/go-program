package httpExample

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strconv"
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}
// 	w.Write([]byte("Hello from snippetBox"))
// }

// // Add snippetView Handler function.
// func snippetView(w http.ResponseWriter, r *http.Request) {
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id < 1 {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
// 	// w.Write([]byte("Display a specific snippet..."))
// }

// // Add a snippetCreate handler function.
// func snippetCreate(w http.ResponseWriter, r *http.Request) {

// 	// Use r.Method to check whether the request is using POST or not.
// 	if r.Method != http.MethodPost {
// 		w.Header().Set("Allow", http.MethodPost)
// 		// w.WriteHeader(405)
// 		// w.Write([]byte("Method not allowed"))
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	w.Write([]byte("Create a new snippet..."))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", home)
// 	mux.HandleFunc("/snippet/view", snippetView)
// 	mux.HandleFunc("/snippet/create", snippetCreate)

// 	log.Println("Starting server on :4000")
// 	err := http.ListenAndServe(":4000", mux)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
