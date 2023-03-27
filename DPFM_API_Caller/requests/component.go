package requests

type Component struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	ProductionOrderSequence              string   `json:"ProductionOrderSequence"`
	ProductionOrderOperation             string   `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations        string   `json:"OrderInternalBillOfOperations"`
	Reservation                          *int     `json:"Reservation"`
	ReservationItem                      *int     `json:"ReservationItem"`
	ComponentProduct                     *string  `json:"ComponentProduct"`
	ComponentProductRequirementDate      *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime      *string  `json:"ComponentProductRequirementTime"`
	ComponentProductIsMarkedForBackflush *bool    `json:"ComponentProductIsMarkedForBackflush"`
	PlannedOrder                         *int     `json:"PlannedOrder"`
	OrderID                              *int     `json:"OrderID"`
	OrderItem                            *int     `json:"OrderItem"`
	SortField                            *string  `json:"SortField"`
	BillOfMaterial                       *int     `json:"BillOfMaterial"`
	BOMItem                              *int     `json:"BOMItem"`
	BOMItemDescription                   *string  `json:"BOMItemDescription"`
	StorageLocation                      *string  `json:"StorageLocation"`
	Batch                                *string  `json:"Batch"`
	GoodsRecipientName                   *string  `json:"GoodsRecipientName"`
	UnloadingPointName                   *string  `json:"UnloadingPointName"`
	ProductCompIsAlternativeItem         *bool    `json:"ProductCompIsAlternativeItem"`
	CostingPolicy                        *string  `json:"CostingPolicy"`
	PriceUnitQty                         *string  `json:"PriceUnitQty"`
	StandardPrice                        *float32 `json:"StandardPrice"`
	MovingAveragePrice                   *float32 `json:"MovingAveragePrice"`
	ComponentScrapInPercent              *float32 `json:"ComponentScrapInPercent"`
	OperationScrapInPercent              *float32 `json:"OperationScrapInPercent"`
	BaseUnit                             *string  `json:"BaseUnit"`
	RequiredQuantity                     *float32 `json:"RequiredQuantity"`
	WithdrawnQuantity                    *float32 `json:"WithdrawnQuantity"`
	ConfirmedAvailableQuantity           *float32 `json:"ConfirmedAvailableQuantity"`
	ProductCompOriginalQuantity          *float32 `json:"ProductCompOriginalQuantity"`
	IsMarkedForDeletion                  *bool    `json:"IsMarkedForDeletion"`
}
