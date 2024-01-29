package database

import (
	campaing "github.com/henrique998/email-N/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (c *CampaignRepository) Create(campaing *campaing.Campaing) error {
	tx := c.Db.Create(campaing)

	return tx.Error
}

func (c *CampaignRepository) Update(campaing *campaing.Campaing) error {
	tx := c.Db.Save(campaing)

	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaing.Campaing, error) {
	var campaigns []campaing.Campaing

	tx := c.Db.Find(&campaigns)

	return campaigns, tx.Error
}

func (c *CampaignRepository) GetById(campaignId string) (*campaing.Campaing, error) {
	var campaign campaing.Campaing

	tx := c.Db.Preload("Contacts").First(&campaign, "id = ?", campaignId)

	return &campaign, tx.Error
}

func (c *CampaignRepository) Delete(campaignData *campaing.Campaing) error {
	tx := c.Db.Select("Contacts").Delete(campaignData)

	return tx.Error
}
