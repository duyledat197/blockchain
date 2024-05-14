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
deploy-contract:
	 go run main.go deployContract
get-accounts:
	docker compose -f ${COMPOSE} logs ganache
start-gateway:
	go run main.go gateway