package service

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/plus100kt/goserver/gag/eclass/model"
	"github.com/plus100kt/goserver/gag/util"
)

func (s *eclassService) Login(body *model.LoginBody) (string, error) {
	// struct to formdata
	ct, formData, err := util.StructToForm(body)
	if err != nil {
		return "", err
	}

	// request
	res, err := http.Post("https://eclass.tukorea.ac.kr/ilos/lo/login.acl", ct, formData)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", err
	}

	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// 성공
	responseString := string(responseBytes)
	if strings.Contains(responseString, `document.location.href="https://eclass.tukorea.ac.kr/ilos/main/main_form.acl"`) {

		// set cookie
		s.cookies = res.Cookies()
		return responseString, err
	}

	// 실패
	if strings.Contains(responseString, "로그인 정보가 일치하지 않습니다.") {
		return "", errors.New("로그인 정보가 일치하지 않습니다.")
	}

	return "", errors.New("로그인 정보가 일치하지 않습니다.")
}
