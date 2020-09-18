package cmnhandler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/repository/CmnRep"
	"github.com/premsynfosys/AMS_WEB/utils"
)

// Notificationss ..
type Notificationss struct {
	Message string `json:"Message"`
	Role    string `json:"Role"`
	UserID  int    `json:"UserID"`
}

//ICommonrep ..
type ICommonrep struct {
	Irepo CmnRep.CmnRepIntrfc
}

//NewCommonHandler ..
func NewCommonHandler(api string) *ICommonrep {
	return &ICommonrep{
		Irepo: CmnRep.NewAPIRepo(api),
	}
}

//Emp ..
func (p *ICommonrep) Emp(w http.ResponseWriter, r *http.Request) {
	usr, _ := utils.GetCookieUser(r)
	status, _ := p.Irepo.GetEmployees(r.Context(), usr.LocationID)
	utils.ExecuteTemplate(w, r, "ITAssetTmpl", status)

}

//Dashboard sfc
func (p *ICommonrep) Dashboard(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	in := CmnModel.AdminDashBoard{}
	in.EmpID = usr.EmployeeID
	if usr.RoleName == "Admin" {
		in.LocID = usr.LocationID
	}

	data, _ := p.Irepo.GetAdminDashBoard(r.Context(), in)

	Mapdata := CmnModel.TemplateData{
		Data: data,
		User: usr,
		Auth: Auth,
	}

	utils.ExecuteTemplate(w, r, "dashboard", Mapdata)
}

func (p *ICommonrep) MyDashBoard(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	in := CmnModel.EmployeeDashboard{}
	in.EmpID = usr.EmployeeID
	in.LocID = usr.LocationID

	data, _ := p.Irepo.GetEmployeeDashboard(r.Context(), in)

	Mapdata := CmnModel.TemplateData{
		Data: data,
		User: usr,
		Auth: Auth,
	}

	utils.ExecuteTemplate(w, r, "MyDashBoard", Mapdata)
}

// Transfer ..
func (p *ICommonrep) Transfer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		usr, Auth := utils.GetCookieUser(r)
		mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
		var RoleID int
		var Grade int
		for i := 0; i < len(mapRoles); i++ {
			if mapRoles[i].FeatureName == "OutWard Request" {
				RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
				Grade = mapRoles[i].MultiLevelApproval_Map.Grade
				break
			}
		}
		data := struct {
			RoleID int
			Grade  int
		}{
			RoleID: RoleID,
			Grade:  Grade,
		}
		Mapdata := CmnModel.TemplateData{
			Data: data,
			User: usr,
			Auth: Auth,
		}
		utils.ExecuteTemplate(w, r, "Transfer", Mapdata)
	}
}

// GetVendors ..
func (p *ICommonrep) GetVendors(w http.ResponseWriter, r *http.Request) {
	data, _ := p.Irepo.GetVendors(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}
func (p *ICommonrep) VendorsDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{

		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "VendorsDetails", Mapdata)
	}
}

// GetCountries ..
func (p *ICommonrep) GetCountries(w http.ResponseWriter, r *http.Request) {
	data, _ := p.Irepo.GetCountries(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}

// GetLocations ..
func (p *ICommonrep) GetLocations(w http.ResponseWriter, r *http.Request) {
	data, _ := p.Irepo.GetLocations(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}

// Location ..
func (p *ICommonrep) Location(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		usr, Auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{
			Data: CmnModel.Locations{},
			User: usr,
			Auth: Auth,
		}
		if r.URL.Query().Get("mode") == "edit" {
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				return
			}
			data, _ := p.Irepo.GetLocations(r.Context())
			for _, item := range data {
				if item.IDLocations == id {
					Mapdata.Data = item
					break
				}
			}
			utils.ExecuteTemplate(w, r, "locationsTmpl", Mapdata)
		} else {
			utils.ExecuteTemplate(w, r, "locationsTmpl", Mapdata)
		}
	} else if r.Method == "POST" {
		mdl := CmnModel.Locations{}
		mdl.Name = r.FormValue("Name")
		mdl.Address1 = r.FormValue("Address1")
		mdl.Address2 = r.FormValue("Address2")
		Countryid, _ := strconv.Atoi(r.FormValue("Country"))
		mdl.Country = Countryid
		stateid, _ := strconv.Atoi(r.FormValue("State"))
		mdl.State = stateid
		mdl.City = r.FormValue("City")
		mdl.Zipcode = r.FormValue("Zipcode")
		mdl.Description = r.FormValue("Description")

		if r.FormValue("IDLocations") != "0" {
			mdl.IDLocations, _ = strconv.Atoi(r.FormValue("IDLocations"))
			err := p.Irepo.UpdateLocations(r.Context(), &mdl)
			if err != nil {
				utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
			} else {
				utils.RespondwithJSON(w, r, http.StatusOK, nil)
			}

		} else {
			data, _ := p.Irepo.CreateLocation(r.Context(), &mdl)
			utils.RespondwithJSON(w, r, http.StatusOK, data)
		}
	}
}
func (p *ICommonrep) LocationsDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{

		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "LocationsDetails", Mapdata)
	}
}

