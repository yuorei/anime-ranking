FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o /docker-anime-ranking
CMD [ "/docker-anime-ranking" ]