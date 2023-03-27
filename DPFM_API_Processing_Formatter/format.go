package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		ProductionOrder:                  data.ProductionOrder,
		HeaderIsReleased:                 data.HeaderIsReleased,
		HeaderIsPartiallyConfirmed:       data.HeaderIsPartiallyConfirmed,
		HeaderIsConfirmed:                data.HeaderIsConfirmed,
		HeaderIsLocked:                   data.HeaderIsLocked,
		HeaderIsMarkedForDeletion:        data.HeaderIsMarkedForDeletion,
		ProductionOrderPlannedStartDate:  data.ProductionOrderPlannedStartDate,
		ProductionOrderPlannedStartTime:  data.ProductionOrderPlannedStartTime,
		ProductionOrderPlannedEndDate:    data.ProductionOrderPlannedEndDate,
		ProductionOrderPlannedEndTime:    data.ProductionOrderPlannedEndTime,
		ProductionOrderActualReleaseDate: data.ProductionOrderActualReleaseDate,
		ProductionOrderActualReleaseTime: data.ProductionOrderActualReleaseTime,
		ProductionOrderActualStartDate:   data.ProductionOrderActualStartDate,
		ProductionOrderActualStartTime:   data.ProductionOrderActualStartTime,
		ProductionOrderActualEndDate:     data.ProductionOrderActualEndDate,
		ProductionOrderActualEndTime:     data.ProductionOrderActualEndTime,
		TotalQuantity:                    data.TotalQuantity,
		PlannedScrapQuantity:             data.PlannedScrapQuantity,
		ConfirmedYieldQuantity:           data.ConfirmedYieldQuantity,
		ProductionOrderHeaderText:        data.ProductionOrderHeaderText,
	}
}

func ConvertToHeaderDocUpdates(headerDocUpdates dpfm_api_input_reader.HeaderDoc) *HeaderDocUpdates {
	data := headerDocUpdates

	return &HeaderDocUpdates{
		FileName: data.FileName,
		FilePath: data.FilePath,
	}
}

func ConvertToItemUpdates(itemUpdates dpfm_api_input_reader.Item) *ItemUpdates {
	data := itemUpdates

	return &ItemUpdates{
		ProductionOrder:                       data.ProductionOrder,
		ProductionOrderItem:                   data.ProductionOrderItem,
		ProductionOrderHasGeneratedOperations: data.ProductionOrderHasGeneratedOperations,
		ProductAvailabilityIsNotChecked:       data.ProductAvailabilityIsNotChecked,
		ProductionOrderPlannedStartDate:       data.ProductionOrderPlannedStartDate,
		ProductionOrderPlannedStartTime:       data.ProductionOrderPlannedStartTime,
		ProductionOrderPlannedEndDate:         data.ProductionOrderPlannedEndDate,
		ProductionOrderPlannedEndTime:         data.ProductionOrderPlannedEndTime,
		ProductionOrderActualReleaseDate:      data.ProductionOrderActualReleaseDate,
		ProductionOrderActualReleaseTime:      data.ProductionOrderActualReleaseTime,
		ProductionOrderActualStartDate:        data.ProductionOrderActualStartDate,
		ProductionOrderActualStartTime:        data.ProductionOrderActualStartTime,
		ProductionOrderActualEndDate:          data.ProductionOrderActualEndDate,
		ProductionOrderActualEndTime:          data.ProductionOrderActualEndTime,
		TotalQuantity:                         data.TotalQuantity,
		PlannedScrapQuantity:                  data.PlannedScrapQuantity,
		ConfirmedYieldQuantity:                data.ConfirmedYieldQuantity,
		ProductionOrderItemText:               data.ProductionOrderItemText,
	}
}

