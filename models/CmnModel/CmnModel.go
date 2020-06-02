package CmnModel

type EmployeeDashboard struct {
	EmpID                  int `json:"EmpID"`
	LocID                  int `json:"LocID"`
	ITAssetsAssigned       int `json:"ITAssetsAssigned"`
	NonITAssetsAssigned    int `json:"NonITAssetsAssigned"`
	ITAssetRequests        int `json:"ITAssetRequests"`
	NonITAssetRequests     int `json:"NonITAssetRequests"`
	ITAssetServiceRequests int `json:"ITAssetServiceRequests"`
}
type AdminDashBoard struct {
	EmpID                  int `json:"EmpID"`
	LocID                  int `json:"LocID"`
	ActivationPendingUsers int `json:"ActivationPendingUsers"`
	InActiveUsers          int `json:"InActiveUsers"`
	ITAssetWarrentyExpired int `json:"ITAssetWarrentyExpired"`

	ITAssetApprovals    int `json:"ITAssetApprovals"`
	NonITAssetApprovals int `json:"NonITAssetApprovals"`

	ITAssetsAvailable      int `json:"ITAssetsAvailable"`
	ITAssetsAssigned       int `json:"ITAssetsAssigned"`
	NonITAssetThreshold    int `json:"NonITAssetThreshold"`
	ConsumableThreshold    int `json:"ConsumableThreshold"`
	OutwardApproval        int `json:"OutwardApproval"`
	ReadyToShip            int `json:"ReadyToShip"`
	InWardAssets           int `json:"InWardAssets"`
	ITAssetServiceRequests int `json:"ITAssetServiceRequests"`
}

//TemplateData ..
type TemplateData struct {
	User LoginModel      `json:"User"`
	Data interface{}     `json:"Data"`
	Auth map[string]bool `json:"Auth"`
}
type Authorization_Create struct {
	RoleID           int
	CreatedBy        int
	Features_List_ID []int
}

//Role ..
type Role struct {
	IDRoles  int    `json:"IDRoles"`
	RoleName string `json:"RoleName"`
}

//Designation ..
type Designation struct {
	IDDesignation   int    `json:"IDDesignation"`
	DesignationName string `json:"DesignationName"`
	CreatedOn       string `json:"CreatedOn"`
}

//Educations ..
type Educations struct {
	IDEducations int    `json:"IDEducations"`
	Name         string `json:"Name"`
	CreatedOn    string `json:"CreatedOn"`
}

//Authorization ..
type Authorization struct {
	IDAuthorization  *int           `json:"IDAuthorization"`
	RoleID           *int           `json:"RoleID"`
	Features_List_ID *int           `json:"Features_List_ID"`
	CreatedOn        *string        `json:"CreatedOn"`
	CreatedBy        *int           `json:"CreatedBy"`
	Features_List    *Features_List `json:"Features_List"`
	Role             *Role          `json:"Role"`
}

//Features_List ..
type Features_List struct {
	IDFeatures_list int    `json:"IDFeatures_list"`
	Feature_Name    string `json:"Feature_Name"`
	Module          string `json:"Module"`
	CreatedOn       string `json:"CreatedOn"`
}

// Employees type details
type Employees struct {
	IDEmployees     int         `json:"IDEmployees"`
	FirstName       string      `json:"FirstName"`
	LastName        string      `json:"LastName"`
	DOB             string      `json:"DOB"`
	Email           string      `json:"Email"`
	EmpCode         string      `json:"EmpCode"`
	Mobile          string      `json:"Mobile"`
	Address         string      `json:"Address"`
	PrmntAddress    string      `json:"PrmntAddress"`
	Status          string      `json:"Status"`
	Image           string      `json:"Image"`
	Education       int         `json:"Education"`
	ExperienceMonth int         `json:"ExperienceMonth"`
	ExperienceYear  int         `json:"ExperienceYear"`
	Designation     int         `json:"Designation"`
	DOJ             string      `json:"DOJ"`
	User            *User       `json:"User"`
	Location        int         `json:"Location"`
	Gender          string      `json:"Gender"`
	ModifiedBy      int         `json:"ModifiedBy"`
	CreatedBy       int         `json:"CreatedBy"`
	CreatedOn       string      `json:"CreatedOn"`
	EducationData   Educations  `json:"EducationData"`
	DesignationData Designation `json:"DesignationData"`
	CreatedByData   *Employees  `json:"CreatedByData"`
	ActionPerformed string      `json:"ActionPerformed"`
}
type LoginModel struct {
	IDUsers     int
	EmployeeID  int
	UserName    string
	RoleID      int
	IDEmployees int
	FirstName   string
	LastName    string
	Email       string
	IDRoles     int
	RoleName    string
	LocationID  int
	Image       string
}
type ActivityLog struct {
	IDHistory           int    `json:"IDHistory"`
	IDMaintable         int    `json:"IDMaintable"`
	CreatedOn           string `json:"CreatedOn"`
	ActionPerformed     string `json:"ActionPerformed"`
	CreatedBy           int    `json:"CreatedBy"`
	Module              string `json:"Module"`
	Name                string `json:"Name"`
	ActionedByFirstName string `json:"ActionedByFirstName"`
	ActionedByLastName  string `json:"ActionedByLastName"`
}

