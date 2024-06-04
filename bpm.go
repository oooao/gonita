package gonita

import (
	"encoding/json"
	"log"
)

//  GetProcessInstanceId
//  @Description: 取得流程ID
//  @receiver b
//  @return []byte
func (b *BPMClient) GetProcessInstanceId() []byte {
	uri := b.apiUri + "bpm/process?c=100"
	log.Println("GetProcessInstanceId()- uri", uri)

	//Http request
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("GetProcessInstanceId() - Status Code:", resp.StatusCode())
	log.Printf("GetProcessInstanceId() - resp:\n %+v", string(resp.Body()))

	return resp.Body()
}

// Create a new caseID(process instance)
// ref: https://documentation.bonitasoft.com/bonita/2021.2/api/bpm-api#start-a-process-using-an-instantiation-contract
//
//  CreateProcessCase
//  @Description: 啟單並取得caseID(process instance)
//  @receiver b
//  @param processId [表單ID]
//  @param jsonBody 只需要提供"內層"結構(轉成string)
//  @return caseId
//
func (b *BPMClient) CreateProcessCase(processId string, jsonBody ...string) (caseId []byte) {
	var s string
	//部分流程不需要在啟單時輸入資料，以下自動設定格式
	if len(jsonBody) != 0 {
		//將資料進行包裝成API要求的格式
		s = ModelInputStringToRawJson(jsonBody[0])
	} else {
		//該流程無需輸入相關資料
		s = "{}"
	}
	log.Println("CreateProcessCase() - StringToRawJson(): ", s)

	uri := b.apiUri + "bpm/process/" + processId + "/instantiation"
	log.Println("CreateProcessCase()- uri", uri)

	//Http request
	resp, err := b.request.SetBody(s).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("CreateProcessCase() - Status Code:", resp.StatusCode())
	log.Printf("CreateProcessCase() - resp:\n %+v", resp.Body())

	return resp.Body()
}

//包裝流程資料格式
func ModelInputStringToRawJson(s string) string {
	s1 := []byte(s)
	var s2 *interface{}

	//json->struct
	err := json.Unmarshal(s1, &s2)
	if err != nil {
		log.Print(err)
	}
	log.Printf("StringToRawJson() - json.Unmarshal(s1, &s2) %+v: ", s2)

	//外層加上modelInput結構，此為流程存取資料所用之格式
	s3 := &FormInput{
		ModelInput: s2,
	}

	//struct->json
	s4, err := json.Marshal(s3)
	if err != nil {
		log.Print(err)
	}

	return string(s4)
}

//  GetStateCaseList
//  @Description: 取得人員有關的任務列表
//  @receiver b
//  @param rows 顯示資料量
//  @param state  |ready|
//  @param userId
//  @return string
func (b *BPMClient) GetStateCaseList(rows string, state string, userId string) []byte {

	uri := b.apiUri + "bpm/humanTask?c=" + rows + "&f=state=" + state + "&f=user_id=" + userId
	log.Println("GetStateCaseList()- uri", uri)

	//Http request
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("GetStateCaseList() - Status Code:", resp.StatusCode())
	log.Printf("CreateProcessCase() - resp:\n %+v", string(resp.Body()))

	return resp.Body()
}

//ExecuteTask
//@Description:審核任務
//@receiver b
//@param taskId
//@param jsonBody 只需要提供"內層"結構(轉成string)
//@return ResponseStatusCode (204為成功)
func (b *BPMClient) ExecuteTask(taskId string, jsonBody ...string) int {

	var s string
	if len(jsonBody) != 0 {
		//該任務需要輸入相關資料，包裝Json格式
		s = ModelInputStringToRawJson(jsonBody[0])
	} else {
		//該任務無須輸入資料。
		s = "{}"
	}

	log.Println("ExecuteTask() - StringToRawJson(): ", s)

	uri := b.apiUri + "bpm/userTask/" + taskId + "/execution?assign=true"
	log.Println("ExecuteTask()- uri", uri)

	//Http request
	resp, err := b.request.SetBody(s).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("ExecuteTask() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//  GetStateCaseList
//  @Description: 顯示該單待執行任務詳細資料
//  @receiver b
//  @param caseId
//  @return []byte
func (b *BPMClient) GetCasePendingTaskDetail(caseId string) []byte {

	uri := b.apiUri + "bpm/humanTask?f=caseId=" + caseId
	log.Println("GetCasePendingTaskDetail() -uri", uri)

	//Http request
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("GetCasePendingTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

// GetCaseArchivedTaskDetail
// @Description: 顯示該單已完成任務詳細資料
// @receiver b
// @parm caseId
// @return string
func (b *BPMClient) GetCaseArchivedTaskDetail(caseId string) []byte {

	uri := b.apiUri + "bpm/archivedTask?f=caseId=" + caseId
	log.Println("GetCaseArchivedTaskDetail() -uri", uri)

	//Http request
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("GetCaseArchivedTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

// GetCaseArchivedTaskDetail
// @Description: 顯示該任務完成後之詳細資料
// @receiver b
// @parm caseId
// @return string
func (b *BPMClient) GetArchivedTaskDetail(sourceObjectId string) []byte {

	uri := b.apiUri + "bpm/archivedHumanTask?f=sourceObjectId=" + sourceObjectId
	log.Println("GetArchivedTaskDetail() -uri", uri)
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}
	log.Printf("GetArchivedTaskDetail() - b.request:\n %+v", b.request)
	log.Println("GetArchivedTaskDetail() - Status Code:", resp.StatusCode())
	return resp.Body()
}

//GetProcessAllCaseList
// @Desctiption: 取得該流程所有單況
// @receiver b
// @parm caseId
// @return string
func (b *BPMClient) GetProcessAllCaseList(rows string, processId string) []byte {

	uri := b.apiUri + "bpm/case?c=" + rows + "&f=processDefinitionId=" + processId
	log.Println("GetCaseArchivedTaskDetail() -uri", uri)

	//Http request
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("GetCaseArchivedTaskDetail() - Status Code:", resp.StatusCode())

	return resp.Body()
}

//UpdateAssignedId
//@Description:更改受指派者（任務轉移）
//@receiver b
//@param taskId
//@param userId
//@return ResponseStatusCode (200為成功)
func (b *BPMClient) UpdateAssignedId(userId string, taskId string) int {

	//轉移給指定人員
	jsonBody := `{"assigned_id" : "` + userId + `"}`

	log.Println("UpdateAssignedId() - StringToRawJson(): ", jsonBody)

	uri := b.apiUri + "bpm/humanTask/" + taskId
	log.Println("UpdateAssignedId()- uri", uri)

	//Http request
	resp, err := b.request.SetBody(jsonBody).Put(uri)
	if err != nil || resp.StatusCode() == 403 {
		//若該任務已存在指派指定人員，須先將該人員移除，再進行指派
		resp, err = b.request.SetBody(`{"assigned_id":""}`).Put(uri)
		//指派指定人員
		resp, err = b.request.SetBody(jsonBody).Put(uri)
	}

	log.Println("UpdateAssignedId() - body: ", jsonBody)
	log.Printf("UpdateAssignedId() - b.request:\n %+v", b.request)
	log.Println("UpdateAssignedId() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}
