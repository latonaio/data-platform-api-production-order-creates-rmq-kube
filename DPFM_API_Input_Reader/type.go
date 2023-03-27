package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey    string   `json:"connection_key"`
	Result           bool     `json:"result"`
	RedisKey         string   `json:"redis_key"`
	Filepath         string   `json:"filepath"`
	APIStatusCode    int      `json:"api_status_code"`
	RuntimeSessionID string   `json:"runtime_session_id"`
	BusinessPartner  *int     `json:"business_partner"`
	ServiceLabel     string   `json:"service_label"`
	APIType          string   `json:"api_type"`
	Header           Header   `json:"ProductionOrder"`
	APISchema        string   `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	Deleted          bool     `json:"deleted"`
}

type Header struct {
	ProductionOrder                     int         `json:"ProductionOrder"`
	ProductionOrderType                 *string     `json:"ProductionOrderType"`
	CreationDate                        *string     `json:"CreationDate"`
	LastChangeDate                      *string     `json:"LastChangeDate"`
	HeaderIsReleased                    *bool       `json:"HeaderIsReleased"`
	HeaderIsPartiallyConfirmed          *bool       `json:"HeaderIsPartiallyConfirmed"`
	HeaderIsConfirmed                   *bool       `json:"HeaderIsConfirmed"`
	HeaderIsLocked                      *bool       `json:"HeaderIsLocked"`
	HeaderIsMarkedForDeletion           *bool       `json:"HeaderIsMarkedForDeletion"`
	Product                             *string     `json:"Product"`
	OwnerProductionPlant                *string     `json:"OwnerProductionPlant"`
	OwnerProductionPlantBusinessPartner *int        `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlantStorageLocation *string     `json:"OwnerProductionPlantStorageLocation"`
	MRPArea                             *string     `json:"MRPArea"`
	MRPController                       *string     `json:"MRPController"`
	ProductionSupervisor                *string     `json:"ProductionSupervisor"`
	ProductionVersion                   *string     `json:"ProductionVersion"`
	PlannedOrder                        *int        `json:"PlannedOrder"`
	OrderID                             *int        `json:"OrderID"`
	OrderItem                           *int        `json:"OrderItem"`
	ProductionOrderPlannedStartDate     *string     `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime     *string     `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate       *string     `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime       *string     `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate    *string     `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime    *string     `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate      *string     `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime      *string     `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate        *string     `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime        *string     `json:"ProductionOrderActualEndTime"`
	ProductionUnit                      *string     `json:"ProductionUnit"`
	TotalQuantity                       *float32    `json:"TotalQuantity"`
	PlannedScrapQuantity                *float32    `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity              *float32    `json:"ConfirmedYieldQuantity"`
	ProductionOrderHeaderText           *string     `json:"ProductionOrderHeaderText"`
	HeaderDoc                           []HeaderDoc `json:"HeaderDoc"`
	Item                                []Item      `json:"Item"`
}

