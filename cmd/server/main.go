package main

import "fmt"

// App - the struct which contains stuff like pointers to db connections.
type App struct {
}

// Run - sets up our apps.
func (app *App) Run() error {
	fmt.Println("This sets up out app")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting our REST API")
		fmt.Println(err)
	}
}
