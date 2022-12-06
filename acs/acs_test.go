package acs

import (
	"testing"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func TestACS(t *testing.T) {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	acs := &ACS{}
	err := acs.CreateAuthnResult(ctx).
		WithNotLoginTicket(ctx, "not_login_ticket_test").
		WithVerifyTicket(ctx, "verify_ticket").
		Submit(ctx)
	if err != nil {
		t.Error(err)
	}
}
