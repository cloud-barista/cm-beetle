package software

import "github.com/docker/docker/api/types"

type Podman struct {
	Containers []types.Container
	//Images     []types.ImageMetadata
}
