# Image to run application
FROM alpine

# Copy binary to container
COPY ./update/update /update
COPY ./create/create /create
