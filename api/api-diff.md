## DeepDiffGo Report
**Diff between:**
- Old: `prev_spec.yaml [v0.5.8]`
- New: `api/swagger.yaml [main(d18b3ba)]`

### Modified APIs

#### `GET` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}/healthz
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/middleware/ns/{nsId}/objectStorage
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/middleware/ns/{nsId}/objectStorage/{osId}
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `HEAD` /migration/middleware/ns/{nsId}/objectStorage/{osId}
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `GET` /migration/middleware/ns/{nsId}/objectStorage/{osId}/object
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `DELETE` /migration/middleware/ns/{nsId}/objectStorage/{osId}/object/{objectKey}
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `HEAD` /migration/middleware/ns/{nsId}/objectStorage/{osId}/object/{objectKey}
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `POST` /recommendation/k8sControlPlane
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `POST` /recommendation/k8sNodeGroup
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `POST` /recommendation/resources/securityGroups
- `Response (503, body, model.ApiResponse-any, object)`: Response added

#### `POST` /recommendation/resources/vNet
- `Response (503, body, model.ApiResponse-any, object)`: Response added

### Added APIs

#### `POST` /recommendation/multiInfra

#### `POST` /recommendation/multiInfraWithNlb

