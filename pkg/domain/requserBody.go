package domain

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	Id          int64   `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string  `json:"name" gorm:"not null;unique" binding:"required"`
	Description string  `json:"description"  binding:"required,min=15"`
	Price       float32 `json:"price" gorm:"not null;unique" binding:"required"`
	Stock       int64   `json:"stock" gorm:"not null;unique" binding:"required"`
}

type AddProduct struct {
	ProductId int64 `json:"productid"`
	Quantity  int   `json:"quantity"`
}
