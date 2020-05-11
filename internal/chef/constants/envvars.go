package constants

const (
	// FirebaseCredential firebase credential key
	FirebaseCredential = "FIREBASE_CREDENTIAL"
	// Port port key
	Port = "PORT"
)

// GetEnVarKeys return list of env var keys needed
func GetEnVarKeys() []string {
	return []string{FirebaseCredential, Port}
}
