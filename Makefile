dep:
	go mod tidy

up:
	docker-compose up -d 

down:
	docker-compose down