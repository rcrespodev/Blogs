FROM golang:1.18-alpine

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Download Go modules
RUN go mod download

# Build
RUN GOARCH=amd64 GOOS=linux go build -o ./bin/main-linux cmd/main.go

# To actually open the port, runtime parameters
# must be supplied to the docker command.
EXPOSE 8080

# (Optional) environment variable that our dockerised
# application can make use of. The value of environment
# variables can also be set via parameters supplied
# to the docker command on the command line.
#ENV HTTP_PORT=8081

# Run
CMD [ "./bin/main-linux" ]









































#Makefile

#BINARY_PATH=bin
#BINARY_NAME=main
#ENTRY_POINT=cmd/web/main.go

#tidy:
 #       go mod tidy

#build_app:
 #       GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH}/${BINARY_NAME}-linux ${ENTRY_POINT}
#       GOARCH=amd64 GOOS=darwin go build -o ${BINARY_PATH}/${BINARY_NAME}-darwin ${ENTRY_POINT}
        #go build -o bin/main cmd/web/main.go

#run_app:
 #       ./${BINARY_PATH}/${BINARY_NAME}-linux
        #go run cmd/web/main.go

#clean:
  #      go clean
 #       rm ${BINARY_PATH}/${BINARY_NAME}-linux

#test:
 #       go test ./... #init server before. The integrations test do requests on http:localhost/PORT

#build_and_run:  tidy build_app run_app

