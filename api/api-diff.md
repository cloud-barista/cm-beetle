## DeepDiffGo Report
**Diff between:**
- Old: `prev_spec.yaml [v0.5.9]`
- New: `api/swagger.yaml [main(8713735)]`

### Modified APIs

#### `GET` /migration/middleware/ns/{nsId}/infra/{infraId}/nlb
- `Response (200, body, model.ApiResponse-array_cloudmodel_NLBInfo, object) .data[].connectionConfig.verifiedMessage`: Property added

#### `POST` /migration/ns/{nsId}/infra
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetInfra.nodeGroups[].subnetIds`: Property added
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetInfra.postCommand`: Property removed
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetInfra.postCommandAsync`: Property added
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetInfra.postCommands`: Property added
- `Request (body, infraInfo, controller.MigrateInfraRequest, object) .targetK8sCluster`: Property added

#### `POST` /migration/ns/{nsId}/infraWithDefaults
- `Request (body, infraInfo, controller.MigrateInfraWithDefaultsRequest, object) .postCommand`: Property removed
- `Request (body, infraInfo, controller.MigrateInfraWithDefaultsRequest, object) .nodeGroups[].nodeUserPassword`: Property removed
- `Request (body, infraInfo, controller.MigrateInfraWithDefaultsRequest, object) .nodeGroups[].distributeSubnets`: Property added
- `Request (body, infraInfo, controller.MigrateInfraWithDefaultsRequest, object) .postCommandAsync`: Property added
- `Request (body, infraInfo, controller.MigrateInfraWithDefaultsRequest, object) .postCommands`: Property added

#### `POST` /migration/ns/{nsId}/resources/securityGroup
- `Response (200, body, model.SecurityGroupInfo, object) .connectionConfig.verifiedMessage`: Property added

#### `GET` /migration/ns/{nsId}/resources/securityGroup/{sgId}
- `Response (200, body, model.SecurityGroupInfo, object) .connectionConfig.verifiedMessage`: Property added

#### `POST` /migration/ns/{nsId}/resources/sshKey
- `Response (200, body, model.SshKeyInfo, object) .connectionConfig.verifiedMessage`: Property added

#### `GET` /migration/ns/{nsId}/resources/sshKey/{sshKeyId}
- `Response (200, body, model.SshKeyInfo, object) .connectionConfig.verifiedMessage`: Property added

#### `POST` /migration/ns/{nsId}/resources/vNet
- `Response (201, body, model.VNetInfo, object) .subnetInfoList[].connectionConfig.verifiedMessage`: Property added
- `Response (201, body, model.VNetInfo, object) .connectionConfig.verifiedMessage`: Property added

#### `GET` /migration/ns/{nsId}/resources/vNet/{vNetId}
- `Response (200, body, model.VNetInfo, object) .subnetInfoList[].connectionConfig.verifiedMessage`: Property added
- `Response (200, body, model.VNetInfo, object) .connectionConfig.verifiedMessage`: Property added

#### `POST` /naming/alignment
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.nodeGroups[].subnetIds`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommand`: Property removed
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommandAsync`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommands`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetK8sCluster`: Property added

#### `POST` /naming/preview
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.nodeGroups[].subnetIds`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommand`: Property removed
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommandAsync`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommands`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetK8sCluster`: Property added

#### `POST` /naming/validation
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.nodeGroups[].subnetIds`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommand`: Property removed
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommands`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetInfra.postCommandAsync`: Property added
- `Request (body, UserInfra, cloudmodel.RecommendedInfra, object) .targetK8sCluster`: Property added

#### `POST` /recommendation/infra
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.nodeGroups[].subnetIds`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommand`: Property removed
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommandAsync`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommands`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetK8sCluster`: Property added

#### `POST` /recommendation/infraWithNlb
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommand`: Property removed
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.nodeGroups[].subnetIds`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommandAsync`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommands`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetK8sCluster`: Property added

#### `POST` /recommendation/multiInfra
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.nodeGroups[].subnetIds`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommand`: Property removed
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommandAsync`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommands`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetK8sCluster`: Property added

#### `POST` /recommendation/multiInfraWithNlb
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.nodeGroups[].subnetIds`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommand`: Property removed
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommandAsync`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetInfra.postCommands`: Property added
- `Response (200, body, model.ApiResponse-array_cloudmodel_RecommendedInfra, object) .data[].targetK8sCluster`: Property added

#### `POST` /validation/ns/{nsId}/infra
- `Request (body, infraInfo, controller.ValidateInfraRequest, object) .targetInfra.nodeGroups[].subnetIds`: Property added
- `Request (body, infraInfo, controller.ValidateInfraRequest, object) .targetInfra.postCommand`: Property removed
- `Request (body, infraInfo, controller.ValidateInfraRequest, object) .targetInfra.postCommands`: Property added
- `Request (body, infraInfo, controller.ValidateInfraRequest, object) .targetInfra.postCommandAsync`: Property added
- `Request (body, infraInfo, controller.ValidateInfraRequest, object) .targetK8sCluster`: Property added

### Added APIs

#### `GET` /migration/ns/{nsId}/k8sCluster

#### `POST` /migration/ns/{nsId}/k8sCluster

#### `DELETE` /migration/ns/{nsId}/k8sCluster/{k8sClusterId}

#### `GET` /migration/ns/{nsId}/k8sCluster/{k8sClusterId}

#### `POST` /recommendation/k8sCluster

#### `POST` /recommendation/resources/k8sNodeGroupImages

#### `POST` /recommendation/resources/k8sNodeGroupSpecs

### Removed APIs

#### `POST` /recommendation/k8sControlPlane

#### `POST` /recommendation/k8sNodeGroup

