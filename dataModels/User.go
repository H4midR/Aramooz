package dataModels

//User model
type User struct {
	Uid      string `json:"uid,omitempty" form:"Uid"`
	Name     string `json:"name,omitempty" form:"Name"`
	Mobile   string `json:"mobile,omitempty" form:"Mobile"`
	Token    string `json:"token,omitempty" form:"Token"`
	Password string `json:"password,omitempty" form:"Password"`
	Acl      string `json:"acl,omitempty" form:"Acl"`
	Kind     string `json:"kind,omitempty"`
}

const UserType = "User"
