package dpfm_api_processing_formatter

type HeaderUpdates struct {
	ProductionOrder                  int      `json:"ProductionOrder"`
	HeaderIsReleased                 *bool    `json:"HeaderIsReleased"`
	HeaderIsPartiallyConfirmed       *bool    `json:"HeaderIsPartiallyConfirmed"`
	HeaderIsConfirmed                *bool    `json:"HeaderIsConfirmed"`
	HeaderIsLocked                   *bool    `json:"HeaderIsLocked"`
	HeaderIsMarkedForDeletion        *bool    `json:"HeaderIsMarkedForDeletion"`
	ProductionOrderPlannedStartDate  *string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime  *string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate    *string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime    *string  `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate *string  `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime *string  `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate   *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime   *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate     *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime     *string  `json:"ProductionOrderActualEndTime"`
	TotalQuantity                    float32  `json:"TotalQuantity"`
	PlannedScrapQuantity             *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity           *float32 `json:"ConfirmedYieldQuantity"`
	ProductionOrderHeaderText        *string  `json:"ProductionOrderHeaderText"`
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
	ProductionOrder                       int      `json:"ProductionOrder"`
	ProductionOrderItem                   int      `json:"ProductionOrderItem"`
	ProductionOrderHasGeneratedOperations *bool    `json:"ProductionOrderHasGeneratedOperations"`
	ProductAvailabilityIsNotChecked       *bool    `json:"ProductAvailabilityIsNotChecked"`
	MinimumLotSizeQuantity                *float32 `json:"MinimumLotSizeQuantity"`
	StandardLotSizeQuantity               *float32 `json:"StandardLotSizeQuantity"`
	LotSizeRoundingQuantity               *float32 `json:"LotSizeRoundingQuantity"`
	MaximumLotSizeQuantity                *float32 `json:"MaximumLotSizeQuantity"`
	LotSizeIsFixed                        *bool    `json:"LotSizeIsFixed"`
	ProductionOrderPlannedStartDate       *string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime       *string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate         *string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime         *string  `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate      *string  `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime      *string  `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate        *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime        *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate          *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime          *string  `json:"ProductionOrderActualEndTime"`
	TotalQuantity                         float32  `json:"TotalQuantity"`
	PlannedScrapQuantity                  *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity                *float32 `json:"ConfirmedYieldQuantity"`
	ProductionOrderItemText               *string  `json:"ProductionOrderItemText"`
}

type ItemComponentUpdates struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	Operations                           int      `json:"Operations"`
	OperationsItem                       int      `json:"OperationsItem"`
	BillOfMaterial                       int      `json:"BillOfMaterial"`
	BillOfMaterialItem                   int      `json:"BillOfMaterialItem"`
	ComponentProductRequirementDate      *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime      *string  `json:"ComponentProductRequirementTime"`
	ComponentProductIsMarkedForBackflush *bool    `json:"ComponentProductIsMarkedForBackflush"`
	ComponentProductBusinessPartner      *int     `json:"ComponentProductBusinessPartner"`
	StockConfirmationPlant               *string  `json:"StockConfirmationPlant"`
	SortField                            *string  `json:"SortField"`
	BOMItemDescription                   *string  `json:"BOMItemDescription"`
	GoodsRecipientName                   *string  `json:"GoodsRecipientName"`
	UnloadingPointName                   *string  `json:"UnloadingPointName"`
	ProductCompIsAlternativeItem         *bool    `json:"ProductCompIsAlternativeItem"`
	CostingPolicy                        *string  `json:"CostingPolicy"`
	PriceUnitQty                         *string  `json:"PriceUnitQty"`
	StandardPrice                        *float32 `json:"StandardPrice"`
	MovingAveragePrice                   *float32 `json:"MovingAveragePrice"`
	ComponentScrapInPercent              *float32 `json:"ComponentScrapInPercent"`
	OperationScrapInPercent              *float32 `json:"OperationScrapInPercent"`
	RequiredQuantity                     *float32 `json:"RequiredQuantity"`
	WithdrawnQuantity                    *float32 `json:"WithdrawnQuantity"`
	ConfirmedAvailableQuantity           *float32 `json:"ConfirmedAvailableQuantity"`
	IsMarkedForDeletion                  *bool    `json:"IsMarkedForDeletion"`
}

type ItemComponentStockConfirmationUpdates struct {
	ProductionOrder     int   `json:"ProductionOrder"`
	ProductionOrderItem int   `json:"ProductionOrderItem"`
	Operations          int   `json:"Operations"`
	OperationsItem      int   `json:"OperationsItem"`
	BillOfMaterial      int   `json:"BillOfMaterial"`
	BillOfMaterialItem  int   `json:"BillOfMaterialItem"`
	IsMarkedForDeletion *bool `json:"IsMarkedForDeletion"`
}

type ItemComponentCostingUpdates struct {
	ProductionOrder     int      `json:"ProductionOrder"`
	ProductionOrderItem int      `json:"ProductionOrderItem"`
	Operations          int      `json:"Operations"`
	OperationsItem      int      `json:"OperationsItem"`
	CostingAmount       *float32 `json:"CostingAmount"`
	IsMarkedForDeletion *bool    `json:"IsMarkedForDeletion"`
}

type ItemDocUpdates struct {
	ProductionOrder     int     `json:"ProductionOrder"`
	ProductionOrderItem int     `json:"ProductionOrderItem"`
	DocType             string  `json:"DocType"`
	DocVersionID        int     `json:"DocVersionID"`
	DocID               string  `json:"DocID"`
	FileName            *string `json:"FileName"`
	FilePath            *string `json:"FilePath"`
}

type ItemOperationsUpdates struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	Operations                           int      `json:"Operations"`
	OperationsItem                       int      `json:"OperationsItem"`
	OperationsText                       *string  `json:"OperationsText"`
	SequenceText                         *string  `json:"SequenceText"`
	OperationIsPartiallyConfirmed        *bool    `json:"OperationIsPartiallyConfirmed"`
	OperationIsConfirmed                 *bool    `json:"OperationIsConfirmed"`
	OperationIsClosed                    *bool    `json:"OperationIsClosed"`
	OperationErlstSchedldExecStrtDte     *string  `json:"OperationErlstSchedldExecStrtDte"`
	OperationErlstSchedldExecStrtTme     *string  `json:"OperationErlstSchedldExecStrtTme"`
	OperationErlstSchedldExecEndDate     *string  `json:"OperationErlstSchedldExecEndDate"`
	OperationErlstSchedldExecEndTme      *string  `json:"OperationErlstSchedldExecEndTme"`
	OperationActualExecutionStartDate    *string  `json:"OperationActualExecutionStartDate"`
	OperationActualExecutionStartTime    *string  `json:"OperationActualExecutionStartTime"`
	OperationActualExecutionEndDate      *string  `json:"OperationActualExecutionEndDate"`
	OperationActualExecutionEndTime      *string  `json:"OperationActualExecutionEndTime"`
	ErlstSchedldExecDurnInWorkdays       *string  `json:"ErlstSchedldExecDurnInWorkdays"`
	OperationActualExecutionDays         *string  `json:"OperationActualExecutionDays"`
	OperationPlannedTotalQuantity        *float32 `json:"OperationPlannedTotalQuantity"`
	OperationTotalConfirmedYieldQuantity *float32 `json:"OperationTotalConfirmedYieldQuantity"`
}
