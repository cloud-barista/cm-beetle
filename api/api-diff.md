## DeepDiffGo Report
**Comparing:**
- Old: `prev_spec.yaml`
- New: `api/swagger.yaml`

### Modified APIs

#### `POST` /migration/ns/{nsId}/mci
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].resourceType`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].commandHistory`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].sourceVmUid`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].sourceCspImageName`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].cspImageId`: Property added

#### `POST` /recommendation/mci
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].resourceType`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].cspImageId`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].sourceCspImageName`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].sourceVmUid`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].commandHistory`: Property added

#### `POST` /recommendation/resources/vmOsImages
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.sourceVmUid`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.sourceCspImageName`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.commandHistory`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.resourceType`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.cspImageId`: Property added

#### `POST` /recommendation/vmInfra
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVmInfra, object)`: Structure changed (From: `model.ApiResponse-cloudmodel_RecommendedVmInfra`, To: `model.ApiResponse-array_cloudmodel_RecommendedVmInfra`)
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVmInfra, object) .item`: Property removed
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVmInfra, object) .items`: Property removed
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedVmInfra, object) .data`: Property added
- `Response (400, body, model.ApiResponse-any, object) .item`: Property removed
- `Response (400, body, model.ApiResponse-any, object) .items`: Property removed
- `Response (400, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .item`: Property removed
- `Response (500, body, model.ApiResponse-any, object) .items`: Property removed
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added

