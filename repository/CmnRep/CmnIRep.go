package CmnRep

import (
	"context"

	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
)

//CmnRepIntrfc explain method def...
type CmnRepIntrfc interface {
	GetAuthorizationList_ByRole(ctx context.Context, RoleID int) ([]*CmnModel.Authorization, error)
	GetFeatures_List(ctx context.Context) ([]*CmnModel.Features_List, error)
	Login(ctx context.Context, usr *CmnModel.User) (*CmnModel.User, error)
	Resend_Activation_Link(ctx context.Context, EmpID int) error
	GetEmployees(ctx context.Context, LocID int) ([]*CmnModel.Employees, error)
	GetEmployeeByID(ctx context.Context, id int) (*CmnModel.Employees, error)
	CreateEmployee(ctx context.Context, usr *CmnModel.Employees) (int64, error)
	CreateUser(ctx context.Context, usr *CmnModel.User) (int64, error)
	UpdateEmployee(ctx context.Context, p *CmnModel.Employees) (*CmnModel.Employees, error)
	UpdateUser(ctx context.Context, usr *CmnModel.User) error
	DeleteEmployee(ctx context.Context, id int) error
	Employees_Bulk_Insert(ctx context.Context, mdl []*CmnModel.Employees) error
	GetRoles(ctx context.Context) ([]*CmnModel.Role, error)
	GetDesignations(ctx context.Context) ([]*CmnModel.Designation, error)
	GetEducations(ctx context.Context) ([]*CmnModel.Educations, error)
	GetUsers(ctx context.Context, LocID int) ([]*CmnModel.User, error)
	GetApprovers(ctx context.Context,LocID int , RoleID int) ([]*CmnModel.User, error) 
	Send_ResetPasswordLink(ctx context.Context, empid int) error
	GetStatus(ctx context.Context, typeName string) ([]*CmnModel.Status, error)
	GetVendors(ctx context.Context) ([]*CmnModel.Vendors, error)
	GetCountries(ctx context.Context) ([]*CmnModel.Countries, error)
	GetStatesByCountry(ctx context.Context, id int) ([]*CmnModel.States, error)
	//GenerateUniqueID(ctx context.Context, modulename string) (int64, error)
	GetLocations(ctx context.Context) ([]*CmnModel.Locations, error)
	CreateVendors(ctx context.Context, usr *CmnModel.Vendors) (int64, error)
	CreateLocation(ctx context.Context, usr *CmnModel.Locations) (int64, error)
	UpdateLocations(ctx context.Context, usr *CmnModel.Locations) error
	UpdateVendors(ctx context.Context, usr *CmnModel.Vendors) error
	GetNotifications(ctx context.Context) ([]*CmnModel.Notifications, error)

	CreateInWardOutWard(ctx context.Context, usr *CmnModel.InWardOutWard) error
	CreateOutWardCart(ctx context.Context, obj []*CmnModel.OutWardCart) error
	GetOutWardCart(ctx context.Context, SenderEmpID int) ([]*CmnModel.OutWardCart, error)
	DeleteOutWardCart(ctx context.Context, listInt interface{}) error
	GetOutWardDetailsByEmp(ctx context.Context, SenderEmpID int) ([]*CmnModel.InWardOutWard, error)
	GetInWardDetailsByEmp(ctx context.Context, RcvrEmpID int) ([]*CmnModel.InWardOutWard, error)
	GetOutWardAssetDetailsByIwOwID(ctx context.Context, IwOwID int) ([]*CmnModel.InWardOutWardAsset, error)
	GetAssetdIDsNotEligbleForTransfer(ctx context.Context, trnsfr []*CmnModel.Transfer) ([]*CmnModel.Transfer, error)
	IWOWDetailsByAprvr(ctx context.Context, AprvrEmpID int) ([]*CmnModel.InWardOutWard, error)
	TransferApprovReject(ctx context.Context, idInWardOutWard int, status string,comments string) error
	ReceiveIWAssets(ctx context.Context, obj *CmnModel.InWardOutWard) error
	OwStatusChange(OWid int, Status int) error
	User_Inactive(ctx context.Context, UserID int) error
	UpdateIsMsngStcksRslvdMain(ctx context.Context, IDInWardOutWard int) error
	Authorization_Create(ctx context.Context, data CmnModel.Authorization_Create) error
	User_Active(ctx context.Context, UserID int) error
	Check_Unique_Email_Mobile(ctx context.Context, data CmnModel.Employees) (*CmnModel.Employees, error)
	Check_Unique_UserName(ctx context.Context, UserName string) (*string, error)
	GetEmployee_History_ByEmpID(ctx context.Context, id int) ([]*CmnModel.Employees, error)
	Get_UsersHistory_ByEmpID(ctx context.Context, id int) ([]*CmnModel.Employees, error)
	Activivty_Log_List(ctx context.Context, id int, FromDate string, ToDate string) ([]*CmnModel.ActivityLog, error)
	GetMultiLevelApproval(ctx context.Context) ([]*CmnModel.MultiLevelApproval_Main, error)
	MultiLevelApproval_config(ctx context.Context, mdl []*CmnModel.MultiLevelApproval_Main) error
	GetMultiLevelApproval_Map(ctx context.Context) ([]*CmnModel.MultiLevelApproval_Main, error) 
	VendorsAssetDetails(ctx context.Context, VendorID int) ([]*CmnModel.VendorsAssetDetails, error)
	VednorsAssetMapInsert(ctx context.Context, mdl CmnModel.Vendors_consumablemaster_map) error 
	VednorsAssetMapUpdate(ctx context.Context, mdl CmnModel.Vendors_consumablemaster_map) error 
	IWOW_ApprovalStatusList(ctx context.Context, IDinwardoutward int) ([]*CmnModel.InWardOutWardApproval, error)
	InwardOutwardReqForward(ctx context.Context, mdl CmnModel.InWardOutWardApproval) error 
	GetAdminDashBoard(ctx context.Context, mdl CmnModel.AdminDashBoard) (*CmnModel.AdminDashBoard, error)
	GetEmployeeDashboard(ctx context.Context, mdl CmnModel.EmployeeDashboard) (*CmnModel.EmployeeDashboard, error) 
	PurchaseOrders_RequestsInsert(ctx context.Context, mdl CmnModel.PurchaseOrders_Requests) error
	GetPurchaseOrderUniqueID() (*string, error)
	GetPODetailsByReqstrID(ctx context.Context,ReqstrID int) ([]*CmnModel.PurchaseOrders_Requests, error) 
	PODetailsByIDPO(ctx context.Context,IDPO int) (*CmnModel.PurchaseOrders_Requests, error)
	POAssetDetailsByIDPO(ctx context.Context, IDPO int) ([]*CmnModel.PurchaseOrders_Assets, error) 
	PO_ApprovalStatusList(ctx context.Context, IDPO int) ([]*CmnModel.POApproval, error) 
	GetPODetailsByApprover(ctx context.Context, ApprvrID int) ([]*CmnModel.PurchaseOrders_Requests, error) 
	POReqApproved(ctx context.Context, mdl CmnModel.POApproval) error 
	POReqForward(ctx context.Context, mdl CmnModel.POApproval) error 
	POReqRejected(ctx context.Context, mdl CmnModel.POApproval) error 
	PurchaseOrders_RequestsUpdate(ctx context.Context, mdl CmnModel.PurchaseOrders_Requests) error
	POStatusChange(IDPO int, Status int) error
	Requisition_RequestsInsert(ctx context.Context, mdl CmnModel.Requisition_Requests) error 
	GetRequisitionDetailsByReqstrID(ctx context.Context, ReqstrID int) ([]*CmnModel.Requisition_Requests, error)
	RequisitionDetailsByID(ctx context.Context, ID int) (*CmnModel.Requisition_Requests, error) 
	RequisitionAssetDetailsByID(ctx context.Context, ID int) ([]*CmnModel.Requisition_Assets, error)
	Requisition_ApprovalStatusList(ctx context.Context, ID int) ([]*CmnModel.RequisitionApproval, error) 
	GetRequisitionDetailsByApprover(ctx context.Context, ApprvrID int) ([]*CmnModel.Requisition_Requests, error)
	Requisition_RequestsUpdate(ctx context.Context, mdl CmnModel.Requisition_Requests) error
}
