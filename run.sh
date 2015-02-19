# Compile client
GLOBIGNORE="server.go"
gopherjs build *.go -o static/client.js

# Run server
GLOBIGNORE="client.go"
go run *.go
