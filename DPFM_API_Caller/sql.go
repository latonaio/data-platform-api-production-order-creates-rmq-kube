package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-production-order-creates-rmq-kube/sub_func_complementer"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item
	var itemComponent *[]dpfm_api_output_formatter.ItemComponent
	var itemComponentDeliveryScheduleLine *[]dpfm_api_output_formatter.ItemComponentDeliveryScheduleLine
	var itemComponentPricingElement *[]dpfm_api_output_formatter.ItemComponentPricingElement
	var itemComponentCosting *[]dpfm_api_output_formatter.ItemComponentCosting
	var itemOperation *[]dpfm_api_output_formatter.ItemOperation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "Item":
			item = c.itemCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponent":
			itemComponent = c.itemComponentCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponentDeliveryScheduleLine":
			itemComponentDeliveryScheduleLine = c.itemComponentDeliveryScheduleLineCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponentPricingElement":
			itemComponentPricingElement = c.itemComponentPricingElementCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemComponentCosting":
			itemComponentCosting = c.itemComponentCostingCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		case "ItemOperation":
			itemOperation = c.itemOperationCreateSql(nil, mtx, input, output, subfuncSDC, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                         	header,
		Item:                           	item,
		ItemComponent:                  	itemComponent,
		ItemComponentDeliveryScheduleLine:	itemComponentDeliveryScheduleLine,
		ItemComponentPricingElement:       	itemComponentPricingElement,
		ItemComponentCosting:           	itemComponentCosting,
		ItemOperation:                  	itemOperation,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header								*dpfm_api_output_formatter.Header
	var item								*[]dpfm_api_output_formatter.Item
	var itemComponent						*[]dpfm_api_output_formatter.ItemComponent
	var itemComponentDeliveryScheduleLine	*[]dpfm_api_output_formatter.ItemComponentDeliveryScheduleLine
	var itemComponentPricingElement			*[]dpfm_api_output_formatter.ItemComponentPricingElement
	var itemComponentCosting				*[]dpfm_api_output_formatter.ItemComponentCosting
	var itemOperation						*[]dpfm_api_output_formatter.ItemOperation
	for _, fn := range accepter {
		switch fn {
		case "Header":
			header = c.headerUpdateSql(mtx, input, output, errs, log)
		case "Item":
			item = c.itemUpdateSql(mtx, input, output, errs, log)
		case "ItemComponent":
			itemComponent = c.itemComponentUpdateSql(mtx, input, output, errs, log)
		case "ItemComponentDeliveryScheduleLine":
			itemComponentDeliveryScheduleLine = c.itemComponentDeliveryScheduleLineUpdateSql(mtx, input, output, errs, log)
		case "ItemComponentPricingElement":
			itemComponentPricingElement = c.itemComponentPricingElementUpdateSql(mtx, input, output, errs, log)
		case "ItemComponentCosting":
			itemComponentCosting = c.itemComponentCostingUpdateSql(mtx, input, output, errs, log)
		case "ItemOperation":
			itemOperation = c.itemOperationUpdateSql(mtx, input, output, errs, log)
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header:                         	header,
		Item:                           	item,
		ItemComponent:                	    itemComponent,
		ItemComponentDeliveryScheduleLine:  itemComponentDeliveryScheduleLine,
		ItemComponentPricingElement:   	    itemComponentPricingElement,
		ItemComponentCosting:           	itemComponentCosting,
		ItemOperation:                  	itemOperation,
	}

	return data
}

func (c *DPFMAPICaller) headerCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID

	headerData := subfuncSDC.Message.Header
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "ProductionOrderHeader", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "Header Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemData := range *subfuncSDC.Message.Item {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "ProductionOrderItem", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentData := range *subfuncSDC.Message.ItemComponent {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentData, "function": "ProductionOrderItemComponent", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ItemComponent Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentDeliveryScheduleLineCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentDeliveryScheduleLine {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentDeliveryScheduleLineData := range *subfuncSDC.Message.ItemComponentDeliveryScheduleLine {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentDeliveryScheduleLineData, "function": "ProductionOrderItemComponentDeliveryScheduleLine", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ItemComponentDeliveryScheduleLine Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentDeliveryScheduleLineCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentPricingElementCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentPricingElement {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentPricingElementData := range *subfuncSDC.Message.ItemComponentPricingElement {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentPricingElementData, "function": "ProductionOrderItemComponentPricingElement", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ItemComponentPricingElement Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentPricingElementCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentCostingCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentCosting {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemComponentCostingData := range *subfuncSDC.Message.ItemComponentCosting {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentCostingData, "function": "ProductionOrderItemComponentCosting", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Item Component Costing Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCostingCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemOperationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	subfuncSDC *sub_func_complementer.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for _, itemOperationData := range *subfuncSDC.Message.ItemOperation {
		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemOperationData, "function": "ProductionOrderItemOperation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "ItemOperation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemOperationCreates(subfuncSDC)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) headerUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Header {
	header := input.Header
	headerData := dpfm_api_processing_formatter.ConvertToHeaderUpdates(header)

	sessionID := input.RuntimeSessionID
	if headerIsUpdate(headerData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": headerData, "function": "ProductionOrderHeader", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "Header Data cannot update"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToHeaderUpdates(header)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Item {
	req := make([]dpfm_api_processing_formatter.ItemUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		itemData := *dpfm_api_processing_formatter.ConvertToItemUpdates(item)

		if itemIsUpdate(&itemData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemData, "function": "ProductionOrderItem", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "Item Data cannot update"
				return nil
			}
		}
		req = append(req, itemData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponent {
	req := make([]dpfm_api_processing_formatter.ItemComponentUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			itemComponentData := *dpfm_api_processing_formatter.ConvertToItemComponentUpdates(itemComponent)

			if itemComponentIsUpdate(&itemComponentData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentData, "function": "ProductionOrderItemComponent", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Component Data cannot update"
					return nil
				}
			}
			req = append(req, itemComponentData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentDeliveryScheduleLineUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentDeliveryScheduleLine {
	req := make([]dpfm_api_processing_formatter.ItemComponentDeliveryScheduleLineUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			for _, itemComponentDeliveryScheduleLine := range itemComponent.ItemComponentDeliveryScheduleLine {
				itemComponentDeliveryScheduleLineData := *dpfm_api_processing_formatter.ConvertToItemComponentDeliveryScheduleLineUpdates(itemComponentDeliveryScheduleLine)

				if itemComponentDeliveryScheduleLineIsUpdate(&itemComponentDeliveryScheduleLineData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentDeliveryScheduleLineData, "function": "ProductionOrderItemComponentDeliveryScheduleLine", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "ItemComponentDeliveryScheduleLine Data cannot update"
						return nil
					}
				}
				req = append(req, itemComponentDeliveryScheduleLineData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentDeliveryScheduleLineUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentPricingElementUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentPricingElement {
	req := make([]dpfm_api_processing_formatter.ItemComponentPricingElementUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			for _, itemComponentPricingElement := range itemComponent.ItemComponentPricingElement {
				itemComponentPricingElementData := *dpfm_api_processing_formatter.ConvertToItemComponentPricingElementUpdates(itemComponentPricingElement)

				if itemComponentPricingElementIsUpdate(&itemComponentPricingElementData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentPricingElementData, "function": "ProductionOrderItemComponentPricingElement", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "ItemComponentPricingElement Data cannot update"
						return nil
					}
				}
				req = append(req, itemComponentPricingElementData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentPricingElementUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemComponentCostingUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemComponentCosting {
	req := make([]dpfm_api_processing_formatter.ItemComponentCostingUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemComponent := range item.ItemComponent {
			for _, itemComponentCosting := range itemComponent.ItemComponentCosting {
				itemComponentCostingData := *dpfm_api_processing_formatter.ConvertToItemComponentCostingUpdates(itemComponentCosting)

				if itemComponentCostingIsUpdate(&itemComponentCostingData) {
					res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemComponentCostingData, "function": "ProductionOrderItemComponentCosting", "runtime_session_id": sessionID})
					if err != nil {
						err = xerrors.Errorf("rmq error: %w", err)
						*errs = append(*errs, err)
						return nil
					}
					res.Success()
					if !checkResult(res) {
						output.SQLUpdateResult = getBoolPtr(false)
						output.SQLUpdateError = "ItemComponentCosting Data cannot update"
						return nil
					}
				}
				req = append(req, itemComponentCostingData)
			}
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemComponentCostingUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) itemOperationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.ItemOperation {
	req := make([]dpfm_api_processing_formatter.ItemOperationUpdates, 0)
	sessionID := input.RuntimeSessionID

	header := input.Header
	for _, item := range header.Item {
		for _, itemOperation := range item.ItemOperation {
			itemOperationData := *dpfm_api_processing_formatter.ConvertToItemOperationUpdates(itemOperation)

			if itemOperationIsUpdate(&itemOperationData) {
				res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": itemOperationData, "function": "ProductionOrderItemOperation", "runtime_session_id": sessionID})
				if err != nil {
					err = xerrors.Errorf("rmq error: %w", err)
					*errs = append(*errs, err)
					return nil
				}
				res.Success()
				if !checkResult(res) {
					output.SQLUpdateResult = getBoolPtr(false)
					output.SQLUpdateError = "Operation Data cannot update"
					return nil
				}
			}
			req = append(req, itemOperationData)
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToItemOperationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func headerIsUpdate(header *dpfm_api_processing_formatter.HeaderUpdates) bool {
	productionOrder := header.ProductionOrder

	return !(productionOrder == 0)
}

func itemIsUpdate(item *dpfm_api_processing_formatter.ItemUpdates) bool {
	productionOrder := item.ProductionOrder
	productionOrderItem := item.ProductionOrderItem

	return !(productionOrder == 0 || productionOrderItem == 0)
}

func itemComponentIsUpdate(itemComponent *dpfm_api_processing_formatter.ItemComponentUpdates) bool {
	productionOrder := itemComponent.ProductionOrder
	productionOrderItem := itemComponent.ProductionOrderItem
	billOfMaterial := itemComponent.BillOfMaterial
	billOfMaterialItem := itemComponent.BillOfMaterialItem

	return !(productionOrder == 0 || productionOrderItem == 0 || billOfMaterial == 0 || billOfMaterialItem == 0)
}

func itemComponentDeliveryScheduleLineIsUpdate(itemComponentDeliveryScheduleLine *dpfm_api_processing_formatter.ItemComponentDeliveryScheduleLineUpdates) bool {
	productionOrder := itemComponentDeliveryScheduleLine.ProductionOrder
	productionOrderItem := itemComponentDeliveryScheduleLine.ProductionOrderItem
	billOfMaterial := itemComponentDeliveryScheduleLine.BillOfMaterial
	billOfMaterialItem := itemComponentDeliveryScheduleLine.BillOfMaterialItem
	scheduleLine := itemComponentDeliveryScheduleLine.ScheduleLine

	return !(productionOrder == 0 || productionOrderItem == 0 || billOfMaterial == 0 || billOfMaterialItem == 0 || scheduleLine == 0)
}

func itemComponentPricingElementIsUpdate(itemComponentPricingElement *dpfm_api_processing_formatter.ItemComponentPricingElementUpdates) bool {
	productionOrder := itemComponentPricingElement.ProductionOrder
	productionOrderItem := itemComponentPricingElement.ProductionOrderItem
	billOfMaterial := itemComponentPricingElement.BillOfMaterial
	billOfMaterialItem := itemComponentPricingElement.BillOfMaterialItem
	pricingProcedureCounter := itemComponentPricingElement.PricingProcedureCounter

	return !(productionOrder == 0 || productionOrderItem == 0 || billOfMaterial == 0 || billOfMaterialItem == 0 || pricingProcedureCounter == 0)
}

func itemComponentCostingIsUpdate(itemComponentCosting *dpfm_api_processing_formatter.ItemComponentCostingUpdates) bool {
	productionOrder := itemComponentCosting.ProductionOrder
	productionOrderItem := itemComponentCosting.ProductionOrderItem
	billOfMaterial := itemComponentCosting.BillOfMaterial
	billOfMaterialItem := itemComponentCosting.BillOfMaterialItem

	return !(productionOrder == 0 || productionOrderItem == 0 || billOfMaterial == 0 || billOfMaterialItem == 0)
}

func itemOperationIsUpdate(itemOperation *dpfm_api_processing_formatter.ItemOperationUpdates) bool {
	productionOrder := itemOperation.ProductionOrder
	productionOrderItem := itemOperation.ProductionOrderItem
	operations := itemOperation.Operations
	operationsItem := itemOperation.OperationsItem
	operationID := itemOperation.OperationID

	return !(productionOrder == 0 || productionOrderItem == 0 || operations == 0 || operationsItem == 0 || operationID == 0)
}
