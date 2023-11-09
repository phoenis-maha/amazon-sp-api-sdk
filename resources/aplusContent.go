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
	params.APIPath = fmt.Sprintf("/aplus/2020-11-01/contentDocuments/%s", params.PathParams["contentReferenceKey"])
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func UpdateContentDocument(params *SellingPartnerParams) error {
	if present := params.Query.Has("contentReferenceKey"); !present {
		return fmt.Errorf("query param 'contentReferenceKey' not present")
	}
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}
	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/aplus/2020-11-01/contentDocuments/%s", params.PathParams["contentReferenceKey"])
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func ListContentDocumentAsinRelations(params *SellingPartnerParams) error {
	if _, present := params.PathParams["contentReferenceKey"]; !present {
		return fmt.Errorf("path param 'contentReferenceKey' not present")
	}
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}
	params.Method = "GET"
	params.APIPath = fmt.Sprintf("/aplus/2020-11-01/contentDocuments/%s/asins", params.PathParams["contentReferenceKey"])
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func PostContentDocumentAsinRelations(params *SellingPartnerParams) error {
	if _, present := params.PathParams["contentReferenceKey"]; !present {
		return fmt.Errorf("path param 'contentReferenceKey' not present")
	}
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/aplus/2020-11-01/contentDocuments/%s/asins", params.PathParams["contentReferenceKey"])
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func ValidateContentDocumentAsinRelations(params *SellingPartnerParams) error {
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/aplus/2020-11-01/contentAsinValidations"
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func SearchContentPublishRecords(params *SellingPartnerParams) error {
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}
	if present := params.Query.Has("asin"); !present {
		return fmt.Errorf("query param 'asin' not present")
	}

	params.Method = "GET"
	params.APIPath = "/aplus/2020-11-01/contentPublishRecords"
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func PostContentDocumentApprovalSubmission(params *SellingPartnerParams) error {
	if _, present := params.PathParams["contentReferenceKey"]; !present {
		return fmt.Errorf("path param 'contentReferenceKey' not present")
	}
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/aplus/2020-11-01/contentDocuments/%s/approvalSubmissions", params.PathParams["contentReferenceKey"])
	params.RestoreRate = 100 * time.Millisecond

	return nil
}

func PostContentDocumentSuspendSubmission(params *SellingPartnerParams) error {
	if _, present := params.PathParams["contentReferenceKey"]; !present {
		return fmt.Errorf("path param 'contentReferenceKey' not present")
	}
	if present := params.Query.Has("marketplaceId"); !present {
		return fmt.Errorf("query param 'marketplaceId' not present")
	}

	params.Method = "POST"
	params.APIPath = fmt.Sprintf("/aplus/2020-11-01/contentDocuments/%s/suspendSubmissions", params.PathParams["contentReferenceKey"])
	params.RestoreRate = 100 * time.Millisecond

	return nil
}
