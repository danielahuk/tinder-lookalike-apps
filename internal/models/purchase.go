package models

type Feature struct {
	ID        int     `json:"id" db:"id"`
	Name      string  `json:"name" db:"name"`
	Price     float64 `json:"price" db:"price"`
	Value     string  `json:"value" db:"value"`
	Status    int     `json:"status" db:"status"`
	CreatedAt string  `json:"created_at" db:"created_at"`
}

type PurchaseRequest struct {
	PackageID int `json:"package_id"`
	Qty       int `json:"qty"`
}

type Purchase struct {
	ID             int    `json:"id" db:"id"`
	Date           string `json:"date" db:"date"`
	Status         int    `json:"status" db:"status"`
	CreatedAt      string `json:"created_at" db:"created_at"`
	PurchaseDetail PurchaseDetail
}

type PurchaseDetail struct {
	TransactionID int     `json:"transaction_id" db:"transaction_id"`
	FeatureID     int     `json:"feature_id" db:"feature_id"`
	Qty           int     `json:"qty" db:"qty"`
	Price         float64 `json:"price" db:"price"`
}
