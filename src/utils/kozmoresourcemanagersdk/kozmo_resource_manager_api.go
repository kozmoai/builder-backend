package kozmoresourcemanagersdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/kozmoai/builder-backend/src/utils/config"
	"github.com/kozmoai/builder-backend/src/utils/resourcelist"
	"github.com/kozmoai/builder-backend/src/utils/tokenvalidator"
)

const (
	BASEURL = "http://127.0.0.1:8008/api/v1"
	// api route part
	GET_AI_AGENT_INTERNAL_API                    = "/api/v1/aiAgent/%d"
	RUN_AI_AGENT_INTERNAL_API                    = "/api/v1/aiAgent/%d/run"
	DELETE_TEAM_ALL_AI_AGENT_INTERNAL_API        = "/api/v1/teams/%d/aiAgent/all"
	FORK_MARKETPLACE_AI_AGENT_INTERNAL_API       = "/api/v1/aiAgent/%d/forkTo/teams/%d/by/users/%d"
	PUBLISH_AI_AGENT_TO_MARKETPLACE_INTERNAL_API = "/api/v1/teams/%d/aiAgent/%d"
)

const (
	PRODUCT_TYPE_AIAGENTS = "aiAgents"
	PRODUCT_TYPE_APPS     = "apps"
	PRODUCT_TYPE_HUBS     = "hubs"
)

type kozmoResourceManagerRestAPI struct {
	Config *config.Config
	Debug  bool `json:"-"`
}

func NewkozmoResourceManagerRestAPI() (*kozmoResourceManagerRestAPI, error) {
	return &kozmoResourceManagerRestAPI{
		Config: config.GetInstance(),
	}, nil
}

func (r *kozmoResourceManagerRestAPI) CloseDebug() {
	r.Debug = false
}

func (r *kozmoResourceManagerRestAPI) OpenDebug() {
	r.Debug = true
}

func (r *kozmoResourceManagerRestAPI) GetResource(resourceType int, resourceID int) (map[string]interface{}, error) {
	// self-hist need skip this method.
	if !r.Config.IsCloudMode() {
		return nil, nil
	}
	switch resourceType {
	case resourcelist.TYPE_AI_AGENT_ID:
		return r.GetAIAgent(resourceID)
	default:
		return nil, errors.New("Invalied resource type: " + resourcelist.GetResourceIDMappedType(resourceType))
	}
}

func (r *kozmoResourceManagerRestAPI) RunResource(resourceType int, resourceID int, req map[string]interface{}) (*RunResourceResult, error) {
	// self-hist need skip this method.
	if !r.Config.IsCloudMode() {
		return nil, nil
	}
	switch resourceType {
	case resourcelist.TYPE_AI_AGENT_ID:
		return r.RunAIAgent(req)
	default:
		return nil, errors.New("Invalied resource type: " + resourcelist.GetResourceIDMappedType(resourceType))
	}
}

func (r *kozmoResourceManagerRestAPI) GetAIAgent(aiAgentID int) (map[string]interface{}, error) {
	// self-hist need skip this method.
	if !r.Config.IsCloudMode() {
		return nil, nil
	}
	client := resty.New()
	tokenValidator := tokenvalidator.NewRequestTokenValidator()
	uri := r.Config.GetkozmoResourceManagerInternalRestAPI() + fmt.Sprintf(GET_AI_AGENT_INTERNAL_API, aiAgentID)
	resp, errInPost := client.R().
		SetHeader("Request-Token", tokenValidator.GenerateValidateToken(strconv.Itoa(aiAgentID))).
		Get(uri)
	if r.Debug {
		log.Printf("[kozmoResourceManagerRestAPI.GetAiAgent()]  uri: %+v \n", uri)
		log.Printf("[kozmoResourceManagerRestAPI.GetAiAgent()]  response: %+v, err: %+v \n", resp, errInPost)
		log.Printf("[kozmoResourceManagerRestAPI.GetAiAgent()]  resp.StatusCode(): %+v \n", resp.StatusCode())
	}
	if errInPost != nil {
		return nil, errInPost
	}
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(resp.String())
	}

	var aiAgent map[string]interface{}
	errInUnMarshal := json.Unmarshal([]byte(resp.String()), &aiAgent)
	if errInUnMarshal != nil {
		return nil, errInUnMarshal
	}
	return aiAgent, nil
}

