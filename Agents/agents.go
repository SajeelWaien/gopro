package agents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sajeelwaien/gopro/database"
	"github.com/sajeelwaien/gopro/models"
)

type PersonInput struct {
	Name string `json:name`
	Ult  string `json:ult`
}

func AddAgent(r http.ResponseWriter, req *http.Request) {
	var agent models.Agent

	err := json.NewDecoder(req.Body).Decode(&agent)

	if err != nil {
		fmt.Printf("ERROR %+v", agent)
		http.Error(r, err.Error(), http.StatusBadRequest)
		return
	}

	// fmt.Printf("%+v\n", newAgent)

	newAgent := models.Agent{Name: agent.Name, Ult: agent.Ult, Abilities: agent.Abilities}

	result := database.DBCon.Create(&newAgent)

	if result.Error != nil {
		http.Error(r, result.Error.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(r, "%+v\n", agent)
}
