package requests

type ItemComponentStockConfirmation struct {
	ProductionOrder           int     `json:"ProductionOrder"`
	ProductionOrderItem       int     `json:"ProductionOrderItem"`
	Operations                int     `json:"Operations"`
	OperationsItem            int     `json:"OperationsItem"`
	BillOfMaterial            int     `json:"BillOfMaterial"`
	BillOfMaterialItem        int     `json:"BillOfMaterialItem"`
	InventoryStockType        *string `json:"InventoryStockType"`
	InventorySpecialStockType *string `json:"InventorySpecialStockType"`
	AvailableProductStock     float32 `json:"AvailableProductStock"`
	IsMarkedForDeletion       *bool   `json:"IsMarkedForDeletion"`
}
