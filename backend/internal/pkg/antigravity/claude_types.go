package antigravity

import (
	"encoding/json"
	"os"
	"strings"
)

// Claude 请求/响应类型定义

// ClaudeRequest Claude Messages API 请求
type ClaudeRequest struct {
	Model       string          `json:"model"`
	Messages    []ClaudeMessage `json:"messages"`
	MaxTokens   int             `json:"max_tokens,omitempty"`
	System      json.RawMessage `json:"system,omitempty"` // string 或 []SystemBlock
	Stream      bool            `json:"stream,omitempty"`
	Temperature *float64        `json:"temperature,omitempty"`
	TopP        *float64        `json:"top_p,omitempty"`
	TopK        *int            `json:"top_k,omitempty"`
	Tools       []ClaudeTool    `json:"tools,omitempty"`
	Thinking    *ThinkingConfig `json:"thinking,omitempty"`
	Metadata    *ClaudeMetadata `json:"metadata,omitempty"`
}

// ClaudeMessage Claude 消息
type ClaudeMessage struct {
	Role    string          `json:"role"` // user, assistant
	Content json.RawMessage `json:"content"`
}

// ThinkingConfig Thinking 配置
type ThinkingConfig struct {
	Type         string `json:"type"`                    // "enabled" / "adaptive" / "disabled"
	BudgetTokens int    `json:"budget_tokens,omitempty"` // thinking budget
}

// ClaudeMetadata 请求元数据
type ClaudeMetadata struct {
	UserID string `json:"user_id,omitempty"`
}

// ClaudeTool Claude 工具定义
// 支持两种格式：
// 1. 标准格式: { "name": "...", "description": "...", "input_schema": {...} }
// 2. Custom 格式 (MCP): { "type": "custom", "name": "...", "custom": { "description": "...", "input_schema": {...} } }
type ClaudeTool struct {
	Type        string          `json:"type,omitempty"` // "custom" 或空（标准格式）
	Name        string          `json:"name"`
	Description string          `json:"description,omitempty"`  // 标准格式使用
	InputSchema map[string]any  `json:"input_schema,omitempty"` // 标准格式使用
	Custom      *CustomToolSpec `json:"custom,omitempty"`       // custom 格式使用
}

// CustomToolSpec MCP custom 工具规格
type CustomToolSpec struct {
	Description string         `json:"description,omitempty"`
	InputSchema map[string]any `json:"input_schema"`
}

// ClaudeCustomToolSpec 兼容旧命名（MCP custom 工具规格）
type ClaudeCustomToolSpec = CustomToolSpec

// SystemBlock system prompt 数组形式的元素
type SystemBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// ContentBlock Claude 消息内容块（解析后）
type ContentBlock struct {
	Type string `json:"type"`
	// text
	Text string `json:"text,omitempty"`
	// thinking
	Thinking  string `json:"thinking,omitempty"`
	Signature string `json:"signature,omitempty"`
	// tool_use
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Input any    `json:"input,omitempty"`
	// tool_result
	ToolUseID string          `json:"tool_use_id,omitempty"`
	Content   json.RawMessage `json:"content,omitempty"`
	IsError   bool            `json:"is_error,omitempty"`
	// image
	Source *ImageSource `json:"source,omitempty"`
}

// ImageSource Claude 图片来源
type ImageSource struct {
	Type      string `json:"type"`       // "base64"
	MediaType string `json:"media_type"` // "image/png", "image/jpeg" 等
	Data      string `json:"data"`
}

// ClaudeResponse Claude Messages API 响应
type ClaudeResponse struct {
	ID           string              `json:"id"`
	Type         string              `json:"type"` // "message"
	Role         string              `json:"role"` // "assistant"
	Model        string              `json:"model"`
	Content      []ClaudeContentItem `json:"content"`
	StopReason   string              `json:"stop_reason,omitempty"`   // end_turn, tool_use, max_tokens
	StopSequence *string             `json:"stop_sequence,omitempty"` // null 或具体值
	Usage        ClaudeUsage         `json:"usage"`
}

// ClaudeContentItem Claude 响应内容项
type ClaudeContentItem struct {
	Type string `json:"type"` // text, thinking, tool_use

	// text
	Text string `json:"text,omitempty"`

	// thinking
	Thinking  string `json:"thinking,omitempty"`
	Signature string `json:"signature,omitempty"`

	// tool_use
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Input any    `json:"input,omitempty"`
}

// ClaudeUsage Claude 用量统计
type ClaudeUsage struct {
	InputTokens              int `json:"input_tokens"`
	OutputTokens             int `json:"output_tokens"`
	CacheCreationInputTokens int `json:"cache_creation_input_tokens,omitempty"`
	CacheReadInputTokens     int `json:"cache_read_input_tokens,omitempty"`
	ImageOutputTokens        int `json:"image_output_tokens,omitempty"`
}

// ClaudeError Claude 错误响应
type ClaudeError struct {
	Type  string      `json:"type"` // "error"
	Error ErrorDetail `json:"error"`
}

// ErrorDetail 错误详情
type ErrorDetail struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// modelDef Antigravity 模型定义（内部使用）
type modelDef struct {
	ID          string
	DisplayName string
	CreatedAt   string // 仅 Claude API 格式使用
	IsReasoning bool
}

