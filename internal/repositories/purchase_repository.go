package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"

	"tinder-apps/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type PurchaseRepository struct {
	db *sqlx.DB
}

func NewPurchaseRepository(db *sqlx.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (p *PurchaseRepository) GetFeatureList() ([]models.Feature, error) {
	var featureList []models.Feature

	statement := sq.Select().
		From("features")

	statement = statement.Columns("id, name, price, value, status, created_at")
	sql, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var featureItem models.Feature
		if err := rows.Scan(&featureItem.ID, &featureItem.Name, &featureItem.Price, &featureItem.Value, &featureItem.Status, &featureItem.CreatedAt); err != nil {
			return nil, err
		}

		featureList = append(featureList, featureItem)
	}

	return featureList, nil
}

func (p *PurchaseRepository) GetFeatureById(id int) (models.Feature, error) {
	var feature models.Feature

	statement := sq.Select().
		From("features").
		Where(sq.Eq{
			"id": id,
		})

	statement = statement.Columns("id, name, price, value, status, created_at")
	sql, args, err := statement.ToSql()
	if err != nil {
		return feature, err
	}

	rows, err := p.db.Query(sql, args...)
	if err != nil {
		return feature, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&feature.ID, &feature.Name, &feature.Price, &feature.Value, &feature.Status, &feature.CreatedAt); err != nil {
			return feature, err
		}
	}

	return feature, nil
}

func (p *PurchaseRepository) CreatePurchase(purchaseRequest models.PurchaseRequest, memberId int, price float64) error {
	today := time.Now()
	currentDate := today.Format("2006-01-02")

	statement := sq.Insert("transactions").
		Columns("date", "status", "member_id").
		Values(currentDate, 1, memberId)

	sql, args, err := statement.ToSql()
	if err != nil {
		return err
	}

	result, err := p.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	statementDetail := sq.Insert("transaction_details").
		Columns("transaction_id", "feature_id", "qty", "price", "status").
		Values(lastInsertId, purchaseRequest.PackageID, purchaseRequest.Qty, price, 1)

	sqlDetail, argsDetail, err := statementDetail.ToSql()
	if err != nil {
		return err
	}

	_, err = p.db.Exec(sqlDetail, argsDetail...)

	return err
}
