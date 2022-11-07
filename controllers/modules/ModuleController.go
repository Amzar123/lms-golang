package modules

import (
	// "mini-project/app/middlewares"
	"mini-project/businesses/modules"
	controller "mini-project/controllers"
	"mini-project/controllers/modules/request"
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

func (ctrl *ModuleController) CreateModule(c echo.Context) error {
	moduleInput := request.Module{}

	if err := c.Bind(&moduleInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := moduleInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	module := ctrl.moduleUseCase.Create(moduleInput.ToDomain())

	return c.JSON(http.StatusCreated, response.FromDomain(module))
}

func (ctrl *ModuleController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	module := ctrl.moduleUseCase.GetByID(id)

	if module.IdModule == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "assignment not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note found", response.FromDomain(module))
}

func (ctrl *ModuleController) Update(c echo.Context) error {
	input := request.Module{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var moduleId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	module := ctrl.moduleUseCase.Update(moduleId, input.ToDomain())

	if module.IdModule == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "Assignment not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "Assignment updated", response.FromDomain(module))
}

func (ctrl *ModuleController) Delete(c echo.Context) error {
	var moduleId string = c.Param("id")

	isSuccess := ctrl.moduleUseCase.Delete(moduleId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "assignment not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "assignment deleted", "")
}