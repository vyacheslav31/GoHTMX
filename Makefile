.PHONY: install
install:
	@go install github.com/cosmtrek/air@latest
	@go install github.com/a-h/templ/cmd/templ@latest
	@go mod download

.PHONY: start
start:
	@air

.PHONY: generate
generate:
	@go generate ./...
	@templ generate

.PHONY: test
test:
	@go test ./...

.PHONY: coverage
coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

.PHONY: css
css:
	@npx tailwindcss -i internal/app/css/tailwind.css -o internal/app/dist/styles.css --minify
