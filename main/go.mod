module github.com/SrBigotones/proxy-challenge/main

go 1.22.5

replace github.com/SrBigotones/proxy-challenge/controllers/proxy => ../controllers/proxy

replace github.com/SrBigotones/proxy-challenge/controllers/stats => ../controllers/stats

replace github.com/SrBigotones/proxy-challenge/persistance/mongo_client => ../persistance/mongo_client

replace github.com/SrBigotones/proxy-challenge/persistance/redis_client => ../persistance/redis_client

replace github.com/SrBigotones/proxy-challenge/model/user_stats => ../model/user_stats

require (
	github.com/SrBigotones/proxy-challenge/controllers/proxy v0.0.0-00010101000000-000000000000
	github.com/SrBigotones/proxy-challenge/controllers/stats v0.0.0-00010101000000-000000000000
)

require (
	github.com/SrBigotones/proxy-challenge/model/user_stats v0.0.0-00010101000000-000000000000 // indirect
	github.com/SrBigotones/proxy-challenge/persistance/mongo_client v0.0.0-00010101000000-000000000000 // indirect
	github.com/SrBigotones/proxy-challenge/persistance/redis_client v0.0.0-00010101000000-000000000000 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/redis/go-redis/v9 v9.6.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.mongodb.org/mongo-driver v1.16.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
