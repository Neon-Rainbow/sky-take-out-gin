package DTO

type AddressBookRequestDTO struct {
	CityCode     string `json:"cityCode"`
	CityName     string `json:"cityName"`
	Consignee    string `json:"consignee"`
	Detail       string `json:"detail"`
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
	IsDefault    bool   `json:"isDefault"`
	Label        string `json:"label"`
	Phone        string `json:"phone"`
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
	Sex          string `json:"sex"`
	UserID       uint   `json:"userId"`
}

type AddAddressBookResponseDTO struct {
}
