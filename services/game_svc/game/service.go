package game

import (

)

// Service is service
type Service struct {

}

// NewService returns game service
func NewService() *Service {
	svc := Service{}
	return &svc
}

func (svc *Service) Play() {
	
}