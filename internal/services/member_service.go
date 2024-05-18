package services

import (
	"tinder-apps/internal/models"
	"tinder-apps/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type MemberService struct {
	repo *repositories.MemberRepository
}

func NewMemberService(repo *repositories.MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) GetAllMembers() ([]models.Member, error) {
	return s.repo.GetAllMembers()
}

func (s *MemberService) GetMemberById(id int) (models.Member, error) {
	return s.repo.GetMemberById(id)
}

func (s *MemberService) GetMemberByEmailPassword(email string, password string) (models.Member, error) {
	member, err := s.repo.GetMemberByEmail(email)
	if err != nil {
		return member, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password)); err != nil {
		return member, err
	}

	return member, nil
}

func (s *MemberService) CreateMember(member *models.Member) error {
	return s.repo.CreateMember(member)
}

func (s *MemberService) UpdateMember(member *models.Member) error {
	return s.repo.UpdateMember(member)
}
