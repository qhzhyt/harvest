package components

import (
	"fmt"
	"github.com/qhzhyt/harvest/examples/simple/components/sub"
	"net"
)

// User 用户
// @component
type User struct {
	// @auto
	ID          int64       `json:"id"`
	Name, Label string      `1000`
	Collection  *Collection `json:"collection"`
	// @ip
	IP net.IP
	// @sub
	Sub sub0.Sub
	// @service
	Service UserService
	// @func
	HandlerFunc func(user *User) (*User, error)
	// @struct
	Struct struct {
		Name string
	}
	// @interface
	Handler interface {
		Run(user *User) (*User, error)
	}
}

type NameStruct struct {
	Name string
}

type NameStruct0 struct {
	Name string
}

// SuperUser
// @component
type SuperUser User

// NewUser 新建用户
// @constructor
func NewUser(uStruct struct{ Name string }) *User {
	var stringFunc func() string
	var nameStruct NameStruct
	var nameStruct0 NameStruct0
	nameStruct = uStruct
	nameStruct0 = NameStruct0(nameStruct)
	user := &User{
		Struct: nameStruct0,
	}
	stringFunc = user.GetName
	stringFunc()
	stringFunc = String
	stringFunc()
	return user
}

// UserService
type UserService interface {
	// @get
	Get() // get a user
}

// @component
type int0 int

// GetName
// @component
func (u User) GetName() string {
	return u.Name
}

// SetName
// @component
func (u *User) SetName(name string) {
	u.Name = name
}

// @component(name = userService)
var UserServiceImpl UserService

// @type
type UserHandlerFunc func(user *User) (*User, error)

// @method
func (i *int0) String() string {
	return fmt.Sprint(i)
}

// @func
func String() string {
	return "string"
}
