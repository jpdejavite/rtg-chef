package constants

const (
	// AppName app name
	AppName = "chef"
	// DatabaseURL postgres database url
	DatabaseURL = "databaseUrl"
)

// GetConfigKeys return list of config keys
func GetConfigKeys() []string {
	return []string{DatabaseURL}
}
