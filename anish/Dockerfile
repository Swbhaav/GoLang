# up: decide base image for this application
From golang:1.24.2-alpine3.21
# selecting the workspace directory inside the container
WORKDIR /app
# copying all the folder and files from this project to inside app working directory,COPY src dest
COPY . .
#installing all packages
RUN go mod download
# building the go project
RUN go build -o secbapi .
# set command to run secbapi executable file
CMD ["./secbapi"]
# exposing host port
EXPOSE 8060
# set command to run secbapi executable file
CMD ["./secbapi"]