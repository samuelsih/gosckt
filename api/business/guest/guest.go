package guest

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kokizzu/gotro/S"
	"github.com/samuelsih/gosckt/api/business"
	"github.com/samuelsih/gosckt/model"
)

// Dependency
type Guest struct {
	Psql *pgxpool.Pool
	Rdb  *redis.Client
}

type Guest_LoginIn struct {
	business.CommonRequest
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Guest_LoginOut struct {
	business.CommonResponse
	User *model.User `json:"user,omitempty"`
}

func (g *Guest) Login(in *Guest_LoginIn) (out Guest_LoginOut) {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	
	if len(in.Email) < 5 {
		out.SetError(http.StatusBadRequest, "invalid email")
		return
	}

	user := model.NewUser(g.Psql)

	if !user.FindByEmail(in.Email) {
		out.SetError(http.StatusNotFound, "user not found")
		return
	}

	if !user.PasswordMatch(in.Password) {
		out.SetError(http.StatusBadRequest, "Email or password doesn't match")
		return
	}

	var buffer bytes.Buffer

	if err := gob.NewEncoder(&buffer).Encode(user); err != nil {
		out.SetError(http.StatusInternalServerError, "cant gob user")
		return
	}

	token := S.RandomCB63(32)

	err := g.Rdb.Set(ctx, token, buffer.Bytes(), 25 * time.Second).Err()
	if err != nil {
		log.Println(err)
		out.SetError(http.StatusInternalServerError, "cant create session for this user")
		return
	}

	out.SetCookie = token
	out.StatusMsg = "success"
	out.StatusCode = http.StatusOK

	out.User = user.Clean()

	return
}

type Guest_RegisterIn struct {
	business.CommonRequest
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Guest_RegisterOut struct {
	business.CommonResponse
	User *model.User `json:"user,omitempty"`
}

func (g *Guest) Register(in *Guest_RegisterIn) (out Guest_RegisterOut) {
	err := ValidateSignIn(*in)
	if err != nil {
		out.SetError(http.StatusBadRequest, err.Error())
		return
	}

	user := model.NewUser(g.Psql)

	if user.HasEmail(in.Email) {
		out.SetError(http.StatusBadRequest, "email already used")
		return
	}

	user.Email = in.Email
	user.Name = in.Name
	user.SetPassword(in.Password)

	if !user.Insert() {
		out.SetError(http.StatusInternalServerError, `failed to create this user`)
		log.Println(err)
		return
	}

	out.StatusCode = http.StatusOK
	out.StatusMsg = "success"

	out.User = user.Clean()

	return
}
