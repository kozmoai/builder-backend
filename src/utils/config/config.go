package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/caarlos0/env"
)

const DEPLOY_MODE_SELF_HOST = "self-host"
const DEPLOY_MODE_CLOUD = "cloud"
const DEPLOY_MODE_CLOUD_TEST = "cloud-test"
const DEPLOY_MODE_CLOUD_BETA = "cloud-beta"
const DEPLOY_MODE_CLOUD_PRODUCTION = "cloud-production"
const DRIVE_TYPE_AWS = "aws"
const DRIVE_TYPE_DO = "do"
const DRIVE_TYPE_MINIO = "minio"
const PROTOCOL_WEBSOCKET = "ws"
const PROTOCOL_WEBSOCKET_OVER_TLS = "wss"

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		var err error
		if instance == nil {
			instance, err = getConfig() // not thread safe
			if err != nil {
				panic(err)
			}
		}
	})
	return instance
}

type Config struct {
	// server config
	ServerHost         string `env:"KOZMO_SERVER_HOST" envDefault:"0.0.0.0"`
	ServerPort         string `env:"KOZMO_SERVER_PORT" envDefault:"8001"`
	InternalServerPort string `env:"KOZMO_SERVER_INTERNAL_PORT" envDefault:"9005"`
	ServerMode         string `env:"KOZMO_SERVER_MODE" envDefault:"debug"`
	DeployMode         string `env:"KOZMO_DEPLOY_MODE" envDefault:"self-host"`
	SecretKey          string `env:"KOZMO_SECRET_KEY" envDefault:"8xEMrWkBARcDDYQ"`

	// websocket config
	WebsocketServerHost                       string `env:"KOZMO_WEBSOCKET_SERVER_HOST" envDefault:"0.0.0.0"`
	WebsocketServerPort                       string `env:"KOZMO_WEBSOCKET_SERVER_PORT" envDefault:"8002"`
	WebsocketServerConnectionHost             string `env:"KOZMO_WEBSOCKET_CONNECTION_HOST" envDefault:"0.0.0.0"`
	WebsocketServerConnectionPort             string `env:"KOZMO_WEBSOCKET_CONNECTION_PORT" envDefault:"80"`
	WebsocketServerConnectionHostSouthAsia    string `env:"KOZMO_WEBSOCKET_CONNECTION_HOST_SOUTH_ASIA" envDefault:"0.0.0.0"`
	WebsocketServerConnectionPortSouthAsia    string `env:"KOZMO_WEBSOCKET_CONNECTION_PORT_SOUTH_ASIA" envDefault:"80"`
	WebsocketServerConnectionHostEastAsia     string `env:"KOZMO_WEBSOCKET_CONNECTION_HOST_EAST_ASIA" envDefault:"0.0.0.0"`
	WebsocketServerConnectionPortEastAsia     string `env:"KOZMO_WEBSOCKET_CONNECTION_PORT_EAST_ASIA" envDefault:"80"`
	WebsocketServerConnectionHostCenterEurope string `env:"KOZMO_WEBSOCKET_CONNECTION_HOST_CENTER_EUROPE" envDefault:"0.0.0.0"`
	WebsocketServerConnectionPortCenterEurope string `env:"KOZMO_WEBSOCKET_CONNECTION_PORT_CENTER_EUROPE" envDefault:"80"`
	WSSEnabled                                string `env:"KOZMO_WSS_ENABLED" envDefault:"false"`

	// key for idconvertor
	RandomKey string `env:"KOZMO_RANDOM_KEY"  envDefault:"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"`
	// storage config
	PostgresAddr     string `env:"KOZMO_PG_ADDR" envDefault:"localhost"`
	PostgresPort     string `env:"KOZMO_PG_PORT" envDefault:"5432"`
	PostgresUser     string `env:"KOZMO_PG_USER" envDefault:"kozmo_builder"`
	PostgresPassword string `env:"KOZMO_PG_PASSWORD" envDefault:"71De5JllWSetLYU"`
	PostgresDatabase string `env:"KOZMO_PG_DATABASE" envDefault:"kozmo_builder"`
	// cache config
	RedisAddr     string `env:"KOZMO_REDIS_ADDR" envDefault:"localhost"`
	RedisPort     string `env:"KOZMO_REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"KOZMO_REDIS_PASSWORD" envDefault:"kozmo2022"`
	RedisDatabase int    `env:"KOZMO_REDIS_DATABASE" envDefault:"0"`
	// drive config
	DriveType             string `env:"KOZMO_DRIVE_TYPE" envDefault:""`
	DriveAccessKeyID      string `env:"KOZMO_DRIVE_ACCESS_KEY_ID" envDefault:""`
	DriveAccessKeySecret  string `env:"KOZMO_DRIVE_ACCESS_KEY_SECRET" envDefault:""`
	DriveRegion           string `env:"KOZMO_DRIVE_REGION" envDefault:""`
	DriveEndpoint         string `env:"KOZMO_DRIVE_ENDPOINT" envDefault:""`
	DriveSystemBucketName string `env:"KOZMO_DRIVE_SYSTEM_BUCKET_NAME" envDefault:"kozmo-cloud"`
	DriveTeamBucketName   string `env:"KOZMO_DRIVE_TEAM_BUCKET_NAME" envDefault:"kozmo-cloud-team"`
	DriveUploadTimeoutRaw string `env:"KOZMO_DRIVE_UPLOAD_TIMEOUT" envDefault:"30s"`
	DriveUploadTimeout    time.Duration
	// supervisor API
	KozmoSupervisorInternalRestAPI string `env:"KOZMO_SUPERVISOR_INTERNAL_API" envDefault:"http://127.0.0.1:9001/api/v1"`

	// peripheral API
	KozmoPeripheralAPI string `env:"KOZMO_PERIPHERAL_API" envDefault:"https://peripheral-api.kozmoai.com/v1/"`
	// resource manager API
	kozmoResourceManagerRestAPI         string `env:"KOZMO_RESOURCE_MANAGER_API" envDefault:"http://kozmo-resource-manager-backend:8006"`
	kozmoResourceManagerInternalRestAPI string `env:"KOZMO_RESOURCE_MANAGER_INTERNAL_API" envDefault:"http://kozmo-resource-manager-backend-internal:9004"`
	// kozmo marketplace config
	KozmoMarketplaceInternalRestAPI string `env:"KOZMO_MARKETPLACE_INTERNAL_API" envDefault:"http://kozmo-marketplace-backend-internal:9003/api/v1"`
	// token for internal api
	ControlToken string `env:"KOZMO_CONTROL_TOKEN" envDefault:""`
	// google config
	KozmoGoogleSheetsClientID     string `env:"KOZMO_GS_CLIENT_ID" envDefault:""`
	KozmoGoogleSheetsClientSecret string `env:"KOZMO_GS_CLIENT_SECRET" envDefault:""`
	KozmoGoogleSheetsRedirectURI  string `env:"KOZMO_GS_REDIRECT_URI" envDefault:""`
	// toke for ip zone detector
	KozmoIPZoneDetectorToken string `env:"KOZMO_IP_ZONE_DETECTOR_TOKEN" envDefault:""`
	// kozmo drive config
	KozmoDriveRestAPI string `env:"KOZMO_DRIVE_API" envDefault:"http://kozmo-drive-backend:8004"`
}

