package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	db "server-tp/servidor-tpespecial/db/sqlc"

	_ "github.com/lib/pq"
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
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
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

type EquiposAPI struct {
	Queries *db.Queries //puntero a la estructura generada por sqlc
}

func main() {
	//creamos la conexion con la base de datos
	dsn := "postgres://tp_user:tp_password@localhost:5432/tp_db?sslmode=disable"

	//intentamos conectarnos a la base
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al preparar la conexión: %v", err)
	}

	//si logramos conectaros con la base, antes de que se cierre cualquier tipo de conexion, la cerramos
	defer conn.Close()

	//hacemos un ping para ver si podemos interactuar con la base de datos
	if err := conn.Ping(); err != nil {
		log.Fatalf("La BD rechazó la conexión: %v", err)
	}

	//si el ping pasa y no hay error, la conexion es un exito
	fmt.Println("Conexión a PostgreSQL establecida con éxito.")

	//instanciamos un repositorio usando el codigo autogenerado por sqlc
	queries := db.New(conn)

	//le damos a la API que creamos el repositorio que instanciamos para usar los datos de la bd
	apiEquipos := &EquiposAPI{Queries: queries}

	_ = apiEquipos //

	staticDir := "./static"
	fileServer := http.FileServer(http.Dir(staticDir))

	// Registramos la ruta raíz "/" para que sirva la página web.
	// (Asegurate de tener tu función manejador404 en el proyecto)
	http.Handle("/", manejador404(fileServer))

	port := ":8080"
	fmt.Printf("Servidor Completo escuchando en http://localhost%s\n", port)
	fmt.Printf("Sirviendo frontend desde: %s\n", staticDir)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
