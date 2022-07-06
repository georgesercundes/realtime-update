run-websocket:
	cd server/websocket/ && go run *.go

run-sse:
	cd server/sse/ && go run *.go

run-client:
	cd client && yarn start