func getConfig() (*Config, error) {
	// fetch
	cfg := &Config{}
	err := env.Parse(cfg)
	// process data
	var errInParseDuration error
	cfg.DriveUploadTimeout, errInParseDuration = time.ParseDuration(cfg.DriveUploadTimeoutRaw)
	if errInParseDuration != nil {
		return nil, errInParseDuration
	}
	// ok
	fmt.Printf("----------------\n")
	fmt.Printf("run by following config: %+v\n", cfg)
	fmt.Printf("parse config error info: %+v\n", err)

	return cfg, err
}

func (c *Config) IsSelfHostMode() bool {
	return c.DeployMode == DEPLOY_MODE_SELF_HOST
}

func (c *Config) IsCloudMode() bool {
	if c.DeployMode == DEPLOY_MODE_CLOUD || c.DeployMode == DEPLOY_MODE_CLOUD_TEST || c.DeployMode == DEPLOY_MODE_CLOUD_BETA || c.DeployMode == DEPLOY_MODE_CLOUD_PRODUCTION {
		return true
	}
	return false
}

func (c *Config) IsCloudTestMode() bool {
	return c.DeployMode == DEPLOY_MODE_CLOUD_TEST
}

func (c *Config) IsCloudBetaMode() bool {
	return c.DeployMode == DEPLOY_MODE_CLOUD_BETA
}

func (c *Config) IsCloudProductionMode() bool {
	return c.DeployMode == DEPLOY_MODE_CLOUD_PRODUCTION
}

