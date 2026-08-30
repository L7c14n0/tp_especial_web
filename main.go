package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func manejador404(fileServer http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// El home solamente acepta GET
		if r.URL.Path == "/" && r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Construimos la ruta del archivo solicitado
		ruta := filepath.Join("./static", r.URL.Path)

		// Comprobamos si existe
		_, err := os.Stat(ruta)

		if err != nil {
			if os.IsNotExist(err) {
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, "./static/404.html")
				return
			}

			http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
			return
		}

		// Si existe y el método es válido, usamos FileServer
		fileServer.ServeHTTP(w, r)
	})
}

func main() {
	// 1. Define el directorio que contiene los archivos estáticos.
	staticDir := "./static"

	fileServer := http.FileServer(http.Dir(staticDir))

	// 3. Registra el manejador para que atienda todas las peticiones ("/").
	// Usamos http.Handle porque fileServer es un http.Handler.

	http.Handle("/", manejador404(fileServer))

	// 4. Define el puerto y muestra un mensaje.
	port := ":8080"
	fmt.Printf("Servidor ESTÁTICO escuchando en http://localhost%s\n", port)
	fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)
	// 5. Inicia el servidor.
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
