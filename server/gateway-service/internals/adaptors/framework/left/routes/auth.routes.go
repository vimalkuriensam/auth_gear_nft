package routes

import "github.com/go-chi/chi/v5"

func (ra *Adaptor) AuthRoutes(r chi.Router) {
	r.Get("/getUser/{id}", ra.authApi.GetUserApi)
	r.Post("/loginUser", ra.authApi.LoginUserApi)
	r.Post("/registerUser", ra.authApi.RegisterUserApi)
	r.Patch("/updateUser", ra.authApi.UpdateUserApi)
	r.Delete("/deleteUser/{id}", ra.authApi.DeleteUserApi)
}
