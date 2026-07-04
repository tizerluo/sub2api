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

// DefaultAntigravityModelMapping 是 Antigravity 平台的默认模型映射
// 当账号未配置 model_mapping 时使用此默认值
// 实测 2026-07-04：Google AI Pro 订阅实际支持的 6 个模型
var DefaultAntigravityModelMapping = map[string]string{
	// === 实际支持的 6 个模型 ===
	// Gemini
	"gemini-3.1-pro-high": AntigravityGemini31ProAgentModel,
	"gemini-3.1-pro-low":  "gemini-3.1-pro-low",
	"gemini-3-flash":      "gemini-3-flash",
	// Claude
	"claude-sonnet-4-6-thinking": "claude-sonnet-4-6-thinking",
	"claude-opus-4-6-thinking":   "claude-opus-4-6-thinking",
	// GPT-OSS
	"gpt-oss-120b-medium": "gpt-oss-120b-medium",

	// === 旧模型名兼容映射（重定向到实际支持的模型）===
	// 旧 Claude → 当前支持的
	"claude-sonnet-4-6":          "claude-sonnet-4-6-thinking",
	"claude-opus-4-6":            "claude-opus-4-6-thinking",
	"claude-sonnet-4-5":          "claude-sonnet-4-6-thinking",
	"claude-sonnet-4-5-thinking": "claude-sonnet-4-6-thinking",
	"claude-opus-4-5-thinking":   "claude-opus-4-6-thinking",
	"claude-opus-4-7":            "claude-opus-4-6-thinking",
	"claude-opus-4-8":            "claude-opus-4-6-thinking",
	"claude-fable-5":             "claude-opus-4-6-thinking",
	"claude-haiku-4-5":           "claude-sonnet-4-6-thinking",
	// 旧 Claude 版本号 ID
	"claude-sonnet-4-5-20250929": "claude-sonnet-4-6-thinking",
	"claude-opus-4-5-20251101":   "claude-opus-4-6-thinking",
	"claude-haiku-4-5-20251001":  "claude-sonnet-4-6-thinking",
	// 旧 Gemini → 当前支持的
	"gemini-2.5-flash":          "gemini-3-flash",
	"gemini-2.5-flash-lite":     "gemini-3-flash",
	"gemini-2.5-pro":            AntigravityGemini31ProAgentModel,
	"gemini-3-pro-high":         AntigravityGemini31ProAgentModel,
	"gemini-3-pro-low":          "gemini-3.1-pro-low",
	AntigravityGemini31ProAgentModel: AntigravityGemini31ProAgentModel,
	"gemini-3.1-pro":                 AntigravityGemini31ProAgentModel,
	"gemini-3.1-pro-preview":         AntigravityGemini31ProAgentModel,
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