func (c *Config) GetWebScoketServerListenAddress() string {
	return c.WebsocketServerHost + ":" + c.WebsocketServerPort
}

func (c *Config) GetWebScoketServerConnectionAddress() string {
	return c.WebsocketServerConnectionHost + ":" + c.WebsocketServerConnectionPort
}

func (c *Config) GetWebsocketProtocol() string {
	if c.WSSEnabled == "true" {
		return PROTOCOL_WEBSOCKET_OVER_TLS
	}
	return PROTOCOL_WEBSOCKET
}

func (c *Config) GetRuntimeEnv() string {
	if c.IsCloudBetaMode() {
		return DEPLOY_MODE_CLOUD_BETA
	} else if c.IsCloudProductionMode() {
		return DEPLOY_MODE_CLOUD_PRODUCTION
	} else {
		return DEPLOY_MODE_CLOUD_TEST
	}
}

func (c *Config) GetSecretKey() string {
	return c.SecretKey
}

func (c *Config) GetRandomKey() string {
	return c.RandomKey
}

func (c *Config) GetPostgresAddr() string {
	return c.PostgresAddr
}

func (c *Config) GetPostgresPort() string {
	return c.PostgresPort
}

func (c *Config) GetPostgresUser() string {
	return c.PostgresUser
}

func (c *Config) GetPostgresPassword() string {
	return c.PostgresPassword
}

func (c *Config) GetPostgresDatabase() string {
	return c.PostgresDatabase
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddr
}

func (c *Config) GetRedisPort() string {
	return c.RedisPort
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

func (c *Config) GetRedisDatabase() int {
	return c.RedisDatabase
}

func (c *Config) GetDriveType() string {
	return c.DriveType
}

func (c *Config) IsAWSTypeDrive() bool {
	if c.DriveType == DRIVE_TYPE_AWS || c.DriveType == DRIVE_TYPE_DO {
		return true
	}
	return false
}

func (c *Config) IsMINIODrive() bool {
	return c.DriveType == DRIVE_TYPE_MINIO
}

func (c *Config) GetAWSS3Endpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetAWSS3AccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetAWSS3AccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetAWSS3Region() string {
	return c.DriveRegion
}

func (c *Config) GetAWSS3SystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetAWSS3TeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetAWSS3Timeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetMINIOAccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetMINIOAccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetMINIOEndpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetMINIOSystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetMINIOTeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetMINIOTimeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetControlToken() string {
	return c.ControlToken
}

func (c *Config) GetKozmoSupervisorInternalRestAPI() string {
	return c.KozmoSupervisorInternalRestAPI
}

func (c *Config) GetKozmoPeripheralAPI() string {
	return c.KozmoPeripheralAPI
}

func (c *Config) GetkozmoResourceManagerRestAPI() string {
	return c.kozmoResourceManagerRestAPI
}

func (c *Config) GetkozmoResourceManagerInternalRestAPI() string {
	return c.kozmoResourceManagerInternalRestAPI
}

func (c *Config) GetKozmoMarketplaceInternalRestAPI() string {
	return c.KozmoMarketplaceInternalRestAPI
}

func (c *Config) GetKozmoGoogleSheetsClientID() string {
	return c.KozmoGoogleSheetsClientID
}

func (c *Config) GetKozmoGoogleSheetsClientSecret() string {
	return c.KozmoGoogleSheetsClientSecret
}

func (c *Config) GetKozmoGoogleSheetsRedirectURI() string {
	return c.KozmoGoogleSheetsRedirectURI
}

func (c *Config) GetIPZoneDetectorToken() string {
	return c.KozmoIPZoneDetectorToken
}

func (c *Config) GetWebScoketServerConnectionAddressSouthAsia() string {
	return c.WebsocketServerConnectionHostSouthAsia + ":" + c.WebsocketServerConnectionPortSouthAsia
}

func (c *Config) GetWebScoketServerConnectionAddressEastAsia() string {
	return c.WebsocketServerConnectionHostEastAsia + ":" + c.WebsocketServerConnectionPortEastAsia
}

func (c *Config) GetWebScoketServerConnectionAddressCenterEurope() string {
	return c.WebsocketServerConnectionHostCenterEurope + ":" + c.WebsocketServerConnectionPortCenterEurope
}

func (c *Config) GetKozmoDriveAPIForSDK() string {
	return c.KozmoDriveRestAPI
}
