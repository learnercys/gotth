tailwind-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

tailwind-build:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

templ-generate:
	templ generate

dev:
	go build -o ./tmp/$(APP_NAME) && ./cmd/$(APP_NAME)/main.go && air

build:
	make tailwind-build

.PHONY: build tailwind-build tailwind-watch templ-generate dev
