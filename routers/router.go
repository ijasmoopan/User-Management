package routers

import (
	"fmt"
	"net/http"

	a "github.com/ijasmoopan/usermanagement/admin/adminHandlers"
	am "github.com/ijasmoopan/usermanagement/admin/adminMiddleware"
	m "github.com/ijasmoopan/usermanagement/middleware"
	h "github.com/ijasmoopan/usermanagement/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/jwtauth/v5"
)



func Router() {
	router := chi.NewRouter()
	
	router.Use(middleware.Logger)

	routerAdmin := router.Group(nil)
	// routerAdmin.Use(am.HaveToken)

	// ------------Admin---------------

	routerAdmin.Get("/adminlogin", a.AdminLogin)
	routerAdmin.Post("/adminvalidation", a.AdminLoginValidation)

	// router.With(am.IsAuthorized).Get("/userlist", a.AccessingUserList)
	// router.Get("/userlist", a.AccessingUserList)

	subRouter := routerAdmin.Group(nil)
	subRouter.Use(am.IsAuthorized)
	subRouter.Get("/userlist", a.AccessingUserList)

	subRouter.Post("/userlist/edit/{id}", a.EditHandler)
	subRouter.Post("/userlist/edit/{id}/editing", a.EditingUserHandler)

	subRouter.Post("/userlist/status/{id}-{status}", a.StatusHandler)

	subRouter.Post("/userlist/delete/{id}", a.DeleteHandler)

	subRouter.Post("/userlist/create", a.CreateHandler)
	subRouter.Post("/userlist/creating", a.CreateUserHandler)

	subRouter.Post("/userlist/search", a.SearchHandler)

	router.Post("/adminlogout", a.AdminLogOut)
	// router.With(am.DeleteToken, am.NoCache).Post("/adminlogout", a.AdminLogOut)

	// -----------User------------------

	routerUser := router.Group(nil)
	// routerUser.Use(m.HaveToken)
	routerUser.Use(m.NoCache)

	routerUser.Get("/signup", h.SignUpIndex)
	routerUser.Post("/signupvalidation", h.SignUpHandler)

	// routerUser.With(m.DeleteToken).Get("/login", h.LogInIndex)
	routerUser.Get("/login", h.LogInIndex)
	// routerUser.With(m.DeleteToken).Post("/loginvalidation", h.LogInHandler)
	routerUser.Post("/loginvalidation", h.LogInHandler)

	// routerUser.With(m.DeleteToken, m.NoCache).Post("/logout", h.LogOutHandler)
	routerUser.Get("/logout", h.LogOutHandler)

	// routerUser.With(jwtauth.Verifier(tokenAuth), jwtauth.Authenticator).Get("/home/{username}", h.HomePage)
	routerUser.With(m.IsAuthorized).Get("/home/{username}", h.HomePage)

	// subRouter := routerUser.Group(nil)
	// subRouter.Use(m.IsAuthorized)
	// subRouter.Get("/home/{username}", h.HomePage)

	fmt.Println("Starting server at port: 8080")
	http.ListenAndServe(":8080", router)
}