type HeaderDoc struct {
	ProductionOrder          int     `json:"ProductionOrder"`
	DocType                  string  `json:"DocType"`
	DocVersionID             int     `json:"DocVersionID"`
	DocID                    string  `json:"DocID"`
	FileExtension            *string `json:"FileExtension"`
	FileName                 *string `json:"FileName"`
	FilePath                 *string `json:"FilePath"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
}

type Item struct {
	ProductionOrder                       int         `json:"ProductionOrder"`
	ProductionOrderItem                   int         `json:"ProductionOrderItem"`
	CreationDate                          *string     `json:"CreationDate"`
	LastChangeDate                        *string     `json:"LastChangeDate"`
	ItemIsReleased                        *bool       `json:"ItemIsReleased"`
	ItemIsPartiallyConfirmed              *bool       `json:"ItemIsPartiallyConfirmed"`
	ItemIsConfirmed                       *bool       `json:"ItemIsConfirmed"`
	ItemIsLocked                          *bool       `json:"ItemIsLocked"`
	ItemIsMarkedForDeletion               *bool       `json:"ItemIsMarkedForDeletion"`
	ProductionOrderHasGeneratedOperations *bool       `json:"ProductionOrderHasGeneratedOperations"`
	ProductAvailabilityIsNotChecked       *bool       `json:"ProductAvailabilityIsNotChecked"`
	Product                               *string     `json:"Product"`
	ProductionPlant                       *string     `json:"ProductionPlant"`
	ProductionPlantBusinessPartner        *int        `json:"ProductionPlantBusinessPartner"`
	ProductionPlantStorageLocation        *string     `json:"ProductionPlantStorageLocation"`
	InternalBillOfOperations              *int        `json:"InternalBillOfOperations"`
	MRPArea                               *string     `json:"MRPArea"`
	MRPController                         *string     `json:"MRPController"`
	ProductionSupervisor                  *string     `json:"ProductionSupervisor"`
	ProductionVersion                     *string     `json:"ProductionVersion"`
	PlannedOrder                          *int        `json:"PlannedOrder"`
	OrderID                               *int        `json:"OrderID"`
	OrderItem                             *int        `json:"OrderItem"`
	ProductionOrderPlannedStartDate       *string     `json:"ProductionOrderPlannedStartDate"`
	ProductionOrderPlannedStartTime       *string     `json:"ProductionOrderPlannedStartTime"`
	ProductionOrderPlannedEndDate         *string     `json:"ProductionOrderPlannedEndDate"`
	ProductionOrderPlannedEndTime         *string     `json:"ProductionOrderPlannedEndTime"`
	ProductionOrderActualReleaseDate      *string     `json:"ProductionOrderActualReleaseDate"`
	ProductionOrderActualReleaseTime      *string     `json:"ProductionOrderActualReleaseTime"`
	ProductionOrderActualStartDate        *string     `json:"ProductionOrderActualStartDate"`
	ProductionOrderActualStartTime        *string     `json:"ProductionOrderActualStartTime"`
	ProductionOrderActualEndDate          *string     `json:"ProductionOrderActualEndDate"`
	ProductionOrderActualEndTime          *string     `json:"ProductionOrderActualEndTime"`
	ProductionUnit                        *string     `json:"ProductionUnit"`
	TotalQuantity                         *float32    `json:"TotalQuantity"`
	PlannedScrapQuantity                  *float32    `json:"PlannedScrapQuantity"`
	ConfirmedYieldQuantity                *float32    `json:"ConfirmedYieldQuantity"`
	ProductionOrderItemText               *string     `json:"ProductionOrderItemText"`
	Component                             []Component `json:"Component"`
}

type Component struct {
	ProductionOrder                      int                          `json:"ProductionOrder"`
	ProductionOrderItem                  int                          `json:"ProductionOrderItem"`
	ProductionOrderSequence              string                       `json:"ProductionOrderSequence"`
	ProductionOrderOperation             string                       `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations        string                       `json:"OrderInternalBillOfOperations"`
	Reservation                          *int                         `json:"Reservation"`
	ReservationItem                      *int                         `json:"ReservationItem"`
	ComponentProduct                     *string                      `json:"ComponentProduct"`
	ComponentProductRequirementDate      *string                      `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime      *string                      `json:"ComponentProductRequirementTime"`
	ComponentProductIsMarkedForBackflush *bool                        `json:"ComponentProductIsMarkedForBackflush"`
	PlannedOrder                         *int                         `json:"PlannedOrder"`
	OrderID                              *int                         `json:"OrderID"`
	OrderItem                            *int                         `json:"OrderItem"`
	SortField                            *string                      `json:"SortField"`
	BillOfMaterial                       *int                         `json:"BillOfMaterial"`
	BOMItem                              *int                         `json:"BOMItem"`
	BOMItemDescription                   *string                      `json:"BOMItemDescription"`
	StorageLocation                      *string                      `json:"StorageLocation"`
	Batch                                *string                      `json:"Batch"`
	GoodsRecipientName                   *string                      `json:"GoodsRecipientName"`
	UnloadingPointName                   *string                      `json:"UnloadingPointName"`
	ProductCompIsAlternativeItem         *bool                        `json:"ProductCompIsAlternativeItem"`
	CostingPolicy                        *string                      `json:"CostingPolicy"`
	PriceUnitQty                         *string                      `json:"PriceUnitQty"`
	StandardPrice                        *float32                     `json:"StandardPrice"`
	MovingAveragePrice                   *float32                     `json:"MovingAveragePrice"`
	ComponentScrapInPercent              *float32                     `json:"ComponentScrapInPercent"`
	OperationScrapInPercent              *float32                     `json:"OperationScrapInPercent"`
	BaseUnit                             *string                      `json:"BaseUnit"`
	RequiredQuantity                     *float32                     `json:"RequiredQuantity"`
	WithdrawnQuantity                    *float32                     `json:"WithdrawnQuantity"`
	ConfirmedAvailableQuantity           *float32                     `json:"ConfirmedAvailableQuantity"`
	ProductCompOriginalQuantity          *float32                     `json:"ProductCompOriginalQuantity"`
	IsMarkedForDeletion                  *bool                        `json:"IsMarkedForDeletion"`
	ComponentStockConfirmation           []ComponentStockConfirmation `json:"ComponentStockConfirmation"`
	ComponentCosting                     []ComponentCosting           `json:"ComponentCosting"`
	Operation                            []Operation                  `json:"Operation"`
}

type ComponentStockConfirmation struct {
	ProductionOrder                 int      `json:"ProductionOrder"`
	ProductionOrderItem             int      `json:"ProductionOrderItem"`
	ProductionOrderSequence         string   `json:"ProductionOrderSequence"`
	ProductionOrderOperation        string   `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations   string   `json:"OrderInternalBillOfOperations"`
	ComponentProduct                string   `json:"ComponentProduct"`
	ComponentProductRequirementDate *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime *string  `json:"ComponentProductRequirementTime"`
	InventoryStockType              *string  `json:"InventoryStockType"`
	InventorySpecialStockType       *string  `json:"InventorySpecialStockType"`
	AvailableProductStock           *float32 `json:"AvailableProductStock"`
	IsMarkedForDeletion             *bool    `json:"IsMarkedForDeletion"`
}

