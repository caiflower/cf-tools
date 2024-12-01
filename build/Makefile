PWD="../"
CMD_NAME="cf"

.PHONY: linux

linux:
	env CGO_ENABLED=0 GO111MODULE=on GOOS=linux  GOARCH=amd64 go build -v -o $(CMD_NAME) $(PWD)/main.go
	chmod +x $(CMD_NAME)

install:
	env CGO_ENABLED=0 GO111MODULE=on go build -v -o $(CMD_NAME) $(PWD)/main.go
	chmod +x $(CMD_NAME)
	echo "sudo mv ./cf /usr/local/bin"

clean:
	rm -rf $(CMD_NAME)