// Vendors ..
func (p *ICommonrep) Vendors(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Query().Get("mode") == "edit" {
			var Loc *CmnModel.Vendors
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				return
			}
			data, _ := p.Irepo.GetVendors(r.Context())
			for _, item := range data {
				if item.Idvendors == id {
					Loc = item
					break
				}
			}
			utils.ExecuteTemplate(w, r, "vendorTmpl", Loc)
		} else {
			utils.ExecuteTemplate(w, r, "vendorTmpl", nil)
		}
	} else if r.Method == "POST" {
		mdl := CmnModel.Vendors{}
		mdl.Name = r.FormValue("Name")
		mdl.Description = r.FormValue("Description")
		mdl.Websites = r.FormValue("Websites")
		mdl.Address = r.FormValue("Address")
		mdl.Email = r.FormValue("Email")
		mdl.ContactPersonName = r.FormValue("ContactPersonName")
		mdl.Phone = r.FormValue("Phone")
		if r.FormValue("Idvendors") != "" {
			mdl.Idvendors, _ = strconv.Atoi(r.FormValue("Idvendors"))
			err := p.Irepo.UpdateVendors(r.Context(), &mdl)
			if err != nil {
				utils.RespondwithJSON(w, r, http.StatusInternalServerError, nil)
			} else {
				utils.RespondwithJSON(w, r, http.StatusOK, nil)
			}

		} else {
			data, _ := p.Irepo.CreateVendors(r.Context(), &mdl)
			utils.RespondwithJSON(w, r, http.StatusOK, data)
		}
	}
}

// GetStatesByCountry ..
func (p *ICommonrep) GetStatesByCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	data, _ := p.Irepo.GetStatesByCountry(r.Context(), id)
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}

// GetNotifications ..
func (p *ICommonrep) GetNotifications(w http.ResponseWriter, r *http.Request) {
	data, _ := p.Irepo.GetNotifications(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, data)
}

//GetOutWardCart ..
func (p *ICommonrep) GetOutWardCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["SenderEmpid"]
	SenderEmpid, _ := strconv.Atoi(ID)
	body, _ := p.Irepo.GetOutWardCart(r.Context(), SenderEmpid)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//DeleteOutWardCart ..
func (p *ICommonrep) DeleteOutWardCart(w http.ResponseWriter, r *http.Request) {
	ListIWA := []*CmnModel.InWardOutWardAsset{}
	var empid int
	lst := []byte(r.FormValue("ListInWardOutWardAsset"))
	id := []byte(r.FormValue("EmpID"))
	json.Unmarshal(lst, &ListIWA)
	json.Unmarshal(id, &empid)
	data := struct {
		ListInWardOutWardAsset []*CmnModel.InWardOutWardAsset
		EmpID                  int
	}{
		ListInWardOutWardAsset: ListIWA,
		EmpID:                  empid,
	}
	err := p.Irepo.DeleteOutWardCart(r.Context(), data)
	utils.RespondwithJSON(w, r, http.StatusOK, err)
}

// OutWardDetails ..
func (p *ICommonrep) OutWardDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	requstmap := []*CmnModel.MultiLevelApproval_Main{}
	for _, itm := range mapRoles {
		if itm.FeatureName == "OutWard Request" {
			requstmap = append(requstmap, itm)
		}
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: requstmap,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "OutWardDetails", Mapdata)
	}
}

