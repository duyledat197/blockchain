GEN_COMPOSE := ./developments/docker-compose.gen.yml

gen-contract:
	docker compose -f ${GEN_COMPOSE} up generate_contract --build
gen-proto:
	docker compose -f ${GEN_COMPOSE} up generate_proto 