package existence_conf

type Returns struct {
	ConnectionKey              string                     `json:"connection_key"`
	Result                     bool                       `json:"result"`
	RedisKey                   string                     `json:"redis_key"`
	RuntimeSessionID           string                     `json:"runtime_session_id"`
	BusinessPartnerID          *int                       `json:"business_partner"`
	Filepath                   string                     `json:"filepath"`
	ServiceLabel               string                     `json:"service_label"`
	ProductMasterGeneralReturn ProductMasterGeneralReturn `json:"ProductMasterGeneral"`
	BPGeneralReturn            BPGeneralReturn            `json:"BusinessPartnerGeneral"`
	PlantGeneralReturn         PlantGeneralReturn         `json:"PlantGeneral"`
	CurrencyReturn             CurrencyReturn             `json:"Currency"`
	BatchReturn                BatchReturn                `json:"Batch"`
	StorageLocationReturn      StorageLocationReturn      `json:"StorageLocation"`
	APISchema                  string                     `json:"api_schema"`
	Accepter                   []string                   `json:"accepter"`
	Deleted                    bool                       `json:"deleted"`
}

type ProductMasterGeneralReturn struct {
	Product string `json:"Product"`
}

type BPGeneralReturn struct {
	BusinessPartner int `json:"BusinessPartner"`
}

type PlantGeneralReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
}

type CurrencyReturn struct {
	Currency string `json:"Currency"`
}

type BatchReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Product         string `json:"Product"`
	Plant           string `json:"Plant"`
	Batch           string `json:"Batch"`
}

type StorageLocationReturn struct {
	BusinessPartner int    `json:"BusinessPartner"`
	Plant           string `json:"Plant"`
	StorageLocation string `json:"StorageLocation"`
}
