package CmnRep

import (
	"context"
	"encoding/json"
	"fmt"

	CmnModel "github.com/premsynfosys/AMS_WEB/models/CmnModel"
	utils "github.com/premsynfosys/AMS_WEB/utils"
)

// NewAPIRepo retunrs implement of post repository interface
func NewAPIRepo(api string) CmnRepIntrfc {
	return &APIRepo{
		APIConn: api,
	}
}

//APIRepo ..
type APIRepo struct {
	APIConn string
}

//GetEmployees ..
func (m *APIRepo) GetEmployees(ctx context.Context, LocID int) ([]*CmnModel.Employees, error) {
	url := fmt.Sprintf(m.APIConn+"/employees?LocID=%d", LocID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Employees, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetEmployeeByID ..
func (m *APIRepo) GetEmployeeByID(ctx context.Context, id int) (*CmnModel.Employees, error) {
	url := fmt.Sprintf(m.APIConn+"/employee/%d", id)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := new(CmnModel.Employees)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//CreateEmployee ..
func (m *APIRepo) CreateEmployee(ctx context.Context, usr *CmnModel.Employees) (int64, error) {
	url := fmt.Sprintf(m.APIConn + "/employee/create")
	j, err := json.Marshal(usr)
	if err != nil {
		return 0, err
	}
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return 0, err
	}
	var data map[string]int64
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return 0, err
	} else {
		return data["id"], nil
	}
}

//CreateUser ..
func (m *APIRepo) CreateUser(ctx context.Context, usr *CmnModel.User) (int64, error) {
	url := fmt.Sprintf(m.APIConn + "/user/create")
	j, err := json.Marshal(usr)
	if err != nil {
		return 0, err
	}
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return 0, err
	}
	var data map[string]int64
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return 0, err
	} else {
		return data["id"], nil
	}
}

//UpdateEmployee ..
func (m *APIRepo) UpdateEmployee(ctx context.Context, usr *CmnModel.Employees) (*CmnModel.Employees, error) {
	url := fmt.Sprintf(m.APIConn + "/employee/update")
	j, err := json.Marshal(usr)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return nil, err
	}
	data := new(CmnModel.Employees)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

//DeleteEmployee ..
func (m *APIRepo) DeleteEmployee(ctx context.Context, id int) error {
	url := fmt.Sprintf(m.APIConn+"/employeeDelete/%d", id)
	_, err := utils.DeleteRequest(url)

	return err

}

