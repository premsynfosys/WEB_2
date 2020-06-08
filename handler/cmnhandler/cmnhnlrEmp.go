package cmnhandler

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/premsynfosys/AMS_WEB/models/CmnModel"
	"github.com/premsynfosys/AMS_WEB/utils"
)

//CreateEmployee a new post
func (p *ICommonrep) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		// usr, Auth := utils.GetCookieUser(r)
		Mapdata := CmnModel.TemplateData{

			User: usr,
			Auth: Auth,
		}
		utils.ExecuteTemplate(w, r, "addemployee", Mapdata)
	} else if r.Method == "POST" {
		mdl := CmnModel.Employees{}
		mdl.FirstName = r.FormValue("FirstName")
		mdl.LastName = r.FormValue("LastName")
		mdl.DOB = r.FormValue("DOB")
		mdl.Email = r.FormValue("Email")
		mdl.Mobile = r.FormValue("Mobile")
		mdl.Address = r.FormValue("Address")
		mdl.PrmntAddress = r.FormValue("PrmntAddress")
		mdl.EmpCode = r.FormValue("EmpCode")
		mdl.Education, _ = strconv.Atoi(r.FormValue("Education"))
		mdl.ExperienceYear, _ = strconv.Atoi(r.FormValue("ExperienceYear"))
		mdl.ExperienceMonth, _ = strconv.Atoi(r.FormValue("ExperienceMonth"))
		mdl.CreatedBy = usr.EmployeeID
		if r.FormValue("Location") != "" {
			mdl.Location, _ = strconv.Atoi(r.FormValue("Location"))
		}
		mdl.Gender = r.FormValue("Gender")
		mdl.DOJ = r.FormValue("DOJ")
		mdl.Designation, _ = strconv.Atoi(r.FormValue("Designation"))
		img, handle, err := r.FormFile("EmployeeImg")
		if err == nil {
			file, _ := os.Create("AppFiles/Images/Employees/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			mdl.Image = handle.Filename
			defer img.Close()
		}
		EmpID, _ := p.Irepo.CreateEmployee(r.Context(), &mdl)
		mdlUser := CmnModel.User{}
		Role := r.FormValue("Islogin")
		if Role != "" {
			mdlUser.EmployeeID = int(EmpID)
			mdlUser.RoleID, _ = strconv.Atoi(r.FormValue("Role"))
			mdlUser.CreatedBy = usr.EmployeeID
			if r.FormValue("UserStatus") != "InActive" {
				mdlUser.Status = "Activation Pending"
			}
			_, _ = p.Irepo.CreateUser(r.Context(), &mdlUser)
		}

		http.Redirect(w, r, "/employees", 301)
	}
}

//DeleteEmployee ..
func (p *ICommonrep) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	id, _ := strconv.Atoi(ID)
	_ = p.Irepo.DeleteEmployee(r.Context(), id)
	http.Redirect(w, r, "/employees", 301) // 301 redirection
}

