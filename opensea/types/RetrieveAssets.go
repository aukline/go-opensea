package types

import (
	"strings"
	"time"
)

type RetrieveAssetsRequest struct {
	Owner                  string   `json:"owner"`
	TokenIDs               []string `json:"token_ids"`
	AssetContractAddress   string   `json:"asset_contract_address"`
	AssetContractAddresses []string `json:"asset_contract_addresses"`
	OrderBy                string   `json:"order_by"`
	OrderDirection         string   `json:"order_direction"` // desc or asc
	OnSale                 string   `json:"on_sale"`         // true for only items which are on sale, false for items which are not
	Offset                 string   `json:"offset"`
	Limit                  string   `json:"limit"`
	Collection             string   `json:"collection"`
}

type RetrieveAssetsResponse struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	// TODO: background_color, animation_url, animation_original_url
	// TODO: auctions, last_sale, top_bid, current_price, current_escrow_price
	// TODO: listing_date, transfer_fee_payment_token, transfer_fee, sell_orders
	TokenID           string        `json:"token_id"`
	NumSales          int           `json:"num_sales"`
	ImageURL          string        `json:"image_url"`
	ImagePreviewURL   string        `json:"image_preview_url"`
	ImageThumbnailURL string        `json:"image_thumbnail_url"`
	ImageOriginalURL  string        `json:"image_original_url"`
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	ExternalLink      string        `json:"external_link"`
	AssetContract     AssetContract `json:"asset_contract"`
	Owner             Owner         `json:"owner"`
	Permalink         string        `json:"permalink"`
	Traits            []Trait       `json:"traits"`
	Decimals          int           `json:"decimals"`
	IsPresale         bool          `json:"is_presale"`
}

type AssetContract struct {
	// TODO: opensea_version
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 OpenSeaTime `json:"created_date"`
	Name                        string      `json:"name"`
	NFTVersion                  string      `json:"nft_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 string      `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                string      `json:"external_link"`
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
	Collection                  Collection  `json:"collection"`
}

type Owner struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
	DiscordID     string `json:"discord_id"`
}

type User struct {
	Username string `json:"username"`
}

type Trait struct {
	// TODO need values for display_type, max_value, order
	TraitType  string      `json:"trait_type"`
	Value      interface{} `json:"value"`
	TraitCount int64       `json:"trait_count"`
}

type Collection struct {
	// TODO: banner_image_url, chat_url, wiki_url
	CreatedDate                 OpenSeaTime `json:"created_date"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	Description                 string      `json:"description"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string      `json:"dev_seller_fee_basis_points"`
	ExternalURL                 string      `json:"external_url"`
	Featured                    bool        `json:"featured"`
	FeaturedImgURL              string      `json:"featured_img_url"`
	Hidden                      bool        `json:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status"`
	ImageURL                    string      `json:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	LargeImageURL               string      `json:"large_image_url"`
	Name                        string      `json:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
	RequireEmail                bool        `json:"require_email"`
	ShortDescription            string      `json:"short_description"`
	Slug                        string      `json:"slug"`
	PayoutAddress               string      `json:"payout_address"`
	DisplayData                 DisplayData `json:"display_data"`
}

type DisplayData struct {
	Images           []string `json:"images"`
	CardDisplayStyle string   `json:"card_display_style"`
}

type OpenSeaTime struct {
	time.Time
}

func (ost *OpenSeaTime) UnmarshalJSON(input []byte) error {
	layout := "2006-01-02T15:04:05.000000"
	trimmed := strings.Trim(string(input), `"`)
	t, err := time.Parse(layout, trimmed)
	if err != nil {
		return err
	}
	ost.Time = t
	return nil
}
