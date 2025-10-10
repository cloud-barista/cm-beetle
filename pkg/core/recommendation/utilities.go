package recommendation

import (
	cloudmodel "github.com/cloud-barista/cm-model/infra/cloud-model"
)

// transposeMatrix transposes a 2D matrix
func transposeMatrix[T any](matrix [][]T) [][]T {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]T{}
	}

	rows := len(matrix)
	cols := len(matrix[0])

	result := make([][]T, cols)
	for i := range result {
		result[i] = make([]T, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// Function to check overall status for the entire list of VMs
func checkOverallSubGroupStatus(subgroups []cloudmodel.CreateSubGroupDynamicReq) string {
	allOk := true
	allNone := true

	for _, subgroup := range subgroups {
		if subgroup.ImageId == "" || subgroup.SpecId == "" {
			allOk = false // At least one VM is not fully populated
		}
		if subgroup.ImageId != "" || subgroup.SpecId != "" {
			allNone = false // At least one VM has a value
		}
	}

	// Determine overall status
	if allNone {
		return string(NothingRecommended)
	} else if allOk {
		return string(FullyRecommended)
	} else {
		return string(PartiallyRecommended)
	}
}
