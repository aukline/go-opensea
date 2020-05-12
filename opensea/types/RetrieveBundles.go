package types

type RetrieveBundlesRequest struct {
	OnSale                 string   `json:"on_sale"` // empty string, true, or false
	Owner                  string   `json:"owner"`
	AssetContractAddress   string   `json:"asset_contract_address"`
	AssetContractAddresses []string `json:"asset_contract_addresses"`
	TokenIDs               []string `json:"token_ids"`
	Limit                  string   `json:"limit"`
	Offset                 string   `json:"offset"`
}

type RetrieveBundlesResponse struct {
	Bundles []Bundle `json:"bundles"`
}

type Bundle struct {
	// TODO: need an example for sell_orders
	Maker         Owner         `json:"maker"`
	Slug          string        `json:"slug"`
	Assets        []Asset       `json:"assets"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	ExternalLink  string        `json:"external_link"`
	AssetContract AssetContract `json:"asset_contract"`
	PermaLink     string        `json:"perma_link"`
}