//User ..
type User struct {
	IDUsers           int             `json:"IDUsers"`
	EmployeeID        int             `json:"EmployeeID"`
	UserName          string          `json:"UserName"`
	Password          []byte          `json:"Password"`
	Status            string          `json:"Status"`
	RoleID            int             `json:"RoleID"`
	Employee          Employees       `json:"Employee"`
	ListAuthorization []Authorization `json:"ListAuthorization"`
	LinkGeneratedOn   *string         `json:"LinkGeneratedOn"`
	Role              *Role           `json:"Role"`
	Educations        Educations      `json:"Educations"`
	Designation       Designation     `json:"Designation"`
	ModifiedBy        int             `json:"ModifiedBy"`
	CreatedBy         int             `json:"CreatedBy"`
	CreatedOn         string          `json:"CreatedOn"`
}

// Status type details
type Status struct {
	IDStatus   int    `json:"IDStatus,omitempty"`
	StatusName string `json:"StatusName,omitempty"`
	Type       string `json:"Type"`
}

//Retire ..
type Retire struct {
	IDRetire   int    `json:"IDRetire"`
	AssetID    int    `json:"AssetID"`
	Reason     int    `json:"Reason"`
	RetireDate string `json:"RetireDate"`
	Commnets   string `json:"Commnets"`
	RetiredBy  int    `json:"RetiredBy"`
}

//Countries ..
type Countries struct {
	ID       int    `json:"ID,omitempty"`
	Sortname string `json:"Sortname,omitempty"`
	Name     string `json:"Name,omitempty"`
}

//States ..
type States struct {
	ID        int    `json:"ID,omitempty"`
	Name      string `json:"Name,omitempty"`
	Countryid string `json:"Countryid,omitempty"`
}

//Vendors ..
type Vendors struct {
	Idvendors         int    `json:"Idvendors"`
	Name              string `json:"Name"`
	Description       string `json:"Description"`
	Websites          string `json:"Websites"`
	Address           string `json:"Address"`
	Email             string `json:"Email"`
	ContactPersonName string `json:"ContactPersonName"`
	Phone             string `json:"Phone"`
	Status            string `json:"Status"`
	CreatedBy         int    `json:"CreatedBy"`
	CreatedOn         string `json:"CreatedOn"`
	ModifiedBy        int    `json:"ModifiedBy"`
	ModifiedOn        string `json:"ModifiedOn"`
}

//Locations ..
type Locations struct {
	IDLocations int    `json:"IDLocations"`
	Name        string `json:"Name"`
	Code        string `json:"Code"`
	Address1    string `json:"Address1"`
	Address2    string `json:"Address2"`
	Country     int    `json:"Country"`
	State       int    `json:"State"`
	City        string `json:"City"`
	Zipcode     string `json:"Zipcode"`
	Description string `json:"Description"`
	Countryname string `json:"Countryname"`
	Statename   string `json:"Statename"`
	CreatedBy   int
	CreatedOn   string
	ModifiedBy  int
	ModifiedOn  string
	Status      string `json:"Status"`
}

//Notifications ..
type Notifications struct {
	IDNotifications int    `json:"IDNotifications"`
	EmpID           int    `json:"EmpID"`
	Message         string `json:"Message"`
	MessageType     string `json:"MessageType"`
	RoleID          int    `json:"RoleID"`
	CreatedOn       string `json:"CreatedOn"`
	Name            string `json:"Name"`
}

//InWardOutWard ..
type InWardOutWard struct {
	IDInWardOutWard int    `json:"IDInWardOutWard"`
	TransactionID   string `json:"TransactionID"`
	ToLocationID    int    `json:"ToLocationID"`
	FromLocationID  int    `json:"FromLocationID"`
	SenderEmpID     int    `json:"SenderEmpID"`
	ReceiverEmpID   int    `json:"ReceiverEmpID"`

	Description            string                `json:"Description"`
	TransferStatus         int                   `json:"TransferStatus"`
	CreatedOn              string                `json:"CreatedOn"`
	StatusUpdatedOn        string                `json:"StatusUpdatedOn"`
	ListInWardOutWardAsset []*InWardOutWardAsset `json:"ListInWardOutWardAsset"`
	TotalItems             int                   `json:"TotalItems"`
	IsMsngStcksRslvdMain   []uint8               `json:"IsMsngStcksRslvdMain"`
	SenderEmployee         Employees             `json:"SenderEmployee"`
	ReceiverEmployee       Employees             `json:"ReceiverEmployee"`
	ApproverEmployee       Employees             `json:"ApproverEmployee"`
	FromLocation           Locations             `json:"FromLocation"`
	ToLocation             Locations             `json:"ToLocation"`
	Status                 Status                `json:"Status"`
	CreatedBy              int                   `json:"CreatedBy"`
	InWardOutWardApproval  InWardOutWardApproval `json:"InWardOutWardApproval"`
}

