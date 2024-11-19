package usecases

import (
	"beer/module/notifies/models"
	"beer/module/notifies/repositories"
	"fmt"
)

type NotifyUsecase interface {
	NotifyDataProcessing(in *models.CreateNotifyGo) error
}

type NotifyUsecaseImpl struct {
	notifyRepository repositories.NotifyRepository
}

func NewNotifyUsecaseImpl(notifyRepository repositories.NotifyRepository) NotifyUsecase {
	return &NotifyUsecaseImpl{
		notifyRepository: notifyRepository,
	}
}

func (n *NotifyUsecaseImpl) NotifyDataProcessing(in *models.CreateNotifyGo) error {
	if in == nil {
		return fmt.Errorf("input data is nil")
	}

	err := n.notifyRepository.CreateNotifyData(in)
	if err != nil {
		return fmt.Errorf("failed to upload data: %w", err)
	}

	return nil
}
