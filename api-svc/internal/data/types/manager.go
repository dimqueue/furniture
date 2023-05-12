package types

type Manager struct {
	UUID         string `db:"uuid" json:"uuid"`
	Login        string `db:"login" json:"login"`
	ManagerEmail string `db:"manager_email" json:"manager_email"`
}
