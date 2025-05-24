build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/http-catch-all-linux-amd64
	GOOS=linux GOARCH=arm64 go build -o ./dist/http-catch-all-linux-arm64
	GOOS=darwin GOARCH=amd64 go build -o ./dist/http-catch-all-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o ./dist/http-catch-all-darwin-arm64
	GOOS=windows GOARCH=amd64 go build -o ./dist/http-catch-all-windows-amd64
	GOOS=windows GOARCH=arm64 go build -o ./dist/http-catch-all-windows-arm64