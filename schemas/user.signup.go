package schemas

type SignUp struct {
	ID       int
	Name     string
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" validate:"required,email"`
	Password string `gorm:"type:varchar(255);not null" validate:"required,min=8,max=72"`
}