func ConvertToComponentUpdates(componentUpdates dpfm_api_input_reader.Component) *ComponentUpdates {
	data := componentUpdates

	return &ComponentUpdates{
		ProductionOrder:                      data.ProductionOrder,
		ProductionOrderItem:                  data.ProductionOrderItem,
		ProductionOrderSequence:              data.ProductionOrderSequence,
		ProductionOrderOperation:             data.ProductionOrderOperation,
		OrderInternalBillOfOperations:        data.OrderInternalBillOfOperations,
		ComponentProductRequirementDate:      data.ComponentProductRequirementDate,
		ComponentProductRequirementTime:      data.ComponentProductRequirementTime,
		ComponentProductIsMarkedForBackflush: data.ComponentProductIsMarkedForBackflush,
		SortField:                            data.SortField,
		BillOfMaterial:                       data.BillOfMaterial,
		BOMItemDescription:                   data.BOMItemDescription,
		GoodsRecipientName:                   data.GoodsRecipientName,
		UnloadingPointName:                   data.UnloadingPointName,
		ProductCompIsAlternativeItem:         data.ProductCompIsAlternativeItem,
		CostingPolicy:                        data.CostingPolicy,
		PriceUnitQty:                         data.PriceUnitQty,
		StandardPrice:                        data.StandardPrice,
		MovingAveragePrice:                   data.MovingAveragePrice,
		ComponentScrapInPercent:              data.ComponentScrapInPercent,
		OperationScrapInPercent:              data.OperationScrapInPercent,
		BaseUnit:                             data.BaseUnit,
		RequiredQuantity:                     data.RequiredQuantity,
		WithdrawnQuantity:                    data.WithdrawnQuantity,
		ConfirmedAvailableQuantity:           data.ConfirmedAvailableQuantity,
		IsMarkedForDeletion:                  data.IsMarkedForDeletion,
	}
}

func ConvertToComponentStockConfirmationUpdates(componentStockConfirmationUpdates dpfm_api_input_reader.ComponentStockConfirmation) *ComponentStockConfirmationUpdates {
	data := componentStockConfirmationUpdates

	return &ComponentStockConfirmationUpdates{
		ProductionOrder:                 data.ProductionOrder,
		ProductionOrderItem:             data.ProductionOrderItem,
		ProductionOrderSequence:         data.ProductionOrderSequence,
		ProductionOrderOperation:        data.ProductionOrderOperation,
		OrderInternalBillOfOperations:   data.OrderInternalBillOfOperations,
		ComponentProduct:                data.ComponentProduct,
		ComponentProductRequirementDate: data.ComponentProductRequirementDate,
		ComponentProductRequirementTime: data.ComponentProductRequirementTime,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}
}

func ConvertToComponentCostingUpdates(componentCostingUpdates dpfm_api_input_reader.ComponentCosting) *ComponentCostingUpdates {
	data := componentCostingUpdates

	return &ComponentCostingUpdates{
		ProductionOrder:               data.ProductionOrder,
		ProductionOrderItem:           data.ProductionOrderItem,
		ProductionOrderSequence:       data.ProductionOrderSequence,
		ProductionOrderOperation:      data.ProductionOrderOperation,
		OrderInternalBillOfOperations: data.OrderInternalBillOfOperations,
		ComponentProduct:              data.ComponentProduct,
		CostingAmount:                 data.CostingAmount,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}
}

func ConvertToOperationUpdates(operationUpdates dpfm_api_input_reader.Operation) *OperationUpdates {
	data := operationUpdates

	return &OperationUpdates{
		ProductionOrder:                      data.ProductionOrder,
		ProductionOrderItem:                  data.ProductionOrderItem,
		ProductionOrderSequence:              data.ProductionOrderSequence,
		ProductionOrderOperation:             data.ProductionOrderOperation,
		OrderInternalBillOfOperations:        data.OrderInternalBillOfOperations,
		OrderIntBillOfOperationsItem:         data.OrderIntBillOfOperationsItem,
		ProductionOrderSequenceText:          data.ProductionOrderSequenceText,
		ProductionOrderOperationText:         data.ProductionOrderOperationText,
		OperationIsReleased:                  data.OperationIsReleased,
		OperationIsPartiallyConfirmed:        data.OperationIsPartiallyConfirmed,
		OperationIsConfirmed:                 data.OperationIsConfirmed,
		OperationIsClosed:                    data.OperationIsClosed,
		OperationErlstSchedldExecStrtDte:     data.OperationErlstSchedldExecStrtDte,
		OperationErlstSchedldExecStrtTme:     data.OperationErlstSchedldExecStrtTme,
		OperationErlstSchedldExecEndDate:     data.OperationErlstSchedldExecEndDate,
		OperationErlstSchedldExecEndTme:      data.OperationErlstSchedldExecEndTme,
		OperationActualExecutionStartDate:    data.OperationActualExecutionStartDate,
		OperationActualExecutionStartTime:    data.OperationActualExecutionStartTime,
		OperationActualExecutionEndDate:      data.OperationActualExecutionEndDate,
		OperationActualExecutionEndTime:      data.OperationActualExecutionEndTime,
		ErlstSchedldExecDurnInWorkdays:       data.ErlstSchedldExecDurnInWorkdays,
		OperationActualExecutionDays:         data.OperationActualExecutionDays,
		OperationPlannedTotalQuantity:        data.OperationPlannedTotalQuantity,
		OperationTotalConfirmedYieldQuantity: data.OperationTotalConfirmedYieldQuantity,
	}
}
