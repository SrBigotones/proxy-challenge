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

	// r.HandleFunc("/", getRequest)

	r.PathPrefix("/categories").HandlerFunc(getCategories)

	r.PathPrefix("/items").HandlerFunc(getItem)

	http.ListenAndServe(":8080", r)
}

func middleware(next http.Handler) http.Handler {
	fmt.Println("Request Middleware")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentMinute := time.Now().Minute()

		redisKey := fmt.Sprintf("%s:%d", strings.Split(r.RemoteAddr, ":")[0], currentMinute)

		reqCheck, err := client.Incr(ctx, redisKey).Result()

		fmt.Println(reqCheck)
		if err != nil {
			panic(err)
		}

		if reqCheck > limitPerIP {
			fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
		} else {
			next.ServeHTTP(w, r)
		}

	})
}

/*
*
ipSolicitante:NumMinuto -> Request count
*/

func getCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Categories")
	fmt.Println(r.RemoteAddr)
	currentMinute := time.Now().Minute()

	redisKey := fmt.Sprintf("categories:%d", currentMinute)

	reqCheck, err := client.Incr(ctx, redisKey).Result()

	fmt.Println(reqCheck)
	if err != nil {
		panic(err)
	}

	if reqCheck > limitCat {
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		fmt.Fprintf(w, "Your content: herererere")
	}
}

func getItem(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Request Item")
	currentMinute := time.Now().Minute()

	redisKey := fmt.Sprintf("%s/items:%d", strings.Split(r.RemoteAddr, ":")[0], currentMinute)

	reqCheck, err := client.Incr(ctx, redisKey).Result()

	fmt.Println(reqCheck)
	if err != nil {
		panic(err)
	}

	if reqCheck > limitItems {
		fmt.Fprintf(w, "Alcanzo el limite de request, espere un minuto")
	} else {
		fmt.Fprintf(w, "Your content: herererere")
	}

}
