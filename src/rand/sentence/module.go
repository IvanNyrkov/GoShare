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

// NewModule creates struct that encapsulates the module
func NewModule(config ModuleConfig) (Module, error) {
	service, err := NewService(
		getAdjectivesFilename(config.AdjectivesFile),
		getNounsFilename(config.NounsFile),
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
