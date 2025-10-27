#!/bin/bash

# create-test-data.sh
# Creates test data with various directory depths and file types for testing object storage filtering

set -e

# Default test data directory
TEST_DATA_DIR="${1:-/tmp/transx-test-data}"

echo "=========================================="
echo "Creating Test Data for Object Storage"
echo "=========================================="
echo "Target directory: ${TEST_DATA_DIR}"
echo ""

# Clean up existing test data
if [ -d "${TEST_DATA_DIR}" ]; then
    echo "Removing existing test data..."
    rm -rf "${TEST_DATA_DIR}"
fi

# Create base directory
mkdir -p "${TEST_DATA_DIR}"

echo "Creating directory structure with multiple depths..."

# Root level files (depth 0)
echo "# Root README" > "${TEST_DATA_DIR}/README.md"
echo "root-level data" > "${TEST_DATA_DIR}/root-file.txt"
echo '{"root": "config"}' > "${TEST_DATA_DIR}/config.json"
echo "root,data,values" > "${TEST_DATA_DIR}/data.csv"
echo "debug log at root" > "${TEST_DATA_DIR}/debug.log"
echo "temp file at root" > "${TEST_DATA_DIR}/temp.tmp"

# Level 1: Main directories
echo "Creating level 1 directories..."

## src/ directory - source code
mkdir -p "${TEST_DATA_DIR}/src"
echo "package main" > "${TEST_DATA_DIR}/src/main.go"
echo "// Main utility" > "${TEST_DATA_DIR}/src/utils.go"
echo '{"version": "1.0.0"}' > "${TEST_DATA_DIR}/src/package.json"
echo "build log" > "${TEST_DATA_DIR}/src/build.log"

## docs/ directory - documentation
mkdir -p "${TEST_DATA_DIR}/docs"
echo "# Documentation" > "${TEST_DATA_DIR}/docs/README.md"
echo "User guide content" > "${TEST_DATA_DIR}/docs/user-guide.txt"
echo "API reference" > "${TEST_DATA_DIR}/docs/api-reference.txt"
cat > "${TEST_DATA_DIR}/docs/tutorial.pdf" << 'EOF'
%PDF-1.4
1 0 obj<</Type/Catalog/Pages 2 0 R>>endobj
2 0 obj<</Type/Pages/Count 1/Kids[3 0 R]>>endobj
3 0 obj<</Type/Page/MediaBox[0 0 612 792]/Parent 2 0 R/Resources<<>>>>endobj
xref
0 4
trailer<</Size 4/Root 1 0 R>>
startxref
149
%%EOF
EOF

## data/ directory - data files
mkdir -p "${TEST_DATA_DIR}/data"
echo "id,name,value" > "${TEST_DATA_DIR}/data/dataset1.csv"
echo "x,y,z" > "${TEST_DATA_DIR}/data/dataset2.csv"
echo '{"data": [1,2,3]}' > "${TEST_DATA_DIR}/data/results.json"
echo '{"config": "data"}' > "${TEST_DATA_DIR}/data/config.json"
echo "data processing log" > "${TEST_DATA_DIR}/data/process.log"

## logs/ directory - log files
mkdir -p "${TEST_DATA_DIR}/logs"
echo "[INFO] Application started" > "${TEST_DATA_DIR}/logs/app.log"
echo "[ERROR] Connection failed" > "${TEST_DATA_DIR}/logs/error.log"
echo "[DEBUG] Debug info" > "${TEST_DATA_DIR}/logs/debug.log"
echo "2025-10-27 00:00:00" > "${TEST_DATA_DIR}/logs/access.log"

## temp/ directory - temporary files
mkdir -p "${TEST_DATA_DIR}/temp"
echo "temporary data 1" > "${TEST_DATA_DIR}/temp/file1.tmp"
echo "temporary data 2" > "${TEST_DATA_DIR}/temp/file2.tmp"
echo "cache data" > "${TEST_DATA_DIR}/temp/cache.tmp"

