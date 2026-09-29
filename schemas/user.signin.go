package schemas

type SignIn struct {
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null" validate:"required,email"`
	Password string `gorm:"type:varchar(255);not null;->:false;<-" validate:"required,min=8,max=72"`
}