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

func ConvertToItemOperationsCreates(subfuncSDC *sub_func_complementer.SDC) (*[]ItemOperations, error) {
	itemOperations := make([]ItemOperations, 0)

	for _, data := range *subfuncSDC.Message.ItemOperations {
		itemOperation, err := TypeConverter[*ItemOperations](data)
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

func ConvertToItemOperationsUpdates(itemOperationsUpdates *[]dpfm_api_processing_formatter.ItemOperationsUpdates) (*[]ItemOperations, error) {
	itemOperations := make([]ItemOperations, 0)

	for _, data := range *itemOperationsUpdates {
		itemoperation, err := TypeConverter[*ItemOperations](data)
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
