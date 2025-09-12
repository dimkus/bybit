package bybit

import (
	"context"
	"net/url"
	"time"
)

// V5UserServiceI :
type V5UserServiceI interface {
	GetAPIKey(context.Context) (*V5APIKeyResponse, error)
}

// V5UserService :
type V5UserService struct {
	client *Client
}

// V5APIKeyResponse :
type V5APIKeyResponse struct {
	CommonV5Response `json:",inline"`
	Result           V5ApiKeyResult `json:"result"`
}

// V5ApiKeyResult :
type V5ApiKeyResult struct {
	CreatedAt     time.Time `json:"createdAt"`
	ExpiredAt     time.Time `json:"expiredAt"`
	Note          string    `json:"note"`
	APIKey        string    `json:"apiKey"`
	MktMakerLevel string    `json:"mktMakerLevel"`
	Secret        string    `json:"secret"`
	VipLevel      string    `json:"vipLevel"`
	ID            string    `json:"id"`
	Permissions   struct {
		ContractTrade []string `json:"ContractTrade"`
		Spot          []string `json:"Spot"`
		Wallet        []string `json:"Wallet"`
		Options       []string `json:"Options"`
		Derivatives   []string `json:"Derivatives"`
		CopyTrading   []string `json:"CopyTrading"`
		BlockTrade    []string `json:"BlockTrade"`
		Exchange      []string `json:"Exchange"`
		Nft           []string `json:"NFT"`
	} `json:"permissions"`
	Ips         []string `json:"ips"`
	DeadlineDay int      `json:"deadlineDay"`
	Unified     int      `json:"unified"`
	Uta         int      `json:"uta"`
	UserID      int      `json:"userID"`
	InviterID   int      `json:"inviterID"`
	Type        int      `json:"type"`
	ReadOnly    int      `json:"readOnly"`
	AffiliateID int      `json:"affiliateID"`
}

// GetAPIKey :
func (s *V5UserService) GetAPIKey(ctx context.Context) (*V5APIKeyResponse, error) {
	var (
		res V5APIKeyResponse
	)

	if err := s.client.getV5Privately(ctx, "/v5/user/query-api", url.Values{}, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
