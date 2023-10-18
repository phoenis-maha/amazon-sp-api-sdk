package resources

import (
	"fmt"
	"time"
)

func GetListingsItem(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerId"]; !present {
		return fmt.Errorf("path param 'sellerId' not present")
	}
	if _, present := params.PathParams["sku"]; !present {
		return fmt.Errorf("path param 'sku' not present")
	}
	if present := params.Query.Has("marketplaceIds"); !present {
		return fmt.Errorf("query param 'marketplaceIds' not present")
	}
	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/listings/2021-08-01/items/%s/%s", params.PathParams["sellerId"], params.PathParams["sku"])
	params.RestoreRate = 1 * time.Second

	return nil
}

func PutListingsItem(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerId"]; !present {
		return fmt.Errorf("path param 'sellerId' not present")
	}
	if _, present := params.PathParams["sku"]; !present {
		return fmt.Errorf("path param 'sku' not present")
	}
	if present := params.Query.Has("marketplaceIds"); !present {
		return fmt.Errorf("query param 'marketplaceIds' not present")
	}
	params.Method = "PUT"
	params.APIPath = fmt.Sprintf("/listings/2021-08-01/items/%s/%s", params.PathParams["sellerId"], params.PathParams["sku"])
	params.RestoreRate = 1 * time.Second

	return nil
}

func PatchListingsItem(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerId"]; !present {
		return fmt.Errorf("path param 'sellerId' not present")
	}
	if _, present := params.PathParams["sku"]; !present {
		return fmt.Errorf("path param 'sku' not present")
	}
	if present := params.Query.Has("marketplaceIds"); !present {
		return fmt.Errorf("query param 'marketplaceIds' not present")
	}
	params.Method = "PATCH"
	params.APIPath = fmt.Sprintf("/listings/2021-08-01/items/%s/%s", params.PathParams["sellerId"], params.PathParams["sku"])
	params.RestoreRate = 1 * time.Second

	return nil
}

func DeleteListingsItem(params *SellingPartnerParams) error {
	if _, present := params.PathParams["sellerId"]; !present {
		return fmt.Errorf("path param 'sellerId' not present")
	}
	if _, present := params.PathParams["sku"]; !present {
		return fmt.Errorf("path param 'sku' not present")
	}
	if present := params.Query.Has("marketplaceIds"); !present {
		return fmt.Errorf("query param 'marketplaceIds' not present")
	}
	params.Method = "DELETE"
	params.APIPath = fmt.Sprintf("/listings/2021-08-01/items/%s/%s", params.PathParams["sellerId"], params.PathParams["sku"])
	params.RestoreRate = 1 * time.Second

	return nil
}
