package handlers

import (
	"accountsmgr/domain"
	"accountsmgr/gen/models"

	"github.com/go-openapi/swag"
)

func asErrorResponse(e *domain.Error) *models.Error {
	er := &models.Error{
		Code:    int64(e.Code),
		Message: swag.String(e.Message),
	}
	return er
}
