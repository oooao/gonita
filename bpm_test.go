package gonita

import (
	"bytes"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestBPMClient_CreateProcessCase(t *testing.T) {

	b := New("")

	formInput1 := `{
		"account":"",
		"pm":"", 
		"tm":["",""] 
	}`

	type fields struct {
		serverUri  string
		apiUri     string
		username   string
		password   string
		request    *resty.Request
		token      string
		jSessionId string
	}
	type args struct {
		processId string
		jsonBody  string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCaseId []byte
	}{
		{
			name: "",
			fields: fields{
				serverUri:  b.serverUri,
				apiUri:     b.apiUri,
				username:   b.username,
				password:   b.password,
				request:    b.request,
				token:      b.token,
				jSessionId: b.jSessionId,
			},
			args: args{
				processId: "",
				jsonBody:  formInput1,
			},
			wantCaseId: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BPMClient{
				serverUri:  tt.fields.serverUri,
				apiUri:     tt.fields.apiUri,
				username:   tt.fields.username,
				password:   tt.fields.password,
				request:    tt.fields.request,
				token:      tt.fields.token,
				jSessionId: tt.fields.jSessionId,
			}
			if gotCaseId := b.CreateProcessCase(tt.args.processId, tt.args.jsonBody); string(gotCaseId) != string(tt.wantCaseId) {
				t.Errorf("CreateProcessCase() = %v, want %v", gotCaseId, tt.wantCaseId)
			}
		})
	}
}

func TestModelInputStringToRawJson(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `{"pm": , "tm": []} TO {"modelInput":{"pm":,"tm":[]}}`,
			args: args{
				s: `{"pm": , "tm": []}`,
			},
			want: `{"modelInput":{"pm":,"tm":[]}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModelInputStringToRawJson(tt.args.s); got != tt.want {
				t.Errorf("StringToRawJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetStateCaseList(t *testing.T) {

	b := New("")
	

	type fields struct {
		serverUri  string
		apiUri     string
		username   string
		password   string
		request    *resty.Request
		token      string
		jSessionId string
	}
	type args struct {
		rows   string
		state  string
		userId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		{
			name: `GetReadyCase("","","")`,
			fields: fields{
				serverUri:  b.serverUri,
				apiUri:     b.apiUri,
				username:   b.username,
				password:   b.password,
				request:    b.request,
				token:      b.token,
				jSessionId: b.jSessionId,
			},
			args: args{
				rows:   "",
				state:  "",
				userId: "",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BPMClient{
				serverUri:  tt.fields.serverUri,
				apiUri:     tt.fields.apiUri,
				username:   tt.fields.username,
				password:   tt.fields.password,
				request:    tt.fields.request,
				token:      tt.fields.token,
				jSessionId: tt.fields.jSessionId,
			}
			if got := b.GetStateCaseList(tt.args.rows, tt.args.state, tt.args.userId); !bytes.Equal(got, tt.want) {
				t.Errorf("GetStateCaseList() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestBPMClient_ExecuteTask(t *testing.T) {

	//身分驗證
	b := New("")
	//建立該函式參數結構
	type args struct {
		taskId   string
		jsonBody string
	}
	//建立測試模型
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: `ExecuteTask`,
			args: args{
				taskId:   "",
				jsonBody: `{"dstaff":[,],"gmapprovalstatus":true}`,
			},
			want: 0,
		},
	}
	//執行測試
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.ExecuteTask(tt.args.taskId, tt.args.jsonBody); got != tt.want {
				t.Errorf("ExecuteTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetCasePendingTaskDetail(t *testing.T) {
	b := New("")
	type args struct {
		caseId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				caseId: "",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetCasePendingTaskDetail(tt.args.caseId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetCasePendingTaskDetail() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestBPMClient_GetCaseArchivedTaskDetail(t *testing.T) {
	b := New("")
	type args struct {
		caseId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				caseId: "",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetCaseArchivedTaskDetail(tt.args.caseId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetCaseArchivedTaskDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetArchivedTaskDetail(t *testing.T) {
	b := New("")
	type args struct {
		sourceObjectId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				sourceObjectId: "",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetArchivedTaskDetail(tt.args.sourceObjectId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetArchivedTaskDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_GetProcessAllCaseList(t *testing.T) {

	b := New("")
	type args struct {
		rows      string
		processId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				rows:      "",
				processId: "",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetProcessAllCaseList(tt.args.rows, tt.args.processId); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.GetProcessAllCaseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_UpdateAssignedId(t *testing.T) {
	b := New("")
	type args struct {
		userId string
		taskId string
	}
	tests := []struct {
		name string
		b    *BPMClient
		args args
		want int
	}{
		{
			args: args{
				userId: "",
				taskId: "",
			},
			want: 0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.UpdateAssignedId(tt.args.userId, tt.args.taskId); got != tt.want {
				t.Errorf("BPMClient.UpdateAssignedId() = %v, want %v", got, tt.want)
			}
		})
	}
}
