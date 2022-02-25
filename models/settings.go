package models

type Settings struct {
	AppParams      Params           `json:"app"`
	PostgresParams PostgresSettings `json:"postgres_params"`
}

type Params struct {
	ServerURL         string `json:"server_url"`
	Login             string `json:"login"`
	OfferLink         string `json:"offer_link"`
	ServerName        string `json:"server_name"`
	AppVersion        string `json:"app_version"`
	PortRun           string `json:"port_run"`
	LogInfo           string `json:"log_info"`
	LogError          string `json:"log_error"`
	LogDebug          string `json:"log_debug"`
	LogWarning        string `json:"log_warning"`
	LogMachineHWID    string `json:"log_machine_hw_id"`
	LogMaxSize        int    `json:"log_max_size"`
	LogMaxBackups     int    `json:"log_max_backups"`
	LogMaxAge         int    `json:"log_max_age"`
	LogCompress       bool   `json:"log_compress"`
	AuthServiceURL    string `json:"auth_service_url"`
	SecretKey         string `json:"secret_key"`
	PaymentServiceURL string `json:"payment_service_url"`
}

type PostgresSettings struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     int    `json:"port"`
	DataBase string `json:"database"`
}
