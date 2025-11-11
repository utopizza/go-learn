package agent

import (
	"context"
	"testing"
)

func TestTemplateDemo(t *testing.T) {
	res, err := TemplateDemo(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
}
