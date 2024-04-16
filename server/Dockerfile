FROM golang:1.16

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/sirupsen/logrus
RUN go get -u github.com/spf13/viper
RUN go get -u golang.org/x/crypto/bcrypt

COPY . .

RUN go build -o main ./cmd/main.go

EXPOSE 8080

CMD ["./main"]