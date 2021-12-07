package agents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sajeelwaien/gopro/models"
)

func AddAgent(r http.ResponseWriter, req *http.Request) {
	var agent models.Agent

	_ = json.NewDecoder(req.Body).Decode(&agent)

	// fmt.Printf("%+v\n", newAgent)

	fmt.Fprintf(r, "%+v\n", req.Body)
	// newAgent := models.Agent{}

	// database.DBCon.Create()
}
