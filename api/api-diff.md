## DeepDiffGo Report
**Diff between:**
- Old: `prev_spec.yaml [v0.5.1]`
- New: `api/swagger.yaml [main(c3afb3d)]`

### Modified APIs

#### `POST` /migration/ns/{nsId}/infra
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetOsImageList[].isBasicGpuImage`: Property added

#### `POST` /naming/alignment
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetOsImageList[].isBasicGpuImage`: Property added

#### `POST` /naming/preview
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetOsImageList[].isBasicGpuImage`: Property added

#### `POST` /naming/validation
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetOsImageList[].isBasicGpuImage`: Property added

#### `POST` /recommendation/infra
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetOsImageList[].isBasicGpuImage`: Property added

