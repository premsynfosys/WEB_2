package ConsumableModel

import (
	"time"

	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
)

//Consumablemaster ..
type Consumablemaster struct {
	IDconsumableMaster int                `json:"IDconsumableMaster"`
	ConsumableName     string             `json:"ConsumableName"`
	GroupID            int                `json:"GroupID"`
	GroupName          string             `json:"GroupName"`
	ConsumableGroups   []*ConsumableGroup `json:"ConsumableGroup"`
	SubGroupID         int                `json:"SubGroupID"`
	Description        string             `json:"Description"`
	CreatedOn          time.Time          `json:"CreatedOn"`
	ModifiedOn         time.Time          `json:"ModifiedOn"`
}

//Consumables ..
type Consumables struct {
	IDConsumables      int                `json:"IDConsumables"`
	IDconsumableMaster int                `json:"IDconsumableMaster"`
	IdentificationNo   string             `json:"IdentificationNo"`
	Img                string             `json:"Img"`
	Description        string             `json:"Description"`
	TotalQnty          int                `json:"TotalQnty"`
	ThresholdQnty      int                `json:"ThresholdQnty"`
	ReOrderStockPrice  float64            `json:"ReOrderStockPrice"`
	ReOrderQuantity    int                `json:"ReOrderQuantity"`
	Warranty           time.Time             `json:"Warranty"`
	StatusID           int                `json:"StatusID"`
	LocationID         int                `json:"LocationID"`
	CreatedOn          time.Time             `json:"CreatedOn"`
	CreatedByName      string             `json:"CreatedByName"`
	ModifiedOn         time.Time             `json:"ModifiedOn"`
	CreatedBy          int                `json:"CreatedBy"`
	ModifiedBy         int                `json:"ModifiedBy"`
	CustomFields1      string             `json:"CustomFields1"`
	CustomFields1Value string             `json:"CustomFields1Value"`
	CustomFields1Type  string             `json:"CustomFields1Type"`
	CustomFields2      string             `json:"CustomFields2"`
	CustomFields2Value string             `json:"CustomFields2Value"`
	CustomFields2Type  string             `json:"CustomFields2Type"`
	CustomFields3      string             `json:"CustomFields3"`
	CustomFields3Value string             `json:"CustomFields3Value"`
	CustomFields3Type  string             `json:"CustomFields3Type"`
	CustomFields4      string             `json:"CustomFields4"`
	CustomFields4Value string             `json:"CustomFields4Value"`
	CustomFields4Type  string             `json:"CustomFields4Type"`
	CustomFields5      string             `json:"CustomFields5"`
	CustomFields5Value string             `json:"CustomFields5Value"`
	CustomFields5Type  string             `json:"CustomFields5Type"`
	Location           CmnModel.Locations `json:"Location"`
	Status             CmnModel.Status    `json:"Status"`
	ConsumableOprtns   ConsumableOprtns   `json:"ConsumableOprtns"`
	Consumablemaster   Consumablemaster   `json:"Consumablemaster"`
	ListStatus         []*CmnModel.Status `json:"ListStatus"`
}

//ConsumableOprtns ..
type ConsumableOprtns struct {
	IDconsumableOprtns int              `json:"IDconsumableOprtns"`
	ConsumableID       int              `json:"ConsumableID"`
	Quantity           int              `json:"Quantity"`
	UnitPrice          float64          `json:"UnitPrice"`
	VendorID           int              `json:"VendorID"`
	OrderedBy          int              `json:"OrderedBy"`
	Comments           string           `json:"Comments"`
	CreataedOn         time.Time           `json:"CreataedOn"`
	StatusID           int              `json:"StatusID"`
	Vendor             CmnModel.Vendors `json:"Vendor"`
	Status             *CmnModel.Status `json:"Status"`
	CreatedByName      string           `json:"CreatedByName"`
}

// ConsumableGroup type details
type ConsumableGroup struct {
	IDconsumablegroups  int    `json:"IDconsumablegroups"`
	ConsumableGroupName string `json:"ConsumableGroupName"`
}

//Consumables_Retire .. 
type Consumables_Retire struct {
	IDConsumables_Retire int    `json:"IDConsumables_Retire"`
	AssetID              int    `json:"AssetID"`
	ReasonID             int    `json:"ReasonID"`
	RetireDate           time.Time `json:"RetireDate"`
	Comments             string `json:"Comments"`
}
