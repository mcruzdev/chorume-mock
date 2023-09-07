install:
	go build -o chorumemock cmd/chorumemock/main.go
	chmod +x chorumemock
	mv chorumemock /usr/local/bin/
