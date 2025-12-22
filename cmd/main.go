package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("%s /users", http.MethodGet), func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"id":   "1",
			"name": "John Doe",
		})

	})
	mux.HandleFunc(fmt.Sprintf("%s /users/{id}", http.MethodGet), func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()

		for key, value := range queries {
			fmt.Print(key, value, "\n")
			// fmt.Printf("Query param: %s=%s\n", key, value[0])
		}

		fmt.Printf("Handling GET /users/%s request\n", r.PathValue("id"))
		w.Header().Set("Content-Type", "application/json")

		//devolver el libro creado
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"id":   "1",
			"name": "John Doe",
		})

	})
	mux.HandleFunc(fmt.Sprintf("%s /users", http.MethodPost), func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc(fmt.Sprintf("%s /users/{id}", http.MethodPatch), func(w http.ResponseWriter, r *http.Request) {})
	root := http.NewServeMux()
	root.Handle("/v1/", http.StripPrefix("/v1", mux))

	fmt.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", root); err != nil {
		fmt.Println("Failed to start server:", err)
	}

}
