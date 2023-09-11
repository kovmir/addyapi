package addyapi

import (
	"encoding/json"
	"strings"
	"time"
)

type AliasSortCond string

const (
	AliasSortLocalPart  AliasSortCond = "local_part"
	AliasSortDomain     AliasSortCond = "domain"
	AliasSortEmail      AliasSortCond = "email"
	AliasSortEmailsFwd  AliasSortCond = "emails_forwarded"
	AliasSortEmailsBlkd AliasSortCond = "emails_blocked"
	AliasSortEmailsRepl AliasSortCond = "emails_replied"
	AliasSortEmailsSent AliasSortCond = "emails_sent"
	AliasSortActive     AliasSortCond = "active"
	AliasSortCreatedAt  AliasSortCond = "created_at"
	AliasSortUpdatedAt  AliasSortCond = "updated_at"
	AliasSortDeletedAt  AliasSortCond = "deleted_at"
)

// This type specifies how to parse timestamps in responses from addy.
type AddyTime time.Time

// Response error
type ErrAddyResponse struct {
	Code    int
	Message string
}

func (at *AddyTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(time.DateTime, s)
	if err != nil {
		return err
	}
	*at = AddyTime(t)
	return nil
}

func (at AddyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(at))
}

func (at AddyTime) String() string {
	return time.Time(at).String()
}

func (ear ErrAddyResponse) Error() string {
	return ear.Message
}

type FailedDelivery struct {
	ID             string   `json:"id"`
	UserID         string   `json:"user_id"`
	RecipientID    string   `json:"recipient_id"`
	RecipientEmail string   `json:"recipient_email"`
	AliasID        string   `json:"alias_id"`
	AliasEmail     string   `json:"alias_email"`
	BounceType     string   `json:"bounce_type"`
	RemoteMta      string   `json:"remote_mta"`
	Sender         string   `json:"sender"`
	EmailType      string   `json:"email_type"`
	Status         string   `json:"status"`
	Code           string   `json:"code"`
	AttemptedAt    AddyTime `json:"attempted_at"`
	CreatedAt      AddyTime `json:"created_at"`
	UpdatedAt      AddyTime `json:"updated_at"`
}

type Domain struct {
	ID                      string    `json:"id"`
	UserID                  string    `json:"user_id"`
	Domain                  string    `json:"domain"`
	Description             string    `json:"description"`
	Aliases                 []Alias   `json:"aliases"`
	DefaultRecipient        string    `json:"default_recipient"`
	Active                  bool      `json:"active"`
	CatchAll                bool      `json:"catch_all"`
	DomainVerifiedAt        *AddyTime `json:"domain_verified_at"`
	DomainMxValidatedAt     *AddyTime `json:"domain_mx_validated_at"`
	DomainSendingVerifiedAt *AddyTime `json:"domain_sending_verified_at"`
	CreatedAt               AddyTime  `json:"created_at"`
	UpdatedAt               *AddyTime `json:"updated_at"`
}

type DomainOptions struct {
	Data               []string `json:"data"`
	DefaultAliasDomain string   `json:"defaultAliasDomain"`
	DefaultAliasFormat string   `json:"defaultAliasFormat"`
}

type AppVersion struct {
	Version string `json:"version"`
	Major   int    `json:"major"`
	Minor   int    `json:"minor"`
	Patch   int    `json:"patch"`
}

type Recipient struct {
	ID               string    `json:"id"`
	UserID           string    `json:"user_id"`
	Email            string    `json:"email"`
	CanReplySend     bool      `json:"can_reply_send"`
	ShouldEncrypt    bool      `json:"should_encrypt"`
	InlineEncryption bool      `json:"inline_encryption"`
	ProtectedHeaders bool      `json:"protected_headers"`
	Fingerprint      string    `json:"fingerprint"`
	EmailVerifiedAt  *AddyTime `json:"email_verified_at"`
	Aliases          []Alias   `json:"aliases"`
	CreatedAt        AddyTime  `json:"created_at"`
	UpdatedAt        *AddyTime `json:"updated_at"`
}

