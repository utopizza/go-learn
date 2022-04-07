package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func testGetting(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func testSettingCookies() {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest(http.MethodGet, "/hello/something", nil)
	cookieName := "ck"
	ctx.Request.AddCookie(&http.Cookie{
		Name:     cookieName,
		Value:    "pornhub",
		Path:     "/",
		Domain:   "www.wys.com",
		Secure:   true,
		HttpOnly: false,
	})
	val, _ := ctx.Cookie(cookieName)
	fmt.Println(val)
}

func testParseForm() {
	data := url.Values{}
	data.Set("name", "foo")
	data.Set("surname", "bar")

	req, _ := http.NewRequest("POST", "https://www.baidu.com", strings.NewReader(data.Encode()))
	req.Header.Set("content-type", "application/json")
	//req.Header.Set("content-type", "application/x-www-form-urlencoded")

	err := req.ParseForm()
	if err != nil {
		fmt.Printf("parse form error: %+v\n", err)
	}
	fmt.Printf("request form: %+v\n", req.Form)

	req.Form.Set("address", "beijing")
	fmt.Printf("setting form: %+v\n", req.Form)
}

func main() {
	testParseForm()
}
