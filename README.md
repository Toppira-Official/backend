# Toppira

## Commands

- Generate Documents:

  ```sh
    swag init -o ./docs -g ./cmd/http/main.go --pd
  ```

- Generate Repositories:

  ```sh
    go run ./cmd/gen
  ```

- Run Linters:

  ```sh
    golangci-lint run
  ```
