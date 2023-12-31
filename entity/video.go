package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=200" validate:"is-cool"`
	Description string `json:"description" binding:"max=200"`
	Url         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

type Person struct {
	FirstName string `json:"firstname" binding:"min=2,max=10"`
	LastName  string `json:"lastname" binding:"min=2,max=10"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}
