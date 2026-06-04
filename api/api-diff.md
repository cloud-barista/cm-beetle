## DeepDiffGo Report
**Diff between:**
- Old: `prev_spec.yaml [v0.5.0]`
- New: `api/swagger.yaml [main(45baffa)]`

### Modified APIs

#### `GET` /httpVersion
- `Response (200, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-string`)
- `Response (200, body, model.ApiResponse-string, object) .data`: Property added
- `Response (200, body, model.ApiResponse-string, object) .error`: Property added
- `Response (200, body, model.ApiResponse-string, object) .success`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Response removed
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added

#### `POST` /migration/data
- `Request (body, reqBody, transx.DataMigrationModel, object) .sourceTransferOptions`: Property removed
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination`: Type changed (From: ``, To: `object`)
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination`: Structure changed (From: ``, To: `transx.DataLocation`)
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.postCmd`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.preCmd`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.storageType`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.filesystem`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.filter`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.objectStorage`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destination.path`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .destinationTransferOptions`: Property removed
- `Request (body, reqBody, transx.DataMigrationModel, object) .source`: Type changed (From: ``, To: `object`)
- `Request (body, reqBody, transx.DataMigrationModel, object) .source`: Structure changed (From: ``, To: `transx.DataLocation`)
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.postCmd`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.preCmd`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.storageType`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.filesystem`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.filter`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.objectStorage`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .source.path`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .strategy`: Property added
- `Request (body, reqBody, transx.DataMigrationModel, object) .encryptionKeyId`: Property added
- `Request (header, X-Request-Id)`: Parameter 'X-Request-Id' (header) added
- `Response (200, body, common.SimpleMsg, object)`: Response removed
- `Response (400, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (400, body, model.ApiResponse-any, object) .data`: Property added
- `Response (400, body, model.ApiResponse-any, object) .error`: Property added
- `Response (400, body, model.ApiResponse-any, object) .success`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Response removed
- `Response (500, body, common.SimpleMsg, object)`: Response removed
- `Response (202, body, model.ApiResponse-model_AsyncJobResponse, object)`: Response added

#### `DELETE` /migration/ns/{nsId}/resources/securityGroup
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Response (200, body, model.IdList, object)`: Structure changed (From: `model.IdList`, To: `model.ResourceDeleteResults`)
- `Response (200, body, model.IdList, object) .output`: Property removed
- `Response (200, body, model.ResourceDeleteResults, object) .total`: Property added
- `Response (200, body, model.ResourceDeleteResults, object) .failedCount`: Property added
- `Response (200, body, model.ResourceDeleteResults, object) .results`: Property added
- `Response (200, body, model.ResourceDeleteResults, object) .successCount`: Property added

#### `GET` /migration/ns/{nsId}/resources/securityGroup
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added

#### `POST` /migration/ns/{nsId}/resources/securityGroup
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Response (200, body, model.SecurityGroupInfo, object) .connectionConfig.regionDetail.representativeZone`: Property added

#### `DELETE` /migration/ns/{nsId}/resources/securityGroup/{sgId}
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added

#### `GET` /migration/ns/{nsId}/resources/securityGroup/{sgId}
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Response (200, body, model.SecurityGroupInfo, object) .connectionConfig.regionDetail.representativeZone`: Property added

#### `GET` /migration/ns/{nsId}/resources/sshKey
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added

#### `POST` /migration/ns/{nsId}/resources/sshKey
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Response (200, body, model.SshKeyInfo, object) .connectionConfig.regionDetail.representativeZone`: Property added

#### `DELETE` /migration/ns/{nsId}/resources/sshKey/{sshKeyId}
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added

#### `GET` /migration/ns/{nsId}/resources/sshKey/{sshKeyId}
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Response (200, body, model.SshKeyInfo, object) .connectionConfig.regionDetail.representativeZone`: Property added

#### `GET` /migration/ns/{nsId}/resources/vNet
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Request (query, option)`: Parameter 'option' (query) added
- `Request (query, filterKey)`: Parameter 'filterKey' (query) added
- `Request (query, filterVal)`: Parameter 'filterVal' (query) added

