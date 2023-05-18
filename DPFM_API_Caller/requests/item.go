package requests

type Item struct {
	ProductionOrder                       int      `json:"ProductionOrder"`
	ProductionOrderItem                   int      `json:"ProductionOrderItem"`
	CreationDate                          string   `json:"CreationDate"`
	LastChangeDate                        string   `json:"LastChangeDate"`
	ItemIsReleased                        *bool    `json:"ItemIsReleased"`
	ItemIsPartiallyConfirmed              *bool    `json:"ItemIsPartiallyConfirmed"`
	ItemIsConfirmed                       *bool    `json:"ItemIsConfirmed"`
	ItemIsLocked                          *bool    `json:"ItemIsLocked"`
	ItemIsMarkedForDeletion               *bool    `json:"ItemIsMarkedForDeletion"`
	ProductionOrderHasGeneratedOperations *bool    `json:"ProductionOrderHasGeneratedOperations"`
	ProductAvailabilityIsNotChecked       *bool    `json:"ProductAvailabilityIsNotChecked"`
	PrecedingItem                         *int     `json:"PrecedingItem"`
	FollowingItem                         *int     `json:"FollowingItem"`
	Product                               *string  `json:"Product"`
	ProductionPlant                       string   `json:"ProductionPlant"`
	ProductionPlantBusinessPartner        int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlantStorageLocation        *string  `json:"ProductionPlantStorageLocation"`
	MRPArea                               *string  `json:"MRPArea"`
	MRPController                         *string  `json:"MRPController"`
	ProductionSupervisor                  *string  `json:"ProductionSupervisor"`
	ProductionVersion                     *int     `json:"ProductionVersion"`
	PlannedOrder                          *int     `json:"PlannedOrder"`
	OrderID                               *int     `json:"OrderID"`
	OrderItem                             *int     `json:"OrderItem"`
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
	ProductionUnit                        *string  `json:"ProductionUnit"`
	TotalQuantity                         float32  `json:"TotalQuantity"`
	PlannedScrapQuantity                  *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity                *float32 `json:"ConfirmedYieldQuantity"`
	ProductionOrderItemText               *string  `json:"ProductionOrderItemText"`
}
