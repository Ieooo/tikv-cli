version=$$(git tag|head -n 1)
commit=$$(git log --pretty=format:%h |head -n 1)

build:
	go build -ldflags "-X main.version=$(version) -X main.commit=$(commit)" -o bin/tikv-cli . 
run-tikv: 
	docker-compose up -d
stop-tikv:
	docker-compose down

.PHONY: build