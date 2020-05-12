package opensea

import (
	"fmt"
	"testing"

	. "github.com/stretchr/testify/assert"

	"austinkline/go-opensea/opensea/types"
)

const (
	ContractSandbox = "0x50f5474724e0ee42d9a4e711ccfb275809fd6d4a"
	ContractGU = "0x0e3a2a1f2146d86a604adc220b4967a898d7fe07"
)

func TestClient_RetrieveBundles_By_Owner(t *testing.T) {
	limit := 1
	owner := "0xd387a6e4e84a6c86bd90c158c6028a58cc8ac459"
	req := types.RetrieveBundlesRequest{Limit: fmt.Sprintf("%d", limit), Owner: owner}
	res, err := c.RetrieveBundles(req)
	Nil(t, err)
	NotNil(t, res)
	Equal(t, limit, len(res.Bundles))
}

func TestClient_RetrieveBundles_By_Contract_Address(t *testing.T) {
	req := types.RetrieveBundlesRequest{AssetContractAddress: ContractSandbox}
	res, err := c.RetrieveBundles(req)
	Nil(t, err)
	for _, bundle := range res.Bundles {
		Equal(t, ContractSandbox, bundle.AssetContract.Address)
	}
}

func TestClient_RetrieveBundles_By_Contract_Addresses(t *testing.T) {
	contractAddresses := []string{ContractGU, ContractSandbox}

	req := types.RetrieveBundlesRequest{AssetContractAddresses: contractAddresses}
	res, err := c.RetrieveBundles(req)
	Nil(t, err)
	for _, bundle := range res.Bundles {
		Contains(t, contractAddresses, bundle.AssetContract.Address)
	}
}
