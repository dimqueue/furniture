package types

type CreateProduct struct {
	Title      string `db:"title" json:"title"`
	Price      int64  `db:"price" json:"price"`
	MaterialId int64  `db:"material_id" json:"material_id"`
}

type Product struct {
	Id       int64    `db:"id" json:"id"`
	StatusID StatusID `db:"status_id" json:"status_id"`
	CreateProduct
}

type UpdateProduct struct {
	Title      *string  `db:"title" json:"title"`
	Price      *int64   `db:"price" json:"price"`
	MaterialId *int64   `db:"material_id" json:"material_id"`
	Id         int64    `db:"-" json:"id"`
	StatusID   StatusID `db:"-" json:"status_id"`
}

func (u *UpdateProduct) SetData() map[string]interface{} {
	updated := make(map[string]interface{})

	updated["status_id"] = u.StatusID
	//updated["id"] = u.Id

	if u.Title != nil {
		updated["title"] = u.Title
	}
	if u.Price != nil {
		updated["price"] = u.Price
	}
	if u.MaterialId != nil {
		updated["material_id"] = u.MaterialId
	}

	return updated
}