// InWardDetails ..
func (p *ICommonrep) InWardDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "InWardDetails", Mapdata)
	}
}

//GetOutWardDetailsByEmp ..
func (p *ICommonrep) GetOutWardDetailsByEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["SenderEmpid"]
	SenderEmpid, _ := strconv.Atoi(ID)
	body, _ := p.Irepo.GetOutWardDetailsByEmp(r.Context(), SenderEmpid)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//GetInWardDetailsByEmp ..
func (p *ICommonrep) GetInWardDetailsByEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["RcvrEmpID"]
	SenderEmpid, _ := strconv.Atoi(ID)
	body, _ := p.Irepo.GetInWardDetailsByEmp(r.Context(), SenderEmpid)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//GetOutWardAssetDetailsByIwOwID ..
func (p *ICommonrep) GetOutWardAssetDetailsByIwOwID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["IwOwID"]
	IwOwID, _ := strconv.Atoi(ID)
	body, err := p.Irepo.GetOutWardAssetDetailsByIwOwID(r.Context(), IwOwID)

	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, body)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//GetAssetdIDsNotEligbleForTransfer ..
func (p *ICommonrep) GetAssetdIDsNotEligbleForTransfer(w http.ResponseWriter, r *http.Request) {
	res := []*CmnModel.Transfer{}
	json.NewDecoder(r.Body).Decode(&res)
	body, err := p.Irepo.GetAssetdIDsNotEligbleForTransfer(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, body)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) DeleteVendors(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.Vendors{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.DeleteVendors(r.Context(), &res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//IWOWDetailsByAprvr ..
func (p *ICommonrep) IWOWDetailsByAprvr(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	AprvrEmp := params["AprvrEmpID"]
	AprvrEmpID, _ := strconv.Atoi(AprvrEmp)
	body, _ := p.Irepo.IWOWDetailsByAprvr(r.Context(), AprvrEmpID)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

// OWApprovalDeatils ..
func (p *ICommonrep) OWApprovalDeatils(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "OWApprovalDeatils", Mapdata)
	}
}

// OWApproval ..
func (p *ICommonrep) OWApproval(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	params := mux.Vars(r)
	AprvrEmp := params["AprvrEmpID"]
	AprvrEmpID, _ := strconv.Atoi(AprvrEmp)
	IDInWardOutWard := params["IDInWardOutWard"]
	idInWardOutWard, _ := strconv.Atoi(IDInWardOutWard)
	OWDetailsByAprvr, _ := p.Irepo.IWOWDetailsByAprvr(r.Context(), AprvrEmpID)
	var res *CmnModel.InWardOutWard
	if len(OWDetailsByAprvr) > 0 {
		for _, itm := range OWDetailsByAprvr {
			if itm.IDInWardOutWard == idInWardOutWard {
				res = itm
				data, _ := p.Irepo.GetOutWardAssetDetailsByIwOwID(r.Context(), itm.IDInWardOutWard)
				res.ListInWardOutWardAsset = data
				break
			}
		}
	}
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	requstmap := []*CmnModel.MultiLevelApproval_Main{}
	for _, itm := range mapRoles {
		if itm.FeatureName == "OutWard Request" {
			requstmap = append(requstmap, itm)
		}
	}
	data := struct {
		IwOw       *CmnModel.InWardOutWard
		ListApprvl []*CmnModel.MultiLevelApproval_Main
	}{
		IwOw:       res,
		ListApprvl: requstmap,
	}
	Mapdata := CmnModel.TemplateData{
		Data: data,
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "OWApproval", Mapdata)
	}
}

// IWReceive ..
func (p *ICommonrep) IWReceive(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	params := mux.Vars(r)
	IDInWardOutWard := params["IDInWardOutWard"]
	idInWardOutWard, _ := strconv.Atoi(IDInWardOutWard)
	ID := params["RcvrEmpID"]
	RcvrEmpID, _ := strconv.Atoi(ID)
	InWardDetails, _ := p.Irepo.GetInWardDetailsByEmp(r.Context(), RcvrEmpID)

	var res *CmnModel.InWardOutWard
	if len(InWardDetails) > 0 {
		for _, itm := range InWardDetails {
			if itm.IDInWardOutWard == idInWardOutWard {
				res = itm
				data, _ := p.Irepo.GetOutWardAssetDetailsByIwOwID(r.Context(), itm.IDInWardOutWard)
				res.ListInWardOutWardAsset = data
				break
			}
		}
	}
	Mapdata := CmnModel.TemplateData{
		Data: res,
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "IWReceive", Mapdata)
	}
}

