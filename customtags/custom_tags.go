package customtags

type User struct {
	Username string `validate:"required,minlen=3,maxlen=20" display:"userName"`
	Email    string `validate:"required,email" display:"email"`
	Age      int    `validate:"min=18,max=100" display:"age"`
	Website  string `validate:"url" display:"webURL"`
	Bio      string `validate:"maxlen=500" display:"biography"`
}
