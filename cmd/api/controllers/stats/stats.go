package stats

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SrBigotones/proxy-challenge/cmd/api/persistance/mongo_client"
	"github.com/gorilla/mux"
)

func RegisterRouter(r *mux.Router) {

	r.PathPrefix("/stats/").HandlerFunc(GetStatsPerIp)
	r.PathPrefix("/stats").HandlerFunc(GetAllStats)

}

func GetStatsPerIp(w http.ResponseWriter, r *http.Request) {

	hIp := r.URL.Query().Get("ip")

	result, _ := mongo_client.FindByIP(hIp)

	if len(result) == 0 || result == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	body2d, _ := json.Marshal(result)

	fmt.Fprint(w, string(body2d))
}

func GetAllStats(w http.ResponseWriter, r *http.Request) {
	result, err := mongo_client.FindAll()

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
