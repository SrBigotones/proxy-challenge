package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

const limitPerIP = 2 //1000
const limitCat = 3   //10_000
const limitItems = 10

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
	fmt.Println("Request Middleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqCheck := readContraintValue(strings.Split(r.RemoteAddr, ":")[0], limitPerIP)

		if !reqCheck {
			fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Categories")

	reqCheck := readContraintValue("categories", limitCat)

	if !reqCheck {
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		fmt.Fprintf(w, "Your content: herererere")
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Item")
	reqCheck := readContraintValue("items", limitItems)

	if !reqCheck {
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		fmt.Fprintf(w, "Your content: herererere")
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
