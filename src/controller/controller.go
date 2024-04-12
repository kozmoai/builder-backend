package controller

import (
	"github.com/kozmoai/builder-backend/src/cache"
	"github.com/kozmoai/builder-backend/src/drive"
	"github.com/kozmoai/builder-backend/src/storage"
	"github.com/kozmoai/builder-backend/src/utils/accesscontrol"
	"github.com/kozmoai/builder-backend/src/utils/tokenvalidator"
)

type Controller struct {
	Storage               *storage.Storage
	Cache                 *cache.Cache
	Drive                 *drive.Drive
	RequestTokenValidator *tokenvalidator.RequestTokenValidator
	AttributeGroup        *accesscontrol.AttributeGroup
}

func NewControllerForBackend(storage *storage.Storage, cache *cache.Cache, drive *drive.Drive, validator *tokenvalidator.RequestTokenValidator, attrg *accesscontrol.AttributeGroup) *Controller {
	return &Controller{
		Storage:               storage,
		Cache:                 cache,
		Drive:                 drive,
		RequestTokenValidator: validator,
		AttributeGroup:        attrg,
	}
}

func NewControllerForBackendInternal(storage *storage.Storage, drive *drive.Drive, validator *tokenvalidator.RequestTokenValidator, attrg *accesscontrol.AttributeGroup) *Controller {
	return &Controller{
		Storage:               storage,
		Drive:                 drive,
		RequestTokenValidator: validator,
		AttributeGroup:        attrg,
	}
}
