build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/getClients clients/getClients.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/getClientById clients/getClientById.go