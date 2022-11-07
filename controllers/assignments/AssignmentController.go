package assignments

import (
	// "mini-project/app/middlewares"
	"mini-project/businesses/assignments"
	controller "mini-project/controllers"
	"mini-project/controllers/assignments/request"
	"mini-project/controllers/assignments/response"
	"net/http"

	// "github.com/golang-jwt/jwt"
	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AssignmentController struct {
	assignmentUseCase assignments.Usecase
}

func NewAssignmentController(assignmentUC assignments.Usecase) *AssignmentController {
	return &AssignmentController{
		assignmentUseCase: assignmentUC,
	}
}

func (ctrl *AssignmentController) GetAssignments(c echo.Context) error {
	assignmentsData := ctrl.assignmentUseCase.GetAssignments()

	assignments := []response.Assignment{}

	for _, assignment := range assignmentsData {
		assignments = append(assignments, response.FromDomain(assignment))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all assignment", assignments)
}

func (ctrl *AssignmentController) CreateAssignment(c echo.Context) error {
	assignmentInput := request.Assignment{}
	id := (guuid.New()).String() 

	assignmentInput.IdAssignment = id;

	if err := c.Bind(&assignmentInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := assignmentInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	assignment := ctrl.assignmentUseCase.Create(assignmentInput.ToDomain())

	return c.JSON(http.StatusCreated, response.FromDomain(assignment))
}

func (ctrl *AssignmentController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	assignment := ctrl.assignmentUseCase.GetByID(id)

	if assignment.IdAssignment == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "assignment not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note found", response.FromDomain(assignment))
}

func (ctrl *AssignmentController) Update(c echo.Context) error {
	input := request.Assignment{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var assignmentId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	assignment := ctrl.assignmentUseCase.Update(assignmentId, input.ToDomain())

	if assignment.IdAssignment == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "Assignment not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "Assignment updated", response.FromDomain(assignment))
}

func (ctrl *AssignmentController) Delete(c echo.Context) error {
	var assignmentId string = c.Param("id")

	isSuccess := ctrl.assignmentUseCase.Delete(assignmentId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "assignment not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "assignment deleted", "")
}