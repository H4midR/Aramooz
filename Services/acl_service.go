// file: services/movie_service.go
/*eslint-disable */
package services

type AclVal struct {
	Value   uint64 `json:"value"`
	PValue  int32  `json:"pvalue"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
}
type AclService interface {
	Acl() *aclService
	Allow() bool
	Privileges()
}
type aclService struct {
	ACLDef map[string]AclVal
}

var ACLList = map[string]AclVal{

	"User": AclVal{
		Value:   0x1,
		Title:   "کاربر تایید شده",
		Comment: "کاربر تایید شده ، ثبت نام کرده",
	},

	"Admin": AclVal{
		Value:   0x2,
		Title:   "مدیر",
		Comment: "مدیر",
	},
}

//---------------------------------------------------------------------------------
// acl coding 0xXXYYY : XX represent to the first array and YYY represent to the second array
func Acl() *aclService {
	myacl := aclService{ACLDef: ACLList}
	return &myacl

}

//SuperAdminACL := 0xfffffffff
//DefaultUser := 0xf00f1d7f
func (a *aclService) Allow(userAcl uint64, acl AclVal) bool {
	return userAcl&acl.Value != 0
}
func (a *aclService) Privileges(Useracl uint64) map[string]AclVal {
	//ctx.Writef("%#x", Useracl)
	res := make(map[string]AclVal)
	for pname, pel := range a.ACLDef {
		res[pname] = AclVal{}
		if a.Allow(Useracl, pel) {
			res[pname] = pel
		}
	}
	return res
}
