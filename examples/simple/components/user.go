package components

// User 用户
// @component
type User struct {
    ID         int64
    Name       string
    Collection *Collection
}

// NewUser 新建用户
// @constructor
func NewUser() {

}

// UserService
// @component
type UserService interface {
    Get()
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
