package model

/*
Phone struct type
*/
type Phone struct {
	ID        int
	Name      string
	Price     float64
	Companyid Company
	Type      string
}

/*
Company struct type
*/
type Company struct {
	ID   int
	Name string
}
