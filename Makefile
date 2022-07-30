migrateup: 
  migrate -path db/migration -database "postgresql://root@roach1:26257/defaultdb?sslmode=disable" -verbose up
migratedown: 
  migrate -path db/migration -database "postgresql://root@roach1:26257/defaultdb?sslmode=disable" -verbose down
sqlc:
  sqlc generate
test: 
  go test -v -cover ./..
server: 
  go run main.go

.PHONY: migrateup migratedown sqlc server test