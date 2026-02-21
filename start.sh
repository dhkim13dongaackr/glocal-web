#!/bin/bash

# Define the port
PORT=20080

echo "Stopping any existing server on port $PORT..."
# Find PID listening on the port and kill it
fuser -k $PORT/tcp 2>/dev/null

echo "Building the server..."
go build -o server .

if [ $? -eq 0 ]; then
    echo "Server build successful."
    echo "Starting server on http://localhost:$PORT ..."
    ./server &
else
    echo "Build failed. Please check your code."
fi
