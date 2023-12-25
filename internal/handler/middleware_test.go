package handler

import (
	"ToDoList/internal/service"
	mock_service "ToDoList/internal/service/mocks"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandlerUserId(t *testing.T) {
	type mockBehaviour func(*mock_service.MockAuthorization, string)
	testTable := []struct {
		name                string
		headerName          string
		headerVal           string
		token               string
		mockbehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:       "OK",
			headerName: "Authorization",
			headerVal:  "Bearer token",
			token:      "token",
			mockbehaviour: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "1",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockbehaviour(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/smth", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/smth", nil)
			req.Header.Set(testCase.headerName, testCase.headerVal)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}

}
