package NonITAssets_mdl

import (
	"time"

	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
)

//NonITAssets ..
type NonITAssets struct {
	IDNonITAsset                 int                           `json:"IDNonITAsset"`
	NonITAssets_Master_ID        int                           `json:"NonITAssets_Master_ID"`
	ModelNo                      string                        `json:"ModelNo"`
	IdentificationNo             string                        `json:"IdentificationNo"`
	Description                  string                        `json:"Description"`
	Img                          string                        `json:"Img"`
	TotalQnty                    int                           `json:"TotalQnty"`
	AvailableQnty                int                           `json:"AvailableQnty"`
	InUseQnty                    int                           `json:"InUseQnty"`
	ThresholdQnty                int                           `json:"ThresholdQnty"`
	ReOrderStockPrice            float64                       `json:"ReOrderStockPrice"`
	ReOrderQuantity              int                           `json:"ReOrderQuantity"`
	StatusID                     int                           `json:"StatusID"`
	LocationID                   int                           `json:"LocationID"`
	Created_On                   time.Time                     `json:"Created_On"`
	Modified_On                  time.Time                     `json:"Modified_On"`
	Created_By                   int                           `json:"Created_By"`
	Modified_By                  int                           `json:"Modified_By"`
	CustomFields1                string                        `json:"CustomFields1"`
	CustomFields1Value           string                        `json:"CustomFields1Value"`
	CustomFields1Type            string                        `json:"CustomFields1Type"`
	CustomFields2                string                        `json:"CustomFields2"`
	CustomFields2Value           string                        `json:"CustomFields2Value"`
	CustomFields2Type            string                        `json:"CustomFields2Type"`
	CustomFields3                string                        `json:"CustomFields3"`
	CustomFields3Value           string                        `json:"CustomFields3Value"`
	CustomFields3Type            string                        `json:"CustomFields3Type"`
	CustomFields4                string                        `json:"CustomFields4"`
	CustomFields4Value           string                        `json:"CustomFields4Value"`
	CustomFields4Type            string                        `json:"CustomFields4Type"`
	CustomFields5                string                        `json:"CustomFields5"`
	CustomFields5Value           string                        `json:"CustomFields5Value"`
	CustomFields5Type            string                        `json:"CustomFields5Type"`
	Locations                    CmnModel.Locations            `json:"Locations"`
	Vendors                      CmnModel.Vendors              `json:"Vendors"`
	Status                       CmnModel.Status               `json:"Status"`
	NonITAssets_Groups           NonITAssets_Groups            `json:"NonITAssets_Groups"`
	NonITAssets_Master           NonITAssets_Master            `json:"NonITAssets_Master"`
	NonITAssets_checkout_checkin *NonITAssets_checkout_checkin `json:"NonITAssets_checkout_checkin"`
	NonITAssets_Oprtns           NonITAssets_Oprtns            `json:"NonITAssets_Oprtns"`
	RecordStatus                 string                        `json:"RecordStatus"`
}

//NonITAssets_Master ..
type NonITAssets_Master struct {
	IDNonITAssets_Master   int                `json:"IDNonITAssets_Master"`
	NonITAssets_Name       string             `json:"NonITAssets_Name"`
	NonITAssets_GroupID    int                `json:"NonITAssets_GroupID"`
	NonITAssets_GroupName  string             `json:"NonITAssets_GroupName"`
	NonITAssets_SubGroupID int                `json:"NonITAssets_SubGroupID"`
	Description            string             `json:"Description"`
	Created_On             time.Time          `json:"Created_On"`
	Modified_On            time.Time          `json:"Modified_On"`
	Created_By             int                `json:"Created_By"`
	Modified_By            int                `json:"Modified_By"`
	Record_Status          string             `json:"Record_Status"`
	NonITAssets_Groups     NonITAssets_Groups `json:"NonITAssets_Groups"`
}

// NonITAssets_Groups ..
type NonITAssets_Groups struct {
	IDNonITAssets_Groups  int       `json:"IDNonITAssets_Groups"`
	NonITAssets_GroupName string    `json:"NonITAssets_GroupName"`
	CreatedOn             time.Time `json:"CreatedOn"`
	CreatedBy             int       `json:"CreatedBy"`
	ModifiedOn            time.Time `json:"ModifiedOn"`
	ModifiedBy            int       `json:"ModifiedBy"`
}

