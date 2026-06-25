package provider

import (
	"QuotesService/internal/config"
	"QuotesService/internal/models"
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func New(cfg *config.Config) *Postgres {
	return &Postgres{
		db: GetConnection(cfg),
	}
}

func GetConnection(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName, cfg.DbSSL)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	return db
}

func (p *Postgres) AutoMigrate() error {
	return p.db.AutoMigrate(&models.Quote{})
}

func (p *Postgres) GetRandomQuote() (*models.Quote, error) {
	var quote models.Quote
	if err := p.db.Order("RANDOM()").First(&quote).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no quotes found")
		}
		return nil, fmt.Errorf("error retrieving quote: %v", err)
	}
	return &quote, nil
}

func (p *Postgres) GetAllQuotes() ([]models.Quote, error) {
	var quotes []models.Quote
	if err := p.db.Find(&quotes).Error; err != nil {
		return nil, fmt.Errorf("error retrieving quotes: %v", err)
	}
	return quotes, nil
}

func (p *Postgres) GetQuoteByID(id uint) (*models.Quote, error) {
	var quote models.Quote
	if err := p.db.First(&quote, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("quote with ID %d not found", id)
		}
		return nil, fmt.Errorf("error retrieving quote: %v", err)
	}
	return &quote, nil
}

func (p *Postgres) CreateQuote(quote *models.Quote) error {
	if err := p.db.Create(quote).Error; err != nil {
		return fmt.Errorf("error creating quote: %v", err)
	}
	return nil
}

func (p *Postgres) CountQuotes() (int64, error) {
	var count int64
	if err := p.db.Model(&models.Quote{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("error counting quotes: %v", err)
	}
	return count, nil
}
