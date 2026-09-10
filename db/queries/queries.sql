-- =========================
-- EQUIPO
-- =========================

-- name: CreateEquipo :one
INSERT INTO equipo (nombre_equipo, division)
VALUES ($1, $2)
RETURNING *;

-- name: GetEquipo :one
SELECT *
FROM equipo
WHERE id_equipo = $1;

-- name: ListEquipos :many
SELECT *
FROM equipo
ORDER BY division ASC, nombre_equipo ASC;

-- name: UpdateEquipo :exec
UPDATE equipo
SET nombre_equipo = $2,
    division = $3
WHERE id_equipo = $1;

-- name: DeleteEquipo :exec
DELETE FROM equipo
WHERE id_equipo = $1;


-- =========================
-- JUGADOR
-- =========================

-- name: CreateJugador :one
INSERT INTO jugador (nombre, apellido, id_equipo)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetJugador :one
SELECT *
FROM jugador
WHERE id_jugador = $1;

-- name: ListJugadores :many
SELECT *
FROM jugador
ORDER BY apellido ASC, nombre ASC;

-- name: UpdateJugador :exec
UPDATE jugador
SET nombre = $2,
    apellido = $3,
    id_equipo = $4
WHERE id_jugador = $1;

-- name: DeleteJugador :exec
DELETE FROM jugador
WHERE id_jugador = $1;


-- =========================
-- PARTIDO
-- =========================

-- name: CreatePartido :one
INSERT INTO partido (
    id_equipo_local,
    id_equipo_visitante,
    goles_local,
    goles_visitante,
    fecha
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPartido :one
SELECT *
FROM partido
WHERE id_partido = $1;

-- name: ListPartidos :many
SELECT *
FROM partido
ORDER BY fecha ASC;

-- name: UpdatePartido :exec
UPDATE partido
SET id_equipo_local = $2,
    id_equipo_visitante = $3,
    goles_local = $4,
    goles_visitante = $5,
    fecha = $6
WHERE id_partido = $1;

-- name: DeletePartido :exec
DELETE FROM partido
WHERE id_partido = $1;