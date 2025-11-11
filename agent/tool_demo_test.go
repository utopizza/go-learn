package agent

import (
	"context"
	"fmt"
	"testing"
)

func TestToolDemo(t *testing.T) {
	res, err := ToolDemo(context.Background())
	if err != nil {
		t.Error(err)
	}
	for _, msg := range res {
		fmt.Println(msg.Content)
	}
}