## backup/ directory - backup files
mkdir -p "${TEST_DATA_DIR}/backup"
echo "backup config" > "${TEST_DATA_DIR}/backup/config-backup.json"
echo "backup data" > "${TEST_DATA_DIR}/backup/data-backup.csv"
echo "old backup" > "${TEST_DATA_DIR}/backup/old-backup.txt"

# Level 2: Nested subdirectories
echo "Creating level 2 directories..."

## src/models/ - nested source
mkdir -p "${TEST_DATA_DIR}/src/models"
echo "package models" > "${TEST_DATA_DIR}/src/models/user.go"
echo "package models" > "${TEST_DATA_DIR}/src/models/product.go"
echo '{"schema": "user"}' > "${TEST_DATA_DIR}/src/models/schema.json"

## src/tests/ - test files
mkdir -p "${TEST_DATA_DIR}/src/tests"
echo "package tests" > "${TEST_DATA_DIR}/src/tests/user_test.go"
echo "test log output" > "${TEST_DATA_DIR}/src/tests/test.log"
echo "coverage report" > "${TEST_DATA_DIR}/src/tests/coverage.txt"

## docs/images/ - image files
mkdir -p "${TEST_DATA_DIR}/docs/images"
echo "PNG image data" > "${TEST_DATA_DIR}/docs/images/diagram.png"
echo "JPG image data" > "${TEST_DATA_DIR}/docs/images/screenshot.jpg"
echo "SVG image data" > "${TEST_DATA_DIR}/docs/images/icon.svg"

## data/raw/ - raw data
mkdir -p "${TEST_DATA_DIR}/data/raw"
echo "raw,data,1" > "${TEST_DATA_DIR}/data/raw/raw1.csv"
echo "raw,data,2" > "${TEST_DATA_DIR}/data/raw/raw2.csv"
echo '{"raw": true}' > "${TEST_DATA_DIR}/data/raw/metadata.json"

## data/processed/ - processed data
mkdir -p "${TEST_DATA_DIR}/data/processed"
echo "processed,data,1" > "${TEST_DATA_DIR}/data/processed/output1.csv"
echo "processed,data,2" > "${TEST_DATA_DIR}/data/processed/output2.csv"
echo '{"processed": true}' > "${TEST_DATA_DIR}/data/processed/summary.json"

## logs/archive/ - archived logs
mkdir -p "${TEST_DATA_DIR}/logs/archive"
echo "archived log 1" > "${TEST_DATA_DIR}/logs/archive/app-2025-01.log"
echo "archived log 2" > "${TEST_DATA_DIR}/logs/archive/app-2025-02.log"
echo "archived errors" > "${TEST_DATA_DIR}/logs/archive/errors-2025-01.log"

# Level 3: Deep nested directories
echo "Creating level 3 directories..."

## src/models/entities/ - deep nesting
mkdir -p "${TEST_DATA_DIR}/src/models/entities"
echo "package entities" > "${TEST_DATA_DIR}/src/models/entities/base.go"
echo "package entities" > "${TEST_DATA_DIR}/src/models/entities/customer.go"

## data/raw/2025/ - date-based directory
mkdir -p "${TEST_DATA_DIR}/data/raw/2025"
echo "jan,data" > "${TEST_DATA_DIR}/data/raw/2025/january.csv"
echo "feb,data" > "${TEST_DATA_DIR}/data/raw/2025/february.csv"
echo '{"month": "jan"}' > "${TEST_DATA_DIR}/data/raw/2025/january.json"

## data/processed/reports/ - nested reports
mkdir -p "${TEST_DATA_DIR}/data/processed/reports"
echo "monthly report" > "${TEST_DATA_DIR}/data/processed/reports/monthly.txt"
echo "quarterly report" > "${TEST_DATA_DIR}/data/processed/reports/quarterly.txt"
echo "report,metrics" > "${TEST_DATA_DIR}/data/processed/reports/metrics.csv"

# Level 4: Very deep nesting
echo "Creating level 4 directories..."

