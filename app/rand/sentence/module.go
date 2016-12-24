package sentence

// Module defines module interface
type Module interface {
	GetService() Service
}
type moduleImpl struct {
	service Service
}

// ModuleConfig contains required configs
type ModuleConfig struct {
	NounsFile      string
	AdjectivesFile string
}

// DefaultConfig is a default config settings
var DefaultConfig = ModuleConfig{
	NounsFile:      defaultNounsFilePath,
	AdjectivesFile: defaultAdjectivesFilePath,
}

// NewModule creates struct that encapsulates the module
func NewModule(config ModuleConfig) (*moduleImpl, error) {
	service, err := NewService(
		config.AdjectivesFile,
		config.NounsFile,
	)
	if err != nil {
		return nil, err
	}
	return &moduleImpl{
		service: service,
	}, nil
}

// GetService returns module service
func (m *moduleImpl) GetService() Service {
	return m.service
}
