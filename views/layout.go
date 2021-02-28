package views

type Layout struct {
	Title string
	// might be: index, blacklist
	Page    string
	Access  int
	Content interface{}
}
