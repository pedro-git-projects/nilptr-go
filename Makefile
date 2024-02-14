run:
	cd ./src/ && go run *.go

watch:
	cd ./src/ && air *.go

test:
	cd ./src/ && go test ./...
