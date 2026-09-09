CREATE TABLE equipo(
    id_equipo SERIAL PRIMARY KEY,
    nombre_equipo VARCHAR(100) NOT NULL,
    cant_jugadores INT NOT NULL,
    division VARCHAR(1) NOT NULL
)