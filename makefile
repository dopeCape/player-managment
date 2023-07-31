start:
	make database && watchexec -e  go -r make dev && make  database_close
dev:
	clear &&  go run cmd/main.go
database:
	mongod --dbpath "db" --fork --logpath "logs" 
database_close:
	mongod --dbpath "db" --shutdown
docker:
	 docker build -t player-managment . && docker run -p 27017:27017 -p 8000:8000 player-managment



