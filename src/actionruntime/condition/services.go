// Copyright 2024 Kozmoai Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package condition

import (
	"errors"
	"fmt"

	"github.com/kozmoai/builder-backend/src/actionruntime/common"
)

type ConditionConnector struct {
	Action ConditionTemplate
}

// server side transformer have no validate resource options method
func (r *ConditionConnector) ValidateResourceOptions(resourceOptions map[string]interface{}) (common.ValidateResult, error) {
	return common.ValidateResult{Valid: true}, nil
}

func (r *ConditionConnector) ValidateActionTemplate(actionOptions map[string]interface{}) (common.ValidateResult, error) {
	fmt.Printf("[DUMP] actionOptions: %+v \n", actionOptions)
	// @todo: check action needed field
	return common.ValidateResult{Valid: true}, nil
}

// server side transformer have no test connection method
func (r *ConditionConnector) TestConnection(resourceOptions map[string]interface{}) (common.ConnectionResult, error) {
	return common.ConnectionResult{Success: false}, errors.New("unsupported type: server side transformer")
}

// server side transformer have no meta info
func (r *ConditionConnector) GetMetaInfo(resourceOptions map[string]interface{}) (common.MetaInfoResult, error) {
	return common.MetaInfoResult{Success: false}, errors.New("unsupported type: server side transformer")
}

func (r *ConditionConnector) Run(resourceOptions map[string]interface{}, actionOptions map[string]interface{}, rawActionOptions map[string]interface{}) (common.RuntimeResult, error) {
	res := common.RuntimeResult{
		Success: false,
		Rows:    []map[string]interface{}{},
		Extra:   map[string]interface{}{},
	}

	fmt.Printf("[DUMP] ConditionConnector.Run() actionOptions: %+v\n", actionOptions)

	fmt.Printf("[DUMP] res: %+v\n", res)
	return res, nil
}
