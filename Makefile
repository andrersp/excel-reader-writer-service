.PHONY: swagger
swagger:
	swag init -g cmd/http/main.go

.PHONY: dev
dev:
	make swagger
	go run cmd/http/main.go