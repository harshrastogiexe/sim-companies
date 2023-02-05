start-db:
	docker compose start

run:
	go run ./cmd/server/.

run-client:
	cd "./client/ng-sim-app";\
	npm start
