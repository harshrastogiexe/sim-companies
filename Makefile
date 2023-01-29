start-db:
	docker compose start

run:
	go run ./app/.

run-client:
	cd "./client/ng-sim-app";\
	npm start
