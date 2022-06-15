package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a JSON to request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}

}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
