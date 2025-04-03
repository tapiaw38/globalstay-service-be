package config

var configService *ConfigurationService

const (
	ReleaseMode GinModeServer = "release"
	DebugMode   GinModeServer = "debug"
)

type (
	GinModeServer string

	ConfigurationService struct {
		ServerConfig ServerConfig
		NoSQLConfig  NoSQLConfig
		GCP          GCPConfig
	}

	ServerConfig struct {
		GinMode   GinModeServer
		Port      string
		Host      string
		JWTSecret string
	}

	NoSQLConfig struct {
		Migrations   *NoSQLCollectionConfig
		Hotels       *NoSQLCollectionConfig
		Locations    *NoSQLCollectionConfig
		Services     *NoSQLCollectionConfig
		Reservations *NoSQLCollectionConfig
	}

	NoSQLCollectionConfig struct {
		DatabaseURI string
		Database    string
		Collection  string
	}

	GCPConfig struct {
		MapsKey string
	}
)

func InitConfigService(config *ConfigurationService) {
	if configService == nil {
		configService = config
	}
}

func GetConfigService() ConfigurationService {
	return *configService
}
