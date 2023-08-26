package dpfm_api_processing_formatter

type HeaderUpdates struct {
	ProductionOrder                                    int      `json:"ProductionOrder"`
	OwnerProductionPlantStorageLocation				   string	`json:"OwnerProductionPlantStorageLocation"`
	DepartureDeliverFromPlantStorageLocation           string   `json:"DepartureDeliverFromPlantStorageLocation"`
	DestinationDeliverToPlantStorageLocation           string   `json:"DestinationDeliverToPlantStorageLocation"`
	ProductionOrderQuantityInBaseUnit                  float32  `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInDepartureProductionUnit   float32  `json:"ProductionOrderQuantityInDepartureProductionUnit"`
	ProductionOrderQuantityInDestinationProductionUnit float32  `json:"ProductionOrderQuantityInDestinationProductionUnit"`
	ProductionOrderQuantityInDepartureDeliveryUnit     float32  `json:"ProductionOrderQuantityInDepartureDeliveryUnit"`
	ProductionOrderQuantityInDestinationDeliveryUnit   float32  `json:"ProductionOrderQuantityInDestinationDeliveryUnit"`
	ProductionOrderPlannedScrapQtyInBaseUnit           *float32 `json:"ProductionOrderPlannedScrapQtyInBaseUnit"`
	ProductionOrderPlannedStartDate                    string   `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime                    string   `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                      string   `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                      string   `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualStartDate                     *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime                     *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate                       *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime                       *string  `json:"ProductionOrderActualEndTime"`
	ProductionOrderHeaderText                          *string  `json:"ProductionOrderHeaderText"`
	IsLocked                                           *bool    `json:"IsLocked"`
}

type HeaderDocUpdates struct {
	ProductionOrder int     `json:"ProductionOrder"`
	DocType         string  `json:"DocType"`
	DocVersionID    int     `json:"DocVersionID"`
	DocID           string  `json:"DocID"`
	FileName        *string `json:"FileName"`
	FilePath        *string `json:"FilePath"`
}

type ItemUpdates struct {
	ProductionOrder                               int      `json:"ProductionOrder"`
	ProductionOrderItem                           int      `json:"ProductionOrderItem"`
	ProductionPlantStorageLocation                string   `json:"ProductionPlantStorageLocation"`
	DeliverFromPlantStorageLocation               string   `json:"DeliverFromPlantStorageLocation"`
	DeliverToPlantStorageLocation                 string   `json:"DeliverToPlantStorageLocation"`
	ProductionOrderQuantityInBaseUnit             float32  `json:"ProductionOrderQuantityInBaseUnit"`
	ProductionOrderQuantityInProductionUnit       float32  `json:"ProductionOrderQuantityInProductionUnit"`
	ProductionOrderQuantityInDeliveryUnit         float32  `json:"ProductionOrderQuantityInDeliveryUnit"`
	ProductionOrderPlannedScrapQtyInBaseUnit      *float32 `json:"ProductionOrderPlannedScrapQtyInBaseUnit"`
	ProductionOrderPlannedStartDate               string   `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime               string   `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate                 string   `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime                 string   `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualStartDate                *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime                *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate                  *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime                  *string  `json:"ProductionOrderActualEndTime"`
	ProductionOrderItemText                       *string  `json:"ProductionOrderItemText"`
	IsLocked                                      *bool    `json:"IsLocked"`
}

type ItemComponentUpdates struct {
	ProductionOrder                                int      `json:"ProductionOrder"`
	ProductionOrderItem                            int      `json:"ProductionOrderItem"`
	BillOfMaterial                                 int      `json:"BillOfMaterial"`
	BillOfMaterialItem                             int      `json:"BillOfMaterialItem"`
	ComponentProductRequirementDate                string   `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime                string   `json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantityInBaseUnit     float32  `json:"ComponentProductRequiredQuantityInBaseUnit"`
	ComponentProductRequiredQuantityInDeliveryUnit float32  `json:"ComponentProductRequiredQuantityInDeliveryUnit"`
	ComponentProductPlannedScrapInPercent          *float32 `json:"ComponentProductPlannedScrapInPercent"`
	IsLocked                                       *bool    `json:"IsLocked"`
}

type ItemComponentDeliveryScheduleLineUpdates struct {
	ProductionOrder                            int      `json:"ProductionOrder"`
	ProductionOrderItem                        int      `json:"ProductionOrderItem"`
	BillOfMaterial                             int      `json:"BillOfMaterial"`
	BillOfMaterialItem                         int      `json:"BillOfMaterialItem"`
	ScheduleLine                               int      `json:"ScheduleLine"`
	ItemScheduleLineDeliveryBlockStatus        *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	IsLocked                                   *bool    `json:"IsLocked"`
}

type ItemComponentCostingUpdates struct {
	ProductionOrder     int     `json:"ProductionOrder"`
	ProductionOrderItem int     `json:"ProductionOrderItem"`
	BillOfMaterial      int     `json:"BillOfMaterial"`
	BillOfMaterialItem  int     `json:"BillOfMaterialItem"`
	IsLocked            *bool   `json:"IsLocked"`
}

type ItemDocUpdates struct {
	ProductionOrder          int     `json:"ProductionOrder"`
	ProductionOrderItem      int     `json:"ProductionOrderItem"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            string  `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type ItemOperationUpdates struct {
	ProductionOrder                                 int      `json:"ProductionOrder"`
	ProductionOrderItem                             int      `json:"ProductionOrderItem"`
	Operations                                      int      `json:"Operations"`
	OperationsItem                                  int      `json:"OperationsItem"`
	OperationID                                     int      `json:"OperationID"`
	SequenceText                                    *string  `json:"SequenceText"`
	OperationText                                   string   `json:"OperationText"`
	OperationPlannedQuantityInBaseUnit              float32  `json:"OperationPlannedQuantityInBaseUnit"`
	OperationPlannedQuantityInProductionUnit        float32  `json:"OperationPlannedQuantityInProductionUnit"`
	OperationPlannedQuantityInOperationUnit         float32  `json:"OperationPlannedQuantityInOperationUnit"`
	OperationPlannedQuantityInDeliveryUnit          float32  `json:"OperationPlannedQuantityInDeliveryUnit"`
	OperationPlannedScrapInPercent                  *float32 `json:"OperationPlannedScrapInPercent"`
	OperationErlstSchedldExecStrtDte                *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme                *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDte                 *string  `json:"OperationErlstSchedldExecEndDte"`
	OperationErlstSchedldExecEndTme                 *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate               *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime               *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate                 *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime                 *string  `json:"OperationActualExecutionEndTime"`
	IsLocked                                        *bool    `json:"IsLocked"`
}

type ItemOperationComponentUpdates struct {
	ProductionOrder                                          int      `json:"ProductionOrder"`
	ProductionOrderItem                                      int      `json:"ProductionOrderItem"`
	Operations                                               int      `json:"Operations"`
	OperationsItem                                           int      `json:"OperationsItem"`
	OperationID                                              int      `json:"OperationID"`
	BillOfMaterial                                           int      `json:"BillOfMaterial"`
	BillOfMaterialItem                                       int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                                int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID                        int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID                   int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	ProductionPlantBusinessPartner                           int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                          string   `json:"ProductionPlant"`
	ComponentProduct                                         string   `json:"ComponentProduct"`
	ComponentProductBuyer                                    int      `json:"ComponentProductBuyer"`
	ComponentProductSeller                                   int      `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                           int      `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                           string   `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty                         int      `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                         string   `json:"ComponentProductDeliverFromPlant"`
	ComponentProductDeliverToPartyInOperation                int      `json:"ComponentProductDeliverToPartyInOperation"`
	ComponentProductDeliverToPlantInOperation                string   `json:"ComponentProductDeliverToPlantInOperation"`
	ComponentProductDeliverFromPartyInOperation              int      `json:"ComponentProductDeliverFromPartyInOperation"`
	ComponentProductDeliverFromPlantInOperation              string   `json:"ComponentProductDeliverFromPlantInOperation"`
	ComponentProductRequirementDateInOperation               string   `json:"ComponentProductRequirementDateInOperation"`
	ComponentProductRequirementTimeInOperation               string   `json:"ComponentProductRequirementTimeInOperation"`
	ComponentProductPlannedQuantityInBaseUnitInOperation     float32  `json:"ComponentProductPlannedQuantityInBaseUnitInOperation"`
	ComponentProductPlannedQuantityInDeliveryUnitInOperation float32  `json:"ComponentProductPlannedQuantityInDeliveryUnitInOperation"`
	ComponentProductPlannedScrapInPercentInOperation         *float32 `json:"ComponentProductPlannedScrapInPercentInOperation"`
	ComponentProductBaseUnit                                 string   `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                             string   `json:"ComponentProductDeliveryUnit"`
	ComponentProductIsMarkedForBackflush                     *bool    `json:"ComponentProductIsMarkedForBackflush"`
	MRPArea                                                  *string  `json:"MRPArea"`
	MRPController                                            *string  `json:"MRPController"`
	ProductionVersion                                        *int     `json:"ProductionVersion"`
	ProductionVersionItem                                    *int     `json:"ProductionVersionItem"`
	ComponentProductReservation                              *int     `json:"ComponentProductReservation"`
	ComponentProductReservationItem                          *int     `json:"ComponentProductReservationItem"`
	CreationDate                                             string   `json:"CreationDate"`
	CreationTime                                             string   `json:"CreationTime"`
	LastChangeDate                                           string   `json:"LastChangeDate"`
	LastChangeTime                                           string   `json:"LastChangeTime"`
	IsReleased                                               *bool    `json:"IsReleased"`
	IsLocked                                                 *bool    `json:"IsLocked"`
	IsCancelled                                              *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                                      *bool    `json:"IsMarkedForDeletion"`
}

type ItemOperationCostingUpdates struct {
	ProductionOrder     int     `json:"ProductionOrder"`
	ProductionOrderItem int     `json:"ProductionOrderItem"`
	Operations          int     `json:"Operations"`
	OperationsItem      int     `json:"OperationsItem"`
	OperationID         int     `json:"OperationID"`
	Currency            string  `json:"Currency"`
	CostingAmount       float32 `json:"CostingAmount"`
	CreationDate        string  `json:"CreationDate"`
	CreationTime        string  `json:"CreationTime"`
	LastChangeDate      string  `json:"LastChangeDate"`
	LastChangeTime      string  `json:"LastChangeTime"`
	IsReleased          *bool   `json:"IsReleased"`
	IsLocked            *bool   `json:"IsLocked"`
	IsCancelled         *bool   `json:"IsCancelled"`
	IsMarkedForDeletion *bool   `json:"IsMarkedForDeletion"`
}
