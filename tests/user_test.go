package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/RVND-90/startup-apps-go/constants"
	"github.com/RVND-90/startup-apps-go/models/dto"
	"github.com/RVND-90/startup-apps-go/server"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)
var app = fiber.New()

func setup() {
	
	go server.Bootstrap(app)
}

func TestUser(t *testing.T) {
	setup()
	defer app.Shutdown()

	type TestCase struct {
		name string
		dto interface{}
		expected dto.ResponseDto
	}

	tests := []TestCase{
		{
			name: "test user - success",
			dto: dto.CreateUserDto{
				Username: "Dummy satu",
				Password: "passwordsa1",
			},
			expected: dto.ResponseDto{
				RC: constants.GetRC(constants.SUCCESS_CODE),
			},
		},
		{
			name: "test user - validation error",
			dto: dto.CreateUserDto{
				Username: "Dummy satu",
				Password: "p",
			},
			expected: dto.ResponseDto{
				RC: constants.GetRC(constants.ERROR_VALIDATION_CODE),
			},
		},

		{
			name: "test user - parse error",
			dto: nil,
			expected: dto.ResponseDto{
				RC: constants.GetRC(constants.ERROR_PARSE_CODE),
			},
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		var err error
		jsonData := []byte("")
		if test.dto == nil {
			jsonData = []byte(`{`)
		}else{
			jsonData, err = json.Marshal(test.dto)
			if err != nil {
				t.Error(err)
				t.Fail()
			}
		}
		
		//post data to server
		req:= httptest.NewRequest("POST", "http://localhost:3000/user", bytes.NewReader(jsonData))
		req.Header.Set("Content-Type", "application/json")
		
		resp, errResp := app.Test(req)
		if errResp != nil {
			t.Error(err)
		}
		
		//
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
			t.Fail()
		}

		t.Logf("%d - %s\n", resp.StatusCode, string(body))
		assert.Equal(t, test.expected.HttpResponseCode, resp.StatusCode)
		assert.Equal(t, test.expected.Code, gjson.Get(string(body), "code").String())
		assert.Equal(t, test.expected.Message, gjson.Get(string(body), "message").String())
	}



	
}