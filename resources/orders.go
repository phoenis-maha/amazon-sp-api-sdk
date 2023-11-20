package resources

import (
	"fmt"
	"time"
)

func GetOrders(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/orders/v0/orders"
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrder(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"]
	params.RestoreRate = 1 * time.Second
	return nil
}

func GetOrderBuyerInfo(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/buyerInfo"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetOrderAddress(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/address"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetOrderItems(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/orderItems"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func GetOrderItemsBuyerInfo(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/orderItems/buyerInfo"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func UpdateShipmentStatus(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/shipment"
	params.RestoreRate = 200 * time.Millisecond
	return nil
}

func GetOrderRegulatedInfo(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "GET"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/regulatedInfo"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func UpdateVerificationStatus(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "PATCH"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/regulatedInfo"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}

func ConfirmShipment(params *SellingPartnerParams) error {
	if _, present := params.PathParams["orderId"]; !present {
		return fmt.Errorf("path param 'orderId' not present")
	}

	params.Method = "POST"
	params.APIPath = "/orders/v0/orders/" + params.PathParams["orderId"] + "/shipmentConfirmation"
	params.RestoreRate = 500 * time.Millisecond
	return nil
}
