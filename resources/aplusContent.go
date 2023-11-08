package resources

import (
	"fmt"
	"time"
)

func SearchContentDocuments(params *SellingPartnerParams) error {
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}
	params.Method = "GET"
	params.APIPath = "/aplus/2020-11-01/contentDocuments"
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func CreateContentDocument(params *SellingPartnerParams) error {
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}
	params.Method = "POST"
	params.APIPath = "/aplus/2020-11-01/contentDocuments"
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func GetContentDocument(params *SellingPartnerParams) error {
	if _, present := params.PathParams["contentReferenceKey"]; !present {
		return fmt.Errorf("path param 'contentReferenceKey' not present")
	}
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}
	if present := params.Query.Has("includedDataSet"); !present {
		return fmt.Errorf("query param 'includedDataSet' not present")
	}

	params.Method = "GET"
	params.APIPath = "/aplus/2020-11-01/contentDocuments"
	params.RestoreRate = 100 * time.Millisecond

	return nil
}
