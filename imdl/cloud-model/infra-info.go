package cloudmodel

// type MigratedVmInfraModel struct {
// 	MigratedVmInfraModel MigratedVmInfraInfo `json:"migratedVmInfraModel" validate:"required"`
// }

type VmInfraInfo struct {
	InfraInfo
}

type InfraInfoList struct {
	Infra []InfraInfo `json:"infra"`
}

type IdList struct {
	IdList []string `json:"idList"`
}
