package main

import (
	"final-project-4/handlers"
	"os"
)

// const port = ":8080"

func main() {
	r := handlers.StartApp()

	r.Run(":" + os.Getenv("PORT"))
}
