## Proxy for API  

Tech used in this challenge:  
- Go  
- Redis, for cache  
- Mongo, analytics persistance  
- nginx, balancer

To run this project simple do:  
` docker-compose up --scale web=[number of instances] --build `

Before runing check the ENV variables inside de docker-compose.yml file  
The exposed routes for this app are:  
- /categories/
- /items/
- /stats
- /stats/?ip=[ip to filter]
