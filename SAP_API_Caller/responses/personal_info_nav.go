package responses

type PersonalInfoNav struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			PersonIDExternal     string      `json:"personIdExternal"`
			StartDate            string      `json:"startDate"`
			LastName             string      `json:"lastName"`
			BusinessLastName     string      `json:"businessLastName"`
			LastModifiedDateTime string      `json:"lastModifiedDateTime"`
			Gender               string      `json:"gender"`
			EndDate              string      `json:"endDate"`
			DisplayName          string      `json:"displayName"`
			CreatedDateTime      string      `json:"createdDateTime"`
			Title                string      `json:"title"`
			CreatedOn            string      `json:"createdOn"`
			BusinessFirstName    string      `json:"businessFirstName"`
			AttachmentID         string      `json:"attachmentId"`
			PreferredName        string      `json:"preferredName"`
			Initials             string      `json:"initials"`
			LastModifiedBy       string      `json:"lastModifiedBy"`
			Script               string      `json:"script"`
			LastModifiedOn       string      `json:"lastModifiedOn"`
			FirstName            string      `json:"firstName"`
			Nationality          string      `json:"nationality"`
			CreatedBy            string      `json:"createdBy"`
			MiddleName           string      `json:"middleName"`
			NativePreferredLang  string      `json:"nativePreferredLang"`
			Salutation           string      `json:"salutation"`
			MaritalStatus        string      `json:"maritalStatus"`
			LocalNavAUS          struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavAUS"`
			LocalNavESP struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavESP"`
			LocalNavUSA struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavUSA"`
			NativePreferredLangNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"nativePreferredLangNav"`
			MaritalStatusNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"maritalStatusNav"`
			PersonNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"personNav"`
			LocalNavFIN struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavFIN"`
			LocalNavARE struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavARE"`
			LocalNavDEU struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavDEU"`
			LocalNavFRA struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavFRA"`
			LocalNavITA struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavITA"`
			LocalNavBRA struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavBRA"`
			LocalNavZAF struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavZAF"`
			TitleNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"titleNav"`
			LocalNavQAT struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavQAT"`
			LocalNavSAU struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavSAU"`
			LocalNavNLD struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavNLD"`
			LocalNavCHE struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavCHE"`
			LocalNavDNK struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavDNK"`
			SalutationNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"salutationNav"`
			LocalNavCHN struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavCHN"`
			WfRequestNav struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"wfRequestNav"`
			LocalNavGBR struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavGBR"`
			LocalNavPRT struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"localNavPRT"`
		} `json:"results"`
	} `json:"d"`
}
