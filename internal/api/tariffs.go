package api

// Example response (json):
// {"vps":[{"id":14,"name":"Regular","cpu":135.0,"ram":324.0,"vram":1.953125,"ip":180.0,"is_active":true},{"id":15,"name":"Premium","cpu":311.0,"ram":324.0,"vram":1.953125,"ip":180.0,"is_active":true},{"id":21,"name":"HugeServer","cpu":130.0,"ram":307.0,"vram":1.953125,"ip":180.0,"is_active":true},{"id":22,"name":"Promo32","cpu":106.0,"ram":252.0,"vram":1.953125,"ip":180.0,"is_active":true},{"id":26,"name":"Powerful","cpu":375.0,"ram":375.0,"vram":1.953125,"ip":180.0,"is_active":true},{"id":40,"name":"RegularEurope","cpu":154.0,"ram":367.0,"vram":1.953125,"ip":180.0,"is_active":true},{"id":41,"name":"PremiumEurope","cpu":356.0,"ram":367.0,"vram":1.953125,"ip":180.0,"is_active":true}],"drive":[{"id":1,"name":"HDD","price":9.0,"is_active":true},{"id":3,"name":"SSD","price":18.5,"is_active":true},{"id":7,"name":"NVMe","price":21.5,"is_active":true},{"id":9,"name":"HDDEurope","price":10.5,"is_active":true},{"id":10,"name":"SSDEurope","price":21.5,"is_active":true}],"additional_drive":[{"id":4,"name":"Huge HDD","price":0.634765625,"is_active":true}],"additional_service":[{"id":1,"name":"Remote Desktop Services Client Access Licenses","price":1081.0,"is_active":true},{"id":2,"name":"Microsoft SQL Server Standard MVL SAL ","price":2642.0,"is_active":true},{"id":4,"name":"DDoS Protection (IP)","price":400.0,"is_active":true},{"id":6,"name":"Гб (бэкап образа диска сервера)","price":6.5,"is_active":true},{"id":7,"name":"Microsoft Office Standard MVL SAL","price":2099.0,"is_active":true},{"id":8,"name":"Microsoft Office Professional Plus\r\n MVL SAL","price":2865.0,"is_active":true},{"id":10,"name":"Microsoft SQL Server Standard 2Lic Core","price":24462.0,"is_active":true},{"id":11,"name":"Microsoft SQL Server Web 2Lic Core","price":1389.0,"is_active":true},{"id":12,"name":"ISPmanager 6 Lite","price":200.0,"is_active":true},{"id":17,"name":"Антивирус Касперского","price":800.0,"is_active":true},{"id":19,"name":"Plesk for VPS Web Admin Edition","price":0.0,"is_active":true},{"id":23,"name":"Снапшот","price":0.8,"is_active":true},{"id":27,"name":"1C License","price":1595.0,"is_active":true}],"payment_period_discount":[{"payment_period":3,"discount":0.05},{"payment_period":4,"discount":0.1},{"payment_period":5,"discount":0.2}]}

type VpsTariff struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	CPU      float32 `json:"cpu"`
	RAM      float32 `json:"ram"`
	VRAM     float32 `json:"vram"`
	IP       float32 `json:"ip"`
	IsActive bool    `json:"is_active"`
}

type DriveTariff struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	IsActive bool    `json:"is_active"`
}

type AdditionalDriveTariff struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	IsActive bool    `json:"is_active"`
}

type AdditionalServiceTariff struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	IsActive bool    `json:"is_active"`
}

type PaymentPeriodDiscount struct {
	PaymentPeriod int32   `json:"payment_period"`
	Discount      float32 `json:"discount"`
}

type TariffsResponse struct {
	VpsTariffs               []VpsTariff               `json:"vps"`
	DriveTariffs             []DriveTariff             `json:"drive"`
	AdditionalDriveTariffs   []AdditionalDriveTariff   `json:"additional_drive"`
	AdditionalServiceTariffs []AdditionalServiceTariff `json:"additional_service"`
	PaymentPeriodDiscounts   []PaymentPeriodDiscount   `json:"payment_period_discount"`
}
