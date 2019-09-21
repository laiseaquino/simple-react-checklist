package main

import (
	"fmt"
	"log"
	"net/http"

	// "./middleware"
	"./router"
)

func main() {
	r := router.Router()

	/* For some reason mongo does not override your previous entries when the status changes,
	   but just to be safe let's not keep trying to fill the database on every run. */
	// middleware.FillExercises(1, 9)
	// middleware.FillExercises(2, 8)
	// middleware.FillExercises(3, 6)
	// middleware.FillExercises(4, 12)
	// middleware.FillExercises(5, 8)
	// middleware.FillExercises(6, 10)
	// middleware.FillExercises(7, 9)
	// middleware.FillExercises(8, 15)
	// middleware.FillExercises(9, 8)
	// middleware.FillExercises(10, 11)
	// middleware.FillExercises(11, 6)
	// middleware.FillExercises(12, 11)
	// middleware.FillExercises(13, 8)
	// middleware.FillExercises(14, 7)
	// middleware.FillExercises(15, 7)
	// middleware.FillExercises(16, 26)
	// middleware.FillExercises(17, 26)

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