#### `POST` /migration/ns/{nsId}/resources/vNet
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Response (200, body, model.VNetInfo, object)`: Response removed
- `Response (400, body, model.SimpleMsg, object)`: Response removed
- `Response (201, body, model.VNetInfo, object)`: Response added
- `Response (404, body, model.SimpleMsg, object)`: Response added
- `Response (409, body, model.SimpleMsg, object)`: Response added

#### `DELETE` /migration/ns/{nsId}/resources/vNet/{vNetId}
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added

#### `GET` /migration/ns/{nsId}/resources/vNet/{vNetId}
- `Request (header, x-request-id)`: Parameter 'x-request-id' (header) added
- `Request (header, x-credential-holder)`: Parameter 'x-credential-holder' (header) added
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].bastionNodes[].mciId`: Property removed
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].bastionNodes[].vmId`: Property removed
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].bastionNodes[].infraId`: Property added
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].bastionNodes[].nodeId`: Property added
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].bastionNodes[].nsId`: Property added
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].connectionConfig.regionDetail.representativeZone`: Property added
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].systemMessage`: Property added
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].conditions`: Property added
- `Response (200, body, model.VNetInfo, object) .connectionConfig.regionDetail.representativeZone`: Property added
- `Response (200, body, model.VNetInfo, object) .systemMessage`: Property added
- `Response (200, body, model.VNetInfo, object) .conditions`: Property added

#### `POST` /recommendation/k8sControlPlane
- `Response (400, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (400, body, model.ApiResponse-any, object) .data`: Property added
- `Response (400, body, model.ApiResponse-any, object) .error`: Property added
- `Response (400, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (200, body, model.K8sClusterDynamicReq, object)`: Structure changed (From: `model.K8sClusterDynamicReq`, To: `model.ApiResponse-model_K8sClusterDynamicReq`)
- `Response (200, body, model.K8sClusterDynamicReq, object) .imageId`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .onAutoScaling`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .connectionName`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .desiredNodeSize`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .label`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .maxNodeSize`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .name`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .rootDiskSize`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .rootDiskType`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .specId`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .minNodeSize`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .nodeGroupName`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .version`: Property removed
- `Response (200, body, model.K8sClusterDynamicReq, object) .description`: Property removed
- `Response (200, body, model.ApiResponse-model_K8sClusterDynamicReq, object) .success`: Property added
- `Response (200, body, model.ApiResponse-model_K8sClusterDynamicReq, object) .data`: Property added
- `Response (200, body, model.ApiResponse-model_K8sClusterDynamicReq, object) .error`: Property added
- `Response (200, body, model.ApiResponse-model_K8sClusterDynamicReq, object) .message`: Property added

#### `POST` /recommendation/k8sNodeGroup
- `Response (200, body, model.K8sNodeGroupReq, object)`: Structure changed (From: `model.K8sNodeGroupReq`, To: `model.ApiResponse-model_K8sNodeGroupReq`)
- `Response (200, body, model.K8sNodeGroupReq, object) .specId`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .description`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .desiredNodeSize`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .minNodeSize`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .name`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .sshKeyId`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .label`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .rootDiskSize`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .maxNodeSize`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .rootDiskType`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .imageId`: Property removed
- `Response (200, body, model.K8sNodeGroupReq, object) .onAutoScaling`: Property removed
- `Response (200, body, model.ApiResponse-model_K8sNodeGroupReq, object) .message`: Property added
- `Response (200, body, model.ApiResponse-model_K8sNodeGroupReq, object) .success`: Property added
- `Response (200, body, model.ApiResponse-model_K8sNodeGroupReq, object) .data`: Property added
- `Response (200, body, model.ApiResponse-model_K8sNodeGroupReq, object) .error`: Property added
- `Response (400, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (400, body, model.ApiResponse-any, object) .success`: Property added
- `Response (400, body, model.ApiResponse-any, object) .data`: Property added
- `Response (400, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added

#### `POST` /recommendation/middleware/objectStorage
- `Request (body, request, controller.RecommendObjectStorageRequest, object) .desiredCloud`: Structure changed (From: `cloudmodel.CloudProperty`, To: `storagemodel.CloudProperty`)
- `Request (body, request, controller.RecommendObjectStorageRequest, object) .sourceObjectStorages[]`: Structure changed (From: `controller.SourceObjectStorageProperty`, To: `storagemodel.SourceObjectStorage`)
- `Request (body, request, controller.RecommendObjectStorageRequest, object) .sourceObjectStorages[].corsRules`: Property removed
- `Request (body, request, controller.RecommendObjectStorageRequest, object) .sourceObjectStorages[].corsRule`: Property added
- `Response (200, body, controller.RecommendObjectStorageResponse, object)`: Structure changed (From: `controller.RecommendObjectStorageResponse`, To: `model.ApiResponse-storagemodel_RecommendedObjectStorage`)
- `Response (200, body, controller.RecommendObjectStorageResponse, object) .description`: Property removed
- `Response (200, body, controller.RecommendObjectStorageResponse, object) .status`: Property removed
- `Response (200, body, controller.RecommendObjectStorageResponse, object) .targetCloud`: Property removed
- `Response (200, body, controller.RecommendObjectStorageResponse, object) .targetObjectStorages`: Property removed
- `Response (200, body, model.ApiResponse-storagemodel_RecommendedObjectStorage, object) .data`: Property added
- `Response (200, body, model.ApiResponse-storagemodel_RecommendedObjectStorage, object) .error`: Property added
- `Response (200, body, model.ApiResponse-storagemodel_RecommendedObjectStorage, object) .message`: Property added
- `Response (200, body, model.ApiResponse-storagemodel_RecommendedObjectStorage, object) .success`: Property added
- `Response (400, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (400, body, model.ApiResponse-any, object) .error`: Property added
- `Response (400, body, model.ApiResponse-any, object) .success`: Property added
- `Response (400, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added

#### `POST` /recommendation/resources/securityGroups
- `Request (body, UserInfra, controller.RecommendVmInfraRequest, object)`: Structure changed (From: `controller.RecommendVmInfraRequest`, To: `controller.RecommendInfraRequest`)
- `Request (body, UserInfra, controller.RecommendVmInfraRequest, object) .onpremiseInfraModel.servers`: Property removed
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nodes`: Property added
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.k8sCluster`: Property added
- `Response (200, body, controller.RecommendSecurityGroupResponse, object)`: Structure changed (From: `controller.RecommendSecurityGroupResponse`, To: `model.ApiResponse-cloudmodel_RecommendedSecurityGroupList`)
- `Response (200, body, controller.RecommendSecurityGroupResponse, object) .count`: Property removed
- `Response (200, body, controller.RecommendSecurityGroupResponse, object) .description`: Property removed
- `Response (200, body, controller.RecommendSecurityGroupResponse, object) .status`: Property removed
- `Response (200, body, controller.RecommendSecurityGroupResponse, object) .targetSecurityGroupList`: Property removed
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedSecurityGroupList, object) .success`: Property added
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedSecurityGroupList, object) .data`: Property added
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedSecurityGroupList, object) .error`: Property added
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedSecurityGroupList, object) .message`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Response removed
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added
- `Response (400, body, model.ApiResponse-any, object)`: Response added

#### `POST` /recommendation/resources/vNet
- `Request (body, UserInfra, controller.RecommendVmInfraRequest, object)`: Structure changed (From: `controller.RecommendVmInfraRequest`, To: `controller.RecommendInfraRequest`)
- `Request (body, UserInfra, controller.RecommendVmInfraRequest, object) .onpremiseInfraModel.servers`: Property removed
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.nodes`: Property added
- `Request (body, UserInfra, controller.RecommendInfraRequest, object) .onpremiseInfraModel.k8sCluster`: Property added
- `Response (200, body, controller.RecommendVNetResponse, object)`: Structure changed (From: `controller.RecommendVNetResponse`, To: `model.ApiResponse-cloudmodel_RecommendedVNetList`)
- `Response (200, body, controller.RecommendVNetResponse, object) .count`: Property removed
- `Response (200, body, controller.RecommendVNetResponse, object) .description`: Property removed
- `Response (200, body, controller.RecommendVNetResponse, object) .targetVNetList`: Property removed
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVNetList, object) .success`: Property added
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVNetList, object) .data`: Property added
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVNetList, object) .error`: Property added
- `Response (200, body, model.ApiResponse-cloudmodel_RecommendedVNetList, object) .message`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Response removed
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added
- `Response (400, body, model.ApiResponse-any, object)`: Response added

#### `DELETE` /request/{reqId}
- `Response (200, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (200, body, model.ApiResponse-any, object) .error`: Property added
- `Response (200, body, model.ApiResponse-any, object) .success`: Property added
- `Response (200, body, model.ApiResponse-any, object) .data`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (404, body, model.ApiResponse-any, object) .data`: Property added
- `Response (404, body, model.ApiResponse-any, object) .error`: Property added
- `Response (404, body, model.ApiResponse-any, object) .success`: Property added

#### `GET` /request/{reqId}
- `Response (200, body, common.RequestDetails, object)`: Structure changed (From: `common.RequestDetails`, To: `model.ApiResponse-common_RequestDetails`)
- `Response (200, body, common.RequestDetails, object) .endTime`: Property removed
- `Response (200, body, common.RequestDetails, object) .errorResponse`: Property removed
- `Response (200, body, common.RequestDetails, object) .requestInfo`: Property removed
- `Response (200, body, common.RequestDetails, object) .responseData`: Property removed
- `Response (200, body, common.RequestDetails, object) .startTime`: Property removed
- `Response (200, body, common.RequestDetails, object) .status`: Property removed
- `Response (200, body, model.ApiResponse-common_RequestDetails, object) .data`: Property added
- `Response (200, body, model.ApiResponse-common_RequestDetails, object) .error`: Property added
- `Response (200, body, model.ApiResponse-common_RequestDetails, object) .message`: Property added
- `Response (200, body, model.ApiResponse-common_RequestDetails, object) .success`: Property added
- `Response (404, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (404, body, model.ApiResponse-any, object) .data`: Property added
- `Response (404, body, model.ApiResponse-any, object) .error`: Property added
- `Response (404, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (500, body, model.ApiResponse-any, object) .success`: Property added
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added

#### `DELETE` /requests
- `Response (200, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (200, body, model.ApiResponse-any, object) .data`: Property added
- `Response (200, body, model.ApiResponse-any, object) .error`: Property added
- `Response (200, body, model.ApiResponse-any, object) .success`: Property added

#### `POST` /summary/source
- `Request (body, Request, controller.GenerateSourceInfraSummaryRequest, object) .onpremiseInfraModel.servers`: Property removed
- `Request (body, Request, controller.GenerateSourceInfraSummaryRequest, object) .onpremiseInfraModel.k8sCluster`: Property added
- `Request (body, Request, controller.GenerateSourceInfraSummaryRequest, object) .onpremiseInfraModel.nodes`: Property added
- `Response (200, body)`: Type changed (From: ``, To: `object`)
- `Response (200, body)`: Structure changed (From: ``, To: `model.ApiResponse-summary_SourceInfraSummary`)
- `Response (200, body, model.ApiResponse-summary_SourceInfraSummary, object) .data`: Property added
- `Response (200, body, model.ApiResponse-summary_SourceInfraSummary, object) .error`: Property added
- `Response (200, body, model.ApiResponse-summary_SourceInfraSummary, object) .message`: Property added
- `Response (200, body, model.ApiResponse-summary_SourceInfraSummary, object) .success`: Property added
- `Response (400, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object)`: Structure changed (From: `github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response`, To: `model.ApiResponse-any`)
- `Response (400, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .object`: Property removed
- `Response (400, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .text`: Property removed
- `Response (400, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .details`: Property removed
- `Response (400, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .list`: Property removed
- `Response (400, body, model.ApiResponse-any, object) .message`: Property added
- `Response (400, body, model.ApiResponse-any, object) .data`: Property added
- `Response (400, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object)`: Structure changed (From: `github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response`, To: `model.ApiResponse-any`)
- `Response (500, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .text`: Property removed
- `Response (500, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .details`: Property removed
- `Response (500, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .list`: Property removed
- `Response (500, body, github_com_cloud-barista_cm-beetle_pkg_api_rest_model_beetle.Response, object) .object`: Property removed
- `Response (500, body, model.ApiResponse-any, object) .data`: Property added
- `Response (500, body, model.ApiResponse-any, object) .error`: Property added
- `Response (500, body, model.ApiResponse-any, object) .message`: Property added

#### `GET` /test/tracing
- `Request (header, traceparent)`: Parameter 'traceparent' (header) added
- `Response (200, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-string`)
- `Response (200, body, model.ApiResponse-string, object) .data`: Property added
- `Response (200, body, model.ApiResponse-string, object) .error`: Property added
- `Response (200, body, model.ApiResponse-string, object) .success`: Property added
- `Response (503, body, common.SimpleMsg, object)`: Structure changed (From: `common.SimpleMsg`, To: `model.ApiResponse-any`)
- `Response (503, body, model.ApiResponse-any, object) .data`: Property added
- `Response (503, body, model.ApiResponse-any, object) .error`: Property added
- `Response (503, body, model.ApiResponse-any, object) .success`: Property added

