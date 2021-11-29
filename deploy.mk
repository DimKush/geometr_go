DEV_PROJECT=geometry_go
DEV_COMPOSE_FILE=./deployments/docker-compose.yaml


.PHONY: install-db
install-db:
	@docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up -d postgres-db