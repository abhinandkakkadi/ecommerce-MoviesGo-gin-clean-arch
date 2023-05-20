package handler

import (
	"github.com/gin-gonic/gin"
	services "github.com/thnkrn/go-gin-clean-arch/pkg/usecase/interface"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

func (cr *AdminHandler) LoginHandler(c *gin.Context) {

	
}