type AccountDetails struct {
	ID                           string    `json:"id"`
	Username                     string    `json:"username"`
	FromName                     string    `json:"from_name"`
	EmailSubject                 string    `json:"email_subject"`
	BannerLocation               string    `json:"banner_location"`
	Bandwidth                    uint      `json:"bandwidth"`
	UsernameCount                uint      `json:"username_count"`
	UsernameLimit                uint      `json:"username_limit"`
	DefaultRecipientID           string    `json:"default_recipient_id"`
	DefaultAliasDomain           string    `json:"default_alias_domain"`
	DefaultAliasFormat           string    `json:"default_alias_format"`
	Subscription                 string    `json:"subscription"`
	SubscriptionEndsAt           *AddyTime `json:"subscription_ends_at"`
	BandwidthLimit               uint      `json:"bandwidth_limit"`
	RecipientCount               uint      `json:"recipient_count"`
	RecipientLimit               uint      `json:"recipient_limit"`
	ActiveDomainCount            uint      `json:"active_domain_count"`
	ActiveDomainLimit            uint      `json:"active_domain_limit"`
	ActiveSharedDomainAliasCount uint      `json:"active_shared_domain_alias_count"`
	ActiveSharedDomainAliasLimit uint      `json:"active_shared_domain_alias_limit"`
	TotalEmailsForwarded         uint      `json:"total_emails_forwarded"`
	TotalEmailsBlocked           uint      `json:"total_emails_blocked"`
	TotalEmailsReplied           uint      `json:"total_emails_replied"`
	TotalEmailsSent              uint      `json:"total_emails_sent"`
	CreatedAt                    AddyTime  `json:"created_at"`
	UpdatedAt                    *AddyTime `json:"updated_at"`
}

type APITokenDetails struct {
	Name      string    `json:"name"`
	CreatedAt AddyTime  `json:"created_at"`
	ExpiresAt *AddyTime `json:"expires_at"`
}

type Alias struct {
	ID              string      `json:"id"`
	UserID          string      `json:"user_id"`
	AliasableID     string      `json:"aliasable_id"`
	AliasableType   string      `json:"aliasable_type"`
	LocalPart       string      `json:"local_part"`
	Extension       string      `json:"extension"`
	Domain          string      `json:"domain"`
	Email           string      `json:"email"`
	Active          bool        `json:"active"`
	Description     string      `json:"description"`
	EmailsForwarded uint        `json:"emails_forwarded"`
	EmailsBlocked   uint        `json:"emails_blocked"`
	EmailsReplied   uint        `json:"emails_replied"`
	EmailsSent      uint        `json:"emails_sent"`
	Recipients      []Recipient `json:"recipients"`
	CreatedAt       AddyTime    `json:"created_at"`
	UpdatedAt       *AddyTime   `json:"updated_at"`
	DeletedAt       *AddyTime   `json:"deleted_at"`
}

// Instead of simply sending the data, addy wraps some of the responses within
// the 'data' JSON field. The following structures are defined to be able to
// parse addy REST API responses.
type AliasResp struct {
	Data Alias `json:"data"`
}
type AliasesResp struct {
	Data []Alias `json:"data"`
}
type DomainResp struct {
	Data Domain `json:"data"`
}
type DomainsResp struct {
	Data []Domain `json:"data"`
}
type FailedDeliveryResp struct {
	Data FailedDelivery `json:"data"`
}
type FailedDeliveriesResp struct {
	Data []FailedDelivery `json:"data"`
}
type RecipientResp struct {
	Data Recipient `json:"data"`
}
type RecipientsResp struct {
	Data []Recipient `json:"data"`
}
type AccountDetailsResp struct {
	Data AccountDetails `json:"data"`
}

type AliasParams struct {
	// Optional:
	// Filter["deleted"] = "with" | "only"
	// Filter["active"] = "true" | "false"
	// Filter["search"] = "search_term"
	Filter     map[string]string
	Sort       AliasSortCond
	PageNumber uint
	PageSize   uint
	// True to get descending sort.
	SortRev bool
}

type RecipientsParams struct {
	// Optional:
	// Filter["verified"] = "true" | "false"
	Filter map[string]string
}
