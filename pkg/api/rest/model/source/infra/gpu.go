package infra

import (
	"github.com/cloud-barista/cm-honeybee/gpu/drm"
	"github.com/cloud-barista/cm-honeybee/gpu/nvidia"
)

type GPU struct {
	NVIDIA []nvidia.NVIDIA `json:"nvidia"`
	DRM    []drm.DRM       `json:"drm"`
}
