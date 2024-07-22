package proxy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/SrBigotones/proxy-challenge/persistance/mongo_client"
	"github.com/SrBigotones/proxy-challenge/persistance/redis_client"

	"github.com/SrBigotones/proxy-challenge/model/user_stats"
)

const limitPerIP = 1000
const limitCat = 10_000
const limitItems = 10
const API_URL = "https://api.mercadolibre.com"
const logStats = true

func RegisterRouter(router *mux.Router) {

	r := router.NewRoute().Subrouter()
	r.Use(middleware)
	r.PathPrefix("/categories").HandlerFunc(getCategories)
	r.PathPrefix("/items").HandlerFunc(getItem)

}

/*
ipSolicitante:NumMinuto -> Request count
*/
func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UnixMilli()
		lrw := NewLogginResponseWrite(w, r)

		fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.URL)

		reqCheck := redis_client.ReadContraintValue(strings.Split(r.RemoteAddr, ":")[0], limitPerIP)

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
			mongo_client.InsertToCollection(stat2D)
		}
	})
}

func getCategories(w http.ResponseWriter, r *http.Request) {

	reqCheck := redis_client.ReadContraintValue("categories", limitCat)

	if !reqCheck {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		getContentFromAPI(r.URL.Path, w)
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {

	reqCheck := redis_client.ReadContraintValue("items", limitItems)

	if !reqCheck {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		getContentFromAPI(r.URL.Path, w)
	}

}

func getContentFromAPI(path string, w http.ResponseWriter) {
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

func NewLogginResponseWrite(w http.ResponseWriter, r *http.Request) *user_stats.UserStats {

	return &user_stats.UserStats{ResponseWriter: w, StatusCode: 200, Ip: strings.Split(r.RemoteAddr, ":")[0], Path: r.URL.Path, ResponseTimeMs: 0}
}
