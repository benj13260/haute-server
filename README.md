# Run local DB with docker
```
docker run --rm \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=postgres   \
  -v pgdata:$HAUTE_PATH/db \
  --name haute-db   \
  --net=haute-net   \
  postgres:15.2
```


  # Migrate db
  ```
  go run db/migrate/migrate.go
  ``` 

  # Run server
  ```
  go mod init ben/haute
  go mod tidy
  go run cmd/api/main.go
  ```