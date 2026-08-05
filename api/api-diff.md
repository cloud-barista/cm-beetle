## DeepDiffGo Report
**Diff between:**
- Old: `prev_spec.yaml [v0.5.7]`
- New: `api/swagger.yaml [main(873ed10)]`

### Modified APIs

#### `GET` /migration/ns/{nsId}/infra
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/ns/{nsId}/infra/{infraId}
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/ns/{nsId}/infra/{infraId}/ssh-ready
- `Response (429, body, model.ApiResponse-controller_CheckSSHReadyResponse, object)`: Structure changed (From: `model.ApiResponse-controller_CheckSSHReadyResponse`, To: `model.ApiResponse-any`)

#### `POST` /migration/ns/{nsId}/infraWithDefaults
- `Request (body, mciInfo, controller.MigrateInfraWithDefaultsRequest, object)`: Parameter 'mciInfo' (body) removed
- `Request (body, infraInfo, controller.MigrateInfraWithDefaultsRequest, object)`: Parameter 'infraInfo' (body) added

#### `GET` /summary/target/ns/{nsId}/infra/{infraId}
- `Request (path, mciId)`: Parameter 'mciId' (path) removed
- `Request (path, infraId)`: Parameter 'infraId' (path) added

