start-db:
	docker compose start

run:
	go run ./cmd/app/.

run-client:
	cd "./client/ng-sim-app";\
	npm start