//NonITAssets_checkout_checkin ..
type NonITAssets_checkout_checkin struct {
	NonITAssets                    *NonITAssets `json:"NonITAssets"`
	IDNonITAssets_Checkout_Checkin int          `json:"IDNonITAssets_Checkout_Checkin"`
	NonITAsset_ID                  int          `json:"NonITAsset_ID"`
	CheckedOutTo                   string       `json:"CheckedOutTo"`
	CheckedOutUserID               int          `json:"CheckedOutUserID"`
	CheckedOutPlace                string       `json:"CheckedOutPlace"`
	CheckedOutDate                 time.Time    `json:"CheckedOutDate"`
	CheckOut_Qnty                  int          `json:"CheckOut_Qnty"`
	InUse                          int          `json:"InUse"`
	CheckOut_Comments              string       `json:"CheckOut_Comments"`
	Record_Status                  string       `json:"Record_Status"`
	Created_On                     time.Time    `json:"Created_On"`
	Created_By                     int          `json:"Created_By"`
	CheckOut_By                    int          `json:"CheckOut_By"`
	Created_ByName                 string       `json:"Created_ByName"`
	CheckOut_ByName                string       `json:"CheckOut_ByName"`
	CheckOut_UserName              string       `json:"CheckOut_UserName"`
	AssetName                      string       `json:"AssetName"`
}

type NonITAssets_checkin struct {
	IDnonitassets_checkin          int       `json:"IDnonitassets_checkin"`
	NonITAssets_Checkout_CheckinID int       `json:"NonITAssets_Checkout_CheckinID"`
	CheckIN_By                     int       `json:"CheckIN_By"`
	CheckIN_ByName                 string    `json:"CheckIN_ByName"`
	CheckIn_Qnty                   int       `json:"CheckIn_Qnty"`
	CheckinDate                    time.Time `json:"CheckinDate"`
	Checkin_Comments               string    `json:"Checkin_Comments"`
}

//NonITAssets_Oprtns ..
type NonITAssets_Oprtns struct {
	IDnonitassets_Oprtns int              `json:"IDnonitassets_Oprtns"`
	NonITAsset_ID        int              `json:"NonITAsset_ID"`
	Quantity             int              `json:"Quantity"`
	Warranty             time.Time        `json:"Warranty"`
	UnitPrice            float64          `json:"UnitPrice"`
	VendorID             int              `json:"VendorID"`
	OrderedBy            int              `json:"OrderedBy"`
	Comments             string           `json:"Comments"`
	StatusID             int              `json:"StatusID"`
	Created_On           time.Time        `json:"Created_On"`
	Created_By           int              `json:"Created_By"`
	CreatedByName        string           `json:"CreatedByName"`
	Vendor               CmnModel.Vendors `json:"Vendor"`
	Status               CmnModel.Status  `json:"Status"`
}

type NonITAssetRequest struct {
	IDNonITAssetrequest       int                       `json:"IDNonITAssetrequest"`
	RequestedBy               int                       `json:"RequestedBy"`
	AssetType                 string                    `json:"AssetType"`
	AssetID                   int                       `json:"AssetID"`
	AssignedQnty              int                       `json:"AssignedQnty"`
	Description               string                    `json:"Description"`
	RequestedOn               time.Time                 `json:"RequestedOn"`
	Priority                  string                    `json:"Priority"`
	ReqStatus                 string                    `json:"ReqStatus"`
	ReqGroupID                int                       `json:"ReqGroupID"`
	CreatedBy                 int                       `json:"CreatedBy"`
	CreatedOn                 time.Time                 `json:"CreatedOn"`
	NonITAssetRequestApproval NonITAssetRequestApproval `json:"NonITAssetRequestApproval"`
}
type NonITAssetRequestApproval struct {
	IDNonitasset_Req_approvals int                `json:"IDNonitasset_Req_approvals"`
	NonITAssetReqGroupID       int                `json:"NonTAssetReqGroupID"`
	RoleID                     int                `json:"RoleID"`
	ApproverID                 int                `json:"ApproverID"`
	Grade                      int                `json:"Grade"`
	CreatedOn                  time.Time          `json:"CreatedOn"`
	Status                     string             `json:"Status"`
	Comments                   string             `json:"Comments"`
	ActionedOn                 time.Time          `json:"ActionedOn"`
	CreatedBy                  int                `json:"CreatedBy"`
	Employee                   CmnModel.Employees `json:"Employee"`
	RoleName                   string             `json:"RoleName"`
	NextRoleID                 int                `json:"NextRoleID"`
	NextApproverID             int                `json:"NextApproverID"`
	NextGrade                  int                `json:"NextGrade"`
}

//ITAssetReqList ..
type NonITAssetReqList struct {
	NonITAssetRequest  NonITAssetRequest  `json:"NonITAssetRequest"`
	NonITAssets_Groups NonITAssets_Groups `json:"NonITAssets_Groups"`
	Approver           CmnModel.Employees `json:"Approver"`
	Requester          CmnModel.Employees `json:"Requester"`
	NonITAssets        NonITAssets        `json:"NonITAssets"`
}
