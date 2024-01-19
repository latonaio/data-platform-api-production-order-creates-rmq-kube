package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-production-order-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"data-platform-api-production-order-creates-rmq-kube/sub_func_complementer"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(subfuncSDC *sub_func_complementer.SDC) (*Header, error) {
	data := subfuncSDC.Message.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(subfuncSDC *sub_func_complementer.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *subfuncSDC.Message.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemComponentCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemComponent, error) {
	itemComponents := make([]ItemComponent, 0)

	for _, data := range *subfuncSDC.Message.ItemComponent {
		itemComponent, err := TypeConverter[*ItemComponent](data)
		if err != nil {
			return nil, err
		}

		itemComponents = append(itemComponents, *itemComponent)
	}

	return &itemComponents, nil
}

func ConvertToItemComponentDeliveryScheduleLineCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemComponentDeliveryScheduleLine, error) {
	itemComponentDeliveryScheduleLines := make([]ItemComponentDeliveryScheduleLine, 0)

	for _, data := range *subfuncSDC.Message.ItemComponentDeliveryScheduleLine {
		itemComponentDeliveryScheduleLine, err := TypeConverter[*ItemComponentDeliveryScheduleLine](data)
		if err != nil {
			return nil, err
		}

		itemComponentDeliveryScheduleLines = append(itemComponentDeliveryScheduleLines, *itemComponentDeliveryScheduleLine)
	}

	return &itemComponentDeliveryScheduleLines, nil
}

func ConvertToItemComponentPricingElementCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemComponentPricingElement, error) {
	itemComponentPricingElements := make([]ItemComponentPricingElement, 0)

	for _, data := range *subfuncSDC.Message.ItemComponentPricingElement {
		itemComponentPricingElement, err := TypeConverter[*ItemComponentPricingElement](data)
		if err != nil {
			return nil, err
		}

		itemComponentPricingElements = append(itemComponentPricingElements, *itemComponentPricingElement)
	}

	return &itemComponentPricingElements, nil
}

func ConvertToItemComponentCostingCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemComponentCosting, error) {
	itemComponentCostings := make([]ItemComponentCosting, 0)

	for _, data := range *subfuncSDC.Message.ItemComponentCosting {
		itemComponentCosting, err := TypeConverter[*ItemComponentCosting](data)
		if err != nil {
			return nil, err
		}

		itemComponentCostings = append(itemComponentCostings, *itemComponentCosting)
	}

	return &itemComponentCostings, nil
}

func ConvertToItemOperationCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemOperation, error) {
	itemOperations := make([]ItemOperation, 0)

	for _, data := range *subfuncSDC.Message.ItemOperation {
		itemOperation, err := TypeConverter[*ItemOperation](data)
		if err != nil {
			return nil, err
		}

		itemOperations = append(itemOperations, *itemOperation)
	}

	return &itemOperations, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemComponentUpdates(itemComponentUpdates *[]dpfm_api_processing_formatter.ItemComponentUpdates) (*[]ItemComponent, error) {
	itemComponents := make([]ItemComponent, 0)

	for _, data := range *itemComponentUpdates {
		itemComponent, err := TypeConverter[*ItemComponent](data)
		if err != nil {
			return nil, err
		}

		itemComponents = append(itemComponents, *itemComponent)
	}

	return &itemComponents, nil
}

func ConvertToItemComponentDeliveryScheduleLineUpdates(itemComponentDeliveryScheduleLineUpdates *[]dpfm_api_processing_formatter.ItemComponentDeliveryScheduleLineUpdates) (*[]ItemComponentDeliveryScheduleLine, error) {
	itemComponentDeliveryScheduleLines := make([]ItemComponentDeliveryScheduleLine, 0)

	for _, data := range *itemComponentDeliveryScheduleLineUpdates {
		itemComponentDeliveryScheduleLine, err := TypeConverter[*ItemComponentDeliveryScheduleLine](data)
		if err != nil {
			return nil, err
		}

		itemComponentDeliveryScheduleLines = append(itemComponentDeliveryScheduleLines, *itemComponentDeliveryScheduleLine)
	}

	return &itemComponentDeliveryScheduleLines, nil
}

func ConvertToItemComponentPricingElementUpdates(itemComponentPricingElementLineUpdates *[]dpfm_api_processing_formatter.ItemComponentPricingElementUpdates) (*[]ItemComponentPricingElement, error) {
	itemComponentPricingElements := make([]ItemComponentPricingElement, 0)

	for _, data := range *itemComponentPricingElementUpdates {
		itemComponentPricingElement, err := TypeConverter[*ItemComponentPricingElement](data)
		if err != nil {
			return nil, err
		}

		itemComponentPricingElements = append(itemComponentPricingElements, *itemComponentPricingElement)
	}

	return &itemComponentPricingElements, nil
}

func ConvertToItemComponentCostingUpdates(itemComponentCostingUpdates *[]dpfm_api_processing_formatter.ItemComponentCostingUpdates) (*[]ItemComponentCosting, error) {
	itemComponentCostings := make([]ItemComponentCosting, 0)

	for _, data := range *itemComponentCostingUpdates {
		itemComponentCosting, err := TypeConverter[*ItemComponentCosting](data)
		if err != nil {
			return nil, err
		}

		itemComponentCostings = append(itemComponentCostings, *itemComponentCosting)
	}

	return &itemComponentCostings, nil
}

func ConvertToItemOperationUpdates(itemOperationUpdates *[]dpfm_api_processing_formatter.ItemOperationUpdates) (*[]ItemOperation, error) {
	itemOperations := make([]ItemOperation, 0)

	for _, data := range *itemOperationUpdates {
		itemoperation, err := TypeConverter[*ItemOperation](data)
		if err != nil {
			return nil, err
		}

		itemOperations = append(itemOperations, *itemoperation)
	}

	return &itemOperations, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
