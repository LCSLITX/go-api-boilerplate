package structs

type Echo struct {
	Echo string
}

type Erro struct {
	Err error
}


// This struct represent an individual row in a MySQL table.
type DBItem struct {
	ColumnA int
	ColumnB string
}
