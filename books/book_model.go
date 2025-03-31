package books

type Book struct {
	Id     int
	Name   string
	Price  int
	Auther string
}

func GetAllBooks() []Book {
	b1 := Book{
		Id:     1,
		Name:   "The Lord of Rings",
		Price:  200,
		Auther: "JRR Tolkien",
	}
	b2 := Book{
		Id:     2,
		Name:   "The Silmarillion",
		Price:  500,
		Auther: "JRR Tolkien",
	}

	return []Book{
		b1,
		b2,
	}

}
