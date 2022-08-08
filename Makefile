base:
	docker run --name=link-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
test:
	go test -v ./...
migrate:
	migrate -path ./migration -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
