# Base image
FROM golang:1.17-alpine

# Create workspace dir in $GOPATH
RUN mkdir -p $GOPATH/src/github.com/vitt-bagal/go-url-shortener
# COPY code to workspace 
COPY . $GOPATH/src/github.com/vitt-bagal/go-url-shortener/
# Build service
RUN cd $GOPATH/src/github.com/vitt-bagal/go-url-shortener && go mod download && go build -o /url-shortner-service

# Expose port to service
EXPOSE 9090
# Default entrypoint to service
CMD [ "/url-shortner-service" ]
# End of Dockerfile