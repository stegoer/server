package controller

// Controller holds the User and Image controllers.
type Controller struct {
	User  interface{ User }
	Image interface{ Image }
}