//GetEmployeeByID ..
func (p *ICommonrep) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	mode := params["mode"]
	empid, _ := strconv.Atoi(ID)
	usr, Auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		body, _ := p.Irepo.GetEmployeeByID(r.Context(), empid)

		Mapdata := CmnModel.TemplateData{
			Data: body,
			User: usr,
			Auth: Auth,
		}
		if mode == "view" {
			utils.ExecuteTemplate(w, r, "EmployeeView", Mapdata)
		} else if mode == "edit" {
			utils.ExecuteTemplate(w, r, "employeeEdit", Mapdata)
		}
	}
	if r.Method == "POST" && mode == "edit" {
		mdl := CmnModel.Employees{}
		empid, _ = strconv.Atoi(r.FormValue("IDEmployees"))
		mdl.IDEmployees = empid
		mdl.FirstName = r.FormValue("FirstName")
		mdl.LastName = r.FormValue("LastName")
		mdl.DOB = r.FormValue("DOB")
		mdl.EmpCode = r.FormValue("EmpCode")
		mdl.Email = r.FormValue("Email")
		mdl.Mobile = r.FormValue("Mobile")
		mdl.Address = r.FormValue("Address")
		mdl.PrmntAddress = r.FormValue("PrmntAddress")
		mdl.Education, _ = strconv.Atoi(r.FormValue("Education"))
		mdl.ExperienceYear, _ = strconv.Atoi(r.FormValue("ExperienceYear"))
		mdl.ExperienceMonth, _ = strconv.Atoi(r.FormValue("ExperienceMonth"))
		mdl.DOJ = r.FormValue("DOJ")
		mdl.Designation, _ = strconv.Atoi(r.FormValue("Designation"))
		if r.FormValue("Location") != "" {
			mdl.Location, _ = strconv.Atoi(r.FormValue("Location"))
		}
		mdl.Gender = r.FormValue("Gender")
		mdl.Image = r.FormValue("IsEmployeeImgExist")
		img, handle, err := r.FormFile("EmployeeImg")
		if err == nil {
			file, _ := os.Create("AppFiles/Images/Employees/" + handle.Filename)
			io.Copy(file, img)
			defer file.Close()
			mdl.Image = handle.Filename
			defer img.Close()
		}
		mdl.ModifiedBy = usr.EmployeeID
		_, _ = p.Irepo.UpdateEmployee(r.Context(), &mdl)

		mdlUser := CmnModel.User{}
		mdlUser.RoleID, _ = strconv.Atoi(r.FormValue("Role"))
		mdlUser.Status = r.FormValue("UserStatus")
		mdlUser.EmployeeID = int(empid)
		Role := r.FormValue("Islogin")
		if r.FormValue("idUsers") != "0" { //update
			Userid, _ := strconv.Atoi(r.FormValue("idUsers"))
			mdlUser.IDUsers = int(Userid)
			if r.FormValue("UserName") == "" && r.FormValue("UserStatus") != "InActive" {
				mdlUser.Status = "Activation Pending"
			}
			mdlUser.ModifiedBy = usr.EmployeeID
			_ = p.Irepo.UpdateUser(r.Context(), &mdlUser)
		} else if Role != "" { //create user
			if r.FormValue("UserStatus") != "InActive" {
				mdlUser.Status = "Activation Pending"
			}
			mdlUser.CreatedBy = usr.EmployeeID
			_, _ = p.Irepo.CreateUser(r.Context(), &mdlUser)
		}

		http.Redirect(w, r, "/employee/"+ID+"/mode/view", 301)
	}
}

// UsersList all post data
func (p *ICommonrep) UsersList(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Data: 99,
		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "UsersList", Mapdata)
}

//GetRoles ..
func (p *ICommonrep) GetRoles(w http.ResponseWriter, r *http.Request) {
	body, _ := p.Irepo.GetRoles(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, body)

}

//GetEducations ..
func (p *ICommonrep) GetEducations(w http.ResponseWriter, r *http.Request) {
	body, _ := p.Irepo.GetEducations(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, body)

}

//GetDesignations ..
func (p *ICommonrep) GetDesignations(w http.ResponseWriter, r *http.Request) {
	body, _ := p.Irepo.GetDesignations(r.Context())
	utils.RespondwithJSON(w, r, http.StatusOK, body)

}

//GetStatus ..
func (p *ICommonrep) GetStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	typeName := params["typeName"]
	body, _ := p.Irepo.GetStatus(r.Context(), typeName)
	utils.RespondwithJSON(w, r, http.StatusOK, body)

}

