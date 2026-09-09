-- name: CreateEquipo :one
INSERT INTO equipo (id_equipo, nombre_equipo, cant_jugadores, division) 
VALUES ($1, $2, $3, $4) 
RETURNING *;

-- name: GetEquipo :one
SELECT * FROM equipo 
WHERE id_equipo = $1;

-- name: ListEquipos :many
SELECT * FROM equipo 
ORDER BY division ASC, nombre_equipo ASC;

-- name: UpdateEquipo :exec
UPDATE equipo 
SET nombre_equipo = $2, cant_jugadores = $3, division = $4 
WHERE id_equipo = $1;

-- name: DeleteEquipo :exec
DELETE FROM equipo 
WHERE id_equipo = $1;