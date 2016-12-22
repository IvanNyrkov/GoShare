package api

// Service is an interface that defines API actions
type Service interface {
}
type serviceImpl struct {
}

// NewService creates new service
func NewService() *serviceImpl {
	return &serviceImpl{}
}
