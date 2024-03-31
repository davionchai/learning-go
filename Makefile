.PHONY: run-hello
run-hello:
	go run ./hello_world

test:
	go test ./hello_world/hello -v

init:
	@unset APP_NAME && \
	read -p "Enter your APP_NAME (required): " APP_NAME && \
	mkdir $$APP_NAME && \
	cd $$APP_NAME && \
	go mod init github.com/davionchai/learning-go/$$APP_NAME && \
	cd .. && \
	go work use $$APP_NAME

tidy:
	go mod tidy
