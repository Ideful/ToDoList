FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY ./ ./

# make wait-for-postgres.sh executable
# RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]