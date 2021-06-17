package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func getPipelineFile(PipelineFile string) string {
	return "/tmp/" + PipelineFile
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Println("Error During file close")
	}
}

func LoadPipelineConfig(pipelineId string) error {
	log.Println("Opening file: ", pipelineId)
	file, err := os.Open(getPipelineFile(pipelineId))
	if err != nil {
		return err
	}
	log.Println("File opened successfully.")
	defer CloseFile(file)

	return err
}

func getPipeline(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set(ContentType, ApplicationJson)
	pipelineId := ps.ByName("pipelineId")
	log.Println("Loading file: ", pipelineId)
	err := LoadPipelineConfig(pipelineId)
	if err != nil {
		log.Printf("Failed to get pipeline:  %s! \n", err)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/rest/v1/pipeline/:pipelineId", getPipeline)

	log.Fatal(http.ListenAndServe(":8080", router))
}