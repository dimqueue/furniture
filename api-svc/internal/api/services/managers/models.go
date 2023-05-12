package managers

import "github.com/dmytroserhiienko02/furniture/api-svc/internal/data/types"

type ProductView struct {
	Product  types.Product  `json:"product"`
	Status   types.Status   `json:"status"`
	Material types.Material `json:"material"`
}
type OrderView struct {
	ProductView ProductView `json:"productView"`
	Order       types.Order `json:"order"`
}
