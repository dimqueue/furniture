package data

import "github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"

type IMaterialRepository interface {
	GetMaterialById(int64) (types.Material, error)
}
