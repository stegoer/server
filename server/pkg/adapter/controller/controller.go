package controller

// Controller holds the controllers across the entire application.
type Controller struct {
	User  interface{ User }
	Image interface{ Image }
}