type ComponentCosting struct {
	ProductionOrder               int      `json:"PlannedOrder"`
	ProductionOrderItem           int      `json:"ProductionOrderItem"`
	ProductionOrderSequence       string   `json:"ProductionOrderSequence"`
	ProductionOrderOperation      string   `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations string   `json:"OrderInternalBillOfOperations"`
	ComponentProduct              string   `json:"ComponentProduct"`
	Currency                      *string  `json:"Currency"`
	CostingAmount                 *float32 `json:"CostingAmount"`
	IsMarkedForDeletion           *bool    `json:"IsMarkedForDeletion"`
}

type Operation struct {
	ProductionOrder                      int      `json:"ProductionOrder"`
	ProductionOrderItem                  int      `json:"ProductionOrderItem"`
	ProductionOrderSequence              string   `json:"ProductionOrderSequence"`
	ProductionOrderOperation             string   `json:"ProductionOrderOperation"`
	OrderInternalBillOfOperations        string   `json:"OrderInternalBillOfOperations"`
	OrderIntBillOfOperationsItem         int      `json:"OrderIntBillOfOperationsItem"`
	ProductionOrderSequenceText          *string  `json:"ProductionOrderSequenceText"`
	ProductionOrderOperationText         *string  `json:"ProductionOrderOperationText"`
	OperationIsReleased                  *bool    `json:"OperationIsReleased"`
	OperationIsPartiallyConfirmed        *bool    `json:"OperationIsPartiallyConfirmed"`
	OperationIsConfirmed                 *bool    `json:"OperationIsConfirmed"`
	OperationIsClosed                    *bool    `json:"OperationIsClosed"`
	OperationIsMarkedForDeletion         *bool    `json:"OperationIsMarkedForDeletion"`
	ProductionPlant                      *string  `json:"ProductionPlant"`
	WorkCenter                           *int     `json:"WorkCenter"`
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
	OperationUnit                        *string  `json:"OperationUnit"`
	OperationPlannedTotalQuantity        *float32 `json:"OperationPlannedTotalQuantity"`
	OperationTotalConfirmedYieldQuantity *float32 `json:"OperationTotalConfirmedYieldQuantity"`
}
