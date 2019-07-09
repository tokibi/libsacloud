package naked

import (
	"time"

	"github.com/sacloud/libsacloud/v2/sacloud/types"
)

// SimpleMonitor シンプル監視
type SimpleMonitor struct {
	ID           types.ID               `json:",omitempty" yaml:"id,omitempty" structs:",omitempty"`
	Name         string                 `json:",omitempty" yaml:"name,omitempty" structs:",omitempty"`
	Description  string                 `json:",omitempty" yaml:"description,omitempty" structs:",omitempty"`
	Tags         []string               `json:"" yaml:"tags"`
	Icon         *Icon                  `json:",omitempty" yaml:"icon,omitempty" structs:",omitempty"`
	CreatedAt    *time.Time             `json:",omitempty" yaml:"created_at,omitempty" structs:",omitempty"`
	ModifiedAt   *time.Time             `json:",omitempty" yaml:"modified_at,omitempty" structs:",omitempty"`
	Availability types.EAvailability    `json:",omitempty" yaml:"availability,omitempty" structs:",omitempty"`
	ServiceClass string                 `json:",omitempty" yaml:"service_class,omitempty" structs:",omitempty"`
	Provider     *Provider              `json:",omitempty" yaml:"provider,omitempty" structs:",omitempty"`
	Settings     *SimpleMonitorSettings `json:",omitempty" yaml:"settings,omitempty" structs:",omitempty"`
	SettingsHash string                 `json:",omitempty" yaml:"settings_hash,omitempty" structs:",omitempty"`
	Status       *SimpleMonitorStatus   `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`
}

// SimpleMonitorSettings シンプル監視セッティング
type SimpleMonitorSettings struct {
	SimpleMonitor *SimpleMonitorSetting `json:",omitempty" yaml:"simple_monitor,omitempty" structs:",omitempty"`
}

// SimpleMonitorSetting シンプル監視セッティング
type SimpleMonitorSetting struct {
	DelayLoop   int                       `json:",omitempty" yaml:"delay_loop,omitempty" structs:",omitempty"`
	HealthCheck *SimpleMonitorHealthCheck `json:",omitempty" yaml:"health_check,omitempty" structs:",omitempty"`
	Enabled     types.StringFlag          `yaml:"enabled"`
	NotifyEmail *SimpleMonitorNotifyEmail `json:",omitempty" yaml:"notify_email,omitempty" structs:",omitempty"`
	NotifySlack *SimpleMonitorNotifySlack `json:",omitempty" yaml:"notify_slack,omitempty" structs:",omitempty"`
}

// SimpleMonitorHealthCheck シンプル監視 ヘルスチェック
type SimpleMonitorHealthCheck struct {
	Protocol          types.ESimpleMonitorProtocol `json:",omitempty" yaml:"protocol,omitempty" structs:",omitempty"`            // プロトコル
	Port              types.StringNumber           `json:",omitempty" yaml:"port,omitempty" structs:",omitempty"`                // ポート
	Path              string                       `json:",omitempty" yaml:"path,omitempty" structs:",omitempty"`                // HTTP/HTTPS監視の場合のリクエストパス
	Status            types.StringNumber           `json:",omitempty" yaml:"status,omitempty" structs:",omitempty"`              // HTTP/HTTPS監視の場合の期待ステータスコード
	SNI               types.StringFlag             `yaml:"sni"`                                                                  // HTTPS監視時のSNI有効/無効
	Host              string                       `json:",omitempty" yaml:"host,omitempty" structs:",omitempty"`                // 対象ホスト(IP or FQDN)
	BasicAuthUsername string                       `json:",omitempty" yaml:"basic_auth_username,omitempty" structs:",omitempty"` // HTTP/HTTPS監視の場合のBASIC認証 ユーザー名
	BasicAuthPassword string                       `json:",omitempty" yaml:"basic_auth_password,omitempty" structs:",omitempty"` // HTTP/HTTPS監視の場合のBASIC認証 パスワード
	QName             string                       `json:",omitempty" yaml:"qname,omitempty" structs:",omitempty"`               // DNS監視の場合の問い合わせFQDN
	ExpectedData      string                       `json:",omitempty" yaml:"expected_data,omitempty" structs:",omitempty"`       // 期待値
	Community         string                       `json:",omitempty" yaml:"community,omitempty" structs:",omitempty"`           // SNMP監視の場合のコミュニティ名
	SNMPVersion       string                       `json:",omitempty" yaml:"snmp_version,omitempty" structs:",omitempty"`        // SNMP監視 SNMPバージョン
	OID               string                       `json:",omitempty" yaml:"oid,omitempty" structs:",omitempty"`                 // SNMP監視 OID
	RemainingDays     int                          `json:",omitempty" yaml:"remaining_days,omitempty" structs:",omitempty"`      // SSL証明書 有効残日数
}

// SimpleMonitorNotifyEmail Eメールでの通知設定
type SimpleMonitorNotifyEmail struct {
	Enabled types.StringFlag `yaml:"enabled"` // 有効/無効
	HTML    types.StringFlag `yaml:"html"`    // メール通知の場合のHTMLメール有効フラグ
}

// SimpleMonitorNotifySlack Slackでの通知設定
type SimpleMonitorNotifySlack struct {
	Enabled             types.StringFlag `yaml:"enabled"`                                                                // 有効/無効
	IncomingWebhooksURL string           `json:",omitempty" yaml:"incoming_webhooks_url,omitempty" structs:",omitempty"` // Slack通知の場合のWebhook URL
}

// SimpleMonitorStatus シンプル監視 設定状況
type SimpleMonitorStatus struct {
	Target string `json:",omitempty" yaml:"target,omitempty" structs:",omitempty"` // 対象のIPアドレス or ホスト名
}

// SimpleMonitorHealthCheckStatus シンプル監視ステータス
type SimpleMonitorHealthCheckStatus struct {
	LastCheckedAt       *time.Time                 `json:",omitempty" yaml:"last_checked_at,omitempty" structs:",omitempty"`
	LastHealthChangedAt *time.Time                 `json:",omitempty" yaml:"last_health_changed_at,omitempty" structs:",omitempty"`
	Health              types.ESimpleMonitorHealth `json:",omitempty" yaml:"health,omitempty" structs:",omitempty"`
}
