# Liga Universitaria - Sistema de Gestión

## Descripción de la aplicación

El sistema tiene como objetivo gestionar y consultar información de la Liga Universitaria de Futsal y Fútbol 11 de la UNICEN.

El usuario podrá consultar:
- Los equipos participantes.
- La división a la que pertenece cada equipo (A o B).
- Los jugadores pertenecientes a cada equipo.
- El nombre y apellido de los jugadores.
- Los partidos disputados.
- El equipo local y el equipo visitante de cada partido.
- El resultado de cada partido.
- La fecha de cada partido.

## Persistencia

La información se almacena en PostgreSQL. El acceso desde Go se realiza mediante `database/sql` y `pgx`.

El esquema y las consultas SQL se encuentran en `db/schema/schema.sql` y `db/queries/queries.sql`. `sqlc` genera el código Go para acceder a la base de datos.

La documentación de persistencia se encuentra en `docs/persistencia.md`.

## Tecnologías

- Go 
- PostgreSQL 
- Docker 
- Docker Compose 
- sqlc
- make
- `database/sql`
- `github.com/jackc/pgx/v5`

## Requisitos

Para ejecutar el proyecto se necesita:
- Go 1.22.2 o superior.
- Docker.
- Docker Compose.
- sqlc
- make

## Ejecución y pruebas

Desde la raíz del proyecto, ejecutar el siguiente comando en bash:

"make test"

Este comando genera el código con sqlc, compila el proyecto, elimina los contenedores y volúmenes anteriores, inicia PostgreSQL mediante Docker Compose, espera a que esté disponible, ejecuta las pruebas y finalmente elimina los contenedores y volúmenes.

Las pruebas se encuentran en db/sqlc/queries_test.go y utilizan el paquete testing de Go. Se verifica el funcionamiento de las operaciones CRUD de equipos, jugadores y partidos, comprobando la creación, consulta, listado, actualización y eliminación de registros.