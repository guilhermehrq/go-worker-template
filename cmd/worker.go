package cmd

import (
	"fmt"
	"log"
	"os"

	"go-worker-template/config/db"
	"go-worker-template/interfaces"
	"go-worker-template/repositories"
	"go-worker-template/usecases"
)

// Worker ...
type Worker struct {
	ExampleUseCase interfaces.ExampleUseCase
}

func (w *Worker) initialization() {
	dbConnection, err := db.InitDb()
	if err != nil {
		log.Printf("Error to init database connection, error: %s\n", err.Error())
		os.Exit(1)
	}

	exampleRepository := repositories.NewExampleRepository(dbConnection)
	w.ExampleUseCase = usecases.NewExampleUseCase(exampleRepository)
}

// StartWorker ...
func (w *Worker) StartWorker() {
	w.initialization()

	examples, err := w.ExampleUseCase.GetExample("teste")
	if err != nil {
		log.Printf("Error to get examples, error: %s\n", err.Error())
		os.Exit(1)
	}

	for _, item := range examples {
		fmt.Printf(`ID: %d, Name: %s, Description: %s`+"\n", item.ID, item.Name, item.Description)
	}

	fmt.Println("Finish!")
}
