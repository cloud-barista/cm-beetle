package infra

import (
	"github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/network"
)

type Infra struct {
	Compute Compute         `json:"compute"`
	Network network.Network `json:"network"`
	GPU     GPU             `json:"gpu"`
	Storage Storage         `json:"storage"`
}
