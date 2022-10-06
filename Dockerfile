FROM golang

WORKDIR /go/src/github.com/tuanbieber/integration-golang

COPY . .

RUN go build -o main -buildvcs=false

CMD ./main
