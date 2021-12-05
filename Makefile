build: # Build all packages in the project
	go build -v ./...

test:  # Run all project tests
	go test -v ./...

compose-up: # Run app via docker-compose (backend, db and web UI for db)
	docker-compose up -d

compose-build-backend: # Build and run backend via docker-compose
	docker-compose up -d --build products-api

docker-container: # Build Docker container with the backend app and tag as latest
	docker build -t egor-products-api:latest .