//TransferApprovReject ..
func (p *ICommonrep) TransferApprovReject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idInWardOutWard := params["idInWardOutWard"]
	status := params["status"]
	comments := r.URL.Query().Get("comments")
	IDInWardOutWard, _ := strconv.Atoi(idInWardOutWard)
	err := p.Irepo.TransferApprovReject(r.Context(), IDInWardOutWard, status, comments)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//ReceiveIWAssets ..
func (p *ICommonrep) ReceiveIWAssets(w http.ResponseWriter, r *http.Request) {
	InWardOutWard := CmnModel.InWardOutWard{}
	err := json.NewDecoder(r.Body).Decode(&InWardOutWard)
	err = p.Irepo.ReceiveIWAssets(r.Context(), &InWardOutWard)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//OwStatusChange ..
func (p *ICommonrep) OwStatusChange(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	OWid := params["OWid"]
	Status := params["Status"]
	owid, _ := strconv.Atoi(OWid)
	status, _ := strconv.Atoi(Status)
	err := p.Irepo.OwStatusChange(owid, status)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//CreateOutWardCart ..
func (p *ICommonrep) CreateOutWardCart(w http.ResponseWriter, r *http.Request) {
	mdl := []*CmnModel.OutWardCart{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.CreateOutWardCart(r.Context(), mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

// //GenerateUniqueID ..
// func (p *ICommonrep) GenerateUniqueID(w http.ResponseWriter, r *http.Request) {
// 	var UniqueID string
// 	params := mux.Vars(r)
// 	modulename := params["modulename"]
// 	maxid, err := p.Irepo.GenerateUniqueID(r.Context(), modulename)
// 	if err != nil {
// 		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
// 		return
// 	}
// 	maxid++
// 	if modulename == "itassets" {
// 		UniqueID = fmt.Sprintf("SYS/ITASSETS/%03d", maxid)
// 	} else if modulename == "consumable" {
// 		UniqueID = fmt.Sprintf("SYS/CNSMBL/%03d", maxid)
// 	} else if modulename == "nonitassets" {
// 		UniqueID = fmt.Sprintf("SYS/NONITASSETS/%03d", maxid)
// 	}

// 	utils.RespondwithJSON(w, r, http.StatusOK, UniqueID)
// }

//CreateInWardOutWard ..
func (p *ICommonrep) CreateInWardOutWard(w http.ResponseWriter, r *http.Request) {
	IW := CmnModel.InWardOutWard{}
	IW.ReceiverEmpID, _ = strconv.Atoi(r.FormValue("Receiver"))
	IW.SenderEmpID, _ = strconv.Atoi(r.FormValue("EmpID"))
	IW.FromLocationID, _ = strconv.Atoi(r.FormValue("EmpLocation"))
	IW.ToLocationID, _ = strconv.Atoi(r.FormValue("Locations"))
	IW.Description = r.FormValue("Description")
	IW.TransferStatus = 9
	transactionid := "TR" + time.Now().Format("02-01-2006 15:04:05")
	rpl := strings.NewReplacer(":", "",
		"-", "",
		" ", "")
	res := rpl.Replace(transactionid)
	IW.TransactionID = res
	//
	ListIWA := []*CmnModel.InWardOutWardAsset{}
	lst := []byte(`` + r.FormValue("ListInWardOutWardAsset") + ``)
	json.Unmarshal(lst, &ListIWA)

	aprvl := CmnModel.InWardOutWardApproval{}
	aprvl.ApproverID, _ = strconv.Atoi(r.FormValue("Approver"))
	aprvl.Grade, _ = strconv.Atoi(r.FormValue("Grade"))
	aprvl.RoleID, _ = strconv.Atoi(r.FormValue("ApprvrRoleID"))

	aprvl.Status = 9

	IW.ListInWardOutWardAsset = ListIWA
	IW.InWardOutWardApproval = aprvl
	err := p.Irepo.CreateInWardOutWard(r.Context(), &IW)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}

}

//UpdateIsMsngStcksRslvdMain ..
func (p *ICommonrep) UpdateIsMsngStcksRslvdMain(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["IDInWardOutWard"]
	ID, _ := strconv.Atoi(id)
	err := p.Irepo.UpdateIsMsngStcksRslvdMain(r.Context(), ID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) MultiApproval(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	data, _ := p.Irepo.GetMultiLevelApproval(r.Context())

	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "ConfigMultiApprvl", Mapdata)
	}
}
func (p *ICommonrep) GetMultiLevelApproval_Map(w http.ResponseWriter, r *http.Request) {

	data, err := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) MultiLevelApproval_config(w http.ResponseWriter, r *http.Request) {
	mdl := []*CmnModel.MultiLevelApproval_Main{}
	json.NewDecoder(r.Body).Decode(&mdl)
	err := p.Irepo.MultiLevelApproval_config(r.Context(), mdl)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) VendorAssetDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	data, _ := p.Irepo.GetMultiLevelApproval(r.Context())

	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "VendorAssetDetails", Mapdata)
	}
}