type InWardOutWardApproval struct {
	IDInwardoutward_Approval int    `json:"IDInwardoutward_Approval"`
	IDinwardoutward          int    `json:"IDinwardoutward"`
	RoleID                   int    `json:"RoleID"`
	RoleName                 string `json:"RoleName"`
	ApproverID               int    `json:"ApproverID"`
	ApproverName             string `json:"ApproverName"`
	Grade                    int    `json:"Grade"`
	Comments                 string `json:"Comments"`
	Status                   int    `json:"Status"`
	StatusName               string `json:"StatusName"`
	CreatedOn                string `json:"CreatedOn"`
	ActionedOn               string `json:"ActionedOn"`
	NextRoleID               int    `json:"NextRoleID"`
	NextApproverID           int    `json:"NextApproverID"`
	NextGrade                int    `json:"NextGrade"`
}

//InWardOutWardAsset ..
type InWardOutWardAsset struct {
	IDinwardoutwardAssets int         `json:"IDinwardoutwardAssets"`
	Inwardoutwardid       int         `json:"Inwardoutwardid"`
	AssetID               int         `json:"AssetID"`
	AssetType             string      `json:"AssetType"`
	Quantity              int         `json:"Quantity"`
	ReceivedQuantity      int         `json:"ReceivedQuantity"`
	Description           string      `json:"Description"`
	TransferStatus        int         `json:"TransferStatus"`
	UpdatedOn             string      `json:"UpdatedOn"`
	IsMsngStcksRslvd      bool        `json:"IsMsngStcksRslvd"`
	ITAsset               interface{} `json:"ITAsset"`
	Consumable            interface{} `json:"Consumable"`
	NonITAsset            interface{} `json:"NonITAsset"`
	Status                Status      `json:"Status"`
	CreatedBy             int         `json:"CreatedBy"`
	CreatedOn             string      `json:"CreatedOn"`
}

//OutWardCart  ..
type OutWardCart struct {
	IDOutWardCart int    `json:"IDOutWardCart"`
	AssetID       int    `json:"AssetID"`
	AssetType     string `json:"AssetType"`
	AssetName     string `json:"AssetName"`
	SenderEmpID   int    `json:"SenderEmpID"`
	CreatedOn     string `json:"CreatedOn"`
	CreatedBy     int    `json:"CreatedBy"`
}

//Transfer ..
type Transfer struct {
	AssetID   int    `json:"AssetID"`
	AssetType string `json:"AssetType"`
	Quantity  int    `json:"Quantity"`
}

type MultiLevelApproval_Main struct {
	IDMultiLevelApproval_Main   int                       `json:"IDMultiLevelApproval_Main"`
	FeatureName                 string                    `json:"FeatureName"`
	Module                      string                    `json:"Module"`
	Levels                      int                       `json:"Levels"`
	CreatedBy                   int                       `json:"CreatedBy"`
	CreatedOn                   string                    `json:"CreatedOn"`
	RoleList                    []Role                    `json:"RoleList"`
	MultiLevelApproval_Map_List []*MultiLevelApproval_Map `json:"MultiLevelApproval_Map_List"`
	MultiLevelApproval_Map      *MultiLevelApproval_Map   `json:"MultiLevelApproval_Map"`
}

type MultiLevelApproval_Map struct {
	IDMultiLevelApproval_Map   int    `json:"IDMultiLevelApproval_Map"`
	MultiLevelApproval_Main_ID int    `json:"MultiLevelApproval_Main_ID"`
	RoleID                     int    `json:"RoleID"`
	Grade                      int    `json:"Grade"`
	CreatedOn                  string `json:"CreatedOn"`
	CreatedBy                  int    `json:"CreatedBy"`
	Role                       Role   `json:"Role"`
}

