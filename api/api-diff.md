# DeepDiffGo Report
**Comparing:**
- Old: `prev_spec.yaml`
- New: `api/swagger.yaml`

### [*] GET `/httpVersion`
- * `Response (500, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)
- * `Response (200, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)
- * `Response (404, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)

### [+] POST `/recommendation/vmInfra` (New API)

### [+] DELETE `/request/{reqId}` (New API)

### [+] GET `/request/{reqId}` (New API)

### [+] DELETE `/requests` (New API)

### [+] GET `/requests` (New API)

### [+] GET `/test/streaming` (New API)

### [*] GET `/test/tracing`
- * `Response (503, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)
- * `Response (200, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)

