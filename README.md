# Setup

1. install Go and postgres
2. create your local database (e.g. `createdb hive-nexus-api`)
3. create a `.env` file and copy out the content of the `.env.example` and change it depending on your settings.
4. you may need to run `go mod download` or `go get ./...` to fetch all the required packages
5. run the application with `go run ./src/main.go`

We are using gofmt as a formatter to format our code. To format the repo use `gofmt -w -s -l .`
