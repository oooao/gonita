package gonita

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBPMClient_AddGroup(t *testing.T) {
	b := New("")
	type args struct {
		bodyInput string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				bodyInput: `{"name": "", "displayname": "", "parent_group_id":""}`,
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.AddGroup(tt.args.bodyInput); !bytes.Equal(got, tt.want) {
				t.Errorf("BPMClient.AddGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_EditGroup(t *testing.T) {
	b := New("")
	type args struct {
		bodyInput string
		groupId   string
	}
	tests := []struct {
		name string
		b    *BPMClient
		args args
		want int
	}{
		{
			args: args{
				bodyInput: `{"name": "", "displayname": "", "parent_group_id":"","country":""}`,
				groupId:   "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.EditGroup(tt.args.bodyInput, tt.args.groupId); got != tt.want {
				t.Errorf("BPMClient.EditGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_DeleteGroup(t *testing.T) {
	b := New("")
	type args struct {
		groupId string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				groupId: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.DeleteGroup(tt.args.groupId); got != tt.want {
				t.Errorf("BPMClient.DeleteGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_AddProfileMember(t *testing.T) {
	b := New("")

	type args struct {
		profile_id string
		id         string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				profile_id: "",
				id:         "",
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.AddProfileMember(tt.args.profile_id, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPMClient.AddprofileMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_AddMembership(t *testing.T) {
	b := New("")

	type args struct {
		bodyinput string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				bodyinput: `{"user_id": "", "group_id": "", "role_id":""}`,
			},
			want: []byte(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.AddMembership(tt.args.bodyinput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPMClient.AddMembership() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_AddProfessionalContactData(t *testing.T) {
	b := New("")

	type args struct {
		id       string
		username string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				id:       "",
				username: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.AddProfessionalContactData(tt.args.id, tt.args.username); got != tt.want {
				t.Errorf("BPMClient.AddProfessionalContactData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_EditUser(t *testing.T) {
	b := New("")

	type args struct {
		userId    string
		bodyinput string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				userId:    "",
				bodyinput: `{"id":"","userName": "}`,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.EditUser(tt.args.userId, tt.args.bodyinput); got != tt.want {
				t.Errorf("BPMClient.EditUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_DeleteMembership(t *testing.T) {
	b := New("")

	type args struct {
		userId  string
		groupId string
		roleId  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				userId:  "",
				groupId: "",
				roleId:  "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.DeleteMembership(tt.args.userId, tt.args.groupId, tt.args.roleId); got != tt.want {
				t.Errorf("BPMClient.DeleteMembership() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_DeleteUser(t *testing.T) {
	b := New("")

	type args struct {
		userId string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				userId: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.DeleteUser(tt.args.userId); got != tt.want {
				t.Errorf("BPMClient.DeleteUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBPMClient_AddUser(t *testing.T) {
	b := New("")
	type args struct {
		bodyInput string
	}
	tests := []struct {
		name  string
		b     *BPMClient
		args  args
		want  int
		want1 bool
	}{
		{
			args: args{
				bodyInput: `{"password":"","userName":"","enabled":""}`,
			},
			want:  10,
			want1: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := b.AddUser(tt.args.bodyInput)
			if got != tt.want {
				t.Errorf("BPMClient.AddUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BPMClient.AddUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBPMClient_GetUserMembership(t *testing.T) {
	b := New("")
	type args struct {
		userId string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				userId: "",
			},
			want: []byte(""),
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.GetUserMembership(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BPMClient.GetUserMembership() = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestBPMClient_EditUserMembership(t *testing.T) {
	b := New("")
	type args struct {
		jsonBody string
		group_id string
		role_id  string
	}
	tests := []struct {
		name string
		b    *BPMClient
		args args
		want bool
	}{
		{
			args: args{
				group_id: "",
				role_id:  "",
				jsonBody: `{"user_id":"",
				"group_id":"", 
				"role_id":"",
				"tm":["",""],
				"pm":""}`,
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := b.EditUserMembership(tt.args.group_id, tt.args.role_id, tt.args.jsonBody); got != tt.want {
				t.Errorf("BPMClient.EditUserMembership() = %v, want %v", got, tt.want)
			}
		})
	}
}
