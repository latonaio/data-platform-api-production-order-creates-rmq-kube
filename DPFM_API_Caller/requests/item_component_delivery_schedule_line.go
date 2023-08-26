package requests

type ItemComponentDeliveryScheduleLine struct {
	ProductionOrder								int			`json:"ProductionOrder"`
	ProductionOrderItem							int			`json:"ProductionOrderItem"`
	BillOfMaterial								int			`json:"BillOfMaterial"`
	BillOfMaterialItem							int			`json:"BillOfMaterialItem"`
	ScheduleLine								int			`json:"ScheduleLine"`
	SupplyChainRelationshipID					int			`json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID		int			`json:"SupplyChainRelationshipStockConfPlantID"`
	ComponentProduct							string		`json:"ComponentProduct"`
	StockConfirmationBusinessPartner			int			`json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant						string		`json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone				*string		`json:"StockConfirmationPlantTimeZone"`
	ComponentProductBatch						*string		`json:"ComponentProductBatch"`
	ComponentProductBatchValidityStartDate		*string		`json:"ComponentProductBatchValidityStartDate"`
	ComponentProductBatchValidityStartTime		*string		`json:"ComponentProductBatchValidityStartTime"`
	ComponentProductBatchValidityEndDate		*string		`json:"ComponentProductBatchValidityEndDate"`
	ComponentProductBatchValidityEndTime		*string		`json:"ComponentProductBatchValidityEndTime"`
	RequestedDeliveryDate						string		`json:"RequestedDeliveryDate"`
	RequestedDeliveryTime						string		`json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate						*string		`json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryTime						*string		`json:"ConfirmedDeliveryTime"`
	OriginalRequiredQuantityInBaseUnit			float32		`json:"OriginalRequiredQuantityInBaseUnit"`
	ConfirmedQuantityByPDTAvailCheckInBaseUnit	*float32	`json:"ConfirmedQuantityByPDTAvailCheckInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit				*float32	`json:"OpenConfirmedQuantityInBaseUnit"`
	DeliveredQuantityInBaseUnit					*float32	`json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit				*float32	`json:"UndeliveredQuantityInBaseUnit"`
	StockIsFullyConfirmed						*bool		`json:"StockIsFullyConfirmed"`
	PlusMinusFlag								*string		`json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus			*bool		`json:"ItemScheduleLineDeliveryBlockStatus"`
	CreationDate                                string  	`json:"CreationDate"`
	CreationTime                                string  	`json:"CreationTime"`
	LastChangeDate                              string  	`json:"LastChangeDate"`
	LastChangeTime                              string  	`json:"LastChangeTime"`
	IsReleased									*bool		`json:"IsReleased"`
	IsLocked									*bool		`json:"IsLocked"`
	IsCancelled									*bool		`json:"IsCancelled"`
	IsMarkedForDeletion							*bool		`json:"IsMarkedForDeletion"`
}
