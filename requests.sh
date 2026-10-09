#!/usr/bin/env bash

set -e

BASE_URL="http://localhost:8080"
ARCHIVO_RESPUESTA=$(mktemp)

trap 'rm -f "$ARCHIVO_RESPUESTA"' EXIT

hacer_pedido() {
    metodo="$1"
    endpoint="$2"
    esperado="$3"
    datos="${4:-}"

    if [ -n "$datos" ]; then
        codigo=$(curl -sS \
            -o "$ARCHIVO_RESPUESTA" \
            -w "%{http_code}" \
            -X "$metodo" \
            "$BASE_URL$endpoint" \
            -H "Content-Type: application/json" \
            -d "$datos")
    else
        codigo=$(curl -sS \
            -o "$ARCHIVO_RESPUESTA" \
            -w "%{http_code}" \
            -X "$metodo" \
            "$BASE_URL$endpoint")
    fi

    echo
    echo "Petición: $metodo $endpoint"
    echo "Estado HTTP: $codigo"
    cat "$ARCHIVO_RESPUESTA"
    echo

    if [ "$codigo" != "$esperado" ]; then
        echo "ERROR: se esperaba HTTP $esperado y se recibió HTTP $codigo."
        exit 1
    fi
}

echo "=== 1. Crear un equipo ==="

hacer_pedido "POST" "/api/equipos" "201" \
    '{"nombre_equipo":"Equipo Prueba API","Division":"A"}'

# Extraer el ID del equipo creado a partir de la respuesta JSON.
ID_EQUIPO=$(python3 -c \
    'import json, sys; print(json.load(open(sys.argv[1]))["id_equipo"])' \
    "$ARCHIVO_RESPUESTA")

echo "ID del equipo creado: $ID_EQUIPO"

echo
echo "=== 2. Listar todos los equipos ==="

hacer_pedido "GET" "/api/equipos" "200"

echo
echo "=== 3. Consultar el equipo creado ==="

hacer_pedido "GET" "/api/equipos/$ID_EQUIPO" "200"

echo
echo "=== 4. Actualizar el equipo ==="

hacer_pedido "PUT" "/api/equipos/$ID_EQUIPO" "200" \
    '{"nombre_equipo":"Equipo Actualizado API","Division":"B"}'

echo
echo "=== 5. Verificar la actualización ==="

hacer_pedido "GET" "/api/equipos/$ID_EQUIPO" "200"

echo
echo "=== 6. Eliminar el equipo ==="

hacer_pedido "DELETE" "/api/equipos/$ID_EQUIPO" "204"

echo
echo "=== 7. Verificar que el equipo ya no existe ==="

hacer_pedido "GET" "/api/equipos/$ID_EQUIPO" "404"

echo
echo "Todas las pruebas finalizaron correctamente."
