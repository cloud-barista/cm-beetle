package migration

import tbmodel "github.com/cloud-barista/cb-tumblebug/src/core/model"

type MciInfoList struct {
	Mci []tbmodel.TbMciInfo `json:"mci"`
}

type IdList struct {
	IdList []string `json:"idList"`
}