func (r *kozmoResourceManagerRestAPI) RunAIAgent(req map[string]interface{}) (*RunResourceResult, error) {
	// self-hist need skip this method.
	if !r.Config.IsCloudMode() {
		return nil, nil
	}
	reqInstance, errInNewReq := NewRunAIAgentRequest(req)
	if errInNewReq != nil {
		return nil, errInNewReq
	}
	client := resty.New()
	uri := r.Config.GetkozmoResourceManagerInternalRestAPI() + fmt.Sprintf(RUN_AI_AGENT_INTERNAL_API, reqInstance.ExportAIAgentIDInInt())
	if r.Debug {
		log.Printf("[kozmoResourceManagerRestAPI.RunAiAgent()]  uri: %+v \n", uri)
	}
	requestToken := reqInstance.ExportRequestToken()
	log.Printf("[kozmoResourceManagerRestAPI.RunAiAgent()]  uri: %+v \n", uri)
	log.Printf("[reqInstance]  reqInstance: %+v \n", reqInstance)
	fmt.Printf("[requestToken] %+v\n", requestToken)
	resp, errInPost := client.R().
		SetHeader("Request-Token", requestToken).
		SetHeader("Authorization", reqInstance.ExportAuthorization()).
		SetBody(req).
		Post(uri)
	if r.Debug {
		log.Printf("[kozmoResourceManagerRestAPI.RunAiAgent()]  response: %+v, err: %+v \n", resp, errInPost)
		log.Printf("[kozmoResourceManagerRestAPI.RunAiAgent()]  resp.StatusCode(): %+v \n", resp.StatusCode())
	}
	if errInPost != nil {
		return nil, errInPost
	}
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(resp.String())
	}

	runResourceResult := NewRunResourceResult()
	errInUnMarshal := json.Unmarshal([]byte(resp.String()), &runResourceResult)
	if errInUnMarshal != nil {
		return nil, errInUnMarshal
	}
	return runResourceResult, nil
}

func (r *kozmoResourceManagerRestAPI) DeleteTeamAllAIAgent(teamID int) error {
	// self-hist need skip this method.
	if !r.Config.IsCloudMode() {
		return nil
	}
	client := resty.New()
	tokenValidator := tokenvalidator.NewRequestTokenValidator()
	uri := r.Config.GetkozmoResourceManagerInternalRestAPI() + fmt.Sprintf(DELETE_TEAM_ALL_AI_AGENT_INTERNAL_API, teamID)
	resp, errInDelete := client.R().
		SetHeader("Request-Token", tokenValidator.GenerateValidateToken(strconv.Itoa(teamID))).
		Delete(uri)
	if r.Debug {
		log.Printf("[kozmoResourceManagerRestAPI.DeleteTeamAllAiAgent()]  uri: %+v \n", uri)
		log.Printf("[kozmoResourceManagerRestAPI.DeleteTeamAllAiAgent()]  response: %+v, err: %+v \n", resp, errInDelete)
		log.Printf("[kozmoResourceManagerRestAPI.DeleteTeamAllAiAgent()]  resp.StatusCode(): %+v \n", resp.StatusCode())
	}
	if errInDelete != nil {
		return errInDelete
	}
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return errors.New(resp.String())
	}

	return nil
}

func (r *kozmoResourceManagerRestAPI) ForkMarketplaceAIAgent(aiAgentID int, toTeamID int, userID int) (*AIAgentForExport, error) {
	// self-hist need skip this method.
	if !r.Config.IsCloudMode() {
		return nil, nil
	}
	client := resty.New()
	tokenValidator := tokenvalidator.NewRequestTokenValidator()
	uri := r.Config.GetkozmoResourceManagerInternalRestAPI() + fmt.Sprintf(FORK_MARKETPLACE_AI_AGENT_INTERNAL_API, aiAgentID, toTeamID, userID)
	resp, errInPost := client.R().
		SetHeader("Request-Token", tokenValidator.GenerateValidateToken(strconv.Itoa(aiAgentID), strconv.Itoa(toTeamID), strconv.Itoa(userID))).
		Post(uri)
	if r.Debug {
		log.Printf("[kozmoResourceManagerRestAPI.ForkMarketplaceAIAgent()]  uri: %+v \n", uri)
		log.Printf("[kozmoResourceManagerRestAPI.ForkMarketplaceAIAgent()]  response: %+v, err: %+v \n", resp, errInPost)
		log.Printf("[kozmoResourceManagerRestAPI.ForkMarketplaceAIAgent()]  resp.StatusCode(): %+v \n", resp.StatusCode())
	}
	if errInPost != nil {
		return nil, errInPost
	}
	if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusCreated {
		return nil, errors.New(resp.String())
	}

	aiAgent := &AIAgentForExport{}
	errInUnMarshal := json.Unmarshal([]byte(resp.String()), &aiAgent)
	if errInUnMarshal != nil {
		return nil, errInUnMarshal
	}
	return aiAgent, nil
}
