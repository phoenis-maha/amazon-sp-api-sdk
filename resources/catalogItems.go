package resources

import (
	"fmt"
	"time"
)

func SearchCatalogItems(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/catalog/2022-04-01/items"
	params.RestoreRate = 1 * time.Second

	return nil
}

func GetCatalogItem(params *SellingPartnerParams) error {
	if _, present := params.PathParams["asin"]; !present {
		return fmt.Errorf("path param 'asin' not present")
	}
	params.Method = "GET"
	params.APIPath = "/catalog/2022-04-01/items/" + params.PathParams["asin"]
	params.RestoreRate = 1 * time.Second

	return nil
}