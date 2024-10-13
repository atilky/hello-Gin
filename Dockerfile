FROM golang:latest
WORKDIR /gindemo1
COPY gindemo1 /gindemo1
EXPOSE 3000
CMD go mod tidy && go build -o main . && ./main