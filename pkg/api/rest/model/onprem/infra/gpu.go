package infra

type NVIDIADeviceAttribute struct {
	GPUUUID             string `json:"gpu_uuid"`
	DriverVersion       string `json:"driver_version"`
	CUDAVersion         string `json:"cuda_version"`
	ProductName         string `json:"product_name"`
	ProductBrand        string `json:"product_brand"`
	ProductArchitecture string `json:"product_architecture"`
}

type NVIDIAPerformance struct {
	GPUUsage        uint32 `json:"gpu_usage"`         // percent
	FBMemoryUsed    uint64 `json:"fb_memory_used"`    // mb
	FBMemoryTotal   uint64 `json:"fb_memory_total"`   // mb
	FBMemoryUsage   uint32 `json:"fb_memory_usage"`   // percent
	Bar1MemoryUsed  uint64 `json:"bar1_memory_used"`  // mb
	Bar1MemoryTotal uint64 `json:"bar1_memory_total"` // mb
	Bar1MemoryUsage uint32 `json:"bar1_memory_usage"` // percent
}

type NVIDIA struct {
	DeviceAttribute NVIDIADeviceAttribute `json:"device_attribute"`
	Performance     NVIDIAPerformance     `json:"performance"`
}

type DRM struct {
	DriverName        string `json:"driver_name"`
	DriverVersion     string `json:"driver_version"`
	DriverDate        string `json:"driver_date"`
	DriverDescription string `json:"driver_description"`
}

type GPU struct {
	NVIDIA []NVIDIA `json:"nvidia"`
	DRM    []DRM    `json:"drm"`
}