### Added APIs

#### `GET` /migration/data/encryptionKey

#### `POST` /migration/data/test/decrypt

#### `POST` /migration/data/test/encrypt

#### `GET` /migration/middleware/ns/{nsId}/objectStorage

#### `POST` /migration/middleware/ns/{nsId}/objectStorage

#### `DELETE` /migration/middleware/ns/{nsId}/objectStorage/{osId}

#### `GET` /migration/middleware/ns/{nsId}/objectStorage/{osId}

#### `HEAD` /migration/middleware/ns/{nsId}/objectStorage/{osId}

#### `GET` /migration/middleware/ns/{nsId}/objectStorage/{osId}/object

#### `DELETE` /migration/middleware/ns/{nsId}/objectStorage/{osId}/object/{objectKey}

#### `HEAD` /migration/middleware/ns/{nsId}/objectStorage/{osId}/object/{objectKey}

#### `GET` /migration/ns/{nsId}/infra

#### `POST` /migration/ns/{nsId}/infra

#### `DELETE` /migration/ns/{nsId}/infra/{infraId}

#### `GET` /migration/ns/{nsId}/infra/{infraId}

#### `POST` /migration/ns/{nsId}/infraWithDefaults

#### `DELETE` /migration/ns/{nsId}/resources/sshKey

