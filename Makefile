PRODUCTION := false
fmt:
	templ fmt .

build-js:
	npm run build 

gen-templ: gen-sqlc
	templ generate

gen-sqlc:
	sqlc generate

build-server: cleanup gen
	go build -o tmp/app ./	

db:
	docker-compose up -d --wait

db-up: db migrate-up
	@echo "Database is up and migrated"

logs:
	docker-compose logs -f

dev: pre-install
	air

gen: fmt build-js gen-templ gen-sqlc

migrate-up:
	migrate -path database/migrations -database "postgresql://user:password@localhost:5432/caas?sslmode=disable" -verbose up
migrate-down:
	migrate -path database/migrations -database "postgresql://user:password@localhost:5432/caas?sslmode=disable" -verbose down

generate-sqlc:
	sqlc generate

cleanup:
	rm -rf tmp

pre-install:
	@command -v templ > /dev/null; \
	if [ $$? -ne 0 ]; then \
		echo "templ not found, installing..."; \
		go install github.com/a-h/templ/cmd/templ@latest; \
		echo "Successfully installed templ"; \
	else \
		echo "templ is already installed"; \
	fi
	@command -v sqlc > /dev/null; \
	if [ $$? -ne 0 ]; then \
		echo "sqlc not found, installing..."; \
		go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; \
		echo "Successfully installed sqlc"; \
	else \
		echo "sqlc is already installed"; \
	fi
	@command -v migrate > /dev/null; \
	if [ $$? -ne 0 ]; then \
		echo "migrate not found, installing..."; \
		go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; \
		echo "Successfully installed migrate"; \
	else \
		echo "migrate is already installed"; \
	fi
	@command -v air > /dev/null; \
	if [ $$? -ne 0 ]; then \
		echo "air not found, installing..."; \
		go install github.com/cosmtrek/air@latest; \
		echo "Successfully installed air"; \
	else \
		echo "air is already installed"; \
	fi
	@command -v npm > /dev/null; \
	if [ $$? -ne 0 ]; then \
		echo "npm not found, exiting..."; \
		exit 1; \
	fi
	npm i

