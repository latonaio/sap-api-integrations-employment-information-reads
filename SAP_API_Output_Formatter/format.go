package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-employment-information-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
	PersonIDExternal:             data.PersonIDExternal,
	UserID:                       data.UserID,
	LastModifiedDateTime:         data.LastModifiedDateTime,
	ServiceDate:                  data.ServiceDate,
	EndDate:                      data.EndDate,
	BonusPayExpirationDate:       data.BonusPayExpirationDate,
	CreatedDateTime:              data.CreatedDateTime,
	IsContingentWorker:           data.IsContingentWorker,
	InitialOptionGrant:           data.InitialOptionGrant,
	InitialStockGrant:            data.InitialStockGrant,
	EmployeeFirstEmployment:      data.EmployeeFirstEmployment,
	EligibleForStock:             data.EligibleForStock,
	EligibleForSalContinuation:   data.EligibleForSalContinuation,
	OkToRehire:                   data.OkToRehire,
	AssignmentIDExternal:         data.AssignmentIDExternal,
	ProfessionalServiceDate:      data.ProfessionalServiceDate,
	SalaryEndDate:                data.SalaryEndDate,
	SeniorityDate:                data.SeniorityDate,
	StartDate:                    data.StartDate,
	HiringNotCompleted:           data.HiringNotCompleted,
	IsECRecord:                   data.IsECRecord,
	PrevEmployeeID:               data.PrevEmployeeID,
	RegretTermination:            data.RegretTermination,
	CreatedOn:                    data.CreatedOn,
	LastDateWorked:               data.LastDateWorked,
	FirstDateWorked:              data.FirstDateWorked,
	OriginalStartDate:            data.OriginalStartDate,
	PayrollEndDate:               data.PayrollEndDate,
	BenefitsEligibilityStartDate: data.BenefitsEligibilityStartDate,
	StockEndDate:                 data.StockEndDate,
	BenefitsEndDate:              data.BenefitsEndDate,
	AssignmentClass:              data.AssignmentClass,
	LastModifiedBy:               data.LastModifiedBy,
	LastModifiedOn:               data.LastModifiedOn,
	CreatedBy:                    data.CreatedBy,
	PersonNav:                    data.PersonNav.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToPersonNav(raw []byte, l *logger.Logger) (*PersonNav, error) {
	pm := &responses.PersonNav{}

	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PersonNav. unmarshal error: %w", err)
	}
	
	return &PersonNav{
	PersonIDExternal:     pm.D.PersonIDExternal,
	LastModifiedDateTime: pm.D.LastModifiedDateTime,
	LastModifiedBy:       pm.D.LastModifiedBy,
	CreatedDateTime:      pm.D.CreatedDateTime,
	DateOfBirth:          pm.D.DateOfBirth,
	PerPersonUUID:        pm.D.PerPersonUUID,
	CreatedOn:            pm.D.CreatedOn,
	LastModifiedOn:       pm.D.LastModifiedOn,
	CountryOfBirth:       pm.D.CountryOfBirth,
	CreatedBy:            pm.D.CreatedBy,
	PersonID:             pm.D.PersonID,
	PersonalInfoNav:      pm.D.PersonalInfoNav.Deferred.URI,
	}, nil
}

func ConvertToPersonalInfoNav(raw []byte, l *logger.Logger) ([]PersonalInfoNav, error) {
	pm := &responses.PersonalInfoNav{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to PersonalInfoNav. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	personalInfoNav := make([]PersonalInfoNav, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		personalInfoNav = append(personalInfoNav, PersonalInfoNav{
	PersonIDExternal:     data.PersonIDExternal,
	StartDate:            data.StartDate,
	LastName:             data.LastName,
	BusinessLastName:     data.BusinessLastName,
	LastModifiedDateTime: data.LastModifiedDateTime,
	Gender:               data.Gender,
	EndDate:              data.EndDate,
	DisplayName:          data.DisplayName,
	CreatedDateTime:      data.CreatedDateTime,
	Title:                data.Title,
	CreatedOn:            data.CreatedOn,
	BusinessFirstName:    data.BusinessFirstName,
	AttachmentID:         data.AttachmentID,
	PreferredName:        data.PreferredName,
	Initials:             data.Initials,
	LastModifiedBy:       data.LastModifiedBy,
	Script:               data.Script,
	LastModifiedOn:       data.LastModifiedOn,
	FirstName:            data.FirstName,
	Nationality:          data.Nationality,
	CreatedBy:            data.CreatedBy,
	MiddleName:           data.MiddleName,
	NativePreferredLang:  data.NativePreferredLang,
	Salutation:           data.Salutation,
	MaritalStatus:        data.MaritalStatus,
		})
	}

	return personalInfoNav, nil
}
