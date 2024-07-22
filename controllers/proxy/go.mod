module github.com/SrBigotones/proxy-challenge/controllers/proxy

go 1.22.5

replace github.com/SrBigotones/proxy-challenge/persistance/mongo_client => ../../persistance/mongo_client

replace github.com/SrBigotones/proxy-challenge/persistance/redis_client => ../../persistance/redis_client

require (
	github.com/SrBigotones/proxy-challenge/persistance/mongo_client v0.0.0-00010101000000-000000000000
	github.com/SrBigotones/proxy-challenge/persistance/redis_client v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/redis/go-redis/v9 v9.6.0 // indirect
	go.mongodb.org/mongo-driver v1.16.0 // indirect
)
