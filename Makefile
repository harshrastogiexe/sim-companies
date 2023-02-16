start-db:
	docker compose start

run:
	cd ./cmd/server/ &&\
	go run .

run-client:
	cd "./client/ng-sim-app";\
	npm start