#### `DELETE` /migration/ns/{nsId}/resources/vNet

#### `POST` /naming/alignment

#### `POST` /naming/preview

#### `POST` /naming/validation

#### `POST` /recommendation/infra

#### `POST` /recommendation/infraWithDefaults

#### `POST` /recommendation/resources/osImages

#### `POST` /recommendation/resources/specs

#### `POST` /report/migration/ns/{nsId}/infra/{infraId}

#### `GET` /summary/target/ns/{nsId}/infra/{infraId}

#### `GET` /test/auth

### Removed APIs

#### `GET` /migration/middleware/objectStorage

#### `POST` /migration/middleware/objectStorage

#### `DELETE` /migration/middleware/objectStorage/{objectStorageName}

#### `GET` /migration/middleware/objectStorage/{objectStorageName}

#### `HEAD` /migration/middleware/objectStorage/{objectStorageName}

#### `GET` /migration/ns/{nsId}/mci

#### `POST` /migration/ns/{nsId}/mci

#### `DELETE` /migration/ns/{nsId}/mci/{mciId}

#### `GET` /migration/ns/{nsId}/mci/{mciId}

#### `POST` /migration/ns/{nsId}/mciWithDefaults

#### `POST` /recommendation/containerInfra

#### `POST` /recommendation/mci

#### `POST` /recommendation/mciWithDefaults

#### `POST` /recommendation/resources/vmOsImages

#### `POST` /recommendation/resources/vmSpecs

#### `POST` /recommendation/vmInfra

#### `POST` /report/migration/ns/{nsId}/mci/{mciId}

#### `GET` /summary/target/ns/{nsId}/mci/{mciId}

