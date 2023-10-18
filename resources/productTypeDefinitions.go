package resources

import (
	"fmt"
	"time"
)

func SearchDefinitionsProductTypes(params *SellingPartnerParams) error {
	if present := params.Query.Has("marketplaceIds"); !present {
		return fmt.Errorf("query param 'marketplaceIds' not present")
	}
	params.Method = "GET"
	params.APIPath = "/definitions/2020-09-01/productTypes"
	params.RestoreRate = 1 * time.Second

	return nil
}

func GetDefinitionsProductType(params *SellingPartnerParams) error {
	if _, present := params.PathParams["productType"]; !present {
		return fmt.Errorf("path param 'productType' not present")
	}
	if present := params.Query.Has("marketplaceIds"); !present {
		return fmt.Errorf("query param 'marketplaceIds' not present")
	}
	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/definitions/2020-09-01/productTypes/%s", params.PathParams["productType"])
	params.RestoreRate = 1 * time.Second

	return nil
}
