package resources

import "time"

func GetScheduledPackage(params *SellingPartnerParams) error {
	params.Method = "GET"
	params.APIPath = "/easyShip/2022-03-23/package"
	params.RestoreRate = 1 * time.Second

	return nil
}

func CreateScheduledPackage(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/easyShip/2022-03-23/package"
	params.RestoreRate = 1 * time.Second

	return nil
}

func CreateScheduledPackageBulk(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/easyShip/2022-03-23/packages/bulk"
	params.RestoreRate = 1 * time.Second

	return nil
}

func UpdateScheduledPackages(params *SellingPartnerParams) error {
	params.Method = "PATCH"
	params.APIPath = "/easyShip/2022-03-23/package"
	params.RestoreRate = 1 * time.Second

	return nil
}

func ListHandoverSlots(params *SellingPartnerParams) error {
	params.Method = "POST"
	params.APIPath = "/easyShip/2022-03-23/timeSlot"
	params.RestoreRate = 1 * time.Second

	return nil
}
