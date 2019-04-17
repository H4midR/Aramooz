package dataModels

import "golang.org/x/crypto/bcrypt"

//User model
type User struct {
	Uid      string `json:"uid,omitempty" form:"Uid"`
	Name     string `json:"name,omitempty" form:"Name"`
	Mobile   string `json:"mobile,omitempty" form:"Mobile"`
	Token    string `json:"token,omitempty" form:"Token"`
	Password string `json:"password,omitempty" form:"Password"`
	ACL      string `json:"acl,omitempty" form:"Acl"`
	Kind     string `json:"kind,omitempty"`
	Email    string `json:"email,omitempty" form:"Email"`
	Address  string `json:"address,omitempty" form:"Address"`
	Bio      string `json:"bio,omitempty" form:"Bio"`
	Exam     []Exam `json:"exam,omitempty"`
}

const UserType = "User"

// GeneratePassword : will generate a hashed password for us based on the
// user's input.
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//CompairPassword compair the given password and hassed password
func CompairPassword(DbHPass string, Password string) error {
	return bcrypt.CompareHashAndPassword([]byte(DbHPass), []byte(Password))
}

// ValidatePassword : will check if passwords are matched.
func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}
