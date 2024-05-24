package routes

import (
	//"net/http"

	"net/http"

	"github.com/gorilla/mux"
	//"github.com/lukaslinardi/fullstack_engineer_sprint_asia/domain/model/general"
	api "github.com/lukaslinardi/xyz_multifinance_api/handler"
)

func getV1(router, routerJWT *mux.Router, handler api.Handler) {
	router.HandleFunc("/v1/signup", handler.Auth.Auth.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/v1/login", handler.Auth.Auth.Login).Methods(http.MethodPost)
	// router.HandleFunc("/v1/psef/forget-password", handler.Auth.Auth.ForgetPassword).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/outlet", handler.User.User.GetOutletList).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/outlet-detail", handler.User.User.GetOutletDetail).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/product", handler.User.User.GetProductList).Methods(http.MethodGet)
	// routerJWT.HandleFunc("/v1/psef/product-detail", handler.User.User.GetProductDetail).Methods(http.MethodGet)
}
