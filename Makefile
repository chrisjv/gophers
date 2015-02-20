all:
	gopherjs build -o static/client.js
	go build -o server
