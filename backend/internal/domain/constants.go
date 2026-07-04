package domain

// Status constants
const (
	StatusActive   = "active"
	StatusDisabled = "disabled"
	StatusError    = "error"
	StatusUnused   = "unused"
	StatusUsed     = "used"
	StatusExpired  = "expired"
)

// Role constants
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// Platform constants
const (
	PlatformAnthropic   = "anthropic"
	PlatformOpenAI      = "openai"
	PlatformGemini      = "gemini"
	PlatformAntigravity = "antigravity"
	PlatformGrok        = "grok"
)

// Account type constants
const (
	AccountTypeOAuth          = "oauth"           // OAuth类型账号（full scope: profile + inference）
	AccountTypeSetupToken     = "setup-token"     // Setup Token类型账号（inference only scope）
	AccountTypeAPIKey         = "apikey"          // API Key类型账号
	AccountTypeUpstream       = "upstream"        // 上游透传类型账号（通过 Base URL + API Key 连接上游）
	AccountTypeBedrock        = "bedrock"         // AWS Bedrock 类型账号（通过 SigV4 签名或 API Key 连接 Bedrock，由 credentials.auth_mode 区分）
	AccountTypeServiceAccount = "service_account" // Google Service Account 类型账号（用于 Vertex AI）
)

// Redeem type constants
const (
	RedeemTypeBalance      = "balance"
	RedeemTypeConcurrency  = "concurrency"
	RedeemTypeSubscription = "subscription"
	RedeemTypeInvitation   = "invitation"
)

// PromoCode status constants
const (
	PromoCodeStatusActive   = "active"
	PromoCodeStatusDisabled = "disabled"
)

// Admin adjustment type constants
const (
	AdjustmentTypeAdminBalance     = "admin_balance"     // 管理员调整余额
	AdjustmentTypeAdminConcurrency = "admin_concurrency" // 管理员调整并发数
)

// Group subscription type constants
const (
	SubscriptionTypeStandard     = "standard"     // 标准计费模式（按余额扣费）
	SubscriptionTypeSubscription = "subscription" // 订阅模式（按限额控制）
)

// Subscription status constants
const (
	SubscriptionStatusActive    = "active"
	SubscriptionStatusExpired   = "expired"
	SubscriptionStatusSuspended = "suspended"
)

// AntigravityGemini31ProAgentModel is the upstream route for Gemini 3.1 Pro High.
const AntigravityGemini31ProAgentModel = "gemini-pro-agent"

// DefaultAntigravityModelMapping 是 Antigravity 平台的默认模型映射。
// 实测 2026-07-04：通过 REST streamGenerateContent 可用的 4 个 Gemini 模型
// （retrieveUserQuota 权威确认）。Claude/GPT-OSS 走 gRPC 路径，REST 不可用。
var DefaultAntigravityModelMapping = map[string]string{
	// === REST 可用的 4 个 Gemini 模型 ===
	"gemini-2.5-pro":       "gemini-2.5-pro",
	"gemini-2.5-flash":     "gemini-2.5-flash",
	"gemini-2.5-flash-lite": "gemini-2.5-flash-lite",
	"gemini-3.1-flash-lite": "gemini-3.1-flash-lite",

	// === agy UI 标签 → REST 模型 ID 映射 ===
	// agy UI 显示这些标签，但 REST 端点用 Gemini 2.5 系列 ID
	"gemini-3.1-pro-high":      "gemini-2.5-pro",
	"gemini-3.1-pro-low":       "gemini-2.5-pro",
	"gemini-3-flash":           "gemini-2.5-flash",
	AntigravityGemini31ProAgentModel: "gemini-2.5-pro",
	"gemini-3.1-pro":                "gemini-2.5-pro",
	"gemini-3.1-pro-preview":        "gemini-2.5-pro",
	"gemini-3-pro-high":             "gemini-2.5-pro",
	"gemini-3-pro-low":              "gemini-2.5-pro",
}

// DefaultBedrockModelMapping 是 AWS Bedrock 平台的默认模型映射
// 将 Anthropic 标准模型名映射到 Bedrock 模型 ID
// 注意：此处的 "us." 前缀仅为默认值，ResolveBedrockModelID 会根据账号配置的
// aws_region 自动调整为匹配的区域前缀（如 eu.、apac.、jp. 等）
var DefaultBedrockModelMapping = map[string]string{
	// Claude Fable
	"claude-fable-5": "anthropic.claude-fable-5",
	// Claude Opus
	"claude-opus-4-8":          "us.anthropic.claude-opus-4-8-v1",
	"claude-opus-4-7":          "us.anthropic.claude-opus-4-7-v1",
	"claude-opus-4-6-thinking": "us.anthropic.claude-opus-4-6-v1",
	"claude-opus-4-6":          "us.anthropic.claude-opus-4-6-v1",
	"claude-opus-4-5-thinking": "us.anthropic.claude-opus-4-5-20251101-v1:0",
	"claude-opus-4-5-20251101": "us.anthropic.claude-opus-4-5-20251101-v1:0",
	"claude-opus-4-1":          "us.anthropic.claude-opus-4-1-20250805-v1:0",
	"claude-opus-4-20250514":   "us.anthropic.claude-opus-4-20250514-v1:0",
	// Claude Sonnet
	"claude-sonnet-5":            "us.anthropic.claude-sonnet-5-v1",
	"claude-sonnet-4-6-thinking": "us.anthropic.claude-sonnet-4-6",
	"claude-sonnet-4-6":          "us.anthropic.claude-sonnet-4-6",
	"claude-sonnet-4-5":          "us.anthropic.claude-sonnet-4-5-20250929-v1:0",
	"claude-sonnet-4-5-thinking": "us.anthropic.claude-sonnet-4-5-20250929-v1:0",
	"claude-sonnet-4-5-20250929": "us.anthropic.claude-sonnet-4-5-20250929-v1:0",
	"claude-sonnet-4-20250514":   "us.anthropic.claude-sonnet-4-20250514-v1:0",
	// Claude Haiku
	"claude-haiku-4-5":          "us.anthropic.claude-haiku-4-5-20251001-v1:0",
	"claude-haiku-4-5-20251001": "us.anthropic.claude-haiku-4-5-20251001-v1:0",
}
