build:
	go build -o bin/pagoda

doc:
	swag init --parseDependency --parseInternal

run:
	CompileDaemon -build="make build" -command="./bin/pagoda"

test:
	go test -v ./...

clean:
	rm -rf bin