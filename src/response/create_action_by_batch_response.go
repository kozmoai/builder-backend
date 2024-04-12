package response

import (
	"github.com/kozmoai/builder-backend/src/model"
)

type CreateActionByBatchResponse struct {
	Actions []*CreateActionResponse `json:"actions"`
}

func NewCreateActionByBatchResponse(actions []*model.Action) *CreateActionByBatchResponse {
	resp := make([]*CreateActionResponse, 0)
	for _, action := range actions {
		resp = append(resp, NewCreateActionResponse(action))
	}
	return &CreateActionByBatchResponse{Actions: resp}
}

func (resp *CreateActionByBatchResponse) ExportForFeedback() interface{} {
	return resp
}
