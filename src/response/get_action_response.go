package response

import (
	"time"

	"github.com/google/uuid"
	"github.com/kozmoai/builder-backend/src/model"
	"github.com/kozmoai/builder-backend/src/utils/idconvertor"
	"github.com/kozmoai/builder-backend/src/utils/resourcelist"
)

type GetActionResponse struct {
	ActionID          string                 `json:"actionID"`
	UID               uuid.UUID              `json:"uid"`
	TeamID            string                 `json:"teamID"`
	AppID             string                 `json:"appID"`
	Version           int                    `json:"version"`
	ResourceID        string                 `json:"resourceID,omitempty"`
	DisplayName       string                 `json:"displayName"`
	ActionType        string                 `json:"actionType"`
	IsVirtualResource bool                   `json:"isVirtualResource"`
	Content           map[string]interface{} `json:"content"`
	Transformer       map[string]interface{} `json:"transformer"`
	TriggerMode       string                 `json:"triggerMode"`
	Config            map[string]interface{} `json:"config"`
	CreatedAt         time.Time              `json:"createdAt,omitempty"`
	CreatedBy         string                 `json:"createdBy,omitempty"`
	UpdatedAt         time.Time              `json:"updatedAt,omitempty"`
	UpdatedBy         string                 `json:"updatedBy,omitempty"`
}

func NewGetActionResponse(action *model.Action) *GetActionResponse {
	actionConfig := action.ExportConfig()
	resp := &GetActionResponse{
		ActionID:          idconvertor.ConvertIntToString(action.ID),
		UID:               action.UID,
		TeamID:            idconvertor.ConvertIntToString(action.TeamID),
		AppID:             idconvertor.ConvertIntToString(action.AppRefID),
		Version:           action.Version,
		ResourceID:        idconvertor.ConvertIntToString(action.ResourceRefID),
		DisplayName:       action.Name,
		ActionType:        resourcelist.GetResourceIDMappedType(action.Type),
		IsVirtualResource: actionConfig.IsVirtualResource,
		Content:           action.ExportTemplateInMap(),
		Transformer:       action.ExportTransformerInMap(),
		TriggerMode:       action.TriggerMode,
		Config:            action.ExportConfigInMap(),
		CreatedAt:         action.CreatedAt,
		CreatedBy:         idconvertor.ConvertIntToString(action.CreatedBy),
		UpdatedAt:         action.UpdatedAt,
		UpdatedBy:         idconvertor.ConvertIntToString(action.UpdatedBy),
	}
	return resp
}

func (resp *GetActionResponse) ExportForFeedback() interface{} {
	return resp
}

func (req *GetActionResponse) AppendVirtualResourceToTemplate(value interface{}) {
	req.Content[model.ACTION_CONFIG_FIELD_VIRTUAL_RESOURCE] = value
}
