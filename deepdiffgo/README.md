# DeepDiffGo

**DeepDiffGo** is a lightweight CLI tool designed to compare Swagger (OpenAPI 2.0) and OpenAPI 3.0 specifications. It performs a "deep diff" by recursively resolving `$ref` references, enabling it to detect changes in actual data structures—even if the reference names differ.

## Installation

**Using wget (Linux/macOS):**

```bash
# Download the binary from the repository
wget https://raw.githubusercontent.com/cloud-barista/cm-beetle/main/deepdiffgo/deepdiffgo -O deepdiffgo

chmod +x deepdiffgo
sudo mv deepdiffgo /usr/local/bin/
```

**Using Go Install:**

```bash
go install github.com/cloud-barista/cm-beetle/deepdiffgo/cmd/deepdiffgo@latest

# Ensure $(go env GOPATH)/bin is in your PATH
```

<details>
<summary><strong>(Optional) Build and use DeepDiffGo</strong></summary>

If you prefer to build from source, ensure you have **Go 1.20+** installed.

1. Clone the repository:

   ```bash
   git clone https://github.com/cloud-barista/cm-beetle.git
   cd cm-beetle/deepdiffgo
   ```

2. Build the binary:

   ```bash
   go build -o deepdiffgo cmd/deepdiffgo/main.go
   ```

3. Move to your PATH:
   ```bash
   sudo mv deepdiffgo /usr/local/bin/
   ```

</details>

## Usage

### Basic Comparison

Compare two local Swagger/OpenAPI files:

```bash
deepdiffgo old_spec.yaml new_spec.yaml
```

### Remote Specification Support

Compare a remote specification with a local file (or two remote URLs):

```bash
deepdiffgo https://example.com/v1/swagger.yaml new_spec.yaml
```

### Output Formats & File Saving

Generate a report in **Markdown** or **JSON** format and save it to a file. This is particularly useful for CI/CD pipelines or GitHub PR comments.

```bash
# Save as Markdown
deepdiffgo old.yaml new.yaml -f markdown -o report.md

# Save as JSON
deepdiffgo old.yaml new.yaml -f json -o report.json
```

### Adding Descriptions to Specs

You can add descriptions (such as version tags, branch names, or commit SHAs) to each specification file. These descriptions will be displayed in the comparison report next to each file name.

```bash
# Add descriptions for version tracking
deepdiffgo old.yaml new.yaml --old-desc "v1.0.0" --new-desc "v1.1.0"

# Add descriptions for branch/commit tracking
deepdiffgo prev_spec.yaml api/swagger.yaml --old-desc "release/v1.0" --new-desc "main(abc123)"

# With Markdown output
deepdiffgo old.yaml new.yaml --old-desc "v1.0.0" --new-desc "v1.1.0" -f markdown -o report.md
```

**Output example:**

```
Diff between:
  - Old: prev_spec.yaml [v1.0.0]
  - New: api/swagger.yaml [main(abc123)]
```

### Help

View all available options:

```bash
deepdiffgo --help
```
