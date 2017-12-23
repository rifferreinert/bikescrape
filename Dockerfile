FROM golang
WORKDIR /go/src/github.com/rifferreinert/bikescrape
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep; dep ensure
RUN go install github.com/rifferreinert/bikescrape/cmd/initdb
RUN go install github.com/rifferreinert/bikescrape/cmd/pullstations
