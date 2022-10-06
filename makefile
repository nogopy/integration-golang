migrate.up:
		migrate -verbose -path "./database/migration" -database "mysql://root:tuanbieber@(localhost:3306)/tuanbieber" up

migrate.down:
		migrate -verbose -path "./database/migration" -database "mysql://root:tuanbieber@(localhost:3306)/tuanbieber" down
test.init:
	docker-compose -f docker-compose.test.yml down
	docker-compose -f docker-compose.test.yml up -d --build
test.start:
	docker exec -it integration-golang-server-test bash -c "sleep 10; go test -tags=integration ./test -v -count=1"