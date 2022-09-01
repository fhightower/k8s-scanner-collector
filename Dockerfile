FROM golang:1.18.0-buster

# Set the working directory to /app
WORKDIR /go/src/code/

# Copy the current directory contents into the container at /app
ADD . /go/src/code/

RUN go get golang.org/x/net/html

# Make port 80 available to the world outside this container
EXPOSE 80

# Run app.py when the container launches
CMD ["go", "run", "collector.go"]

