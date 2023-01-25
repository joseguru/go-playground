import (
	"github.com/jinzhu/gorm"
    "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// Open a connection to the MySQL database
	db, err := gorm.Open("mysql", "root:geranimo@/testgo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Automatically migrate the schema
	db.AutoMigrate(&User{})

	// Create a new user
	user := User{Name: "John Doe", Email: "johndoe@example.com"}

	// Save the user to the database
	db.Create(&user)
}
  