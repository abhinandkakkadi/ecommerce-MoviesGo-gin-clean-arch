package usecase

import (
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
)

type cartUseCase struct {
	cartRepository interfaces.CartRepository
}


func NewCartUseCase(repository interfaces.CartRepository) services.CartUseCase {
	
	return &cartUseCase{
		cartRepository: repository,
	}

}

func (cr *cartUseCase) AddToCart(product_id int,userID int) {

	cr.cartRepository.AddToCart(product_id,userID)
}