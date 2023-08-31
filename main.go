package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"wonderooo/teescrap/uploader"
	"wonderooo/teescrap/webapp"
)

func main() {
	app := webapp.New()
	db := webapp.InitDb()
	app.AddHandle("/", webapp.HandleStaticContent())
	app.AddHandle("/cycles", webapp.HandleCreateCycle(db))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		call := <-ch
		log.Printf("System call: %s", call)
		cancel()
	}()

	if err := app.Serve(ctx); err != nil {
		log.Fatalf("Failed to serve: ", err)
	}
}

func exampleJobs() []uploader.Job {
	/*

	 */
	return []uploader.Job{
		uploader.NewJob("shared/rng.jpeg", "dupa1", "dsc", uploader.ColorChoices{}, "shirtooo1", "shirtooo2", "shirtooo3"),
		uploader.NewJob("shared/rng.jpeg", "dupa2", "dsc", uploader.ColorChoices{}, "shirtooo1", "shirtooo2", "shirtooo3"),
		uploader.NewJob("shared/rng.jpeg", "dupa3", "dsc", uploader.ColorChoices{}, "shirtooo1", "shirtooo2", "shirtooo3"),
	}
}
