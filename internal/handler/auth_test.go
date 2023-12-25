package handler

import (
	models "ToDoList/internal/models"
	"ToDoList/internal/service"
	mock_service "ToDoList/internal/service/mocks"
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehaviour func(*mock_service.MockAuthorization, models.User)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           models.User
		mockbehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"test","password":"test"}`,
			inputUser: models.User{
				Username: "test",
				Password: "test",
			},
			mockbehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		// {
		// 	name:      "Empty Fields",
		// 	inputBody: `{"username":"test"}`,

		// 	mockbehaviour: func(s *mock_service.MockAuthorization, user models.User) {
		// 		s.EXPECT().CreateUser(user).Return(1, nil)
		// 	},
		// 	expectedStatusCode:  400,
		// 	// expectedRequestBody: `{"message":"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		// },
		{
			name:      "Service Failure",
			inputBody: `{"username":"test","password":"test"}`,
			inputUser: models.User{
				Username: "test",
				Password: "test",
			},
			mockbehaviour: func(s *mock_service.MockAuthorization, user models.User) {
				s.EXPECT().CreateUser(user).Return(1, errors.New("service failute"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failute"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockbehaviour(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}

}
