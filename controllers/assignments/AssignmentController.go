package assignments

import (
	// "mini-project/app/middlewares"
	"mini-project/businesses/assignments"
	controller "mini-project/controllers"
	"mini-project/controllers/assignments/request"
	"mini-project/controllers/assignments/response"
	"net/http"

	// "github.com/golang-jwt/jwt"
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