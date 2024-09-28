package oauth

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Long-Plan/longplan-api/config"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/Long-Plan/longplan-api/pkg/requestor"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	User UserDto `json:"user"`
	jwt.RegisteredClaims
}

type errorDto struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type accessTokenDto struct {
	AccessToken string `json:"access_token"`
}

type UserDto struct {
	CmuitaccountName   string `json:"cmuitaccount_name"`
	Cmuitaccount       string `json:"cmuitaccount"`
	StudentID          string `json:"student_id"`
	PrenameID          string `json:"prename_id"`
	PrenameTH          string `json:"prename_TH"`
	PrenameEN          string `json:"prename_EN"`
	FirstnameTH        string `json:"firstname_TH"`
	FirstnameEN        string `json:"firstname_EN"`
	LastnameTH         string `json:"lastname_TH"`
	LastnameEN         string `json:"lastname_EN"`
	OrganizationCode   string `json:"organization_code"`
	OrganizationNameTH string `json:"organization_name_TH"`
	OrganizationNameEN string `json:"organization_name_EN"`
	ItaccounttypeID    string `json:"itaccounttype_id"`
	ItaccounttypeTH    string `json:"itaccounttype_TH"`
	ItaccounttypeEN    string `json:"itaccounttype_EN"`
}

func CmuOauthValidation(code string, isLocalOrigin bool) (*UserDto, error) {
	accessToken, err := getAccessToken(code, isLocalOrigin)
	if err != nil {
		return nil, err
	}
	user, err := getCmuBasicInfo(accessToken.AccessToken)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getAccessToken(code string, isLocalOrigin bool) (*accessTokenDto, error) {
	mode := os.Getenv("mode")
	config := config.Config.CmuOauth
	url := config.CmuOauthToken
	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	redirectURI := config.CmuOauthRedirectURL
	if isLocalOrigin && mode == "dev" {
		redirectURI = config.CmuOauthRedirectURLLocal
	}

	params := map[string]interface{}{
		"code":          code,
		"redirect_uri":  redirectURI,
		"client_id":     config.CmuOauthClientID,
		"client_secret": config.CmuOauthClientSecret,
		"grant_type":    "authorization_code",
	}
	paramsEncode := requestor.BuildQueryParams(params)
	payload := strings.Replace(paramsEncode, "?", "", 1)
	res, statusCode, err := requestor.HttpPost[interface{}](url, header, payload)
	if err != nil {
		return nil, errors.InternalErr(err.Error())
	}

	log.Info("result: ", res)

	statusCodeStr := strconv.Itoa(statusCode)
	if strings.HasPrefix(statusCodeStr, "2") {
		var result accessTokenDto
		lodash.Recast(res, &result)
		return &result, nil
	} else {

		return nil, errors.CmuOauthErr("can't get access_token")
	}
}

func getCmuBasicInfo(accessToken string) (*UserDto, error) {
	config := config.Config.CmuOauth
	url := config.CmuOauthInfo
	header := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %v", accessToken),
	}
	res, statusCode, err := requestor.HttpGet[map[string]interface{}](url, header)
	if err != nil {
		return nil, errors.InternalErr(err.Error())
	}

	statusCodeStr := strconv.Itoa(statusCode)
	if strings.HasPrefix(statusCodeStr, "2") {
		var result UserDto
		lodash.Recast(res, &result)
		return &result, nil
	} else {
		var errorModel errorDto
		lodash.Recast(res, &errorModel)
		return nil, errors.CmuOauthErr(fmt.Sprintf("error: %v, error_desc: %v", errorModel.Error, errorModel.ErrorDescription))
	}
}
