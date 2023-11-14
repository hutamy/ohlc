dep:
	go mod tidy

up:
	docker-compose up -d 

down:
	docker-compose down

run:
	go run cmd/service/main.go

publish:
	go run cmd/publisher/main.go

consume:
	go run cmd/consumer/main.go