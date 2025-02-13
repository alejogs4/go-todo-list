APP_NAME = app
DOCKER_COMPOSE_FILE = docker-compose.yml

.PHONY: all run test test-integration clean

all: run

run:
	@echo "Starting the application..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

test:
	@echo "Running tests..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm $(APP_NAME) go test ./...

test-integration:
	@echo "Running integration tests..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) run --rm $(APP_NAME) sh -c "RUN_INTEGRATION_TESTS=true go test ./... -run '^TestIntegration'"

clean:
	@echo "Cleaning up..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down -v