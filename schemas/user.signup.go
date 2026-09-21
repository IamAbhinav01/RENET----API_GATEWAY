package schemas

type User struct {
	ID       int
	Name     string
	Email    string `gorm:"unique;notNull"`
	Password string
}