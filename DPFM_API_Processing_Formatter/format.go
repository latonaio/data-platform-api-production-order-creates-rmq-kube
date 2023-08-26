package dpfm_api_processing_formatter

import dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		ProductionOrder:									data.ProductionOrder,
		OwnerProductionPlantStorageLocation:				data.OwnerProductionPlantStorageLocation,
		DepartureDeliverFromPlantStorageLocation :			data.DepartureDeliverFromPlantStorageLocation ,
		DestinationDeliverToPlantStorageLocation:			data.DestinationDeliverToPlantStorageLocation,
		ProductBaseUnit:									data.ProductBaseUnit,
		ProductionOrderQuantityInBaseUnit:					data.ProductionOrderQuantityInBaseUnit,
		ProductionOrderQuantityInDepartureProductionUnit:	data.ProductionOrderQuantityInDepartureProductionUnit,
		ProductionOrderQuantityInDestinationProductionUnit:	data.ProductionOrderQuantityInDestinationProductionUnit,
		ProductionOrderQuantityInDepartureDeliveryUnit:		data.ProductionOrderQuantityInDepartureDeliveryUnit,
		ProductionOrderQuantityInDestinationDeliveryUnit:	data.ProductionOrderQuantityInDestinationDeliveryUnit,
		ProductionOrderPlannedScrapQtyInBaseUnit:			data.ProductionOrderPlannedScrapQtyInBaseUnit,
		ProductionOrderPlannedStartDate:					data.ProductionOrderPlannedStartDate,
		ProductionOrderPlannedStartTime:					data.ProductionOrderPlannedStartTime,
		ProductionOrderPlannedEndDate:						data.ProductionOrderPlannedEndDate,
		ProductionOrderPlannedEndTime:						data.ProductionOrderPlannedEndTime,
		ProductionOrderActualStartDate:						data.ProductionOrderActualStartDate,
		ProductionOrderActualStartTime:						data.ProductionOrderActualStartTime,
		ProductionOrderActualEndDate:						data.ProductionOrderActualEndDate,
		ProductionOrderActualEndTime:						data.ProductionOrderActualEndTime,
		ProductionOrderHeaderText:							data.ProductionOrderHeaderText,
		IsLocked:											data.IsLocked,
	}
}	


func ConvertToItemUpdates(itemUpdates dpfm_api_input_reader.Item) *ItemUpdates {
	data := itemUpdates

	return &ItemUpdates{
		ProductionOrder:								data.ProductionOrder,
		ProductionOrderItem:							data.ProductionOrderItem,
		ProductionPlantBusinessPartner:					data.ProductionPlantBusinessPartner,
		ProductionPlantStorageLocation:					data.ProductionPlantStorageLocation,
		DeliverFromPlantStorageLocation:				data.DeliverFromPlantStorageLocation,
		DeliverToPlantStorageLocation:					data.DeliverToPlantStorageLocation,
		ProductionOrderQuantityInBaseUnit:				data.ProductionOrderQuantityInBaseUnit,
		ProductionOrderQuantityInProductionUnit:		data.ProductionOrderQuantityInProductionUnit,
		ProductionOrderQuantityInDeliveryUnit:			data.ProductionOrderQuantityInDeliveryUnit,
		ProductionOrderPlannedScrapQtyInBaseUnit:		data.ProductionOrderPlannedScrapQtyInBaseUnit,
		ProductionOrderPlannedStartDate:				data.ProductionOrderPlannedStartDate,
		ProductionOrderPlannedStartTime:				data.ProductionOrderPlannedStartTime,
		ProductionOrderPlannedEndDate:					data.ProductionOrderPlannedEndDate,
		ProductionOrderPlannedEndTime:					data.ProductionOrderPlannedEndTime,
		ProductionOrderActualStartDate:					data.ProductionOrderActualStartDate,
		ProductionOrderActualStartTime:					data.ProductionOrderActualStartTime,
		ProductionOrderActualEndDate:					data.ProductionOrderActualEndDate,
		ProductionOrderActualEndTime:					data.ProductionOrderActualEndTime,
		ProductionOrderItemText:						data.ProductionOrderItemText,
		IsLocked:										data.IsLocked,
	}
}

