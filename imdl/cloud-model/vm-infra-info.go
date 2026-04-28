package cloudmodel

// type MigratedVmInfraModel struct {
// 	MigratedVmInfraModel MigratedVmInfraInfo `json:"migratedVmInfraModel" validate:"required"`
// }

type VmInfraInfo struct {
	MciInfo
}

type MciInfoList struct {
	Mci []MciInfo `json:"mci"`
}

type IdList struct {
	IdList []string `json:"idList"`
}