func (p *ICommonrep) GetVendorsAssetDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	VendorID := params["vendorID"]
	vendorID, _ := strconv.Atoi(VendorID)
	data, err := p.Irepo.VendorsAssetDetails(r.Context(), vendorID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) VednorsAssetMapInsert(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.Vendors_consumablemaster_map{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.VednorsAssetMapInsert(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) VednorsAssetMapUpdate(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.Vendors_consumablemaster_map{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.VednorsAssetMapUpdate(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) IWOW_ApprovalStatusList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDinwardoutward := params["IDinwardoutward"]
	IDinwardoutwards, _ := strconv.Atoi(IDinwardoutward)
	data, err := p.Irepo.IWOW_ApprovalStatusList(r.Context(), IDinwardoutwards)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) InwardOutwardReqForward(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.InWardOutWardApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.InwardOutwardReqForward(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) PurchaseOrder(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	var RoleID int
	var Grade int
	for i := 0; i < len(mapRoles); i++ {
		if mapRoles[i].FeatureName == "Purchase Order" {
			RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
			Grade = mapRoles[i].MultiLevelApproval_Map.Grade
			break
		}
	}
	data := struct {
		RoleID int
		Grade  int
	}{
		RoleID: RoleID,
		Grade:  Grade,
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "PurchaseOrder", Mapdata)
	}
}
func (p *ICommonrep) PurchaseOrderView(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	params := mux.Vars(r)
	IDPO := params["IDPO"]
	IDPOs, _ := strconv.Atoi(IDPO)
	POR, _ := p.Irepo.PODetailsByIDPO(r.Context(), IDPOs)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	requstmap := []*CmnModel.MultiLevelApproval_Main{}
	for _, itm := range mapRoles {
		if itm.FeatureName == "Purchase Order" {
			requstmap = append(requstmap, itm)
		}
	}
	data := struct {
		POR        *CmnModel.PurchaseOrders_Requests
		ListApprvl []*CmnModel.MultiLevelApproval_Main
	}{
		POR:        POR,
		ListApprvl: requstmap,
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "PurchaseOrderView", Mapdata)
	}
}
func (p *ICommonrep) PurchaseOrderDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)

	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "PurchaseOrderDetails", Mapdata)
	}
}

func (p *ICommonrep) Error(w http.ResponseWriter, r *http.Request) {

	utils.ExecuteTemplate(w, r, "Error", nil)

}

func (p *ICommonrep) PurchaseOrders_RequestsInsert(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.PurchaseOrders_Requests{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.PurchaseOrders_RequestsInsert(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) GetPurchaseOrderUniqueID(w http.ResponseWriter, r *http.Request) {
	NextID, err := p.Irepo.GetPurchaseOrderUniqueID()
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, NextID)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) GetPODetailsByReqstrID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ReqstrID := params["ReqstrID"]
	ReqstrIDs, _ := strconv.Atoi(ReqstrID)
	data, err := p.Irepo.GetPODetailsByReqstrID(r.Context(), ReqstrIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}
func (p *ICommonrep) POAssetDetailsByIDPO(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	POID := params["IDPO"]
	POIDs, _ := strconv.Atoi(POID)
	data, err := p.Irepo.POAssetDetailsByIDPO(r.Context(), POIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}
func (p *ICommonrep) PO_ApprovalStatusList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	POID := params["IDPO"]
	POIDs, _ := strconv.Atoi(POID)
	data, err := p.Irepo.PO_ApprovalStatusList(r.Context(), POIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) POApprovalDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "POApprovalDetails", Mapdata)
	}
}

