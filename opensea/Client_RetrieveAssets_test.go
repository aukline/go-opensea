package opensea

import (
	"testing"

	. "github.com/stretchr/testify/assert"

	"austinkline/go-opensea/opensea/types"
)

var (
	c = NewClient(BaseURLV1)
)

func TestClient_RetrieveAssets(t *testing.T) {
	offset := "0"
	limit := "1"
	request := types.RetrieveAssetsRequest{Offset: offset, Limit: limit}
	res, err := c.RetrieveAssets(request)
	Nil(t, err)
	NotNil(t, res)
	Equal(t, 1, len(res.Assets))
}

func TestClient_RetrieveAssets_By_Contract_Address_And_TokenID(t *testing.T) {
	tokenID := "41469"
	contractAddress := "0x50f5474724e0ee42d9a4e711ccfb275809fd6d4a"
	request := types.RetrieveAssetsRequest{TokenIDs: []string{tokenID}, AssetContractAddress: contractAddress}
	res, err := c.RetrieveAssets(request)
	Nil(t, err)
	Equal(t, 1, len(res.Assets))
	Equal(t, tokenID, res.Assets[0].TokenID)
	Equal(t, contractAddress, res.Assets[0].AssetContract.Address)
}
