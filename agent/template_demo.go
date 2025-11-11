package agent

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func TemplateDemo(ctx context.Context) ([]*schema.Message, error) {
	// 创建模板
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage("你是一个{role}。"),
		schema.MessagesPlaceholder("history_key", false),
		&schema.Message{
			Role:    schema.User,
			Content: "请帮我{task}。",
		},
	)

	// 在 Chain 中使用
	chain := compose.NewChain[map[string]any, []*schema.Message]()
	chain.AppendChatTemplate(template)

	r, err := chain.Compile(ctx)
	if err != nil {
		return nil, err
	}

	// 准备变量
	variables := map[string]any{
		"role": "专业的助手",
		"task": "写一首诗",
		"history_key": []*schema.Message{
			{Role: schema.User, Content: "告诉我油画是什么?"},
			{Role: schema.Assistant, Content: "油画是xxx"},
		},
	}

	output, err := r.Invoke(context.Background(), variables)
	if err != nil {
		return nil, err
	}
	return output, nil
}