func (p *ICommonrep) GetPODetailsByApprover(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ApprvrID := params["ApprvrID"]
	ApprvrIDs, _ := strconv.Atoi(ApprvrID)
	data, err := p.Irepo.GetPODetailsByApprover(r.Context(), ApprvrIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}
func (p *ICommonrep) POReqApproved(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.POApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.POReqApproved(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) POReqForward(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.POApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.POReqForward(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) POReqRejected(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.POApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.POReqRejected(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) PO_Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDPO := params["IDPO"]
	IDPOs, _ := strconv.Atoi(IDPO)
	usr, Auth := utils.GetCookieUser(r)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	POR, _ := p.Irepo.PODetailsByIDPO(r.Context(), IDPOs)
	var RoleID int
	var Grade int
	for i := 0; i < len(mapRoles); i++ {
		if mapRoles[i].FeatureName == "Purchase Order" {
			RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
			Grade = mapRoles[i].MultiLevelApproval_Map.Grade
			break
		}
	}
	data := struct {
		RoleID int
		Grade  int
		POR    *CmnModel.PurchaseOrders_Requests
	}{
		RoleID: RoleID,
		Grade:  Grade,
		POR:    POR,
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "PO_Edit", Mapdata)
	}
}

func (p *ICommonrep) PurchaseOrders_RequestsUpdate(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.PurchaseOrders_Requests{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.PurchaseOrders_RequestsUpdate(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) POStatusChange(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDPO := params["IDPO"]
	Status := params["Status"]
	IDPOs, _ := strconv.Atoi(IDPO)
	status, _ := strconv.Atoi(Status)
	err := p.Irepo.POStatusChange(IDPOs, status)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) Requisition(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	var RoleID int
	var Grade int
	for i := 0; i < len(mapRoles); i++ {
		if mapRoles[i].FeatureName == "Requisition" {
			RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
			Grade = mapRoles[i].MultiLevelApproval_Map.Grade
			break
		}
	}
	data := struct {
		RoleID int
		Grade  int
	}{
		RoleID: RoleID,
		Grade:  Grade,
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "Requisition", Mapdata)
	}
}

func (p *ICommonrep) Requisition_RequestsInsert(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.Requisition_Requests{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.Requisition_RequestsInsert(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) RequisitionDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)

	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "RequisitionDetails", Mapdata)
	}
}

func (p *ICommonrep) GetRequisitionDetailsByReqstrID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ReqstrID := params["ReqstrID"]
	ReqstrIDs, _ := strconv.Atoi(ReqstrID)
	data, err := p.Irepo.GetRequisitionDetailsByReqstrID(r.Context(), ReqstrIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) RequisitionView(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	params := mux.Vars(r)
	ID := params["ID"]
	IDs, _ := strconv.Atoi(ID)
	POR, _ := p.Irepo.RequisitionDetailsByID(r.Context(), IDs)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	requstmap := []*CmnModel.MultiLevelApproval_Main{}
	for _, itm := range mapRoles {
		if itm.FeatureName == "Requisition" {
			requstmap = append(requstmap, itm)
		}
	}
	data := struct {
		POR        *CmnModel.Requisition_Requests
		ListApprvl []*CmnModel.MultiLevelApproval_Main
	}{
		POR:        POR,
		ListApprvl: requstmap,
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "RequisitionView", Mapdata)
	}
}

func (p *ICommonrep) RequisitionAssetDetailsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ID"]
	IDs, _ := strconv.Atoi(ID)
	data, err := p.Irepo.RequisitionAssetDetailsByID(r.Context(), IDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) GetRequisitionHistoryByReqID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ReqID"]
	IDs, _ := strconv.Atoi(ID)
	data, err := p.Irepo.GetRequisitionHistoryByReqID(r.Context(), IDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) Requisition_ApprovalStatusList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	POID := params["ID"]
	POIDs, _ := strconv.Atoi(POID)
	data, err := p.Irepo.Requisition_ApprovalStatusList(r.Context(), POIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) RequisitionApprovalDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "RequisitionApprovalDetails", Mapdata)
	}
}
func (p *ICommonrep) GetRequisitionDetailsByApprover(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ApprvrID := params["ApprvrID"]
	ApprvrIDs, _ := strconv.Atoi(ApprvrID)
	data, err := p.Irepo.GetRequisitionDetailsByApprover(r.Context(), ApprvrIDs)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}

