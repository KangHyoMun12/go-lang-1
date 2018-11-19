package model

/*
ApricotBlossom get constructor model
*/
var ApricotBlossom = Apricotblossom{}

/*
Apricotblossom type struct
*/
type Apricotblossom struct {
	ID      int
	Price   float64
	Imageid int
	Descibe string
}

/*
NewApricotblossom type func to return struct
*/
func NewApricotblossom(id int, price float64, imageid int, descibe string) Apricotblossom {
	return Apricotblossom{id, price, imageid, descibe}
}
