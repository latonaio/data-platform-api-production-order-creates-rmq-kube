package requests

type Header struct {
	ProductionOrder                     int      `json:"ProductionOrder"`
	ProductionOrderType                 *string  `json:"ProductionOrderType"`
	CreationDate                        string   `json:"CreationDate"`
	LastChangeDate                      string   `json:"LastChangeDate"`
	HeaderIsReleased                    *bool    `json:"HeaderIsReleased"`
	HeaderIsPartiallyConfirmed          *bool    `json:"HeaderIsPartiallyConfirmed"`
	HeaderIsConfirmed                   *bool    `json:"HeaderIsConfirmed"`
	HeaderIsLocked                      *bool    `json:"HeaderIsLocked"`
	HeaderIsMarkedForDeletion           *bool    `json:"HeaderIsMarkedForDeletion"`
	Product                             *string  `json:"Product"`
	OwnerProductionPlant                string   `json:"OwnerProductionPlant"`
	OwnerProductionPlantBusinessPartner int      `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantStorageLocation *string  `json:"OwnerProductionPlantStorageLocation"`
	MRPArea                             *string  `json:"MRPArea"`
	MRPController                       *string  `json:"MRPController"`
	ProductionSupervisor                *string  `json:"ProductionSupervisor"`
	ProductionVersion                   *int     `json:"ProductionVersion"`
	PlannedOrder                        *int     `json:"PlannedOrder"`
	OrderID                             *int     `json:"OrderID"`
	OrderItem                           *int     `json:"OrderItem"`
	ProductionOrderPlannedStartDate     *string  `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime     *string  `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate       *string  `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime       *string  `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate    *string  `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime    *string  `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate      *string  `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime      *string  `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate        *string  `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime        *string  `json:"ProductionOrderActualEndTime"`
	ProductionUnit                      *string  `json:"ProductionUnit"`
	TotalQuantity                       float32  `json:"TotalQuantity"`
	PlannedScrapQuantity                *float32 `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity              *float32 `json:"ConfirmedYieldQuantity"`
	ProductionOrderHeaderText           *string  `json:"ProductionOrderHeaderText"`
}
