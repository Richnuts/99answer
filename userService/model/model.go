package model

type User struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt int    `json:"createdAt" db:"created_at"`
	UpdatedAt int    `json:"updatedAt" db:"updated_at"`
}

type UserFilter struct {
	ID *string
}

type Pagination struct {
	Page    int
	PerPage int
}
