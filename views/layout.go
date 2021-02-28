package views

type Layout struct {
	Title string
	// might be: index, blacklists
	Page    string
	Access  int
	Content interface{}
}
