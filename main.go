package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	filePath := filepath.Join("static", "index.html")
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Error interno al leer el archivo HTML", http.StatusInternalServerError)
		log.Printf("Error leyendo archivo: %v", err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(content)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
