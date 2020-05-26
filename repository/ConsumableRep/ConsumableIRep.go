package ConsumableRep

import (
	"context"
	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/models/ConsumableModel"
)

// ConsumableIntrfc explain method def...
type ConsumableIntrfc interface {
	CreateConsumable(ctx context.Context, mdl *ConsumableModel.Consumables) error
	UpdateConsumable(ctx context.Context, mdl *ConsumableModel.Consumables) error
	GetConsumableGroups(ctx context.Context) ([]*ConsumableModel.ConsumableGroup, error)
	GetConsumableList(ctx context.Context,LocID int) ([]*ConsumableModel.Consumables, error)
	GetConsumableMasterList(ctx context.Context) ([]*ConsumableModel.Consumablemaster, error)
	GetConsumableListByID(ctx context.Context, IDConsumable int) (*ConsumableModel.Consumables, error)
	GetConsumableOprtnListByID(ctx context.Context, IDConsumable int) ([]*ConsumableModel.Consumables, error)
	GetConsumableOprtnList(ctx context.Context) ([]*ConsumableModel.Consumables, error)
	
	ConsumableOprtnsAddStock(ctx context.Context, mdl *ConsumableModel.ConsumableOprtns) error
	PostConsumableOprtnsRemovestock(ctx context.Context, mdl *ConsumableModel.ConsumableOprtns) error
	ConsumableBulkEdit(ctx context.Context, usr *ConsumableModel.Consumables, idss *[]string) error
	CnsmblBulkRetire(ctx context.Context, mdl []*CmnModel.InWardOutWardAsset) error 
	CheckDuplicateAssetEntry(ctx context.Context, MasterID int,LocID int) (*int, error) 
	ConsumablemasterInsert(ctx context.Context, mdl *ConsumableModel.Consumablemaster) error
	Check_Unique_Consumable(ctx context.Context, ConsumableName string) (*string, error) 
	GetVendorsByConsumable(ctx context.Context,ConsumableID int) ([]*CmnModel.VendorsAssetDetails, error)
	ConsumableBulkDelete(ctx context.Context, idss []string) error 
}
