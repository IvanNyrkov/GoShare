package store

// Module defines module interface
type Module interface {
	GetService() Service
}
type moduleImpl struct {
	service Service
}

// ModuleConfig contains required configs
type ModuleConfig struct {
	UploadDir string
}

// DefaultConfig is a default config settings
var DefaultConfig = ModuleConfig{
	UploadDir: defaultUploadDir,
}

// NewModule creates struct that encapsulates the module
func NewModule(config ModuleConfig) *moduleImpl {
	service := NewService(config.UploadDir)
	return &moduleImpl{
		service: service,
	}
}

// GetService returns module service
func (m *moduleImpl) GetService() Service {
	return m.service
}
