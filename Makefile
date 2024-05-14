GEN_COMPOSE := ./developments/docker-compose.gen.yml
COMPOSE := ./developments/docker-compose.yml

gen-contract:
	docker compose -f ${GEN_COMPOSE} up generate_contract --build
gen-proto:
	docker compose -f ${GEN_COMPOSE} up generate_proto
gen-mock:
	docker compose -f ${GEN_COMPOSE} up generate_mock
compose:
	docker compose -f ${COMPOSE} up -d --build
get-accounts:
	docker compose -f ${COMPOSE} logs ganache

# using go cmd
start-gateway:
	go run main.go gateway
deploy-contract:
	go run main.go deployContract
start-user:
	go run main.go user
start-watcher:
	go run main.go watcher