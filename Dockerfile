FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ship-it/*.go ./
RUN go build -o /ship-it
EXPOSE 8080
CMD [ "/ship-it" ]