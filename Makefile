.PHONY: run-hello
run-hello:
	go run ./hello_world

.PHONY: test
test:
	go test ./hello_world/hello -v

.PHONY: init
init:
	@unset APP_NAME && \
	read -p "Enter your APP_NAME (required): " APP_NAME && \
	mkdir $$APP_NAME && \
	cd $$APP_NAME && \
	go mod init github.com/davionchai/learning-go/$$APP_NAME && \
	cd .. && \
	go work use $$APP_NAME

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: install
install:
	sudo rm -rf /usr/local/go

# download the desired version from this page https://go.dev/dl/
.PHONY: uninstall
uninstall:
	wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
	sudo tar -C /usr/local -xzf ./go1.23.0.linux-amd64.tar.gz
	go version