type VendorsAssetDetails struct {
	Vendors_consumablemaster_map Vendors_consumablemaster_map `json:"Vendors_consumablemaster_map"`
	Consumablemaster             interface{}                  `json:"Consumablemaster"`
	Consumablegroups             interface{}                  `json:"Consumablegroups"`
	Vendors                      Vendors                      `json:"Vendors"`
	CreatedBy                    string                       `json:"CreatedBy"`
}
type Vendors_consumablemaster_map struct {
	IDVendors_ConsumableMaster_Map int     `json:"IDVendors_ConsumableMaster_Map"`
	ConsumableMasterID             int     `json:"ConsumableMasterID"`
	VendorsID                      int     `json:"VendorsID"`
	PriceperUnit                   float64 `json:"PriceperUnit"`
	ItemType                       string  `json:"ItemType"`
	VendorRfrdAssetName            string  `json:"VendorRfrdAssetName"`
	CreatedBy                      int     `json:"CreatedBy"`
	CreatedOn                      string  `json:"CreatedOn"`
}

type PurchaseOrders_Requests struct {
	IDPurchaseOrders_Requests int    `json:"IDPurchaseOrders_Requests"`
	POID                      string `json:"POID"`
	LocationID                int    `json:"LocationID"`
	VendorID                  int    `json:"VendorID"`
	//	POApproverID              int                     `json:"POApproverID"`
	PORequestedBy             int                     `json:"PORequestedBy"`
	PODescription             string                  `json:"PODescription"`
	ShipmentTerms             string                  `json:"ShipmentTerms"`
	PaymentTerms              string                  `json:"PaymentTerms"`
	TotalAmmount              float64                 `json:"TotalAmmount"`
	StatusID                  int                     `json:"StatusID"`
	CreatedBy                 int                     `json:"CreatedBy"`
	ModifiedBy                int                     `json:"ModifiedBy"`
	CreatedOn                 string                  `json:"CreatedOn"`
	ModifiedOn                string                  `json:"ModifiedOn"`
	RecordStatus              string                  `json:"RecordStatus"`
	PORequestedByName         string                  `json:"PORequestedByName"`
	StatusName                string                  `json:"StatusName"`
	ListPurchaseOrders_Assets []PurchaseOrders_Assets `json:"ListPurchaseOrders_Assets"`
	VendorData                Vendors                 `json:"VendorData"`
	LocationData              Locations               `json:"LocationData"`
	POApproval                POApproval              `json:"POApproval"`
}

type PurchaseOrders_Assets struct {
	IDpurchaseorders_Assets    int     `json:"IDpurchaseorders_Assets"`
	Purchaseorders_requests_ID int     `json:"Purchaseorders_requests_ID"`
	AssetType                  string  `json:"AssetType"`
	AssetName                  string  `json:"AssetName"`
	AssetID                    int     `json:"AssetID"`
	PriceperUnit               float64 `json:"PriceperUnit"`
	Quantity                   int     `json:"Quantity"`
	AssetComments              string  `json:"AssetComments"`
	CreatedBy                  int     `json:"CreatedBy"`
	CreatedOn                  string  `json:"CreatedOn"`
	ModifiedOn                 string  `json:"ModifiedOn"`
	ModifiedBy                 int     `json:"ModifiedBy"`
}

type POApproval struct {
	IDPO_approval             int    `json:"IDPO_approval"`
	PurchaseOrders_RequestsID int    `json:"PurchaseOrders_RequestsID"`
	RoleID                    int    `json:"RoleID"`
	RoleName                  string `json:"RoleName"`
	ApproverID                int    `json:"ApproverID"`
	ApproverName              string `json:"ApproverName"`
	Grade                     int    `json:"Grade"`
	Comments                  string `json:"Comments"`

	StatusName     string `json:"StatusName"`
	CreatedOn      string `json:"CreatedOn"`
	ActionedOn     string `json:"ActionedOn"`
	NextRoleID     int    `json:"NextRoleID"`
	NextApproverID int    `json:"NextApproverID"`
	NextGrade      int    `json:"NextGrade"`
}

type PurchaseOrders_Assets_Received struct {
	IDpurchaseorders_Assets_Received int     `json:"IDpurchaseorders_Assets_Received"`
	Purchaseorders_requests_ID       int     `json:"Purchaseorders_requests_ID"`
	AssetType                        string  `json:"AssetType"`
	AssetName                        string  `json:"AssetName"`
	AssetID                          int     `json:"AssetID"`
	PriceperUnit                     float64 `json:"PriceperUnit"`
	Quantity                         int     `json:"Quantity"`
	AssetComments                    string  `json:"AssetComments"`
	CreatedBy                        int     `json:"CreatedBy"`
	CreatedOn                        string  `json:"CreatedOn"`
	ModelNo                          string  `json:"ModelNo"`
	SerialNo                         string  `json:"SerialNo"`
}

type PO_Bills struct {
	IDPO_Bills                int    `json:"IDPO_Bills"`
	PurchaseOrders_RequestsID int    `json:"PurchaseOrders_RequestsID"`
	InvoiceNumber             string `json:"InvoiceNumber"`
}
