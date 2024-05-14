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
	PRIVATE_KEY="8378c4e4383b0ae04a815eb57208e4512aba0e5a9d0badcded84efef05dd98d7" CHAIN_URL="http://localhost:8545" go run deployments/contract/main.go
get-accounts:
	docker compose -f ${COMPOSE} exec -it ganache ganache-cli --wallet.accounts 