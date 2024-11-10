package data

type Middleware struct {
	IsProtected  bool
	CSRFRequired bool
	HTTPMethod   string
}
