package updatelist

import (
	"sc-profile/models"
	"sc-profile/repository/updatelist"
)

type IService interface {
	AddUpdateList(updateList models.UpdateList) error
}

type Service struct {
	UpdateListRepository updatelist.IRepository
}

func NewService(updateListRepository updatelist.IRepository) *Service {
	return &Service{UpdateListRepository: updateListRepository}
}

func (s *Service) AddUpdateList(updateList models.UpdateList) error {
	if err := s.UpdateListRepository.InsertUpdateList(updateList); err != nil {
		return err
	}

	return nil
}
