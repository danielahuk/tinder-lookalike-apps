package repositories

import (
	"tinder-apps/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type MemberRepository struct {
	db *sqlx.DB
}

func NewMemberRepository(db *sqlx.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) GetAllMembers() ([]models.Member, error) {
	var members []models.Member
	statement := sq.Select().
		From("members").
		PlaceholderFormat(sq.Dollar)

	statement = statement.Columns("id, name, email, gender, label, quota, status, created_at")
	sql, args, err := statement.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var item models.Member
		if err := rows.Scan(&item.ID, &item.Name, &item.Email, &item.Gender, &item.Label, &item.Quota, &item.Status, &item.CreatedAt); err != nil {
			return nil, err
		}

		members = append(members, item)
	}

	return members, nil
}

func (r *MemberRepository) GetMemberById(id int) (models.Member, error) {
	var members models.Member

	statement := sq.Select().
		From("members").
		Where(sq.Eq{
			"id": id,
		})

	statement = statement.Columns("id, name, email, gender, label, quota, status, created_at")
	sql, args, err := statement.ToSql()
	if err != nil {
		return members, err
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return members, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&members.ID, &members.Name, &members.Email, &members.Gender, &members.Label, &members.Quota, &members.Status, &members.CreatedAt); err != nil {
			return members, err
		}
	}

	return members, nil
}

func (r *MemberRepository) GetMemberByEmail(email string) (models.Member, error) {
	var members models.Member

	statement := sq.Select().
		From("members").
		Where(sq.Eq{
			"email": email,
		})

	statement = statement.Columns("id, name, email, password, gender, label, quota, status, created_at")
	sql, args, err := statement.ToSql()
	if err != nil {
		return members, err
	}

	rows, err := r.db.Query(sql, args...)
	if err != nil {
		return members, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&members.ID, &members.Name, &members.Email, &members.Password, &members.Gender, &members.Label, &members.Quota, &members.Status, &members.CreatedAt); err != nil {
			return members, err
		}
	}

	return members, nil
}

func (r *MemberRepository) CreateMember(member *models.Member) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(member.Password), bcrypt.DefaultCost)

	statement := sq.Insert("members").
		Columns("name", "email", "password", "gender", "label", "quota", "status").
		Values(member.Name, member.Email, password, member.Gender, member.Label, member.Quota, member.Status)

	sql, args, err := statement.ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	return err
}

func (r *MemberRepository) UpdateMember(member *models.Member) error {
	statement := sq.Update("members").
		Set("quota", member.Quota).
		Set("label", member.Label).
		Where(sq.Eq{
			"id": member.ID,
		})

	sql, args, err := statement.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(sql, args...)
	return err
}
