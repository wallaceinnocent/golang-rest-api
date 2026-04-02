# Use a tiny base image
FROM alpine:latest

# Security: Add certificates so your app can make HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Just copy the pre-built binary from your host machine into the image
COPY cmd/server/server .

# Ensure the binary has execution permissions
RUN chmod +x ./server

EXPOSE 8080

CMD ["./server"]