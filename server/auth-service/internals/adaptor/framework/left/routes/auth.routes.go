package routes

import "github.com/go-chi/chi/v5"

func (ra *Adaptor) AuthRoutes(r chi.Router) {
	r.Get("/getUser/{id}", ra.api.GetUserApi)
	r.Post("/loginUser", ra.api.LoginUserApi)
	r.Post("/registerUser", ra.api.RegisterUserApi)
	r.Patch("/updateUser", ra.api.UpdateUserApi)
	r.Delete("/deleteUser/{id}", ra.api.DeleteUserApi)
}
