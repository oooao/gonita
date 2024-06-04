package gonita

import "github.com/go-resty/resty/v2"

const (
	URI_BPM = "API/"
)

// TODO: Need chanage to Variables of Environment
var (
	user_ppassword = ""
	server_ip_port = ""
)

//
//  FormInput
//  @Description: CRUD必要的 JSON 外層結構
//
type FormInput struct {
	ModelInput *interface{} `json:"modelInput,omitempty"`
}

type BPMClient struct {
	serverUri  string
	apiUri     string
	username   string
	password   string
	request    *resty.Request
	token      string
	jSessionId string // JSESSIONID
}
