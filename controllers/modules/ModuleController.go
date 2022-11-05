package modules

import (
	// "mini-project/app/middlewares"
	"mini-project/businesses/modules"
	controller "mini-project/controllers"
	// "mini-project/controllers/modules/request"
	"mini-project/controllers/modules/response"
	"net/http"

	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type ModuleController struct {
	moduleUseCase modules.Usecase
}

func NewModuleController(moduleUC modules.Usecase) *ModuleController {
	return &ModuleController{
		moduleUseCase: moduleUC,
	}
}

func (ctrl *ModuleController) GetModules(c echo.Context) error {
	modulesData := ctrl.moduleUseCase.GetModules()

	modules := []response.Module{}

	for _, module := range modulesData {
		modules = append(modules, response.FromDomain(module))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all module", modules)
}