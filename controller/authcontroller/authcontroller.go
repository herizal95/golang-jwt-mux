package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/herizal95/golang-jwt-mux/config"
	"github.com/herizal95/golang-jwt-mux/helpers"
	"github.com/herizal95/golang-jwt-mux/models"
	"github.com/herizal95/golang-jwt-mux/services"
	"github.com/satori/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 400, "Passwords do not match", nil)
		return
	}

	passwordHash, err := services.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Uid:      uuid.NewV4(),
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Successfully to register user!", user)

}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User

	if err := config.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := services.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
	}

	token, err := services.GenerateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Successfully to Login", token)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
