package services

import (
	"database/sql"
	"errors"
	"tinder-apps/internal/models"
	"tinder-apps/internal/repositories"
)

type PartnerService struct {
	repo *repositories.PartnerRepository
}

func NewPartnerService(repo *repositories.PartnerRepository) *PartnerService {
	return &PartnerService{repo: repo}
}

func (s *PartnerService) GetPartnerList(userID int) (models.PartnerList, error) {
	partnerList, err := s.repo.GetPartnerList(userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			partnerList, err = s.repo.GeneratePartner(userID)

			if err != nil {
				return partnerList, err
			}
		}
	}

	return partnerList, err
}

func (s *PartnerService) GetPartnerCount(userID int) (int, error) {
	return s.repo.GetPartnerCount(userID)
}

func (s *PartnerService) GetPartnerCheck(userID int, targetID int) (int, error) {
	return s.repo.GetPartnerCheck(userID, targetID)
}

func (s *PartnerService) UpdatePartnership(sourceID int, memberTarget *models.Member, direction string) error {
	return s.repo.UpdatePartnership(sourceID, memberTarget, direction)
}
