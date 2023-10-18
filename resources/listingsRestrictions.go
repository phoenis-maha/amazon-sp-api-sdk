package resources

import (
	"fmt"
	"time"
)

func GetListingsRestrictions(params *SellingPartnerParams) error {
	if present := params.Query.Has("asin"); !present {
		return fmt.Errorf("query param 'asin' not present")
	}
	params.Method = "GET"
	params.APIPath = "/listings/2021-08-01/restrictions"
	params.RestoreRate = 200 * time.Millisecond

	return nil
}
