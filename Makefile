.PHONY: docker-build
docker-build:
	@docker build -t local/vietlott:latest .

.PHONY: app-docker
app-docker:
	@docker-compose up -d