func ConvertToItemComponentUpdates(itemComponentUpdates dpfm_api_input_reader.ItemComponent) *ItemComponentUpdates {
	data := itemComponentUpdates

	return &ItemComponentUpdates{
		ProductionOrder:								data.ProductionOrder,
		ProductionOrderItem:							data.ProductionOrderItem,
		BillOfMaterial:									data.BillOfMaterial,
		BillOfMaterialItem:								data.BillOfMaterialItem,
		ComponentProductRequirementDate:				data.ComponentProductRequirementDate,
		ComponentProductRequirementTime:				data.ComponentProductRequirementTime,
		ComponentProductRequiredQuantityInBaseUnit:		data.ComponentProductRequiredQuantityInBaseUnit,
		ComponentProductRequiredQuantityInDeliveryUnit:	data.ComponentProductRequiredQuantityInDeliveryUnit,
		ComponentProductPlannedScrapInPercent:			data.ComponentProductPlannedScrapInPercent,
		IsLocked:										data.IsLocked,
	}
}

func ConvertToItemComponentDeliveryScheduleLineUpdates(itemComponentDeliveryScheduleLineUpdates dpfm_api_input_reader.ItemComponentDeliveryScheduleLine) *ItemComponentDeliveryScheduleLineUpdates {
	data := itemComponentDeliveryScheduleLineUpdates

	return &ItemComponentDeliveryScheduleLineUpdates{
		ProductionOrder:						data.ProductionOrder,
		ProductionOrderItem:					data.ProductionOrderItem,
		BillOfMaterial:							data.BillOfMaterial,
		BillOfMaterialItem:						data.BillOfMaterialItem,
		ScheduleLine:							data.ScheduleLine,
		ItemScheduleLineDeliveryBlockStatus:	data.ItemScheduleLineDeliveryBlockStatus,
		IsLocked:								data.IsLocked,
	}
}

func ConvertToItemComponentCostingUpdates(itemComponentCostingUpdates dpfm_api_input_reader.ItemComponentCosting) *ItemComponentCostingUpdates {
	data := itemComponentCostingUpdates

	return &ItemComponentCostingUpdates{
		ProductionOrder:				data.ProductionOrder,
		ProductionOrderItem:			data.ProductionOrderItem,
		BillOfMaterial:					data.BillOfMaterial,
		BillOfMaterialItem:				data.BillOfMaterialItem,
		IsLocked:						data.IsLocked,
	}
}

func ConvertToItemOperationUpdates(itemOperationUpdates dpfm_api_input_reader.ItemOperation) *ItemOperationUpdates{
	data := itemOperationUpdates

	return &ItemOperationUpdates{
		ProductionOrder:									data.ProductionOrder,
		ProductionOrderItem:								data.ProductionOrderItem,
		Operations:											data.Operations,
		OperationsItem:										data.OperationsItem,
		OperationID:										data.OperationID,
		SequenceText:										data.SequenceText,
		OperationText:										data.OperationText,
		OperationPlannedQuantityInBaseUnit:					data.OperationPlannedQuantityInBaseUnit,
		OperationPlannedQuantityInProductionUnit:			data.OperationPlannedQuantityInProductionUnit,
		OperationPlannedQuantityInOperationUnit:			data.OperationPlannedQuantityInOperationUnit,
		OperationPlannedQuantityInDeliveryUnit:				data.OperationPlannedQuantityInDeliveryUnit,
		OperationPlannedScrapInPercent:						data.OperationPlannedScrapInPercent,
		OperationErlstSchedldExecStrtDte:					data.OperationErlstSchedldExecStrtDte,
		OperationErlstSchedldExecStrtTme:					data.OperationErlstSchedldExecStrtTme,
		OperationErlstSchedldExecEndDte:					data.OperationErlstSchedldExecEndDte,
		OperationErlstSchedldExecEndTme:					data.OperationErlstSchedldExecEndTme,
		OperationActualExecutionStartDate:					data.OperationActualExecutionStartDate,
		OperationActualExecutionStartTime:					data.OperationActualExecutionStartTime,
		OperationActualExecutionEndDate:					data.OperationActualExecutionEndDate,
		OperationActualExecutionEndTime:					data.OperationActualExecutionEndTime,
		IsLocked:											data.IsLocked,
	}
}