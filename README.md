# Docker
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


  # Setup db
  ```
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  go run db/migrate/migrate.go
  ``` 

  # Go cmd
  ```
  go mod init ben/haute
  go mod tidy
  go run cmd/api/main.go
  ```