package agent

import (
	"context"
	"testing"
)

func TestChatDemo(t *testing.T) {
	res, err := ChatDemo(context.Background())
	if err != nil {
		t.Error(err)
	}
	t.Logf("res=%+v", res)
}
