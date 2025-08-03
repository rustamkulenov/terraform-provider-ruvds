package api

type NetworkV4 struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
}

type VirtualServer struct {
	ID                      int32       `json:"virtual_server_id"`
	Status                  string      `json:"status"`
	CreateProgress          int32       `json:"create_progress"`
	DataCenterID            int32       `json:"datacenter"`
	TariffID                int32       `json:"tariff_id"`
	PaymentPeriod           int32       `json:"payment_period"`
	OSID                    int32       `json:"os_id"`
	TemplateID              *string     `json:"template_id,omitempty"`
	CPU                     int32       `json:"cpu"`
	RAM                     float64     `json:"ram"`
	VRAM                    int32       `json:"vram"`
	Drive                   int32       `json:"drive"`
	DriveTariffID           int32       `json:"drive_tariff_id"`
	AdditionalDrive         *int32      `json:"additional_drive,omitempty"`
	AdditionalDriveTariffID *int32      `json:"additional_drive_tariff_id,omitempty"`
	IP                      int32       `json:"ip"`
	DDOSProtection          float32     `json:"ddos_protection"`
	UserComment             *string     `json:"user_comment,omitempty"`
	PaidTill                string      `json:"paid_till"`
	NetworkV4               []NetworkV4 `json:"network_v4"`
}

type VirtualServersResponse struct {
	VirtualServers []VirtualServer `json:"servers"`
}
