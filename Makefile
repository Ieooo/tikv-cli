build:
	go build -o bin/tikv-cli ./cmd/tikvcli
run-tikv: 
	docker-compose up -d
stop-tikv:
	docker-compose down

.PHONY: build