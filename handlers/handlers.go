package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GeorgeHN666/proyect-Onlife-Server/middleware"
	"github.com/GeorgeHN666/proyect-Onlife-Server/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handler() {

	// Here we Store the Function Of Mux New Router Into A Variable, That Its Gonna Be Reusable
	route := mux.NewRouter()

	// SignUp - Login section
	route.HandleFunc("/signup", routers.SignUpEndPoint).Methods("POST")
	route.HandleFunc("/login", routers.LoginEndPoint).Methods("POST")
	// End Of Section

	// Profile Section
	route.HandleFunc("/sP", middleware.CheckDB(middleware.ValidToken(routers.SeeProfileEndPoint))).Methods("GET")
	route.HandleFunc("/mP", middleware.CheckDB(middleware.ValidToken(routers.ModifyUserEndPoint))).Methods("PUT")
	// End Of Profile Section

	// Posts Sections
	route.HandleFunc("/nP", middleware.CheckDB(middleware.ValidToken(routers.CreatePostEndPoint))).Methods("POST")
	route.HandleFunc("/rP", middleware.CheckDB(middleware.ValidToken(routers.ReadOwnPostEndPoint))).Methods("GET")
	route.HandleFunc("/dP", middleware.CheckDB(middleware.ValidToken(routers.DeletePostEndPoint))).Methods("DELETE")
	// End Of Post Section

	// Likes Section
	route.HandleFunc("/l", middleware.CheckDB(middleware.ValidToken(routers.LikeEndPoint))).Methods("GET")
	route.HandleFunc("/cL", middleware.CheckDB(middleware.ValidToken(routers.CheckLike))).Methods("GET")
	route.HandleFunc("/dL", middleware.CheckDB(middleware.ValidToken(routers.DeleteLikeEndpoint))).Methods("DELETE")
	// End of section

	// Upload AVATAR/Banner section
	route.HandleFunc("/uA", middleware.CheckDB(middleware.ValidToken(routers.UploadAvatarEndPoint))).Methods("POST")
	route.HandleFunc("/oA", middleware.CheckDB(routers.ObtainAvatarEndpoint)).Methods("GET")
	route.HandleFunc("/uB", middleware.CheckDB(middleware.ValidToken(routers.UploadBannerEndPoint))).Methods("POST")
	route.HandleFunc("/oB", middleware.CheckDB(routers.ObtainBannerEndpoint)).Methods("GET")
	// End Of Section

	// Comment Section
	route.HandleFunc("/pC", middleware.CheckDB(middleware.ValidToken(routers.PostComment))).Methods("POST")
	route.HandleFunc("/sC", middleware.CheckDB(middleware.ValidToken(routers.ReadCommentsEndPoint))).Methods("GET")
	route.HandleFunc("/dC", middleware.CheckDB(middleware.ValidToken(routers.DeleteComment))).Methods("DELETE")
	// End Of Section

	// Follow Section
	route.HandleFunc("/f", middleware.CheckDB(middleware.ValidToken(routers.FollowEndPoint))).Methods("GET")
	route.HandleFunc("/cF", middleware.CheckDB(middleware.ValidToken(routers.CheckFollowEndPoint))).Methods("GET")
	route.HandleFunc("/dF", middleware.CheckDB(middleware.ValidToken(routers.DeleteFollowEndPoint))).Methods("DELETE")
	// End of section

	// All Users Secction
	route.HandleFunc("/sA", middleware.CheckDB(middleware.ValidToken(routers.SeeAllUsers))).Methods("GET")
	route.HandleFunc("/rAP", middleware.CheckDB(middleware.ValidToken(routers.SeeRelatedPost))).Methods("GET")

	// Here we Create a Variable Called PORT That its Gonna take the PORT of our local env
	PORT := os.Getenv("PORT")

	// If The PORT doesn't exist we set the port

	if PORT == "" {
		PORT = "8080"
	}

	// Here we Create The Permissions That Cors Allow us, and pass the server mux
	permissions := cors.AllowAll().Handler(route)

	// Finally we do the Listen and serve
	log.Fatal(http.ListenAndServe(":"+PORT, permissions))

}
