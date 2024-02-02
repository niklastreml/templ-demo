PRODUCTION := false
fmt:
	templ fmt .

build-js:
		npm run build 

gen-templ: gen-sqlc
	templ generate

gen-sqlc:
	sqlc generate

build-server: gen
	go build -o tmp/app ./	


db:
	docker-compose up -d
logs:
	docker-compose logs -f

dev:
	air

gen: fmt build-js gen-templ gen-sqlc

migrate-up:
	migrate -path database/migrations -database "postgresql://user:password@localhost:5432/caas?sslmode=disable" -verbose up
migrate-down:
	migrate -path database/migrations -database "postgresql://user:password@localhost:5432/caas?sslmode=disable" -verbose down

generate-sqlc:
	sqlc generate