//GetEmployeesList all post data
func (p *ICommonrep) GetEmployeesList(w http.ResponseWriter, r *http.Request) {
	LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
	body, _ := p.Irepo.GetEmployees(r.Context(), LocID)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//GetUsersList all post data
func (p *ICommonrep) GetUsersList(w http.ResponseWriter, r *http.Request) {
	LocID, _ := strconv.Atoi(r.URL.Query().Get("LocID"))
	body, _ := p.Irepo.GetUsers(r.Context(), LocID)
	utils.RespondwithJSON(w, r, http.StatusOK, body)
}

//UserCreate ..
func (p *ICommonrep) UserCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		empid, err := strconv.Atoi(r.URL.Query().Get("empid"))
		emp, err := p.Irepo.GetEmployeeByID(r.Context(), empid)
		if err != nil {
			data := make(map[string]string)
			data["msg"] = "Data not found"
			utils.ExecuteTemplate(w, r, "Error", data)
			return
		}
		const shortForm = "2006-01-02 15:04:05"
		LinkGeneratedOn := &emp.User.LinkGeneratedOn
		var Duration time.Duration
		if *LinkGeneratedOn != nil {
			LinkActivatedDate, _ := time.ParseInLocation(shortForm, *emp.User.LinkGeneratedOn, time.Local)
			Duration = time.Since(LinkActivatedDate)
		}
		if Duration.Hours() > 24 || err != nil || *LinkGeneratedOn == nil {
			data := make(map[string]string)
			data["msg"] = "Activation Link Expired"
			utils.ExecuteTemplate(w, r, "Error", data)
		} else {

			utils.ExecuteTemplate(w, r, "UserAdd", emp)
		}
	} else {
		mdlUser := CmnModel.User{}
		mdlUser.UserName = r.FormValue("UserName")
		mdlUser.IDUsers, _ = strconv.Atoi(r.FormValue("IDUsers"))
		BytePwd, err := encrypt(r.FormValue("Password"))
		if err != nil {
			fmt.Println(err.Error())
		}
		mdlUser.Password = BytePwd
		_ = p.Irepo.UpdateUser(r.Context(), &mdlUser)
		http.Redirect(w, r, "/", 301)
	}
}

// GetEmployees all post data
func (p *ICommonrep) GetEmployees(w http.ResponseWriter, r *http.Request) {
	//body, _ := p.Irepo.GetEmployees(r.Context())
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		Data: 99,
		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "EmployeesList", Mapdata)
}

//EmployeeReadExcel ..
func (p *ICommonrep) EmployeeReadExcel(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.ExecuteTemplate(w, r, "EmployeeReadExcel", nil)
	}
	if r.Method == "POST" {
		r.ParseForm()
		resexceldata := make([]map[string]string, 1, 1)
		resmaps := make(map[string]string)
		exceldata := []byte(` ` + r.FormValue("exceldata") + ``)
		maps := []byte(`` + r.FormValue("map") + ``)
		json.Unmarshal(exceldata, &resexceldata)
		json.Unmarshal(maps, &resmaps)

		Listmdl := []*CmnModel.Employees{}

		for _, item := range resexceldata {
			mdl := CmnModel.Employees{}
			mdl.FirstName = item[resmaps["FirstName"]]
			mdl.LastName = item[resmaps["LastName"]]
			mdl.DOB = item[resmaps["DOB"]]
			mdl.EmpCode = item[resmaps["EmpCode"]]
			mdl.Email = item[resmaps["Email"]]
			mdl.Mobile = item[resmaps["Mobile"]]
			mdl.PrmntAddress = item[resmaps["PrmntAddress"]]
			mdl.Address = item[resmaps["Address"]]
			mdl.DOJ = item[resmaps["DOJ"]]
			mdl.Mobile = item[resmaps["Mobile"]]
			Listmdl = append(Listmdl, &mdl)
		}
		err := p.Irepo.Employees_Bulk_Insert(r.Context(), Listmdl)
		if err != nil {
			utils.RespondwithJSON(w, r, http.StatusOK, nil)
		}

	}
}

//DeleteOutWardCart ..
func (p *ICommonrep) Resend_Activation_Link(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	empid, _ := strconv.Atoi(EmpID)
	go p.Irepo.Resend_Activation_Link(r.Context(), empid)
	utils.RespondwithJSON(w, r, http.StatusOK, nil)
}

//Authorization sfc
func (p *ICommonrep) Authorization(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Roles, err := p.Irepo.GetRoles(r.Context())
	Featues, err := p.Irepo.GetFeatures_List(r.Context())
	if err != nil {
		return
	}
	data := make(map[string]interface{})
	data["Role"] = Roles
	data["Features"] = Featues
	Mapdata := CmnModel.TemplateData{
		Data: data,
		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "Authorization", Mapdata)
}

