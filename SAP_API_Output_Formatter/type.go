package sap_api_output_formatter

type Header struct {
	PersonIDExternal             string      `json:"PersonIDExternal"`
	UserID                       string      `json:"UserID"`
	LastModifiedDateTime         string      `json:"LastModifiedDateTime"`
	ServiceDate                  string      `json:"ServiceDate"`
	EndDate                      string      `json:"EndDate"`
	BonusPayExpirationDate       string      `json:"BonusPayExpirationDate"`
	CreatedDateTime              string      `json:"CreatedDateTime"`
	IsContingentWorker           bool        `json:"IsContingentWorker"`
	InitialOptionGrant           string      `json:"InitialOptionGrant"`
	InitialStockGrant            string      `json:"InitialStockGrant"`
	EmployeeFirstEmployment      string      `json:"EmployeeFirstEmployment"`
	EligibleForStock             bool        `json:"EligibleForStock"`
	EligibleForSalContinuation   string      `json:"EligibleForSalContinuation"`
	OkToRehire                   string      `json:"OkToRehire"`
	AssignmentIDExternal         string      `json:"AssignmentIDExternal"`
	ProfessionalServiceDate      string      `json:"ProfessionalServiceDate"`
	SalaryEndDate                string      `json:"SalaryEndDate"`
	SeniorityDate                string      `json:"SeniorityDate"`
	StartDate                    string      `json:"StartDate"`
	HiringNotCompleted           bool        `json:"HiringNotCompleted"`
	IsECRecord                   bool        `json:"IsECRecord"`
	PrevEmployeeID               string      `json:"PrevEmployeeID"`
	RegretTermination            string      `json:"RegretTermination"`
	CreatedOn                    string      `json:"CreatedOn"`
	LastDateWorked               string      `json:"LastDateWorked"`
	FirstDateWorked              string      `json:"FirstDateWorked"`
	OriginalStartDate            string      `json:"OriginalStartDate"`
	PayrollEndDate               string      `json:"PayrollEndDate"`
	BenefitsEligibilityStartDate string      `json:"BenefitsEligibilityStartDate"`
	StockEndDate                 string      `json:"StockEndDate"`
	BenefitsEndDate              string      `json:"BenefitsEndDate"`
	AssignmentClass              string      `json:"AssignmentClass"`
	LastModifiedBy               string      `json:"LastModifiedBy"`
	LastModifiedOn               string      `json:"LastModifiedOn"`
	CreatedBy                    string      `json:"CreatedBy"`
	PersonNav                    string      `json:"PersonNav"`
}

type PersonNav struct {
	PersonIDExternal     string `json:"PersonIDExternal"`
	LastModifiedDateTime string `json:"LastModifiedDateTime"`
	LastModifiedBy       string `json:"LastModifiedBy"`
	CreatedDateTime      string `json:"CreatedDateTime"`
	DateOfBirth          string `json:"DateOfBirth"`
	PerPersonUUID        string `json:"PerPersonUUID"`
	CreatedOn            string `json:"CreatedOn"`
	LastModifiedOn       string `json:"LastModifiedOn"`
	CountryOfBirth       string `json:"CountryOfBirth"`
	CreatedBy            string `json:"CreatedBy"`
	PersonID             string `json:"PersonID"`
	PersonalInfoNav      string `json:"PersonalInfoNav"`	
}

type PersonalInfoNav struct {
	PersonIDExternal     string      `json:"PersonIDExternal"`
	StartDate            string      `json:"StartDate"`
	LastName             string      `json:"LastName"`
	BusinessLastName     string      `json:"BusinessLastName"`
	LastModifiedDateTime string      `json:"LastModifiedDateTime"`
	Gender               string      `json:"Gender"`
	EndDate              string      `json:"EndDate"`
	DisplayName          string      `json:"DisplayName"`
	CreatedDateTime      string      `json:"CreatedDateTime"`
	Title                string      `json:"Title"`
	CreatedOn            string      `json:"CreatedOn"`
	BusinessFirstName    string      `json:"BusinessFirstName"`
	AttachmentID         string      `json:"AttachmentID"`
	PreferredName        string      `json:"PreferredName"`
	Initials             string      `json:"Initials"`
	LastModifiedBy       string      `json:"LastModifiedBy"`
	Script               string      `json:"Script"`
	LastModifiedOn       string      `json:"LastModifiedOn"`
	FirstName            string      `json:"FirstName"`
	Nationality          string      `json:"Nationality"`
	CreatedBy            string      `json:"CreatedBy"`
	MiddleName           string      `json:"MiddleName"`
	NativePreferredLang  string      `json:"NativePreferredLang"`
	Salutation           string      `json:"Salutation"`
	MaritalStatus        string      `json:"MaritalStatus"`
}
