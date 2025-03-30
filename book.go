package main

type Book struct{
	Name string
	Auther  string
	Price int
}

func Fetchallbook()[]Book{

	return []Book{
			{Name: "DAA",Auther: "Pearson",Price: 500,},
		{Name: "DSA",Auther: "Pearson",Price: 5000,},

	}
}