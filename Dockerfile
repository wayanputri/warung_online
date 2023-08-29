FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "warung_online"
RUN go build -o warung_online

# run executable file
CMD ["./warung_online"]