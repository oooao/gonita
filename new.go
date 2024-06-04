package gonita

import (
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

//  New
//  @Description: 傳入 username 創建該人員Client端請求實例.
//  @param username
func New(username string) *BPMClient {

	//建立rest client端結構
	b := &BPMClient{
		serverUri:  "http://" + server_ip_port + "/bonita/",
		apiUri:     "",
		username:   username,
		password:   user_ppassword,
		request:    resty.New().NewRequest(),
		token:      "",
		jSessionId: "",
	}

	//http://server_ip_port/bonita/ + API/
	b.apiUri = b.serverUri + URI_BPM

	//進行身分驗證並取得token
	b.GetLoginToken()

	log.Println("New() - b: X-Bonita-API-Token: ", b.request.Header.Get("X-Bonita-API-Token"))

	return b
}

//  GetLoginToken
//  @Description: Login, /bonita/loginservice
//  @receiver b
// https://documentation.bonitasoft.com/bonita/2021.2/api/rest-api-authentication
func (b *BPMClient) GetLoginToken() {

	uri := b.serverUri + "loginservice"
	log.Println("GetLoginToken() - uri: ", uri)

	//新增一個新的resty客戶端與請求實例。
	r := resty.New().NewRequest()
	// Content-Type : application/x-www-form-urlencoded
	resp, err := r.
		SetFormData(map[string]string{
			"username": b.username,
			"password": b.password,
		}).
		Post(uri)
	if err != nil {
		log.Print(err)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "X-Bonita-API-Token" {
			b.token = cookie.Value // 獲取 token
		}
	}
	log.Println("GetLoginToken() - b.token: ", b.token)

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "JSESSIONID" {
			b.jSessionId = cookie.Value // 獲取 JSESSIONID
		}
	}
	log.Println("GetLoginToken() - b.jSessionId: ", b.jSessionId)

	log.Printf("GetLoginToken() - Request - r:\n %+v", r)

	b.request = b.request.
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			// "Charset":            "utf-8",
			// "Accept":             "application/json",
			"X-Bonita-API-Token": b.token,
			// "JSESSIONID":         b.jSessionId,
		}).SetCookie(&http.Cookie{
		Name:       "JSESSIONID",
		Value:      b.jSessionId,
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	})
	log.Printf("GetLoginToken() - b.request:\n %+v", b.request)
}
