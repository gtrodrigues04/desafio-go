package api

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Route struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Source      Point  `json:"source"`
	Destination Point  `json:"destination"`
}

type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func RoutesAPI(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Get("/routes", GetRoutesHandler(db))
	r.Post("/routes", CreateRouteHandler(db))

	return r
}

func GetRoutesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, source, destination FROM routes")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var routes []Route
		for rows.Next() {
			var route Route
			var sourceJSON, destJSON []byte
			if err := rows.Scan(&route.ID, &route.Name, &sourceJSON, &destJSON); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := json.Unmarshal(sourceJSON, &route.Source); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if err := json.Unmarshal(destJSON, &route.Destination); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			routes = append(routes, route)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(routes)
	}
}

func CreateRouteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var route Route
		err := json.NewDecoder(r.Body).Decode(&route)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sourceJSON, err := json.Marshal(route.Source)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		destJSON, err := json.Marshal(route.Destination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("INSERT INTO routes (name, source, destination) VALUES (?, ?, ?)", route.Name, sourceJSON, destJSON)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
