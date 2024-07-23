package api

import (
	"net/http"

	"github.com/SrBigotones/proxy-challenge/cmd/api/controllers/proxy"
	"github.com/SrBigotones/proxy-challenge/cmd/api/controllers/stats"
	"github.com/SrBigotones/proxy-challenge/cmd/api/persistance/mongo_client"
	"github.com/SrBigotones/proxy-challenge/cmd/api/persistance/redis_client"
	"github.com/gorilla/mux"
)

type API struct {
	addr string
	port string
}

func NewApiServer(addr string, port string) *API {
	return &API{
		addr: addr,
		port: port,
	}
}

func (api *API) Run() {
	println("Starting server")

	mongoSession := mongo_client.NewMognoClient("localhost", "27017", "proxy", "client_stats")
	redisSession := redis_client.NewRedisClient("localhost", "6379", "", 0)
	statController := stats.NewStatController(mongoSession)
	proxyController := proxy.NewProxyController(redisSession, mongoSession)

	r := mux.NewRouter()
	statController.RegisterRouter(r)
	proxyController.RegisterRouter(r)

	http.ListenAndServe((api.addr + ":" + api.port), r)
}
