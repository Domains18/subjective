# services


GO_SERVICE=task-tracker
NODE_SERVICE-diary-api



# Build the go service
build-go:
	docker build -t $(GO_SERVICE) ./task-tracker


build-node:
	docker build -t $(NODE_SERVICE) ./diary-api


# Run the go service

run-go:
	docker run -d -p 8080:8080 $(GO_SERVICE)


run-node:
	docker run -d -p 8081:8081 $(NODE_SERVICE)


# run all
run-all:
	docker-compose up -d