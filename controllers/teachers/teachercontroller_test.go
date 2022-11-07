package teachers

import (
	// "bytes"
	// "encoding/json"
	"net/http"
	"net/http/httptest"
	// "strconv"
	"strings"
	"testing"

	"mini-project/util"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	// "mini-project/controllers/teachers/request"

	_dbDriver "mini-project/drivers/mysql"
)

func InitEchoUser() *echo.Echo {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_TEST_NAME"),
	}

	configDB.InitDB()
	e := echo.New()

	return e
}

// func TestRegisterUser_Success(t *testing.T) {
// 	var testCases = []struct {
// 		name                   string
// 		path                   string
// 		expectedStatus         int
// 		expectedBodyStartsWith string
// 	}{{
// 		name:                   "success",
// 		path:                   "/api/v1/teachers",
// 		expectedStatus:         http.StatusCreated,
// 		expectedBodyStartsWith: "{\"status\":",
// 	},
// 	}

// 	e := InitEchoUser()

// 	teacherInput := request.Teacher{}

// 	teacherInput.Email = "ajizapar080500@gmail.com"
// 	teacherInput.Password = "TEST123"
// 	teacherInput.NIP = "123456789"
// 	teacherInput.Name = "Aji Muhammad Zapar"
// 	teacherInput.RoleId = "2837"

// 	jsonBody, _ := json.Marshal(&teacherInput)
// 	bodyReader := bytes.NewReader(jsonBody)

// 	req := httptest.NewRequest(http.MethodPost, "/api/v1/teachers", bodyReader)
// 	rec := httptest.NewRecorder()

// 	req.Header.Add("Content-Type", "application/json")

// 	c := e.NewContext(req, rec)

// 	var ctrl *TeacherController

// 	// ctrl = &TeacherController{}

// 	for _, testCase := range testCases {
// 		c.SetPath(testCase.path)

// 		if assert.NoError(t, ctrl.CreateTeacher(c)) {
// 			assert.Equal(t, http.StatusCreated, rec.Code)
// 			body := rec.Body.String()

// 			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
// 		}
// 	}

// }

func TestGetAllTeacher_Success(t *testing.T) {
	var testCases = []struct {
		name                   string
		path                   string
		expectedStatus         int
		expectedBodyStartsWith string
	}{{
		name:                   "success",
		path:                   "/api/v1/teachers",
		expectedStatus:         http.StatusOK,
		expectedBodyStartsWith: "{\"status\":",
	},
	}

	e := InitEchoUser()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/teachers", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var ctrl *TeacherController
	ctrl = &TeacherController{}

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, ctrl.GetTeachers(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}
