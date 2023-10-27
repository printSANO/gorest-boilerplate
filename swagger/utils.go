package swagger

import (
	"log"

	"github.com/go-chi/chi/v5"
)

func GenerateJson() {
	log.Println("Generate swagger.json")
}

func GenerateEnpointMap(rts []chi.Route) map[string]string {
	var endpointMap map[string]string
	return endpointMap
}
