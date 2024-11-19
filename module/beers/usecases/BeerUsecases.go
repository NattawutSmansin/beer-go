package usecases

import (
	"beer/module/beers/models"
	"beer/module/beers/repositories"
	"fmt"
)

type BeerUsecase interface {
	BeerCreateDataProcess(in *models.CreateBeer) error
	BeerUpdateDataProcess(in *models.UpdateBeer) error
	BeerDataProcess(beerId uint32) (*models.Beer, error)
	BeerDataDelete(beerId uint32) error
	BeerListDataProcess(name string, page,limit int) ([]*models.Beer, int64, error)
}

type BeerUsecaseImpl struct {
	beerRepository repositories.BeerRepository
}

func NewBeerUsecaseImpl(beerRepository repositories.BeerRepository) BeerUsecase {
	return &BeerUsecaseImpl{
		beerRepository: beerRepository,
	}
}

func (beerUsecaseImpl *BeerUsecaseImpl) BeerCreateDataProcess(in *models.CreateBeer) error {
	if in == nil {
		return fmt.Errorf("input data is nil")
	}

	err := beerUsecaseImpl.beerRepository.Store(in)
	if err != nil {
		return fmt.Errorf("failed to create data: %w", err)
	}

	return nil
}

func (beerUsecaseImpl *BeerUsecaseImpl) BeerListDataProcess(name string, page,limit int) ([]*models.Beer, int64, error) {
	listBeer, total, err := beerUsecaseImpl.beerRepository.List(name, page,limit)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create data: %w", err)
	}

	return listBeer, total, nil
}

func (beerUsecaseImpl *BeerUsecaseImpl) BeerUpdateDataProcess(in *models.UpdateBeer) error {
	if in == nil {
		return fmt.Errorf("input data is nil")
	}

	err := beerUsecaseImpl.beerRepository.Update(in)
	if err != nil {
		return fmt.Errorf("failed to update data: %w", err)
	}

	return nil
}

func (beerUsecaseImpl *BeerUsecaseImpl) BeerDataProcess(beerId uint32) (*models.Beer, error) {
	beer, err := beerUsecaseImpl.beerRepository.Data(beerId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data: %w", err)
	}

	return beer, nil
}

func (beerUsecaseImpl *BeerUsecaseImpl) BeerDataDelete(beerId uint32) error {
	err := beerUsecaseImpl.beerRepository.Delete(beerId)
	if err != nil {
		return fmt.Errorf("failed to retrieve data: %w", err)
	}

	return nil
}
