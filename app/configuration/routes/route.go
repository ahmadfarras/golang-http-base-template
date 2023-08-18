package routes

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"
)

func InitRoute(db *sql.DB) *httprouter.Router {
	router := httprouter.New()

	CategoryRoute(db, router)

	return router
}
