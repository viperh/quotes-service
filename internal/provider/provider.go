package provider

import "QuotesService/internal/models"

type Provider interface {
	GetRandomQuote() (*models.Quote, error)
	GetAllQuotes() ([]models.Quote, error)
	GetQuoteByID(id uint) (*models.Quote, error)
	CreateQuote(quote *models.Quote) error
	CountQuotes() (int64, error)
	AutoMigrate() error
}
