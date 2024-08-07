test:
	go test -timeout 300ms ./...

coverage:
	go test -coverprofile cover.out -timeout 300ms ./... || go tool cover -html=cover.out -o index.html
