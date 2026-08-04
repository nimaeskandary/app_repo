package secure_storage_core

// Config configures secure storage.
type Config struct {
	Namespace string `json:"Namespace" validate:"required"`
}