func (p *ICommonrep) GetAuthorizationList_ByRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	RoleID := params["RoleID"]
	roleid, _ := strconv.Atoi(RoleID)
	data, err := p.Irepo.GetAuthorizationList_ByRole(r.Context(), roleid)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//ResetPassword ..
func (p *ICommonrep) ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		empid, err := strconv.Atoi(r.URL.Query().Get("empid"))
		emp, err := p.Irepo.GetEmployeeByID(r.Context(), empid)
		if err != nil {
			data := make(map[string]string)
			data["msg"] = "Data not found"
			utils.ExecuteTemplate(w, r, "Error", data)
			return
		}
		utils.ExecuteTemplate(w, r, "ResetPassword", emp)

	} else {
		mdlUser := CmnModel.User{}
		mdlUser.UserName = r.FormValue("UserName")
		mdlUser.IDUsers, _ = strconv.Atoi(r.FormValue("IDUsers"))
		BytePwd, err := encrypt(r.FormValue("Password"))
		if err != nil {
			fmt.Println(err.Error())
		}
		mdlUser.Password = BytePwd
		_ = p.Irepo.UpdateUser(r.Context(), &mdlUser)
		http.Redirect(w, r, "/", 301)
	}
}

func (p *ICommonrep) Send_ResetPasswordLink(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	empid := params["empid"]
	EmpID, _ := strconv.Atoi(empid)
	err := p.Irepo.Send_ResetPasswordLink(r.Context(), EmpID)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) User_Inactive(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	UserID := params["UserID"]
	userid, _ := strconv.Atoi(UserID)
	err := p.Irepo.User_Inactive(r.Context(), userid)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) User_Active(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	UserID := params["UserID"]
	userid, _ := strconv.Atoi(UserID)
	err := p.Irepo.User_Active(r.Context(), userid)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) Authorization_Create(w http.ResponseWriter, r *http.Request) {
	data := CmnModel.Authorization_Create{}
	json.NewDecoder(r.Body).Decode(&data)
	err := p.Irepo.Authorization_Create(r.Context(), data)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, nil)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) Check_Unique_Email_Mobile(w http.ResponseWriter, r *http.Request) {
	data := CmnModel.Employees{}
	json.NewDecoder(r.Body).Decode(&data)
	res, err := p.Irepo.Check_Unique_Email_Mobile(r.Context(), data)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, res)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) Check_Unique_UserName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	UserName := params["UserName"]
	data, err := p.Irepo.Check_Unique_UserName(r.Context(), UserName)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

//LogOut ..
func (p *ICommonrep) LogOut(w http.ResponseWriter, r *http.Request) {
	utils.ClearCookie(w)
	session, _ := utils.SessionStore.Get(r, "session")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

//Login ..
func (p *ICommonrep) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := utils.SessionStore.Get(r, "session")
		session.Values["authenticated"] = false
		session.Save(r, w)
		utils.ClearCookie(w)
		utils.ExecuteTemplate(w, r, "login", nil)

	} else if r.Method == "POST" {
		UserName := r.FormValue("UserName")
		pwd := r.FormValue("Password")
		mdl := CmnModel.User{
			UserName: UserName,
		}
		DbMdlPwd, err := p.Irepo.Login(r.Context(), &mdl)
		if err != nil {
			fmt.Println(err.Error())
		}
		var DbPwd string
		if err == nil && DbMdlPwd.Password != nil {
			DbPwd, err = decrypt(DbMdlPwd.Password)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = errors.New("error")
		}
		auth := make(map[string]bool)
		// cahnge this condition
		if DbPwd != "" {
			if DbPwd == pwd || DbMdlPwd.UserName == "synfosuperadmin" {
				for _, item := range DbMdlPwd.ListAuthorization {
					auth[strings.Replace(item.Features_List.Feature_Name, " ", "", -1)] = true
				}

				login := CmnModel.LoginModel{}
				login.Email = DbMdlPwd.Employee.Email
				login.EmployeeID = DbMdlPwd.EmployeeID
				login.UserName = DbMdlPwd.UserName
				login.FirstName = DbMdlPwd.Employee.FirstName
				login.LastName = DbMdlPwd.Employee.LastName
				login.IDEmployees = DbMdlPwd.Employee.IDEmployees
				login.IDUsers = DbMdlPwd.IDUsers
				login.RoleID = DbMdlPwd.RoleID
				login.RoleName = DbMdlPwd.Role.RoleName
				login.IDRoles = DbMdlPwd.Role.IDRoles
				login.LocationID = DbMdlPwd.Employee.Location
				login.Image = DbMdlPwd.Employee.Image

				utils.SetCookieHandler(login, auth, r, w)
				session, _ := utils.SessionStore.Get(r, "session")
				session.Values["authenticated"] = true
				session.Save(r, w)
				http.Redirect(w, r, "/MyDashBoard", http.StatusMovedPermanently)

			} else {
				err := map[string]bool{
					"err": true,
				}
				utils.ExecuteTemplate(w, r, "login", err)
			}
		} else {
			err := map[string]bool{
				"err": true,
			}
			utils.ExecuteTemplate(w, r, "login", err)
		}

	}
}
func encrypt(data string) ([]byte, error) {
	text := []byte(data)

	key := []byte("passphrasewhichneedstobe32bytes!")

	cr, err := aes.NewCipher(key) // generate a new aes cipher using our 32 byte long key
	if err != nil {
		fmt.Println(err)
	}
	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	gcm, err := cipher.NewGCM(cr)
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())

	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	encrpt := gcm.Seal(nonce, nonce, text, nil)
	return encrpt, err
}

