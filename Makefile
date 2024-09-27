.PHONY: swagger
swagger:
	swag init -g cmd/http/main.go

.PHONY: dev
dev:
	air