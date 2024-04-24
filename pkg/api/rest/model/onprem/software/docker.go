package software

import "github.com/docker/docker/api/types"

type Docker struct {
	Containers []types.Container
	//Images     []types.ImageMetadata
}
