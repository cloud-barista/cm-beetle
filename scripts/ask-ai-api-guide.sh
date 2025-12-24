#!/bin/bash

# Check if API path is provided
API_PATH=$1
if [ -z "$API_PATH" ]; then
    echo "Usage: $0 <api_path>"
    echo "Example: $0 /migration/ns/{nsId}/mci"
    exit 1
fi

# Path to the prompt file
PROMPT_FILE=".github/prompts/api-guide.prompt.md"

# Check if prompt file exists
if [ ! -f "$PROMPT_FILE" ]; then
    echo "Error: Prompt file not found at $PROMPT_FILE"
    exit 1
fi

# Read the prompt template
TEMPLATE=$(cat "$PROMPT_FILE")

# Replace {{api_path}} with the provided argument
# Using sed for compatibility. Escaping slashes in API_PATH is important.
ESCAPED_API_PATH=$(echo "$API_PATH" | sed 's/\//\\\//g')
PROMPT=$(echo "$TEMPLATE" | sed "s/{{api_path}}/$ESCAPED_API_PATH/g")

# Check authentication status by running a dummy prompt
echo "Checking GitHub Copilot authentication..."
if ! npx -y @github/copilot -p "hi" --silent &> /dev/null; then
    echo ""
    echo "❌ Error: You are not authenticated with GitHub Copilot."
    echo ""
    echo "👉 Please run the following command to login:"
    echo "   npx @github/copilot"
    echo "   > Then type '/login' and follow the instructions."
    echo ""
    echo "After logging in, run this command again."
    exit 1
fi

# Run copilot using npx (no installation required)
echo "Generating API Guide for: $API_PATH"
echo "----------------------------------------"
npx -y @github/copilot -p "$PROMPT" --model "claude-sonnet-4.5" --allow-tool 'shell'
