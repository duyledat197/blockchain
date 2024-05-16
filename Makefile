GEN_COMPOSE := ./developments/docker-compose.gen.yml
COMPOSE := ./developments/docker-compose.yml
DOCKER_FILE := ./developments/Dockerfile
DOCKER_ALL := ./developments/docker-compose.all.yml

gen-contract:
	docker compose -f ${GEN_COMPOSE} up generate_contract --build
gen-proto:
	docker compose -f ${GEN_COMPOSE} up generate_proto
gen-mock:
	docker compose -f ${GEN_COMPOSE} up generate_mock_internal
	docker compose -f ${GEN_COMPOSE} up generate_mock_pkg
compose:
	docker compose -f ${COMPOSE} up -d --build
get-accounts:
	docker compose -f ${COMPOSE} logs ganache
docker-build:
	docker build -f ${DOCKER_FILE} -t openmyth/blockchain .


start-all:
	docker compose -f ${DOCKER_ALL} up -d --scale deploy_contract=0
deploy-contract:
	docker compose -f ${DOCKER_ALL} up deploy_contract -d

# testing
test-publisher:
	go run tests/kafka/publisher/main.go  
test-subscriber:
	go run tests/kafka/subscriber/main.go 
test-mongo:
	go run tests/mongo/main.go


start-blockchain-algorithm:
	go run tests/blockchain/main.go