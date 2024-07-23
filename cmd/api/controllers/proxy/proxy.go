package proxy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/SrBigotones/proxy-challenge/cmd/api/model/user_stats"
	"github.com/SrBigotones/proxy-challenge/cmd/api/persistance/mongo_client"
	"github.com/SrBigotones/proxy-challenge/cmd/api/persistance/redis_client"
	"github.com/gorilla/mux"
)

const limitPerIP = 1000
const limitCat = 10_000
const limitItems = 10
const API_URL = "https://api.mercadolibre.com"
const logStats = true

type ProxyControler struct {
	redisSession *redis_client.RedisClient
	mongoSession *mongo_client.MongoClient
}

func NewProxyController(redisSession *redis_client.RedisClient, mongoSession *mongo_client.MongoClient) *ProxyControler {
	return &ProxyControler{
		redisSession: redisSession,
		mongoSession: mongoSession,
	}
}

func (controller *ProxyControler) RegisterRouter(router *mux.Router) {

	r := router.NewRoute().Subrouter()
	r.Use(controller.middleware)
	r.PathPrefix("/categories").HandlerFunc(controller.getCategories)
	r.PathPrefix("/items").HandlerFunc(controller.getItem)

}

/*
ipSolicitante:NumMinuto -> Request count
*/
func (controller *ProxyControler) middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UnixMilli()
		lrw := controller.NewLogginResponseWrite(w, r)

		fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.URL)

		reqCheck := controller.redisSession.ReadContraintValue(strings.Split(r.RemoteAddr, ":")[0], limitPerIP)

		if !reqCheck {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
		} else {
			next.ServeHTTP(lrw, r)
		}

		//Persistimos la data en mongoDB
		if logStats {
			endTime := time.Now().UnixMilli() - startTime
			lrw.ResponseTimeMs = endTime
			lrw.Timestamp = time.Now().String()
			stat2D, _ := json.Marshal(lrw)
			fmt.Println(string(stat2D))
			controller.mongoSession.InsertToCollection(stat2D)
		}
	})
}

func (controller *ProxyControler) getCategories(w http.ResponseWriter, r *http.Request) {

	reqCheck := controller.redisSession.ReadContraintValue("categories", limitCat)

	if !reqCheck {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		controller.getContentFromAPI(r.URL.Path, w)
	}
}

func (controller *ProxyControler) getItem(w http.ResponseWriter, r *http.Request) {

	reqCheck := controller.redisSession.ReadContraintValue("items", limitItems)

	if !reqCheck {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		controller.getContentFromAPI(r.URL.Path, w)
	}

}

func (controller *ProxyControler) getContentFromAPI(path string, w http.ResponseWriter) {
	res, err := http.Get(API_URL + path)
	if err != nil {
		panic(err)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(res.StatusCode)
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		fmt.Fprint(w, string(resBody))
	}
}

func (controller *ProxyControler) NewLogginResponseWrite(w http.ResponseWriter, r *http.Request) *user_stats.UserStats {

	return &user_stats.UserStats{ResponseWriter: w, StatusCode: 200, Ip: strings.Split(r.RemoteAddr, ":")[0], Path: r.URL.Path, ResponseTimeMs: 0}
}
