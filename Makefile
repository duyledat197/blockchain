CONTRACT_COMPOSE := ./developments/docker-compose.sol.yml

gen-contract:
	docker compose -f ${CONTRACT_COMPOSE} up generate_contract --build
