package controllers

import (
	"net/http"
	"QuotesService/internal/api/types"
	"QuotesService/internal/provider"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Provider provider.Provider
}

func New(provider provider.Provider) *Controller {
	return &Controller{
		Provider: provider,
	}
}

// GetRoot godoc
//
//	@Summary		Get a random quote
//	@Description	Returns a random weather, nature, or science fact from the database
//	@Tags			quotes
//	@Produce		json
//	@Success		200	{object}	types.APIResponse
//	@Failure		500	{object}	types.APIResponse
//	@Router			/ [get]
func (c *Controller) GetRoot(ctx *gin.Context) {
	quote, err := c.Provider.GetRandomQuote()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, types.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APIResponse{
		Status: http.StatusOK,
		Data:   quote,
	})
}