//GetUsers ..
func (m *APIRepo) GetUsers(ctx context.Context, LocID int) ([]*CmnModel.User, error) {
	url := fmt.Sprintf(m.APIConn+"/GetUsers?LocID=%d", LocID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.User, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetRoles ..
func (m *APIRepo) GetRoles(ctx context.Context) ([]*CmnModel.Role, error) {
	url := fmt.Sprintf(m.APIConn + "/GetRoles")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Role, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetEducations ..
func (m *APIRepo) GetEducations(ctx context.Context) ([]*CmnModel.Educations, error) {
	url := fmt.Sprintf(m.APIConn + "/GetEducations")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Educations, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetDesignations ..
func (m *APIRepo) GetDesignations(ctx context.Context) ([]*CmnModel.Designation, error) {
	url := fmt.Sprintf(m.APIConn + "/GetDesignations")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Designation, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//OwStatusChange ..
func (m *APIRepo) OwStatusChange(OWid int, Status int) error {
	url := fmt.Sprintf(m.APIConn+"/OwStatusChange/%d/%d", OWid, Status)
	_, err := utils.GetRequest(url)
	return err
}

//IWOWDetailsByAprvr ..
func (m *APIRepo) IWOWDetailsByAprvr(ctx context.Context, AprvrEmpID int) ([]*CmnModel.InWardOutWard, error) {
	url := fmt.Sprintf(m.APIConn+"/IWOWDetailsByAprvr/%d", AprvrEmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.InWardOutWard, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//TransferApprovReject ..
func (m *APIRepo) TransferApprovReject(ctx context.Context, idInWardOutWard int, status string, comments string) error {
	url := fmt.Sprintf(m.APIConn+"/TransferApprovReject/%d/%v?comments=%v", idInWardOutWard, status, comments)
	_, err := utils.GetRequest(url)
	return err
}

//ReceiveIWAssets ..
func (m *APIRepo) ReceiveIWAssets(ctx context.Context, obj *CmnModel.InWardOutWard) error {
	url := fmt.Sprintf(m.APIConn + "/ReceiveIWAssets")
	j, err := json.Marshal(obj)
	_, err = utils.PostRequest(url, j)
	return err
}

//Login ..
func (m *APIRepo) Login(ctx context.Context, usr *CmnModel.User) (*CmnModel.User, error) {
	url := fmt.Sprintf(m.APIConn + "/Login")
	mdl := CmnModel.User{
		UserName: usr.UserName,
		Password: usr.Password,
	}
	j, err := json.Marshal(mdl)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return nil, err
	}
	var res *CmnModel.User
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//GetOutWardCart ..
func (m *APIRepo) GetOutWardCart(ctx context.Context, SenderEmpID int) ([]*CmnModel.OutWardCart, error) {
	url := fmt.Sprintf(m.APIConn+"/GetOutWardCart/%d", SenderEmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.OutWardCart, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//DeleteOutWardCart ..
func (m *APIRepo) DeleteOutWardCart(ctx context.Context, listInt interface{}) error {
	url := fmt.Sprintf(m.APIConn + "/DeleteOutWardCart")
	j, err := json.Marshal(listInt)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//GetOutWardDetailsByEmp ..
func (m *APIRepo) GetOutWardDetailsByEmp(ctx context.Context, SenderEmpID int) ([]*CmnModel.InWardOutWard, error) {
	url := fmt.Sprintf(m.APIConn+"/OutWardDetailsByEmp/%d", SenderEmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.InWardOutWard, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetInWardDetailsByEmp ..
func (m *APIRepo) GetInWardDetailsByEmp(ctx context.Context, RcvrEmpID int) ([]*CmnModel.InWardOutWard, error) {
	url := fmt.Sprintf(m.APIConn+"/InWardDetailsByEmp/%d", RcvrEmpID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.InWardOutWard, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetOutWardAssetDetailsByIwOwID ..
func (m *APIRepo) GetOutWardAssetDetailsByIwOwID(ctx context.Context, IwOwID int) ([]*CmnModel.InWardOutWardAsset, error) {
	url := fmt.Sprintf(m.APIConn+"/OutWardAssetDetailsByIwOwID/%d", IwOwID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.InWardOutWardAsset, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetAssetdIDsNotEligbleForTransfer ..
func (m *APIRepo) GetAssetdIDsNotEligbleForTransfer(ctx context.Context, trnsfr []*CmnModel.Transfer) ([]*CmnModel.Transfer, error) {
	url := fmt.Sprintf(m.APIConn + "/GetAssetdIDsNotEligbleForTransfer")
	j, err := json.Marshal(trnsfr)
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return nil, err
	}
	var res []*CmnModel.Transfer
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

//CreateInWardOutWard ..
func (m *APIRepo) CreateInWardOutWard(ctx context.Context, usr *CmnModel.InWardOutWard) error {
	url := fmt.Sprintf(m.APIConn + "/CreateInWardOutWard")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err
}

//CreateVendors ..
func (m *APIRepo) CreateVendors(ctx context.Context, usr *CmnModel.Vendors) (int64, error) {
	url := fmt.Sprintf(m.APIConn + "/CreateVendors")
	j, err := json.Marshal(usr)
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return 0, err
	}
	var data map[string]int64
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return 0, err
	} else {
		return data["id"], nil
	}
}

//UpdateVendors ..
func (m *APIRepo) UpdateVendors(ctx context.Context, usr *CmnModel.Vendors) error {
	url := fmt.Sprintf(m.APIConn + "/UpdateVendors")
	j, err := json.Marshal(usr)
	_, err = utils.PostRequest(url, j)
	if err != nil {
		return err
	}
	return err

}

func (m *APIRepo) DeleteVendors(ctx context.Context, usr *CmnModel.Vendors) error {
	url := fmt.Sprintf(m.APIConn + "/DeleteVendors")
	j, err := json.Marshal(usr)
	_, err = utils.PostRequest(url, j)
	if err != nil {
		return err
	}
	return err

}

//UpdateLocations ..
func (m *APIRepo) UpdateLocations(ctx context.Context, usr *CmnModel.Locations) error {
	url := fmt.Sprintf(m.APIConn + "/UpdateLocations")
	j, err := json.Marshal(usr)
	_, err = utils.PostRequest(url, j)
	if err != nil {
		return err
	}
	return err

}

//CreateOutWardCart ..
func (m *APIRepo) CreateOutWardCart(ctx context.Context, obj []*CmnModel.OutWardCart) error {
	url := fmt.Sprintf(m.APIConn + "/CreateOutWardCart")
	j, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err

}

//GetLocations ..
func (m *APIRepo) GetLocations(ctx context.Context) ([]*CmnModel.Locations, error) {
	url := fmt.Sprintf(m.APIConn + "/GetLocations")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Locations, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetNotifications ..
func (m *APIRepo) GetNotifications(ctx context.Context) ([]*CmnModel.Notifications, error) {
	url := fmt.Sprintf(m.APIConn + "/GetNotifications")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Notifications, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//CreateLocation ..
func (m *APIRepo) CreateLocation(ctx context.Context, usr *CmnModel.Locations) (int64, error) {
	url := fmt.Sprintf(m.APIConn + "/CreateLocations")
	j, err := json.Marshal(usr)
	bytes, err := utils.PostRequest(url, j)
	if err != nil {
		return 0, err
	}
	var data map[string]int64
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return 0, err
	} else {
		return data["id"], nil
	}
}

//GetStatesByCountry ..
func (m *APIRepo) GetStatesByCountry(ctx context.Context, id int) ([]*CmnModel.States, error) {
	url := fmt.Sprintf(m.APIConn+"/GetStatesByCountry/%d", id)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.States, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetStatus ..
func (m *APIRepo) GetStatus(ctx context.Context, typeName string) ([]*CmnModel.Status, error) {
	url := fmt.Sprintf(m.APIConn + "/status/" + typeName)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Status, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetVendors ..
func (m *APIRepo) GetVendors(ctx context.Context) ([]*CmnModel.Vendors, error) {
	url := fmt.Sprintf(m.APIConn + "/GetVendors")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Vendors, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetCountries ..
func (m *APIRepo) GetCountries(ctx context.Context) ([]*CmnModel.Countries, error) {
	url := fmt.Sprintf(m.APIConn + "/GetCountries")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Countries, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//UpdateUser ..
func (m *APIRepo) UpdateUser(ctx context.Context, usr *CmnModel.User) error {
	url := fmt.Sprintf(m.APIConn + "/user/update")
	j, err := json.Marshal(usr)
	if err != nil {
		return err
	}
	_, err = utils.PutRequest(url, j)

	return err

}

// //GenerateUniqueID ..
// func (m *APIRepo) GenerateUniqueID(ctx context.Context, modulename string) (int64, error) {
// 	url := fmt.Sprintf(m.APIConn + "/GetUniqueID/" + modulename)
// 	bytes, err := utils.GetRequest(url)
// 	if err != nil {
// 		return 0, err
// 	}
// 	data := make(map[string]int64)
// 	err = json.Unmarshal(bytes, &data)
// 	if err != nil {
// 		return 0, err
// 	} else {
// 		return data["id"], nil
// 	}

// }

//UpdateIsMsngStcksRslvdMain ..
func (m *APIRepo) UpdateIsMsngStcksRslvdMain(ctx context.Context, IDInWardOutWard int) error {
	url := fmt.Sprintf(m.APIConn+"/UpdateIsMsngStcksRslvdMain/%d", IDInWardOutWard)
	_, err := utils.PutRequest(url, nil)
	return err

}

//Employees_Bulk_Insert ..
func (m *APIRepo) Employees_Bulk_Insert(ctx context.Context, mdl []*CmnModel.Employees) error {
	url := fmt.Sprintf(m.APIConn + "/employee_BulkInsert")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PutRequest(url, j)
	return err
}

//Resend_Activation_Link ..
func (m *APIRepo) Resend_Activation_Link(ctx context.Context, EmpID int) error {
	url := fmt.Sprintf(m.APIConn+"/ResendActivationLink/%d", EmpID)
	_, err := utils.PutRequest(url, nil)
	return err
}

//GetAuthorizationList_ByRole ..
func (m *APIRepo) GetAuthorizationList_ByRole(ctx context.Context, RoleID int) ([]*CmnModel.Authorization, error) {
	url := fmt.Sprintf(m.APIConn+"/GetAuthorizationList_ByRole/%d", RoleID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Authorization, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}

//GetFeatures_List ..
func (m *APIRepo) GetFeatures_List(ctx context.Context) ([]*CmnModel.Features_List, error) {
	url := fmt.Sprintf(m.APIConn + "/GetFeatures_List")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Features_List, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}

}

//Send_ResetPasswordLink ..
func (m *APIRepo) Send_ResetPasswordLink(ctx context.Context, empid int) error {
	url := fmt.Sprintf(m.APIConn+"/Send_ResetPasswordLink/%d", empid)
	_, err := utils.PutRequest(url, nil)
	return err
}

//User_Inactive ..
func (m *APIRepo) User_Inactive(ctx context.Context, UserID int) error {
	url := fmt.Sprintf(m.APIConn+"/User_Inactive/%d", UserID)
	_, err := utils.PutRequest(url, nil)
	return err
}

//Authorization_Create ..
func (m *APIRepo) Authorization_Create(ctx context.Context, data CmnModel.Authorization_Create) error {
	url := fmt.Sprintf(m.APIConn + "/Authorization_Create")
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = utils.PostRequest(url, j)
	return err

}

func (m *APIRepo) User_Active(ctx context.Context, UserID int) error {
	url := fmt.Sprintf(m.APIConn+"/User_Active/%d", UserID)
	_, err := utils.PutRequest(url, nil)
	return err
}

func (m *APIRepo) Check_Unique_Email_Mobile(ctx context.Context, emp CmnModel.Employees) (*CmnModel.Employees, error) {
	url := fmt.Sprintf(m.APIConn + "/Check_Unique_Email_Mobile")
	j, err := json.Marshal(emp)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return nil, err
	}
	data := CmnModel.Employees{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return &data, nil
	}

}
func (m *APIRepo) Check_Unique_UserName(ctx context.Context, UserName string) (*string, error) {
	url := fmt.Sprintf(m.APIConn + "/Check_Unique_UserName/" + UserName)

	bytes, err := utils.PutRequest(url, nil)
	if err != nil {
		return nil, err
	}
	var Usrnm string
	err = json.Unmarshal(bytes, &Usrnm)
	if err != nil {
		return nil, err
	} else {
		return &Usrnm, nil
	}

}

//GetEmployee_History_ByEmpID ..
func (m *APIRepo) GetEmployee_History_ByEmpID(ctx context.Context, id int) ([]*CmnModel.Employees, error) {
	url := fmt.Sprintf(m.APIConn+"/GetEmployee_History_ByEmpID/%d", id)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Employees, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) Get_UsersHistory_ByEmpID(ctx context.Context, id int) ([]*CmnModel.Employees, error) {
	url := fmt.Sprintf(m.APIConn+"/Get_UsersHistory_ByEmpID/%d", id)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.Employees, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type Activity struct {
	FromDate string `json:"FromDate"`
	ToDate   string `json:"ToDate"`
}

func (m *APIRepo) Activivty_Log_List(ctx context.Context, id int, FromDate string, ToDate string) ([]*CmnModel.ActivityLog, error) {
	url := fmt.Sprintf(m.APIConn+"/Activivty_Log_List/%d", id)
	a := Activity{
		FromDate: FromDate,
		ToDate:   ToDate,
	}
	j, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return nil, err
	}

	data := make([]*CmnModel.ActivityLog, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) GetMultiLevelApproval(ctx context.Context) ([]*CmnModel.MultiLevelApproval_Main, error) {
	url := fmt.Sprintf(m.APIConn + "/GetMultiLevelApproval")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.MultiLevelApproval_Main, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) GetMultiLevelApproval_Map(ctx context.Context) ([]*CmnModel.MultiLevelApproval_Main, error) {
	url := fmt.Sprintf(m.APIConn + "/GetMultiLevelApproval_Map")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.MultiLevelApproval_Main, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) MultiLevelApproval_config(ctx context.Context, mdl []*CmnModel.MultiLevelApproval_Main) error {
	url := fmt.Sprintf(m.APIConn + "/MultiLevelApproval_config")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	_, err = utils.PutRequest(url, j)
	return err
}

func (m *APIRepo) GetApprovers(ctx context.Context, LocID int, RoleID int) ([]*CmnModel.User, error) {
	url := fmt.Sprintf(m.APIConn+"/GetApprovers/%d/%d", LocID, RoleID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.User, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) VendorsAssetDetails(ctx context.Context, VendorID int) ([]*CmnModel.VendorsAssetDetails, error) {
	url := fmt.Sprintf(m.APIConn+"/VendorsAssetDetails/%d", VendorID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.VendorsAssetDetails, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) VednorsAssetMapInsert(ctx context.Context, mdl CmnModel.Vendors_consumablemaster_map) error {
	url := fmt.Sprintf(m.APIConn + "/VednorsAssetMapInsert")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	data := CmnModel.Employees{}
	err = json.Unmarshal(bytes, &data)

	return err

}

func (m *APIRepo) VednorsAssetMapUpdate(ctx context.Context, mdl CmnModel.Vendors_consumablemaster_map) error {
	url := fmt.Sprintf(m.APIConn + "/VednorsAssetMapUpdate")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	data := CmnModel.Employees{}
	err = json.Unmarshal(bytes, &data)

	return err

}

func (m *APIRepo) IWOW_ApprovalStatusList(ctx context.Context, IDinwardoutward int) ([]*CmnModel.InWardOutWardApproval, error) {
	url := fmt.Sprintf(m.APIConn+"/IWOW_ApprovalStatusList/%d", IDinwardoutward)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := make([]*CmnModel.InWardOutWardApproval, 0)
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *APIRepo) InwardOutwardReqForward(ctx context.Context, mdl CmnModel.InWardOutWardApproval) error {
	url := fmt.Sprintf(m.APIConn + "/InwardOutwardReqForward")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) GetAdminDashBoard(ctx context.Context, mdl CmnModel.AdminDashBoard) (*CmnModel.AdminDashBoard, error) {
	url := fmt.Sprintf(m.APIConn + "/GetAdminDashBoard")
	j, err := json.Marshal(mdl)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return nil, err
	}
	data := CmnModel.AdminDashBoard{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return &data, nil
	}

}

func (m *APIRepo) GetEmployeeDashboard(ctx context.Context, mdl CmnModel.EmployeeDashboard) (*CmnModel.EmployeeDashboard, error) {
	url := fmt.Sprintf(m.APIConn + "/GetEmployeeDashboard")
	j, err := json.Marshal(mdl)
	if err != nil {
		return nil, err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return nil, err
	}
	data := CmnModel.EmployeeDashboard{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	} else {
		return &data, nil
	}

}

func (m *APIRepo) PurchaseOrders_RequestsInsert(ctx context.Context, mdl CmnModel.PurchaseOrders_Requests) error {
	url := fmt.Sprintf(m.APIConn + "/PurchaseOrders_RequestsInsert")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) GetPurchaseOrderUniqueID() (*string, error) {
	url := fmt.Sprintf(m.APIConn + "/GetPurchaseOrderUniqueID")
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	var NextID string
	err = json.Unmarshal(bytes, &NextID)
	if err != nil {
		return nil, err
	}
	return &NextID, nil

}

func (m *APIRepo) GetPODetailsByReqstrID(ctx context.Context, ReqstrID int) ([]*CmnModel.PurchaseOrders_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/GetPODetailsByReqstrID/%d", ReqstrID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.PurchaseOrders_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
func (m *APIRepo) PODetailsByIDPO(ctx context.Context, IDPO int) (*CmnModel.PurchaseOrders_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/PODetailsByIDPO/%d", IDPO)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := CmnModel.PurchaseOrders_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (m *APIRepo) POAssetDetailsByIDPO(ctx context.Context, IDPO int) ([]*CmnModel.PurchaseOrders_Assets, error) {
	url := fmt.Sprintf(m.APIConn+"/POAssetDetailsByIDPO/%d", IDPO)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.PurchaseOrders_Assets{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *APIRepo) PO_ApprovalStatusList(ctx context.Context, IDPO int) ([]*CmnModel.POApproval, error) {
	url := fmt.Sprintf(m.APIConn+"/PO_ApprovalStatusList/%d", IDPO)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.POApproval{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *APIRepo) GetPODetailsByApprover(ctx context.Context, ApprvrID int) ([]*CmnModel.PurchaseOrders_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/GetPODetailsByApprover/%d", ApprvrID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.PurchaseOrders_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}
func (m *APIRepo) POReqApproved(ctx context.Context, mdl CmnModel.POApproval) error {
	url := fmt.Sprintf(m.APIConn + "/POReqApproved")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}
func (m *APIRepo) POReqForward(ctx context.Context, mdl CmnModel.POApproval) error {
	url := fmt.Sprintf(m.APIConn + "/POReqForward")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}
func (m *APIRepo) POReqRejected(ctx context.Context, mdl CmnModel.POApproval) error {
	url := fmt.Sprintf(m.APIConn + "/POReqRejected")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) PurchaseOrders_RequestsUpdate(ctx context.Context, mdl CmnModel.PurchaseOrders_Requests) error {
	url := fmt.Sprintf(m.APIConn + "/PurchaseOrders_RequestsUpdate")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) POStatusChange(IDPO int, Status int) error {
	url := fmt.Sprintf(m.APIConn+"/POStatusChange/%d/%d", IDPO, Status)
	_, err := utils.GetRequest(url)
	return err
}

func (m *APIRepo) Requisition_RequestsInsert(ctx context.Context, mdl CmnModel.Requisition_Requests) error {
	url := fmt.Sprintf(m.APIConn + "/Requisition_RequestsInsert")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) GetRequisitionDetailsByReqstrID(ctx context.Context, ReqstrID int) ([]*CmnModel.Requisition_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/GetRequisitionDetailsByReqstrID/%d", ReqstrID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.Requisition_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *APIRepo) RequisitionDetailsByID(ctx context.Context, ID int) (*CmnModel.Requisition_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/RequisitionDetailsByID/%d", ID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := CmnModel.Requisition_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (m *APIRepo) RequisitionAssetDetailsByID(ctx context.Context, ID int) ([]*CmnModel.Requisition_Assets, error) {
	url := fmt.Sprintf(m.APIConn+"/RequisitionAssetDetailsByID/%d", ID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.Requisition_Assets{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *APIRepo) Requisition_ApprovalStatusList(ctx context.Context, ID int) ([]*CmnModel.RequisitionApproval, error) {
	url := fmt.Sprintf(m.APIConn+"/Requisition_ApprovalStatusList/%d", ID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.RequisitionApproval{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *APIRepo) GetRequisitionDetailsByApprover(ctx context.Context, ApprvrID int) ([]*CmnModel.Requisition_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/GetRequisitionDetailsByApprover/%d", ApprvrID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.Requisition_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (m *APIRepo) Requisition_RequestsUpdate(ctx context.Context, mdl CmnModel.Requisition_Requests) error {
	url := fmt.Sprintf(m.APIConn + "/Requisition_RequestsUpdate")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}
func (m *APIRepo) RequisitionReqRejected(ctx context.Context, mdl CmnModel.POApproval) error {
	url := fmt.Sprintf(m.APIConn + "/RequisitionReqRejected")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) RequisitionReqApproved(ctx context.Context, mdl CmnModel.RequisitionApproval) error {
	url := fmt.Sprintf(m.APIConn + "/RequisitionReqApproved")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}
func (m *APIRepo) RequisitionReqForward(ctx context.Context, mdl CmnModel.RequisitionApproval) error {
	url := fmt.Sprintf(m.APIConn + "/RequisitionReqForward")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}

func (m *APIRepo) RequisitionStatusChange(ID int, Status int) error {
	url := fmt.Sprintf(m.APIConn+"/RequisitionStatusChange/%d/%d", ID, Status)
	_, err := utils.GetRequest(url)
	return err
}

func (m *APIRepo) RequisitionStcokReceived(ctx context.Context, mdl CmnModel.Requisition_Requests) error {
	url := fmt.Sprintf(m.APIConn + "/RequisitionStcokReceived")
	j, err := json.Marshal(mdl)
	if err != nil {
		return err
	}
	bytes, err := utils.PutRequest(url, j)
	if err != nil {
		return err
	}
	var er string
	err = json.Unmarshal(bytes, &er)

	return err

}


func (m *APIRepo) GetRequisitionHistoryByReqID(ctx context.Context, ID int) ([]*CmnModel.Requisition_Requests, error) {
	url := fmt.Sprintf(m.APIConn+"/GetRequisitionHistoryByReqID/%d", ID)
	bytes, err := utils.GetRequest(url)
	if err != nil {
		return nil, err
	}
	data := []*CmnModel.Requisition_Requests{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil

}