package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

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


	dtoNormal := dto.CreateUserDto{
		Username: "Dummy satu",
		Password: "passwordsa1",
	}

	jsonData, err := json.Marshal(dtoNormal)
	if err != nil {
		t.Error(err)
		t.Fail()
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

	t.Log(string(body))
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "00", gjson.Get(string(body), "code").String())
	t.Log("test user")
}