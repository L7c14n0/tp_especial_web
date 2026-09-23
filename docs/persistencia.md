# Documentación de Persistencia

La información del sistema se almacena en una base de datos relacional **PostgreSQL**, gestionada mediante contenedores de Docker para asegurar un entorno reproducible. La aplicación interactúa con los datos utilizando **sqlc**, lo que nos permite usar consultas SQL puras con validación estricta desde Go.

El esquema de la base de datos está diseñado para garantizar la integridad de los datos mediante restricciones (constraints) y se compone de las siguientes tres tablas:

## 1. Tabla `equipo`
Almacena la información de los clubes que participan.
- **`id_equipo`**: Identificador único y clave primaria.
- **`nombre_equipo`**: Nombre descriptivo del club.
- **`division`**: Categoría en la que juega. Tiene una restricción lógica (`CHECK`) que obliga a que el valor ingresado sea estrictamente 'A' o 'B'.

## 2. Tabla `jugador`
Registra el plantel de deportistas y los vincula con su club.
- **`id_jugador`**: Identificador único y clave primaria.
- **`nombre` y `apellido`**: Datos personales del jugador.
- **`id_equipo`**: Clave foránea (`FOREIGN KEY`) que relaciona al jugador obligatoriamente con un equipo válido existente en la tabla `equipo`.

## 3. Tabla `partido`
Almacena el historial de encuentros, qué equipos se enfrentaron y el resultado final.
- **`id_partido`**: Identificador único y clave primaria.
- **`id_equipo_local` e `id_equipo_visitante`**: Claves foráneas que apuntan a la tabla `equipo` para definir quiénes jugaron.
- **`goles_local` y `goles_visitante`**: Registro del marcador.
- **`fecha`**: Día y hora exacta del encuentro.
- **Reglas de negocio (CHECKs)**: 
  - Impide cargar goles con números negativos.
  - Bloquea la posibilidad de que un equipo juegue contra sí mismo (`id_equipo_local <> id_equipo_visitante`).