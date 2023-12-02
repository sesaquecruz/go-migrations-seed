FROM golang:1.21.4-alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o build/sql-seed main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/build/sql-seed .
CMD [ "./sql-seed" ]
