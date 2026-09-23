package db_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"server-tp/servidor-tpespecial/db/sqlc"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *sqlc.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	connStr := "postgres://tp_user:tp_password@localhost:5432/tp_db?sslmode=disable"
	var err error

	testDB, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("No se pudo abrir la conexión: %v", err)
	}

	if err = testDB.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a PostgreSQL: %v", err)
	}

	testQueries = sqlc.New(testDB)

	code := m.Run()

	testDB.Close()
	os.Exit(code)
}

var equipoTest sqlc.Equipo

func TestCreateEquipo(t *testing.T) {
	ctx := context.Background()
	var err error

	equipoTest, err = testQueries.CreateEquipo(ctx, sqlc.CreateEquipoParams{
		NombreEquipo: "Los Pumas",
		Division:     "A",
	})

	if err != nil {
		t.Fatalf("error al crear equipo: %v", err)
	}
	if equipoTest.NombreEquipo != "Los Pumas" {
		t.Errorf("nombre incorrecto: got %q, want %q", equipoTest.NombreEquipo, "Los Pumas")
	}
}

func TestGetEquipo(t *testing.T) {
	ctx := context.Background()

	obtenido, err := testQueries.GetEquipo(ctx, equipoTest.IDEquipo)
	if err != nil {
		t.Fatalf("error al obtener equipo: %v", err)
	}
	if obtenido.IDEquipo != equipoTest.IDEquipo {
		t.Errorf("ID incorrecto: got %d, want %d", obtenido.IDEquipo, equipoTest.IDEquipo)
	}
}

func TestUpdateEquipo(t *testing.T) {
	ctx := context.Background()

	err := testQueries.UpdateEquipo(ctx, sqlc.UpdateEquipoParams{
		IDEquipo:     equipoTest.IDEquipo,
		NombreEquipo: "Los Pumas Actualizados",
		Division:     "B",
	})
	if err != nil {
		t.Fatalf("error al actualizar equipo: %v", err)
	}
}

func TestDeleteEquipo(t *testing.T) {
	ctx := context.Background()

	err := testQueries.DeleteEquipo(ctx, equipoTest.IDEquipo)
	if err != nil {
		t.Fatalf("error al eliminar equipo: %v", err)
	}
}
