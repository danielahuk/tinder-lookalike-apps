package repositories

import (
	"time"
	"tinder-apps/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type PartnerRepository struct {
	db *sqlx.DB
}

func NewPartnerRepository(db *sqlx.DB) *PartnerRepository {
	return &PartnerRepository{db: db}
}

func (r *PartnerRepository) GetPartnerList(userID int) (models.PartnerList, error) {
	var partner models.PartnerList

	today := time.Now()
	currentDate := today.Format("2006-01-02")

	statement := sq.Select().
		From("partnership p").
		Join("members m ON p.member_id2 = m.id").
		Where(sq.Eq{
			"p.member_id1": userID,
			"p.created_at": currentDate,
		}).
		OrderBy("RAND()")

	statement = statement.Columns("p.member_id1, p.member_id2, m.name AS 'member_target_name', p.status, p.created_at")
	sql, args, err := statement.ToSql()
	if err != nil {
		return partner, err
	}

	var count int
	err = r.db.QueryRow(sql, args...).Scan(&count)

	if count == 0 {
		return partner, err
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return partner, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&partner.MemberSourceID, &partner.MemberTargetID, &partner.MemberTargetName, &partner.Status, &partner.CreatedAt); err != nil {
			return partner, err
		}
	}

	return partner, nil
}

func (r *PartnerRepository) GetPartnerCount(userID int) (int, error) {
	today := time.Now()
	currentDate := today.Format("2006-01-02")

	statement := sq.Select().
		From("partnership").
		Where(sq.Eq{
			"member_id1": userID,
			"created_at": currentDate,
		})

	statement = statement.Columns("member_id1")
	sql, args, err := statement.ToSql()
	if err != nil {
		return 0, err
	}

	var count int
	err = r.db.QueryRow(sql, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PartnerRepository) GetPartnerCheck(userID int, targetID int) (int, error) {
	today := time.Now()
	currentDate := today.Format("2006-01-02")

	statement := sq.Select().
		From("partnership").
		Where(sq.Eq{
			"member_id1":           userID,
			"member_id2":           targetID,
			"LEFT(created_at, 10)": currentDate,
		})

	statement = statement.Columns("member_id1")
	sql, args, err := statement.ToSql()
	if err != nil {
		return 0, err
	}

	var count int
	err = r.db.QueryRow(sql, args...).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *PartnerRepository) GeneratePartner(userID int) (models.PartnerList, error) {
	var partnerList models.PartnerList

	statement := sq.Select().
		From("members").
		Where(sq.Eq{
			"id": userID,
		})

	statement = statement.Columns("id, name, gender, status, quota")
	sql, args, err := statement.ToSql()
	if err != nil {
		return partnerList, err
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return partnerList, err
	}

	defer rows.Close()

	var partner models.Member
	for rows.Next() {
		if err := rows.Scan(&partner.ID, &partner.Name, &partner.Gender, &partner.Status, &partner.Quota); err != nil {
			return partnerList, err
		}
	}

	newStatement := sq.Select().
		From("members m").
		LeftJoin("partnership p ON m.id = p.member_id2").
		Where(sq.NotEq{
			"m.id":     userID,
			"m.gender": partner.Gender,
		}).
		Where(sq.Eq{
			"m.status": 1,
		}).
		Where(sq.Expr("p.member_id2 IS NULL")).
		OrderBy("RAND()").
		Limit(1)

	newStatement = newStatement.Columns("m.id")
	newSql, newArgs, newErr := newStatement.ToSql()
	if newErr != nil {
		return partnerList, newErr
	}

	var count int
	err = r.db.QueryRow(newSql, newArgs...).Scan(&count)

	if count == 0 {
		return partnerList, err
	}

	newStatement = newStatement.Columns("m.name, m.gender, m.status")
	newSql, newArgs, newErr = newStatement.ToSql()
	if newErr != nil {
		return partnerList, newErr
	}

	newRows, newErr := r.db.Query(newSql, newArgs...)
	if newErr != nil {
		return partnerList, newErr
	}

	defer newRows.Close()

	var newPartner models.Member
	for newRows.Next() {
		if err := newRows.Scan(&newPartner.ID, &newPartner.Name, &newPartner.Gender, &newPartner.Status); err != nil {
			return partnerList, err
		}
	}

	return models.PartnerList{
		MemberSourceID:   userID,
		MemberTargetID:   newPartner.ID,
		MemberTargetName: newPartner.Name,
		Status:           0,
	}, nil
}

func (r *PartnerRepository) UpdatePartnership(sourceID int, member *models.Member, direction string) error {
	status := 1

	if direction == "left" {
		status = 2
	}

	statement := sq.Insert("partnership").
		Columns("member_id1", "member_id2", "status").
		Values(sourceID, member.ID, status)

	sql, args, err := statement.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	return err
}
