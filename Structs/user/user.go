package user

import(
	"fmt"
	"errors"
)

type User struct {
	firstName string
	lastName string
	age int
}

type Admin struct {
	User
	email string
	password string
}

// New creates a new User instance with validation
// It returns an error if the input data is invalid.
func NewAdmin(email, password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "",
			lastName: "",
			age: 0,
		},
	}
}

func New(firstName string, lastName string, age int) (*User,error){
	if firstName == "" || lastName == "" || age <= 0 {
		return nil, errors.New("invalid user data")
	}
		return &User{	
				firstName: firstName,
				lastName: lastName,
				age: age,
			}, nil
}

func (u User) GetUserInfo (){
	fmt.Println("First Name:", u.firstName)
	fmt.Println("Last Name:", u.lastName)
	fmt.Println("Age:", u.age)
}

func (u *User) ClearUserInfo() {
	u.firstName = ""
	u.lastName = ""
	u.age = 0
	fmt.Println("User info cleared")
}