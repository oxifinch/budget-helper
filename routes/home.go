package routes

import (
	"budget-helper/database"
	"fmt"
	"log"
	"net/http"
)

func Home(db *database.Database, w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Welcome to Budget Helper!\n")
	if err != nil {
		log.Panic(err)
	}
}
