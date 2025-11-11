package agent

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

func ToolDemo(ctx context.Context) ([]*schema.Message, error) {
	// 初始化 tools
	searchTool, err := duckduckgo.NewTool(ctx, &duckduckgo.Config{})
	todoTools := []tool.BaseTool{
		//getAddTodoTool(), // NewTool 构建
		//updateTool,       // InferTool 构建
		//&ListTodoTool{},  // 实现Tool接口
		searchTool, // 官方封装的工具
	}

	// 创建并配置 ChatModel
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Model:   "qwen3:8b",               // 模型名称
	})
	if err != nil {
		return nil, err
	}

	// 获取工具信息并绑定到 ChatModel
	toolInfos := make([]*schema.ToolInfo, 0, len(todoTools))
	for _, t := range todoTools {
		info, err := t.Info(ctx)
		if err != nil {
			return nil, err
		}
		toolInfos = append(toolInfos, info)
	}
	if err = chatModel.BindTools(toolInfos); err != nil {
		return nil, err
	}

	// 创建 tools 节点
	todoToolsNode, err := compose.NewToolNode(ctx, &compose.ToolsNodeConfig{
		Tools: todoTools,
	})
	if err != nil {
		return nil, err
	}

	// 构建完整的处理链
	chain := compose.NewChain[[]*schema.Message, []*schema.Message]()
	chain.
		AppendChatModel(chatModel, compose.WithNodeName("chat_model")).
		AppendToolsNode(todoToolsNode, compose.WithNodeName("tools"))

	// 编译并运行 chain
	agent, err := chain.Compile(ctx)
	if err != nil {
		return nil, err
	}

	// 运行示例
	resp, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "添加一个学习 Eino 的 TODO，同时搜索一下 cloudwego/eino 的仓库地址",
		},
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
