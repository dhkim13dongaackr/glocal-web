#!/bin/bash

PORT=20080

echo "🔌 Killing process on port $PORT..."

# Try using lsof if available
if command -v lsof >/dev/null 2>&1; then
    PID=$(lsof -t -i:$PORT)
    if [ -n "$PID" ]; then
        kill -9 $PID
        echo "✅ Process $PID killed."
    else
        echo "Example: No process found on port $PORT."
    fi
else
    # Fallback to pkill (kills 'server' or 'main' binaries)
    echo "⚠️ 'lsof' not found. Trying pkill..."
    pkill -f "tmp/main" 2>/dev/null
    pkill -f "./server" 2>/dev/null
fi

echo "🚀 Starting Air..."
# Check if air is in path, otherwise assume ~/go/bin/air
if command -v air >/dev/null 2>&1; then
    air
else
    ~/go/bin/air
fi
