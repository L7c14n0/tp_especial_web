# Liga Universitaria - Sistema de gestión

## Descripción

La aplicación tiene como objetivo gestionar y consultar información de la Liga Universitaria de Futsal y Fútbol 11 de la UNICEN.

El usuario podrá consultar:

- Equipos participantes y su división (A o B).
- Jugadores pertenecientes a cada equipo.
- Datos de los jugadores.
- Partidos disputados.
- Equipos local y visitante de cada partido.
- Resultado y fecha de cada partido.

En este TP se implementa la **capa de persistencia**, utilizando PostgreSQL como base de datos y `sqlc` para generar el código Go de acceso a los datos. La integración con el servidor web se realizará en etapas posteriores.

## Tecnologías

- Go
- PostgreSQL
- Docker y Docker Compose
- SQL
- sqlc
- `database/sql`
- `github.com/jackc/pgx/v5`

## Persistencia

La base de datos está compuesta por tres tablas:

- `equipo`: equipos y división.
- `jugador`: jugadores y equipo al que pertenecen.
- `partido`: partidos, equipos participantes, resultado y fecha.

Relaciones:

- Un equipo puede tener muchos jugadores.
- Un equipo puede participar en muchos partidos como local o visitante.
- Cada jugador pertenece a un equipo.
- Cada partido tiene un equipo local y uno visitante.

Las operaciones CRUD de estas entidades están definidas en `db/queries/queries.sql` y el esquema de la base de datos en `db/schema/schema.sql`.

La documentación detallada de persistencia se encuentra en [`docs/persistencia.md`](docs/persistencia.md).

## Requisitos

Se necesita tener instalado:

- Go 1.22.2 o superior.
- Docker.
- Docker Compose.
- sqlc.

Docker debe estar iniciado para ejecutar las pruebas.

## Ejecución

El proyecto incluye un `Makefile` que automatiza la preparación del entorno y la ejecución de las pruebas.

Desde la raíz del proyecto:

```bash
make test