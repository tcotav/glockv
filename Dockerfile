# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# install SQLite
RUN apt-get update
RUN apt-get install sqlite3 libsqlite3-dev -y

# Copy the local package files to the container's workspace.
COPY src /go/src

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)

#ENTRYPOINT go test /go/src/github.com/tcotav/glockv/model_test/model_test.go

# Run the outyet command by default when the container starts.
#ENTRYPOINT /go/bin/glockv

# Document that the service listens on port 8080.
#EXPOSE 8080
