package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-employee-basic-data-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToBusinessUserCollection(raw []byte, l *logger.Logger) ([]BusinessUserCollection, error) {
	pm := &responses.BusinessUserCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to BusinessUserCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	businessUserCollection := make([]BusinessUserCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		businessUserCollection = append(businessUserCollection, BusinessUserCollection{
			ObjectID:                             data.ObjectID,                          
			ETag:                                 data.ETag,                              
			EmployeeID:                           data.EmployeeID,                        
			EmployeeUUID:                         data.EmployeeUUID,                      
			UserID:                               data.UserID,                            
			TechnicalUserID:                      data.TechnicalUserID,                   
			IdentityUUID:                         data.IdentityUUID,                      
			BusinessPartnerID:                    data.BusinessPartnerID,                 
			BusinessPartnerFormattedName:         data.BusinessPartnerFormattedName,
			DepartmentName:                       data.DepartmentName,                    
			CompanyName:                          data.CompanyName,                       
			ManagerName:                          data.ManagerName,                      
			EmailURI:                             data.EmailURI,                         
			DecimalFormatCode:                    data.DecimalFormatCode,                
			DecimalFormatCodeText:                data.DecimalFormatCodeText,           
			DateFormatCode:                       data.DateFormatCode,                   
			DateFormatCodeText:                   data.DateFormatCodeText,               
			TimeFormatCode:                       data.TimeFormatCode,                   
			TimeFormatCodeText:                   data.TimeFormatCodeText,               
			TimeZoneCode:                         data.TimeZoneCode,                    
			TimeZoneCodeText:                     data.TimeZoneCodeText,                 
			LogonLanguageCode:                    data.LogonLanguageCode,                
			LogonLanguageCodeText:                data.LogonLanguageCodeText,           
			UserValidityStartDate:                data.UserValidityStartDate,           
			UserValidityEndDate:                  data.UserValidityEndDate,             
			UserLockedIndicator:                  data.UserLockedIndicator,              
			UserCountedIndicator:                 data.UserCountedIndicator,             
			PasswordPolicyCode:                   data.PasswordPolicyCode,              
			PasswordPolicyCodeText:               data.PasswordPolicyCodeText,          
			PasswordInactiveIndicator:            data.PasswordInactiveIndicator,       
			PasswordLockedIndicator:              data.PasswordLockedIndicator,         
			UserAccountTypeCode:                  data.UserAccountTypeCode,             
			UserAccountTypeCodeText:              data.UserAccountTypeCodeText,         
			CreatedOn:                            data.CreatedOn,                      
			UserCreatedBy:                        data.UserCreatedBy,                              
			EntityLastChangedOn:                  data.EntityLastChangedOn,              
			UserChangedBy:                        data.UserChangedBy,                    
			UserChangedOn:                        data.UserChangedOn,
			ToBusinessUserBusinessRoleAssignment: data.BusinessUserBusinessRoleAssignment.Deferred.URI,
		})
	}

	return businessUserCollection, nil
}

func ConvertToToBusinessUserBusinessRoleAssignment(raw []byte, l *logger.Logger) ([]ToBusinessUserBusinessRoleAssignment, error) {
	pm := &responses.ToBusinessUserBusinessRoleAssignment{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToBusinessUserBusinessRoleAssignment. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toBusinessUserBusinessRoleAssignment := make([]ToBusinessUserBusinessRoleAssignment, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toBusinessUserBusinessRoleAssignment = append(toBusinessUserBusinessRoleAssignment, ToBusinessUserBusinessRoleAssignment{
			ObjectID:                   data.ObjectID,
			ParentObjectID:      		data.ParentObjectID,
			EmployeeID:          		data.EmployeeID,
			UserID:              		data.UserID,
			BusinessRoleID:      		data.BusinessRoleID,
			EntityLastChangedOn: 		data.EntityLastChangedOn,
		})
	}

	return toBusinessUserBusinessRoleAssignment, nil
}

func ConvertToEmployeeBasicData(raw []byte, l *logger.Logger) ([]EmployeeBasicData, error) {
	pm := &responses.EmployeeBasicData{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to EmployeeBasicData. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	employeeBasicData := make([]EmployeeBasicData, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		employeeBasicData = append(employeeBasicData, EmployeeBasicData{
			ObjectID:                           data.ObjectID,                                           
			ETag:                               data.ETag,                                                            
			EmployeeID:                         data.EmployeeID,                                                      
			EmployeeUUID:                       data.EmployeeUUID,                                                    
			UserID:                             data.UserID,                                                          
			IdentityUUID:                       data.IdentityUUID,                                                    
			BusinessPartnerID:                  data.BusinessPartnerID,                                               
			CurrentInternalEmployeeIndicator:   data.CurrentInternalEmployeeIndicator,                                
			CurrentExternalEmployeeIndicator:   data.CurrentExternalEmployeeIndicator,                                
			FormattedName:                      data.FormattedName,                                                   
			TitleCode:                          data.TitleCode,                                                       
			AcademicTitleCode:                  data.AcademicTitleCode,                                               
			FirstName:                          data.FirstName,                         
			MiddleName:                         data.MiddleName,                                                      
			LastName:                           data.LastName,                                                        
			SecondLastName:                     data.SecondLastName,                                                  
			NickName:                           data.NickName,                                                        
			GenderCode:                         data.GenderCode,                                                      
			LanguageCode:                       data.LanguageCode,                                                    
			FormattedAddress:                   data.FormattedAddress,                                                
			CountryCode:                        data.CountryCode,                                                     
			RegionCode:                         data.RegionCode,                                                      
			AddressLine1:                       data.AddressLine1,                                                    
			AddressLine2:                       data.AddressLine2,                                                    
			HouseNumber:                        data.HouseNumber,                                                     
			Street:                             data.Street,                                                          
			AddressLine4:                       data.AddressLine4,                                                    
			AddressLine5:                       data.AddressLine5,                                                    
			City:                               data.City,                                                            
			PostalCode:                         data.PostalCode,                                                      
			Phone:                              data.Phone,                                                           
			Mobile:                             data.Mobile,                                                          
			Fax:                                data.Fax,                                                             
			Email:                              data.Email,                                                           
			UserValidityStartDate:              data.UserValidityStartDate,                                           
			UserValidityEndDate:                data.UserValidityEndDate,                                             
			UserPasswordPolicyCode:             data.UserPasswordPolicyCode,                                          
			UserLockedIndicator:                data.UserLockedIndicator,                                             
			TimeZoneCode:                       data.TimeZoneCode,                                                    
			ManagerUUID:                        data.ManagerUUID,                                                     
			ManagerFormattedName:               data.ManagerFormattedName,                                            
			JobName:                            data.JobName,                                                         
			CreatedOn:                          data.CreatedOn,                                                       
			CreatedBy:                          data.CreatedBy,                                                       
			ChangedOn:                          data.ChangedOn,                                                       
			ChangedBy:                          data.ChangedBy,                                                       
			EntityLastChangedOn:                data.EntityLastChangedOn,                                             
		})
	}

	return employeeBasicData, nil
}