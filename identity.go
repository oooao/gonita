package gonita

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

//AddGroup
//  @Description: 新增部門
//  @receiver b
//  @param bodyInput
//  @return []byte
func (b *BPMClient) AddGroup(bodyInput string) []byte {

	uri := b.apiUri + "identity/group"
	log.Println("AddGroup() -uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("AddGroup() - Status Code:", resp.StatusCode())
	return resp.Body()
}

//EditGroup
//  @Description: 更新部門
//  @receiver b
//  @param bodyInput
//  @param groupId
//  @return int (200為成功)
func (b *BPMClient) EditGroup(bodyInput string, groupId string) int {

	//s := StringToRawJson(bodyInput)
	uri := b.apiUri + "identity/group/" + groupId
	log.Println("EditGroup() -uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Put(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("EditGroup() - Status Code:", resp.StatusCode())
	return resp.StatusCode()
}

//DeleteGroup
//  @Description: 刪除部門
//  @receiver b
//  @param groupId
//  @return int (200為成功)
func (b *BPMClient) DeleteGroup(groupId string) int {

	uri := b.apiUri + "identity/group/" + groupId
	log.Println("DeleteGroup() -uri", uri)

	//Http request
	resp, err := b.request.Delete(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("DeleteGroup() - Status Code:", resp.StatusCode())
	return resp.StatusCode()
}

//AddUser
//  @Description: 新增人員
//  @receiver b
//  @param bodyInput
//  @return bool
func (b *BPMClient) AddUser(bodyInput string) (int, bool) {

	uri := b.apiUri + "identity/user"
	log.Println("AddUser() -uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("AddUser() - Status Code:", resp.StatusCode())
	if resp.StatusCode() == 200 {
		var s2 map[string]interface{}

		//json->struct
		err2 := json.Unmarshal(resp.Body(), &s2)
		if err2 != nil {
			log.Print(err2)
		}

		//取得該新增人員之userId
		id := fmt.Sprintf("%v", s2["id"])

		//string->int
		f, err3 := strconv.Atoi(id)
		if err3 != nil {
			log.Print(err3)
		}

		//取得該新增人員之userName
		str2 := fmt.Sprintf("%v", s2["userName"])

		//新增人員不代表新增使用者，因此需同時給予登入一般使用者與管理員之權限
		errProfileMember := b.AddProfileMember("1", id)
		errProfileManager := b.AddProfileMember("2", id)

		//新增該人員信箱
		errProfessionalContactData := b.AddProfessionalContactData(id, str2)

		if errProfileMember == 200 && errProfileManager == 200 && errProfessionalContactData == 200 {
			//f為userId
			return f, true
		}
	}
	return 0, false
}

//AddProfileMember
//  @Description: 設定為後台管理員或使用者
//  @receiver b
//  @param bodyInput
//  @return []byte
func (b *BPMClient) AddProfileMember(profile_id string, id string) int {

	//管理員id為2 一般使用者為1
	bodyInput := `{"profile_id":"` + profile_id + `", "member_type":"USER", "user_id":"` + id + `"}`
	uri := b.apiUri + "portal/profileMember"
	log.Println("AddProfileMember() - uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("AddProfileMember() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//AddMembership
//  @Description: 設定隸屬部門
//  @receiver b
//  @param bodyInput
//  @return bool
func (b *BPMClient) AddMembership(bodyInput string) bool {

	uri := b.apiUri + "identity/membership"
	log.Println("AddMembership() - uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("AddMembership() - Status Code:", resp.StatusCode())

	return resp.StatusCode() == 200
}

//AddProfessionalContactData
//  @Description: 新增人員聯繫資訊
//  @receiver b
//  @param id
//  @param username
//  @return int
func (b *BPMClient) AddProfessionalContactData(id string, username string) int {

	bodyInput := `{"id":"` + id + `", "email":"` + username + `@hta.com.tw"}`

	uri := b.apiUri + "identity/professionalcontactdata"
	log.Println("AddProfessionalContactData() - uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Post(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("AddProfessionalContactData() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//EditUser
//  @Description: 編輯人員
//  @receiver b
//  @param bodyInput
//  @return int
func (b *BPMClient) EditUser(userId string, bodyInput string) int {

	uri := b.apiUri + "identity/user/" + userId
	log.Println("EditGroup() - uri", uri)

	//Http request
	resp, err := b.request.SetBody(bodyInput).Put(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("EditUser() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//DeleteMembership
//  @Description: 刪除人員隸屬部門
//  @receiver b
//  @param userId
//  @param groupId
//  @param roleId
//  @return int
func (b *BPMClient) DeleteMembership(userId string, groupId string, roleId string) int {

	uri := b.apiUri + "identity/membership/" + userId + "/" + groupId + "/" + roleId
	log.Println("DeleteMembership() - uri", uri)

	//Http request
	resp, err := b.request.Delete(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("DeleteMembership() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

//DeleteUser
//  @Description: 刪除人員
//  @receiver b
//  @param bodyInput
//  @return int
func (b *BPMClient) DeleteUser(userId string) int {
	uri := b.apiUri + "identity/user/" + userId
	log.Println("DeleteUser() - uri", uri)

	//Http request
	resp, err := b.request.Delete(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("DeleteUser() - Status Code:", resp.StatusCode())

	return resp.StatusCode()
}

// func RebuildJson(jsonBody string, key ...string) string {
// 	var s1 map[string]interface{}

// 	err := json.Unmarshal([]byte(jsonBody), &s1)
// 	if err != nil {
// 		log.Println("反序列化失败", err)
// 	}
// 	s2 := make(map[string]interface{})
// 	for _, keyy := range key {
// 		s2[keyy] = s1[keyy]
// 	}

// 	s3, err := json.Marshal(s2)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	return string(s3)
// }

//GetUserMembership
//  @Description: 查看人員隸屬部門角色
func (b *BPMClient) GetUserMembership(userId string) []byte {
	uri := b.apiUri + "identity/membership?c=25&f=user_id=" + userId
	log.Println("GetUserMembership() -uri", uri)

	//Http request
	resp, err := b.request.Get(uri)
	if err != nil {
		log.Print(err)
	}

	log.Println("GetUserMembership() - Status Code:", resp.StatusCode())
	return resp.Body()
}

//EditUserMembership
//  @Description: 編輯人員隸屬部門（bonita僅存在刪除與新增API)
//  @reciver b
//  @param jsonBody(欲修改之人員 部門 與職稱)
//  @return bool
func (b *BPMClient) EditUserMembership(group_id, role_id, jsonBody string) bool {
	var s1 map[string]interface{}

	//json->struct
	err := json.Unmarshal([]byte(jsonBody), &s1)
	if err != nil {
		log.Print(err)
	}

	//取得userID
	userId := fmt.Sprintf("%v", s1["user_id"])
	//刪除原角色
	err3 := b.DeleteMembership(userId, group_id, role_id)
	if err3 != 200 {
		log.Print("原角色刪除失敗")
	}
	//新增該人員之角色
	err4 := b.AddMembership(jsonBody)
	if err4 == false {
		return false
	}
	return true
}
