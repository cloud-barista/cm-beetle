default:
	go build -o beetle
cc:
	GOOS=linux GOARCH=arm go build -o cm-beetle-arm
run:
	./beetle
# runwithport:
# 	./cm-beetle --port=$(PORT)
clean:
	rm -v beetle