package main

import (
	"fmt"
	"net/http"
)

func main() {
	html404 := `<!DOCTYPE html>
	<html lang="es">
	<head>
    	<meta charset="UTF-8">
    	<meta name="viewport" content="width=device-width, initial-scale=1.0">
    	<title>Página no encontrada</title>
	</head>
	<body>
    	<div class="contenedor">
        	<h1>Error 404</h1>
        	<p>La página que estás buscando no existe.</p>
        	<a href="/">Volver al inicio</a>
    	</div>
	</body>
	</html>`

	// 1. Define el directorio que contiene los archivos estáticos.
	staticDir := "./static"

	fileServer := http.FileServer(http.Dir(staticDir))

	// 3. Registra el manejador para que atienda todas las peticiones ("/").
	// Usamos http.Handle porque fileServer es un http.Handler.

	http.Handle("/", fileServer)

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

func manejador (w http.ResponseWriter, r *http.Request) https.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}
}