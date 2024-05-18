package models

type Member struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	Gender    string `json:"gender" db:"gender"`
	Label     string `json:"label" db:"label"`
	Quota     int    `json:"quota" db:"quota"`
	Status    int    `json:"status" db:"status"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type MemberLogin struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
	Token string `json:"token" db:"token"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