## data/raw/2025/Q1/ - very deep structure
mkdir -p "${TEST_DATA_DIR}/data/raw/2025/Q1"
echo "q1,jan,data" > "${TEST_DATA_DIR}/data/raw/2025/Q1/jan-sales.csv"
echo "q1,feb,data" > "${TEST_DATA_DIR}/data/raw/2025/Q1/feb-sales.csv"
echo '{"quarter": "Q1"}' > "${TEST_DATA_DIR}/data/raw/2025/Q1/summary.json"
echo "Q1 processing log" > "${TEST_DATA_DIR}/data/raw/2025/Q1/process.log"

# Create .git directory (commonly excluded)
echo "Creating .git directory..."
mkdir -p "${TEST_DATA_DIR}/.git/objects"
echo "git object data" > "${TEST_DATA_DIR}/.git/objects/abc123"
echo "ref: refs/heads/main" > "${TEST_DATA_DIR}/.git/HEAD"
echo "git config" > "${TEST_DATA_DIR}/.git/config"

# Create node_modules directory (commonly excluded)
echo "Creating node_modules directory..."
mkdir -p "${TEST_DATA_DIR}/node_modules/package1"
echo '{"name": "package1"}' > "${TEST_DATA_DIR}/node_modules/package1/package.json"
echo "module code" > "${TEST_DATA_DIR}/node_modules/package1/index.js"
mkdir -p "${TEST_DATA_DIR}/node_modules/package2"
echo '{"name": "package2"}' > "${TEST_DATA_DIR}/node_modules/package2/package.json"

# Create hidden files
echo "Creating hidden files..."
echo "environment variables" > "${TEST_DATA_DIR}/.env"
echo "editor config" > "${TEST_DATA_DIR}/.editorconfig"
echo "git ignore rules" > "${TEST_DATA_DIR}/.gitignore"

# Summary
echo ""
echo "=========================================="
echo "Test Data Creation Complete!"
echo "=========================================="
echo ""
echo "Directory structure:"
tree -L 3 "${TEST_DATA_DIR}" 2>/dev/null || find "${TEST_DATA_DIR}" -type d | head -20

echo ""
echo "File statistics:"
echo "  Total files: $(find "${TEST_DATA_DIR}" -type f | wc -l)"
echo "  Total directories: $(find "${TEST_DATA_DIR}" -type d | wc -l)"
echo "  Max depth: 4 levels"
echo ""

echo "File types distribution:"
echo "  .txt files: $(find "${TEST_DATA_DIR}" -name "*.txt" | wc -l)"
echo "  .json files: $(find "${TEST_DATA_DIR}" -name "*.json" | wc -l)"
echo "  .csv files: $(find "${TEST_DATA_DIR}" -name "*.csv" | wc -l)"
echo "  .log files: $(find "${TEST_DATA_DIR}" -name "*.log" | wc -l)"
echo "  .tmp files: $(find "${TEST_DATA_DIR}" -name "*.tmp" | wc -l)"
echo "  .go files: $(find "${TEST_DATA_DIR}" -name "*.go" | wc -l)"
echo "  .md files: $(find "${TEST_DATA_DIR}" -name "*.md" | wc -l)"
echo ""

echo "Test scenarios you can try:"
echo ""
echo "1. Exclude all log files:"
echo "   \"exclude\": [\"*.log\"]"
echo ""
echo "2. Include only JSON and CSV files:"
echo "   \"include\": [\"*.json\", \"*.csv\"]"
echo ""
echo "3. Exclude temporary and log files:"
echo "   \"exclude\": [\"*.tmp\", \"*.log\"]"
echo ""
echo "4. Exclude .git and node_modules directories:"
echo "   \"exclude\": [\".git/*\", \"node_modules/*\"]"
echo ""
echo "5. Include only data files, exclude logs:"
echo "   \"include\": [\"data/*\"],"
echo "   \"exclude\": [\"*.log\"]"
echo ""
echo "6. Include docs, exclude images:"
echo "   \"include\": [\"docs/*\"],"
echo "   \"exclude\": [\"*.png\", \"*.jpg\", \"*.svg\"]"
echo ""
echo "7. Exclude backup and archive directories:"
echo "   \"exclude\": [\"backup/*\", \"*/archive/*\"]"
echo ""
echo "=========================================="
