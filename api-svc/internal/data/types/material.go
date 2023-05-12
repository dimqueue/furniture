package types

type Material struct {
	Id    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}
