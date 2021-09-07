FROM golang:1.16-alpine

WORKDIR /code

ADD ./src ./
RUN go mod download
RUN go build -o /ddgolangstack

EXPOSE 5001
CMD [ "/ddgolangstack" ]