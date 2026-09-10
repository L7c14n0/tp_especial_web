CREATE TABLE equipo(
    id_equipo SERIAL PRIMARY KEY,
    nombre_equipo VARCHAR(100) NOT NULL,
    division VARCHAR(1) NOT NULL CHECK (division IN ('A', 'B'))
);

CREATE TABLE jugador (
    id_jugador SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido VARCHAR(100) NOT NULL,
    id_equipo INT NOT NULL,
    FOREIGN KEY (id_equipo) REFERENCES equipo(id_equipo)
);

CREATE TABLE partido (
    id_partido SERIAL PRIMARY KEY,
    id_equipo_local INT NOT NULL,
    id_equipo_visitante INT NOT NULL,
    goles_local INT NOT NULL,
    goles_visitante INT NOT NULL,
    fecha TIMESTAMP NOT NULL,

    FOREIGN KEY (id_equipo_local) REFERENCES equipo(id_equipo),
    FOREIGN KEY (id_equipo_visitante) REFERENCES equipo(id_equipo),

    CHECK (id_equipo_local <> id_equipo_visitante),
    CHECK (goles_local >= 0),
    CHECK (goles_visitante >= 0)
);
