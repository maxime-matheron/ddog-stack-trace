FROM golang:1.16-alpine

WORKDIR /code

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /dd-golang-stack

EXPOSE 5001

CMD [ "/dd-golang-stack" ]