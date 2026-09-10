package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func conectarDB(t *testing.T) *sql.DB {
	t.Helper()

	connStr := "postgres://tp_user:tp_password@localhost:5432/tp_db?sslmode=disable"

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("error al abrir conexión: %v", err)
	}

	ctx := context.Background()

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("error al conectar con PostgreSQL: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
	})

	return db
}

func TestEquipoCRUD(t *testing.T) {
	db := conectarDB(t)
	queries := New(db)
	ctx := context.Background()

	equipo, err := queries.CreateEquipo(ctx, CreateEquipoParams{
		NombreEquipo: "Los Pumas",
		Division:     "A",
	})
	if err != nil {
		t.Fatalf("error al crear equipo: %v", err)
	}

	if equipo.NombreEquipo != "Los Pumas" {
		t.Errorf("nombre incorrecto: got %q, want %q", equipo.NombreEquipo, "Los Pumas")
	}

	obtenido, err := queries.GetEquipo(ctx, equipo.IDEquipo)
	if err != nil {
		t.Fatalf("error al obtener equipo: %v", err)
	}

	if obtenido.IDEquipo != equipo.IDEquipo {
		t.Errorf("ID incorrecto: got %d, want %d", obtenido.IDEquipo, equipo.IDEquipo)
	}

	equipos, err := queries.ListEquipos(ctx)
	if err != nil {
		t.Fatalf("error al listar equipos: %v", err)
	}

	encontrado := false

	for _, e := range equipos {
		if e.IDEquipo == equipo.IDEquipo {
			encontrado = true
			break
		}
	}

	if !encontrado {
		t.Errorf("el equipo creado no apareció en ListEquipos")
	}

	err = queries.UpdateEquipo(ctx, UpdateEquipoParams{
		IDEquipo:     equipo.IDEquipo,
		NombreEquipo: "Los Pumas Actualizados",
		Division:     "B",
	})
	if err != nil {
		t.Fatalf("error al actualizar equipo: %v", err)
	}

	actualizado, err := queries.GetEquipo(ctx, equipo.IDEquipo)
	if err != nil {
		t.Fatalf("error al obtener equipo actualizado: %v", err)
	}

	if actualizado.NombreEquipo != "Los Pumas Actualizados" {
		t.Errorf("nombre no actualizado correctamente")
	}

	if actualizado.Division != "B" {
		t.Errorf("division no actualizada correctamente")
	}

	err = queries.DeleteEquipo(ctx, equipo.IDEquipo)
	if err != nil {
		t.Fatalf("error al eliminar equipo: %v", err)
	}

	_, err = queries.GetEquipo(ctx, equipo.IDEquipo)
	if err == nil {
		t.Errorf("se esperaba un error al obtener el equipo eliminado")
	}
}

func TestJugadorCRUD(t *testing.T) {
	db := conectarDB(t)
	queries := New(db)
	ctx := context.Background()

	equipo, err := queries.CreateEquipo(ctx, CreateEquipoParams{
		NombreEquipo: "Equipo Jugador Test",
		Division:     "A",
	})
	if err != nil {
		t.Fatalf("error al crear equipo: %v", err)
	}

	t.Cleanup(func() {
		queries.DeleteEquipo(ctx, equipo.IDEquipo)
	})

	jugador, err := queries.CreateJugador(ctx, CreateJugadorParams{
		Nombre:   "Juan",
		Apellido: "Perez",
		IDEquipo: equipo.IDEquipo,
	})
	if err != nil {
		t.Fatalf("error al crear jugador: %v", err)
	}

	if jugador.Nombre != "Juan" || jugador.Apellido != "Perez" {
		t.Errorf("datos incorrectos del jugador creado")
	}

	obtenido, err := queries.GetJugador(ctx, jugador.IDJugador)
	if err != nil {
		t.Fatalf("error al obtener jugador: %v", err)
	}

	if obtenido.IDJugador != jugador.IDJugador {
		t.Errorf("ID incorrecto del jugador")
	}

	jugadores, err := queries.ListJugadores(ctx)
	if err != nil {
		t.Fatalf("error al listar jugadores: %v", err)
	}

	encontrado := false

	for _, j := range jugadores {
		if j.IDJugador == jugador.IDJugador {
			encontrado = true
			break
		}
	}

	if !encontrado {
		t.Errorf("el jugador creado no apareció en ListJugadores")
	}

	err = queries.UpdateJugador(ctx, UpdateJugadorParams{
		IDJugador: jugador.IDJugador,
		Nombre:    "Juan Carlos",
		Apellido:  "Perez Actualizado",
		IDEquipo:  equipo.IDEquipo,
	})
	if err != nil {
		t.Fatalf("error al actualizar jugador: %v", err)
	}

	actualizado, err := queries.GetJugador(ctx, jugador.IDJugador)
	if err != nil {
		t.Fatalf("error al obtener jugador actualizado: %v", err)
	}

	if actualizado.Nombre != "Juan Carlos" {
		t.Errorf("nombre del jugador no actualizado")
	}

	if actualizado.Apellido != "Perez Actualizado" {
		t.Errorf("apellido del jugador no actualizado")
	}

	err = queries.DeleteJugador(ctx, jugador.IDJugador)
	if err != nil {
		t.Fatalf("error al eliminar jugador: %v", err)
	}

	_, err = queries.GetJugador(ctx, jugador.IDJugador)
	if err == nil {
		t.Errorf("se esperaba un error al obtener el jugador eliminado")
	}
}

