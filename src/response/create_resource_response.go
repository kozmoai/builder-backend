package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/kozmoai/builder-backend/src/model"
	"github.com/kozmoai/builder-backend/src/utils/idconvertor"
	"github.com/kozmoai/builder-backend/src/utils/resourcelist"
)

type CreateResourceResponse struct {
	ID        string                 `json:"resourceID"`
	UID       uuid.UUID              `json:"uid"`
	TeamID    string                 `json:"teamID"`
	Name      string                 `json:"resourceName" validate:"required"`
	Type      string                 `json:"resourceType" validate:"required"`
	Options   map[string]interface{} `json:"content" validate:"required"`
	CreatedAt time.Time              `json:"createdAt,omitempty"`
	CreatedBy string                 `json:"createdBy,omitempty"`
	UpdatedAt time.Time              `json:"updatedAt,omitempty"`
	UpdatedBy string                 `json:"updatedBy,omitempty"`
}

func NewCreateResourceResponse(resource *model.Resource) *CreateResourceResponse {
	return &CreateResourceResponse{
		ID:        idconvertor.ConvertIntToString(resource.ID),
		UID:       resource.UID,
		TeamID:    idconvertor.ConvertIntToString(resource.TeamID),
		Name:      resource.Name,
		Type:      resourcelist.GetResourceIDMappedType(resource.Type),
		Options:   resource.ExportOptionsInMap(),
		CreatedAt: resource.CreatedAt,
		CreatedBy: idconvertor.ConvertIntToString(resource.CreatedBy),
		UpdatedAt: resource.UpdatedAt,
		UpdatedBy: idconvertor.ConvertIntToString(resource.UpdatedBy),
	}
}

func (resp *CreateResourceResponse) ExportForFeedback() interface{} {
	return resp
}
