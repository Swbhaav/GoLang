package books

func GetBookNameList() []string {
	booknames := []string{}
	for _, value := range GetAllBooks() {
		booknames = append(booknames, value.Name)
	}
	return booknames

}
