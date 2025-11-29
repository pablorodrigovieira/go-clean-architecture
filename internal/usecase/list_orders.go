package usecase

import (
	"github.com/pablorodrigovieira/go-clean-architecture/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (u *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := u.OrderRepository.List()
	if err != nil {
		return nil, err
	}
	result := make([]OrderOutputDTO, 0, len(orders))
	for _, o := range orders {
		result = append(result, OrderOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return result, nil
}
