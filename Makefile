GEN_COMPOSE := ./developments/docker-compose.gen.yml
COMPOSE := ./developments/docker-compose.yml
DOCKER_FILE := ./developments/Dockerfile
DOCKER_ALL := ./developments/docker-compose.all.yml

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
docker-build:
	docker build -f ${DOCKER_FILE} -t openmyth/blockchain .

# using go cmd
start-gateway:
	go run main.go gateway
deploy-contract:
	SERVICE=deploy_contract go run main.go deployContract
start-user:
	go run main.go user
start-watcher:
	go run main.go watcher

start-all:
	docker compose -f ${DOCKER_ALL} up -d --build

# testing
test-publisher:
	go run tests/kafka/publisher/main.go  
test-subscriber:
	go run tests/kafka/subscriber/main.go  
test-mongo:
	go run tests/mongo/main.go