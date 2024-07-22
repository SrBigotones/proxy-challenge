package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

const limitPerIP = 2 //1000
const limitCat = 1   //10_000
const limitItems = 10
const API_URL = "https://api.mercadolibre.com"
const logStats = true

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})
var ctx = context.Background()

func main() {

	r := mux.NewRouter()

	r.Use(middleware)

	r.PathPrefix("/categories").HandlerFunc(getCategories)
	r.PathPrefix("/items").HandlerFunc(getItem)

	http.ListenAndServe(":8080", r)
}

/*
ipSolicitante:NumMinuto -> Request count
*/
func middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UnixMilli()
		lrw := NewLogginResponseWrite(w, r)

		fmt.Printf("Request from %s to %s\n", r.RemoteAddr, r.URL)

		reqCheck := readContraintValue(strings.Split(r.RemoteAddr, ":")[0], limitPerIP)

		if !reqCheck {
			w.WriteHeader(http.StatusTooManyRequests)
			fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
		} else {
			next.ServeHTTP(lrw, r)
		}

		if logStats {
			endTime := time.Now().UnixMilli() - startTime
			lrw.ResponseTimeMs = endTime
			stat2D, _ := json.Marshal(lrw)
			fmt.Println(string(stat2D))
		}
	})
}

func getCategories(w http.ResponseWriter, r *http.Request) {

	reqCheck := readContraintValue("categories", limitCat)

	if !reqCheck {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		getContentFromAPI(r.URL.Path, w)
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {

	reqCheck := readContraintValue("items", limitItems)

	if !reqCheck {
		w.WriteHeader(http.StatusTooManyRequests)
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		getContentFromAPI(r.URL.Path, w)
	}

}

func readContraintValue(key string, limit int64) bool {
	currentMinute := time.Now().Minute()

	redisKey := fmt.Sprintf("%s:%d", key, currentMinute)
	reqCheck, err := client.Incr(ctx, redisKey).Result()

	if err != nil || reqCheck > limit {
		return false
	}

	return true
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

type ClientStat struct {
	Ip                  string `json:"ip"`
	Path                string `json:"path"`
	ResponseTimeMs      int64  `json:"responseTimeMs"`
	StatusCode          int    `json:"statusCode"`
	http.ResponseWriter `json:"-"`
}

func (stat *ClientStat) WriteHeader(code int) {
	stat.StatusCode = code
	stat.ResponseWriter.WriteHeader(code)
}

func NewLogginResponseWrite(w http.ResponseWriter, r *http.Request) *ClientStat {

	return &ClientStat{ResponseWriter: w, StatusCode: 200, Ip: r.RemoteAddr, Path: r.URL.Path, ResponseTimeMs: 0}
}
