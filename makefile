migrate.up:
		migrate -verbose -path "./database/migration" -database "mysql://root:tuanbieber@(localhost:3306)/tuanbieber" up
migrate.down:
		migrate -verbose -path "./database/migration" -database "mysql://root:tuanbieber@(localhost:3306)/tuanbieber" down
test.it:
	docker-compose -f docker-compose.test.yml down
	docker-compose -f docker-compose.test.yml up -d --build
	sleep 30
	docker exec -it integration-golang-server-test bash -c "go test -tags=integration ./test -v -count=1"
	sleep 5
	docker-compose -f docker-compose.test.yml down
	docker image rm integration-golang_web
	docker image rm mysql
