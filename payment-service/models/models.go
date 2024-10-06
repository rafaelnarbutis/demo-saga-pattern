package models

type Payment struct {
	Price    float32
	Notebook Notebook
	Address  Address
}

type Notebook struct {
	Memory    int16
	Cpu       int16
	Hd        int16
	ScreeSize int16
}

type Address struct {
	Street     string
	Number     string
	Country    string
	Complement string
}
