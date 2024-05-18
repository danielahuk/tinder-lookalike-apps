package services

import (
	"tinder-apps/internal/models"
	"tinder-apps/internal/repositories"
)

type PurchaseService struct {
	repo *repositories.PurchaseRepository
}

func NewPurchaseService(repo *repositories.PurchaseRepository) *PurchaseService {
	return &PurchaseService{repo: repo}
}

func (p *PurchaseService) GetFeaturesList() ([]models.Feature, error) {
	return p.repo.GetFeatureList()
}

func (p *PurchaseService) GetFeatureById(id int) (models.Feature, error) {
	return p.repo.GetFeatureById(id)
}

func (p *PurchaseService) CreatePurchase(purchaseRequest models.PurchaseRequest, memberId int, price float64) error {
	return p.repo.CreatePurchase(purchaseRequest, memberId, price)
}
