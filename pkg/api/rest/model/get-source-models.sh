#!/bin/bash

# Set the repository URL and the target directories
REPO_URL="https://github.com/cloud-barista/cm-honeybee.git"
BRANCH="main"
MODEL_DIR="model"
TARGET_DIR="./source"
README_FILE="./README.md"

# Check if the cm-honeybee directory exists, and delete it if it does
if [ -d "cm-honeybee" ]; then
    rm -rf cm-honeybee
fi

# Then proceed with cloning
git clone -b $BRANCH --single-branch $REPO_URL

# Create target directory if it doesn't exist
mkdir -p $TARGET_DIR

# Copy the model directory contents to the target directory
cp -r cm-honeybee/$MODEL_DIR/* $TARGET_DIR

# Get the latest commit hash
cd cm-honeybee
LATEST_COMMIT_HASH=$(git rev-parse HEAD)
cd ..

# Write synchronization details to README file
{
    echo "## Source Model Synchronization Details"
    echo ""
    echo "### Synchronization Date"
    echo "- Date: $(date)"
    echo ""
    echo "### Source Repository Details"
    echo "- Repository: [cloud-barista/cm-honeybee]($REPO_URL)"
    echo "- Branch: $BRANCH"
    echo "- Latest Commit Hash: $LATEST_COMMIT_HASH"
    echo ""
    echo "### Usage Instructions"
    echo "- Update the models with the command below:"
    echo "  \`\`\`"
    echo "  make source-model"
    echo "  \`\`\`"
    echo "- This command synchronizes the models to the latest versions."
    echo "- Do not edit these files directly as they are managed automatically."
} > $README_FILE

# Remove the cloned repository
rm -rf cm-honeybee

# Confirmation message
echo "Model directory from cm-honeybee has been synchronized to $TARGET_DIR"
echo "Synchronization details recorded in $README_FILE"