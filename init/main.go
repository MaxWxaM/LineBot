package main

func main() {
	app, cleanup, err := InitApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Start(); err != nil {
		panic(err)
	}
}
