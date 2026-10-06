package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	db "server-tp/servidor-tpespecial/db/sqlc"
)

func (api *EquiposAPI) POSTEquipoHandler(w http.ResponseWriter, r *http.Request) {
	//creamos un struct temporal que contiene los datos del request en formato JSON
	var params db.CreateEquipoParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Error en formato: "+err.Error(), http.StatusBadRequest)
		return
	}

	//validaciones del negocio
	if params.NombreEquipo == "" {
		http.Error(w, "El nombre del equipo es requerido", http.StatusBadRequest)
		return
	}

	//creamos usando el codigo generado por sqlc
	nuevoEquipo, err := api.Queries.CreateEquipo(r.Context(), params)
	if err != nil {
		http.Error(w, "Error al crear equipo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	//si llega hasta aca es porque se creo el equipo
	//devolvemos el equipo en JSON y ademas el codigo 201 de Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nuevoEquipo)
}

func (api *EquiposAPI) GETEquiposHandler(w http.ResponseWriter, r *http.Request) {
	//ListEquipos es el metodo autogenerado por sqlc para mostrar los equipos
	equipos, err := api.Queries.ListEquipos(r.Context())
	if err != nil {
		http.Error(w, "Error al listar equipos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipos)
}

// GET de 1 solo equipo
func (api *EquiposAPI) GETEquipoHandler(w http.ResponseWriter, r *http.Request) {
	//extraemos el id de la URL y lo convertimos a string para verificar si se escribio un id
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	//buscamos en la base de datos el id que vino
	equipo, err := api.Queries.GetEquipo(r.Context(), int32(id))
	if err != nil {
		//devolvemos 404 si no se encuentra el equipo
		if err == sql.ErrNoRows {
			http.Error(w, "Equipo no encontrado", http.StatusNotFound)
			return
		}
		//si existe algun otro error, devolvemos 500
		http.Error(w, "Error al buscar el equipo", http.StatusInternalServerError)
		return
	}

	//devolvemos el equipo en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equipo)
}

// PUT
func (api *EquiposAPI) PUTEquipoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var params db.UpdateEquipoParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	params.IDEquipo = int32(id)

	//actualizamos el equipo
	err = api.Queries.UpdateEquipo(r.Context(), params)
	if err != nil {
		http.Error(w, "Error al actualizar el equipo", http.StatusInternalServerError)
		return
	}

	//devolvemos un mensaje de éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"mensaje": "Equipo actualizado correctamente"}`))
}

// DELETE
func (api *EquiposAPI) DELETEEquipoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = api.Queries.DeleteEquipo(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error al eliminar el equipo", http.StatusInternalServerError)
		return
	}

	// Código 204: Operación exitosa, pero no hay contenido para devolver
	w.WriteHeader(http.StatusNoContent)
}
