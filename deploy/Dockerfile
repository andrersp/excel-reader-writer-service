############################
# STEP 1 build executable binary
############################

FROM golang:1.23-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -v -a -installsuffix cgo -o /go/bin/app cmd/http/main.go


############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT [ "/go/bin/app" ]