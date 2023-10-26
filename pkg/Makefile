default:
	go build -o cm-beetle
cc:
	GOOS=linux GOARCH=arm go build -o cm-beetle-arm
swag swagger:
	~/go/bin/swag i -o ./api/rest/docs
run:
	./cm-beetle
runwithport:
	./cm-beetle --port=$(PORT)
clean:
	rm -v cm-beetle cm-beetle-arm
