dep:
	go mod tidy

up:
	docker-compose up -d 

down:
	docker-compose down

run:
	go run cmd/service/main.go

trx:
	go run cmd/transaction/main.go