// Antigravity 通过 REST streamGenerateContent 可用的模型。
// 注意：agy UI 显示更多模型（含 Claude/GPT-OSS），但那些走 gRPC
// CloudCode/GenerateChat 路径，sub2api 的 REST 中继不可用。
// 权威来源：retrieveUserQuota 返回的 modelId（实测 2026-07-04）。
var claudeModels = []modelDef{}
var gptossModels = []modelDef{}

var geminiModels = []modelDef{
	{ID: "gemini-2.5-pro", DisplayName: "Gemini 2.5 Pro", CreatedAt: "2025-01-01T00:00:00Z", IsReasoning: true},
	{ID: "gemini-2.5-flash", DisplayName: "Gemini 2.5 Flash", CreatedAt: "2025-01-01T00:00:00Z"},
	{ID: "gemini-2.5-flash-lite", DisplayName: "Gemini 2.5 Flash Lite", CreatedAt: "2025-01-01T00:00:00Z"},
	{ID: "gemini-3.1-flash-lite", DisplayName: "Gemini 3.1 Flash Lite", CreatedAt: "2026-02-19T00:00:00Z"},
}

// ========== Claude API 格式 (/v1/models) ==========

// ClaudeModel Claude API 模型格式
type ClaudeModel struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	DisplayName string `json:"display_name"`
	CreatedAt   string `json:"created_at"`
}

// AntigravityModelsEnv 是 Antigravity 自定义模型列表的环境变量名。
// 格式：逗号分隔的模型 ID（如 "claude-sonnet-4-6,gemini-3-flash"）。
// 设置后 DefaultModels() 返回自定义列表，覆盖硬编码默认值。
const AntigravityModelsEnv = "ANTIGRAVITY_MODELS"

// DefaultModels 返回 Claude API 格式的模型列表（Claude + Gemini）
// 如果设置了 ANTIGRAVITY_MODELS 环境变量，则使用自定义列表覆盖硬编码默认值。
func DefaultModels() []ClaudeModel {
	// 检查环境变量覆盖
	if customModels := os.Getenv(AntigravityModelsEnv); strings.TrimSpace(customModels) != "" {
		return parseCustomModels(customModels)
	}

	all := append(append(claudeModels, geminiModels...), gptossModels...)
	result := make([]ClaudeModel, len(all))
	for i, m := range all {
		result[i] = ClaudeModel{ID: m.ID, Type: "model", DisplayName: m.DisplayName, CreatedAt: m.CreatedAt}
	}
	return result
}

// parseCustomModels 将逗号分隔的模型 ID 字符串解析为 ClaudeModel 列表
func parseCustomModels(models string) []ClaudeModel {
	parts := strings.Split(models, ",")
	result := make([]ClaudeModel, 0, len(parts))
	for _, p := range parts {
		id := strings.TrimSpace(p)
		if id == "" {
			continue
		}
		result = append(result, ClaudeModel{ID: id, Type: "model", DisplayName: id})
	}
	return result
}

// ========== Gemini v1beta 格式 (/v1beta/models) ==========

// GeminiModel Gemini v1beta 模型格式
type GeminiModel struct {
	Name                       string   `json:"name"`
	DisplayName                string   `json:"displayName,omitempty"`
	SupportedGenerationMethods []string `json:"supportedGenerationMethods,omitempty"`
}

// GeminiModelsListResponse Gemini v1beta 模型列表响应
type GeminiModelsListResponse struct {
	Models []GeminiModel `json:"models"`
}

var defaultGeminiMethods = []string{"generateContent", "streamGenerateContent"}

// DefaultGeminiModels 返回 Gemini v1beta 格式的模型列表（仅 Gemini 模型）
func DefaultGeminiModels() []GeminiModel {
	result := make([]GeminiModel, len(geminiModels))
	for i, m := range geminiModels {
		result[i] = GeminiModel{Name: "models/" + m.ID, DisplayName: m.DisplayName, SupportedGenerationMethods: defaultGeminiMethods}
	}
	return result
}

// FallbackGeminiModelsList 返回 Gemini v1beta 格式的模型列表响应
func FallbackGeminiModelsList() GeminiModelsListResponse {
	return GeminiModelsListResponse{Models: DefaultGeminiModels()}
}

// FallbackGeminiModel 返回单个模型信息（v1beta 格式）
func FallbackGeminiModel(model string) GeminiModel {
	if model == "" {
		return GeminiModel{Name: "models/unknown", SupportedGenerationMethods: defaultGeminiMethods}
	}
	name := model
	if len(model) < 7 || model[:7] != "models/" {
		name = "models/" + model
	}
	return GeminiModel{Name: name, SupportedGenerationMethods: defaultGeminiMethods}
}

// IsGeminiReasoningModel 判断是否为不支持参数和强制 ToolConfig 的 Gemini 推理模型
func IsGeminiReasoningModel(modelID string) bool {
	lowerID := strings.ToLower(modelID)
	for _, m := range geminiModels {
		if strings.Contains(lowerID, m.ID) && m.IsReasoning {
			return true
		}
	}
	return false
}
