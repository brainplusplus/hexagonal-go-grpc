# Inject .env file
ifneq (,$(wildcard .env))
    include .env
    export
endif

# Detect OS
ifeq ($(OS), Windows_NT)
    OS := Windows
else
    OS := $(shell uname -s 2>/dev/null || echo Unix)
endif

# Set TIMESTAMP for both Unix and Windows (Git Bash)
ifeq ($(OS), Windows)
    TIMESTAMP := $(shell powershell -Command "Get-Date -Format 'yyyyMMddHHmmss'")
	TOUCH := type nul > # Windows equivalent of `touch`
    OTEL_CONFIG_PATH := $(shell powershell -Command "[System.IO.Path]::GetFullPath('otel-collector-config.yaml') -replace '\\', '/'")
    PROMETHEUS_CONFIG_PATH := $(shell powershell -Command "[System.IO.Path]::GetFullPath('prometheus.yml') -replace '\\', '/'")
else
    TIMESTAMP := $(shell date +%Y%m%d%H%M%S)
	TOUCH := touch
	OTEL_CONFIG_PATH := $(shell pwd)/otel-collector-config.yaml
	PROMETHEUS_CONFIG_PATH := $(shell pwd)/prometheus.yml
endif

export OTEL_CONFIG_PATH
export PROMETHEUS_CONFIG_PATH

#make tidy
.PHONY: tidy
tidy:
ifeq ($(OS), Windows)
	@set BUILD_START=%TIME% \
	&& go mod tidy \
	&& echo go mod tidy took %TIME%
else
	@BUILD_START=$$(date +%s) \
	;go go mod tidy \
	&& echo "Go mod tidy took $$(($$(date +%s)-BUILD_START)) seconds"
endif

#make vendor
.PHONY: vendor
vendor:
ifeq ($(OS), Windows)
	@set BUILD_START=%TIME% \
	&& go mod tidy \
	&& go mod vendor \
	&& echo go mod tidy and vendor took %TIME%
else
	@BUILD_START=$$(date +%s) \
	;go go mod tidy \
	&& go mod vendor \
	&& echo "Go mod tidy and vendor took $$(($$(date +%s)-BUILD_START)) seconds"
endif

#make build
.PHONY: build
build:
ifeq ($(OS), Windows)
	@set BUILD_START=%TIME% \
	&& set CGO_ENABLED=0 \
	&& go build -o app.exe cmd/app/main.go \
	&& echo Build took %TIME%
else
	@BUILD_START=$$(date +%s) \
	;CGO_ENABLED=0 go build -o ./app ./cmd/app/main.go \
	&& echo "Build took $$(($$(date +%s)-BUILD_START)) seconds"
endif


#make migration name=create_customers_table
.PHONY: migration
migration:
	@$(TOUCH) ./migrations/$(db_driver)/$(TIMESTAMP)_$(name).up.sql \
    && @$(TOUCH) ./migrations/$(db_driver)/$(TIMESTAMP)_$(name).down.sql

#make migrate value=0
.PHONY: migrate
migrate:
	@go run ./cmd/migrate/migrate.go --cmd=migrate --value=$(value)

#make migrate-up
.PHONY: migrate-up
migrate-up:
	@go run ./cmd/migrate/migrate.go --cmd=up

#make migrate-down
.PHONY: migrate-down
migrate-down:
	@go run ./cmd/migrate/migrate.go --cmd=down

#make migrate-force value=0
.PHONY: migrate-force
migrate-force:
	@go run ./cmd/migrate/migrate.go --cmd=force --value=$(value)

#make migrate-steps value=0
.PHONY: migrate-steps
migrate-steps:
	@go run ./cmd/migrate/migrate.go --cmd=steps --value=$(value)


#make proto
.PHONY: proto
proto:
	@protoc --proto_path=proto \
       --go_out=proto/gen --go_opt=paths=source_relative \
       --connect-go_out=proto/gen --connect-go_opt=paths=source_relative \
       api/v1/customer.proto

.PHONY: mock
mock:
	@mockery --recursive --dir=internal/ports --name="^." --with-expecter --inpackage

.PHONY: test
test:
	@go test ./internal/... -v

.PHONY: coverage-test
coverage-test:
	@go test ./internal/... -v -coverprofile=cover.out -cover
	@go tool cover -func cover.out
	@go tool cover -html=cover.out -o cover.html

#make build-run
.PHONY: build-run
build-run: build
ifeq ($(OS), Windows)
	@./app.exe
else
	@./app
endif

#make run
.PHONY: run
run:
	@go run cmd/app/main.go

.PHONY: docker-compose-up
docker-compose-up:
ifeq ($(OS), Windows)
	@set "OTEL_CONFIG_PATH=$(OTEL_CONFIG_PATH)" && set "PROMETHEUS_CONFIG_PATH=$(PROMETHEUS_CONFIG_PATH)" && docker-compose up -d
else
	OTEL_CONFIG_PATH=$(OTEL_CONFIG_PATH) PROMETHEUS_CONFIG_PATH=$(PROMETHEUS_CONFIG_PATH) docker-compose up -d
endif

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose down