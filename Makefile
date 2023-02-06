run: 
	go run main.go
build: 
	go build -o bin/main main.go
migrate_up: 
	soda migrate up 
migrate_down: 
	soda migrate down