package stats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SrBigotones/proxy-challenge/cmd/api/persistance/mongo_client"
	"github.com/gorilla/mux"
)

type StatsController struct {
	mongoSession *mongo_client.MongoClient
}

func NewStatController(mongoSession *mongo_client.MongoClient) *StatsController {
	controller := StatsController{mongoSession: mongoSession}
	return &controller
}

func (controller *StatsController) RegisterRouter(r *mux.Router) {

	r.PathPrefix("/stats/").HandlerFunc(controller.GetStatsPerIp)
	r.PathPrefix("/stats").HandlerFunc(controller.GetAllStats)

}

func (controller *StatsController) GetStatsPerIp(w http.ResponseWriter, r *http.Request) {

	hIp := r.URL.Query().Get("ip")

	result, _ := controller.mongoSession.FindByIP(hIp)

	if len(result) == 0 || result == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	body2d, _ := json.Marshal(result)

	fmt.Fprint(w, string(body2d))
}

func (controller *StatsController) GetAllStats(w http.ResponseWriter, r *http.Request) {
	result, err := controller.mongoSession.FindAll()

	if err != nil {
		panic(err)
	}

	if len(result) == 0 || result == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	body2d, _ := json.Marshal(result)

	fmt.Fprint(w, string(body2d))
}
