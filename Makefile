.PHONY: docker-build
docker-build:
	@docker build -t local/vietlott:latest .

.PHONY: stack-up
stack-up:
	@docker-compose -f docker-compose.dev.yml up -d --force-recreate

.PHONY: stack-down
stack-down:
	@docker-compose -f docker-compose.dev.yml down

.PHONY: stack-down
stack-down-complete:
	@docker-compose -f docker-compose.dev.yml down -v

.PHONY: log-vietlott
log-vietlott:
	@docker logs -f --tail 20 vietlott

.PHONY: log-db
log-db:
	@docker-compose -f docker-compose.dev.yml logs -f --tail 20 mongodb