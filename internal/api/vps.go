package api

type NetworkV4 struct {
	IPAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
}

type VirtualServer struct {
	ID                      int32       `json:"virtual_server_id"`
	Status                  *string     `json:"status,omitempty"`
	CreateProgress          int32       `json:"create_progress"`
	DataCenterID            int32       `json:"datacenter"`
	TariffID                int32       `json:"tariff_id"`
	PaymentPeriod           int32       `json:"payment_period"`
	OSID                    int32       `json:"os_id"`
	TemplateID              *string     `json:"template_id,omitempty"`
	CPU                     int32       `json:"cpu"`
	RAM                     float32     `json:"ram"`
	VRAM                    int32       `json:"vram"`
	Drive                   int32       `json:"drive"`
	DriveTariffID           int32       `json:"drive_tariff_id"`
	AdditionalDrive         *int32      `json:"additional_drive,omitempty"`
	AdditionalDriveTariffID *int32      `json:"additional_drive_tariff_id,omitempty"`
	IP                      int32       `json:"ip"`
	DDOSProtection          float32     `json:"ddos_protection"`
	UserComment             *string     `json:"user_comment,omitempty"`
	PaidTill                string      `json:"paid_till"`
	SShKeyID                *string     `json:"ssh_key_id,omitempty"`
	ComputerName            *string     `json:"computer_name,omitempty"`
	NetworkV4               []NetworkV4 `json:"network_v4"`
}

type VirtualServersResponse struct {
	VirtualServers []VirtualServer `json:"servers"`
}

// CreateVpsRequest creates a new virtual server request with the minimum list of mandatory parameters.
func CreateVpsRequest(datacenterId int32, tariffId int32, paymentPeriod int32, osId int32, cpu int32, ram float32, drive int32, driveTariffId int32, ip int32) VirtualServer {
	return VirtualServer{
		DataCenterID:  datacenterId,
		TariffID:      tariffId,
		PaymentPeriod: paymentPeriod,
		OSID:          osId,
		CPU:           cpu,
		RAM:           ram,
		Drive:         drive,
		DriveTariffID: driveTariffId,
		IP:            ip,
	}
}

type CreateVpsErrorResponse struct {
	Message         string `json:"message"`
	Id              string `json:"id"`
	UserId          string `json:"user_id,omitempty"`
	TwoFactorId     string `json:"two_factor_id,omitempty"`
	TwoFactorSecret string `json:"two_factor_secret,omitempty"`
	TwoFactorOtp    string `json:"two_factor_otp,omitempty"`
	TwoFactorSms    string `json:"two_factor_sms,omitempty"`
	TwoFactorEmail  string `json:"two_factor_email,omitempty"`
}

type CreateVpsOkResponse struct {
	VirtualServerId int32         `json:"virtual_server_id"`
	PaymentPeriod   int32         `json:"payment_period"`
	CostRub         float64       `json:"cost_rub"`
	Password        string        `json:"password"`
	Status          VirtualServer `json:"status"`
}