func TestPartidoCRUD(t *testing.T) {
	db := conectarDB(t)
	queries := New(db)
	ctx := context.Background()

	local, err := queries.CreateEquipo(ctx, CreateEquipoParams{
		NombreEquipo: "Local Test",
		Division:     "A",
	})
	if err != nil {
		t.Fatalf("error al crear equipo local: %v", err)
	}

	visitante, err := queries.CreateEquipo(ctx, CreateEquipoParams{
		NombreEquipo: "Visitante Test",
		Division:     "A",
	})
	if err != nil {
		t.Fatalf("error al crear equipo visitante: %v", err)
	}

	t.Cleanup(func() {
		queries.DeleteEquipo(ctx, local.IDEquipo)
		queries.DeleteEquipo(ctx, visitante.IDEquipo)
	})

	partido, err := queries.CreatePartido(ctx, CreatePartidoParams{
		IDEquipoLocal:     local.IDEquipo,
		IDEquipoVisitante: visitante.IDEquipo,
		GolesLocal:        3,
		GolesVisitante:    1,
		Fecha:             time.Date(2026, 8, 20, 20, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("error al crear partido: %v", err)
	}

	if partido.GolesLocal != 3 || partido.GolesVisitante != 1 {
		t.Errorf("resultado incorrecto del partido")
	}

	obtenido, err := queries.GetPartido(ctx, partido.IDPartido)
	if err != nil {
		t.Fatalf("error al obtener partido: %v", err)
	}

	if obtenido.IDPartido != partido.IDPartido {
		t.Errorf("ID incorrecto del partido")
	}

	partidos, err := queries.ListPartidos(ctx)
	if err != nil {
		t.Fatalf("error al listar partidos: %v", err)
	}

	encontrado := false

	for _, p := range partidos {
		if p.IDPartido == partido.IDPartido {
			encontrado = true
			break
		}
	}

	if !encontrado {
		t.Errorf("el partido creado no apareció en ListPartidos")
	}

	err = queries.UpdatePartido(ctx, UpdatePartidoParams{
		IDPartido:         partido.IDPartido,
		IDEquipoLocal:     local.IDEquipo,
		IDEquipoVisitante: visitante.IDEquipo,
		GolesLocal:        2,
		GolesVisitante:    2,
		Fecha:             time.Date(2026, 8, 21, 21, 0, 0, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("error al actualizar partido: %v", err)
	}

	actualizado, err := queries.GetPartido(ctx, partido.IDPartido)
	if err != nil {
		t.Fatalf("error al obtener partido actualizado: %v", err)
	}

	if actualizado.GolesLocal != 2 || actualizado.GolesVisitante != 2 {
		t.Errorf("resultado del partido no actualizado correctamente")
	}

	err = queries.DeletePartido(ctx, partido.IDPartido)
	if err != nil {
		t.Fatalf("error al eliminar partido: %v", err)
	}

	_, err = queries.GetPartido(ctx, partido.IDPartido)
	if err == nil {
		t.Errorf("se esperaba un error al obtener el partido eliminado")
	}
}
