package project_structs

type From_mal struct {
	Id       int
	Username string
}

type Username_and_id struct {
	Id       int
	Username string
	In_db    bool
}

type Genre_xpath struct {
	Name        string
	Count       int
	Description string
}
