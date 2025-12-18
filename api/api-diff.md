## DeepDiffGo Report
**Comparing:**
- Old: `prev_spec.yaml`
- New: `api/swagger.yaml`

### Modified APIs

#### `GET` /httpVersion
- `Response (200, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)
- `Response (404, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)
- `Response (500, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)

#### `POST` /migration/ns/{nsId}/mci
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].sourceCspImageName`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].resourceType`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].commandHistory`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].cspImageId`: Property added
- `Request (body, mciInfo, controller.MigrateInfraRequest, object) .targetVmOsImageList[].sourceVmUid`: Property added

#### `POST` /recommendation/containerInfra
- `Deprecated`: Deprecated status changed (From: `false`, To: `true`)
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel`: Property removed
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .servers`: Property added
- `Response (200, body, controller.RecommendInfraResponse, object)`: Structure changed (From: `controller.RecommendInfraResponse`, To: `common.SimpleMsg`)
- `Response (200, body, controller.RecommendInfraResponse, object) .status`: Property removed
- `Response (200, body, controller.RecommendInfraResponse, object) .targetInfra`: Property removed
- `Response (200, body, controller.RecommendInfraResponse, object) .description`: Property removed
- `Response (200, body, common.SimpleMsg, object) .message`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Response removed
- `Response (500, body, common.SimpleMsg, object)`: Response removed

#### `POST` /recommendation/mci
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].resourceType`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].commandHistory`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].sourceVmUid`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].sourceCspImageName`: Property added
- `Response (200, body, controller.RecommendVmInfraResponse, object) .targetVmOsImageList[].cspImageId`: Property added

#### `POST` /recommendation/resources/vmOsImages
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.sourceCspImageName`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.cspImageId`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.sourceVmUid`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.commandHistory`: Property added
- `Response (200, body, controller.RecommendVmOsImageResponse, object) .recommendedVmOsImageList[].targetVmOsImage.resourceType`: Property added

#### `GET` /test/tracing
- `Response (200, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)
- `Response (503, body, common.SimpleMessage, object)`: Structure changed (From: `common.SimpleMessage`, To: `common.SimpleMsg`)

### Added APIs

#### `GET` /migration/middleware/objectStorage

#### `POST` /migration/middleware/objectStorage

#### `DELETE` /migration/middleware/objectStorage/{objectStorageName}

#### `GET` /migration/middleware/objectStorage/{objectStorageName}

#### `HEAD` /migration/middleware/objectStorage/{objectStorageName}

#### `POST` /recommendation/k8sControlPlane

#### `POST` /recommendation/k8sNodeGroup

#### `POST` /recommendation/middleware/objectStorage

#### `POST` /recommendation/vmInfra

#### `POST` /report/migration/ns/{nsId}/mci/{mciId}

#### `DELETE` /request/{reqId}

#### `GET` /request/{reqId}

#### `DELETE` /requests

#### `GET` /requests

#### `POST` /summary/source

#### `GET` /summary/target/ns/{nsId}/mci/{mciId}

#### `GET` /test/streaming

