## DeepDiffGo Report
**Diff between:**
- Old: `prev_spec.yaml [v0.5.2]`
- New: `api/swagger.yaml [main(229c59d)]`

### Modified APIs

#### `POST` /migration/ns/{nsId}/infra
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetNlbList`: Property added
- `Request (query, useExisting)`: Parameter 'useExisting' (query) added

#### `POST` /naming/alignment
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetNlbList`: Property added

#### `POST` /naming/preview
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetNlbList`: Property added

#### `POST` /naming/validation
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetNlbList`: Property added

#### `POST` /recommendation/infra
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nlbs`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetNlbList`: Property added

#### `POST` /recommendation/infraWithDefaults
- `Request (body, UserInfra, controller.RecommendInfraWithDefaultsRequest, object) .onpremiseInfraModel.nlbs`: Property added

#### `POST` /recommendation/resources/osImages
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nlbs`: Property added

#### `POST` /recommendation/resources/securityGroups
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nlbs`: Property added

#### `POST` /recommendation/resources/specs
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nlbs`: Property added

#### `POST` /recommendation/resources/vNet
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nlbs`: Property added

#### `POST` /report/migration/ns/{nsId}/infra/{infraId}
- `Request (body, onpremiseInfraModel, controller.GenerateMigrationReportRequest, object) .onpremiseInfraModel.nlbs`: Property added

#### `POST` /summary/source
- `Request (body, Request, controller.GenerateSourceInfraSummaryRequest, object) .onpremiseInfraModel.nlbs`: Property added

### Added APIs

#### `GET` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb

#### `POST` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb

#### `DELETE` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}

#### `GET` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb/{nlbId}

#### `POST` /recommendation/infraWithNlb