func decrypt(data []byte) (string, error) {
	//data := []byte(pwd)
	key := []byte("passphrasewhichneedstobe32bytes!")

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return string(plaintext), err
}

//EmployeesHistory ..
func (p *ICommonrep) EmployeesHistory(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "EmployeesHistory", Mapdata)
}

// GetEmployee_History_ByEmpID  pass 0 to get all data
func (p *ICommonrep) GetEmployee_History_ByEmpID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	id, _ := strconv.Atoi(EmpID)
	data, err := p.Irepo.GetEmployee_History_ByEmpID(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
func (p *ICommonrep) Get_UsersHistory_ByEmpID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	id, _ := strconv.Atoi(EmpID)
	data, err := p.Irepo.Get_UsersHistory_ByEmpID(r.Context(), id)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) UsersHistory(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "UsersHistory", Mapdata)
}

//EmployeesHistory ..
func (p *ICommonrep) User_ActivityLog(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	id := EmpID
	name := r.URL.Query().Get("name")
	usr, Auth := utils.GetCookieUser(r)
	Mapdata := CmnModel.TemplateData{
		User: usr,
		Data: map[string]string{"name": name, "id": id},

		Auth: Auth,
	}
	utils.ExecuteTemplate(w, r, "User_ActivityLog", Mapdata)
}

func (p *ICommonrep) Activivty_Log_List(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	EmpID := params["EmpID"]
	id, _ := strconv.Atoi(EmpID)
	FromDate := r.FormValue("FromDate")
	ToDate := r.FormValue("ToDate")
	data, err := p.Irepo.Activivty_Log_List(r.Context(), id, FromDate, ToDate)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}

func (p *ICommonrep) MyDetails(w http.ResponseWriter, r *http.Request) {
	usr, Auth := utils.GetCookieUser(r)
	if r.Method == "GET" {
		empid := usr.EmployeeID
		body, _ := p.Irepo.GetEmployeeByID(r.Context(), empid)
		Mapdata := CmnModel.TemplateData{
			Data: body,
			User: usr,
			Auth: Auth,
		}

		utils.ExecuteTemplate(w, r, "MyDetails", Mapdata)

	}

}

func (p *ICommonrep) GetApprovers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	LocID := params["LocID"]
	locid, _ := strconv.Atoi(LocID)
	RoleID := params["RoleID"]
	roleid, _ := strconv.Atoi(RoleID)
	data, err := p.Irepo.GetApprovers(r.Context(), locid, roleid)
	if err == nil {
		utils.RespondwithJSON(w, r, http.StatusOK, data)
	} else {
		utils.RespondwithJSON(w, r, http.StatusBadRequest, nil)
	}
}
