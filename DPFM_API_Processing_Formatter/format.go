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
		TotalQuantity:                    *data.TotalQuantity,
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
		MinimumLotSizeQuantity:                data.MinimumLotSizeQuantity,
		StandardLotSizeQuantity:               data.StandardLotSizeQuantity,
		LotSizeRoundingQuantity:               data.LotSizeRoundingQuantity,
		MaximumLotSizeQuantity:                data.MaximumLotSizeQuantity,
		LotSizeIsFixed:                        data.LotSizeIsFixed,
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
		TotalQuantity:                         *data.TotalQuantity,
		PlannedScrapQuantity:                  data.PlannedScrapQuantity,
		ConfirmedYieldQuantity:                data.ConfirmedYieldQuantity,
		ProductionOrderItemText:               data.ProductionOrderItemText,
	}
}

func ConvertToItemComponentUpdates(itemComponentUpdates dpfm_api_input_reader.ItemComponent) *ItemComponentUpdates {
	data := itemComponentUpdates

	return &ItemComponentUpdates{
		ProductionOrder:                      data.ProductionOrder,
		ProductionOrderItem:                  data.ProductionOrderItem,
		Operations:                           data.Operations,
		OperationsItem:                       data.OperationsItem,
		BillOfMaterial:                       data.BillOfMaterial,
		BillOfMaterialItem:                   data.BillOfMaterialItem,
		ComponentProductRequirementDate:      data.ComponentProductRequirementDate,
		ComponentProductRequirementTime:      data.ComponentProductRequirementTime,
		ComponentProductIsMarkedForBackflush: data.ComponentProductIsMarkedForBackflush,
		ComponentProductBusinessPartner:      data.ComponentProductBusinessPartner,
		StockConfirmationPlant:               data.StockConfirmationPlant,
		SortField:                            data.SortField,
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
		RequiredQuantity:                     data.RequiredQuantity,
		WithdrawnQuantity:                    data.WithdrawnQuantity,
		ConfirmedAvailableQuantity:           data.ConfirmedAvailableQuantity,
		IsMarkedForDeletion:                  data.IsMarkedForDeletion,
	}
}

func ConvertToItemComponentStockConfirmationUpdates(itemComponentStockConfirmationUpdates dpfm_api_input_reader.ItemComponentStockConfirmation) *ItemComponentStockConfirmationUpdates {
	data := itemComponentStockConfirmationUpdates

	return &ItemComponentStockConfirmationUpdates{
		ProductionOrder:     data.ProductionOrder,
		ProductionOrderItem: data.ProductionOrderItem,
		Operations:          data.Operations,
		OperationsItem:      data.OperationsItem,
		BillOfMaterial:      data.BillOfMaterial,
		BillOfMaterialItem:  data.BillOfMaterialItem,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToItemComponentCostingUpdates(itemComponentCostingUpdates dpfm_api_input_reader.ItemComponentCosting) *ItemComponentCostingUpdates {
	data := itemComponentCostingUpdates

	return &ItemComponentCostingUpdates{
		ProductionOrder:     data.ProductionOrder,
		ProductionOrderItem: data.ProductionOrderItem,
		Operations:          data.Operations,
		OperationsItem:      data.OperationsItem,
		CostingAmount:       data.CostingAmount,
		IsMarkedForDeletion: data.IsMarkedForDeletion,
	}
}

func ConvertToItemOperationsUpdates(itemOperationsUpdates dpfm_api_input_reader.ItemOperations) *ItemOperationsUpdates {
	data := itemOperationsUpdates

	return &ItemOperationsUpdates{
		ProductionOrder:                      data.ProductionOrder,
		ProductionOrderItem:                  data.ProductionOrderItem,
		Operations:                           data.Operations,
		OperationsItem:                       data.OperationsItem,
		OperationsText:                       data.OperationsText,
		SequenceText:                         data.SequenceText,
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