func (p *ICommonrep) Requisition_Edit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["ID"]
	IDPOs, _ := strconv.Atoi(ID)
	usr, Auth := utils.GetCookieUser(r)
	mapRoles, _ := p.Irepo.GetMultiLevelApproval_Map(r.Context())
	POR, _ := p.Irepo.RequisitionDetailsByID(r.Context(), IDPOs)
	var RoleID int
	var Grade int
	for i := 0; i < len(mapRoles); i++ {
		if mapRoles[i].FeatureName == "Requisition" {
			RoleID = mapRoles[i].MultiLevelApproval_Map.RoleID
			Grade = mapRoles[i].MultiLevelApproval_Map.Grade
			break
		}
	}
	data := struct {
		RoleID int
		Grade  int
		POR    *CmnModel.Requisition_Requests
	}{
		RoleID: RoleID,
		Grade:  Grade,
		POR:    POR,
	}
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
		Data: data,
	}
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "Requisition_Edit", Mapdata)
	}
}

func (p *ICommonrep) Requisition_RequestsUpdate(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.Requisition_Requests{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.Requisition_RequestsUpdate(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) RequisitionReqRejected(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.POApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.RequisitionReqRejected(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) RequisitionReqApproved(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.RequisitionApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.RequisitionReqApproved(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) RequisitionReqForward(w http.ResponseWriter, r *http.Request) {
	res := CmnModel.RequisitionApproval{}
	json.NewDecoder(r.Body).Decode(&res)
	err := p.Irepo.RequisitionReqForward(r.Context(), res)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) RequisitionStatusChange(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	IDPO := params["ID"]
	Status := params["Status"]
	IDPOs, _ := strconv.Atoi(IDPO)
	status, _ := strconv.Atoi(Status)
	err := p.Irepo.RequisitionStatusChange(IDPOs, status)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) RequisitionReceivedStock(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mdl := CmnModel.Requisition_Requests{}
		mdl.IDRequisition_Requests, _ = strconv.Atoi(r.FormValue("IDRequisition_Requests"))
		mdl.ModifiedBy, _ = strconv.Atoi(r.FormValue("CreatedBy"))
		mdl.BillInvoiceNo = r.FormValue("BillInvoiceNo")
		TotalPaidAmmount, _ := strconv.ParseFloat(r.FormValue("TotalPaidAmmount"), 32)
		mdl.TotalPaidAmmount = TotalPaidAmmount
		mdl.LocationID, _ = strconv.Atoi(r.FormValue("LocationID"))
		mdl.VendorID, _ = strconv.Atoi(r.FormValue("VendorID"))
		img, handle, err := r.FormFile("BillImagePath")
		if err == nil {
			file, _ := os.Create("AppFiles/Images/RequisitionBIlls/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			mdl.BillImagePath = handle.Filename
			defer img.Close()
		}

		ListStocks := []CmnModel.Requisition_Assets{}
		customfieldsRaw := []byte(r.FormValue("ReceivedStockList"))
		json.Unmarshal(customfieldsRaw, &ListStocks)
		mdl.ListRequisition_Assets = ListStocks

		err = p.Irepo.RequisitionStcokReceived(r.Context(), mdl)
		if err == nil {
			utils.RespondwithJSON(w, r, http.StatusOK, nil)
		} else {
			utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
		}

	}
}

func (p *ICommonrep) RequisitionHistoryDetails_Partial(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	ID := params["ReqID"]
	IDs, _ := strconv.Atoi(ID)
	utils.ExecuteTemplate(w, r, "RequisitionHistoryDetails_Partial", struct {
		ReqID int
	}{
		ReqID: IDs,
	})
}

func (p *ICommonrep) GetSearchDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	Name := params["Name"]
	locid, _ := strconv.Atoi(LocID)
	data, err := p.Irepo.GetSearchDetails(r.Context(), locid, Name)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, err.Error())
	}
}
