package uranium

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/model/session"
	"github.com/go-uranium/uranium/model/sign_up"
	"github.com/go-uranium/uranium/model/user"
	"github.com/go-uranium/uranium/utils/hash"
	"github.com/go-uranium/uranium/utils/recaptcha"
	"github.com/go-uranium/uranium/utils/token"
	"github.com/go-uranium/uranium/utils/validate"
)

func (ushio *Ushio) HandleLogin(ctx *fiber.Ctx) error {
	ushio.Lock.RLock()
	defer ushio.Lock.RUnlock()
	sessionToken := ctx.Cookies("token")
	ss, err := ushio.Cache.Session(sessionToken)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if ss != nil && ss.Valid() {
		return ctx.Redirect("/home?msg=Already Logged in!", 303)
	}

	return ctx.Render("login", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "login",
		},
	})
}

func (ushio *Ushio) HandlePOSTLogin(ctx *fiber.Ctx) error {
	ushio.Lock.RLock()
	defer ushio.Lock.RUnlock()
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	useEmail := ctx.FormValue("use-email")
	remA := ctx.FormValue("remember")
	reca := ctx.FormValue("g-recaptcha-response")

	s, err := recaptcha.Verify(reca)
	if err != nil {
		return ctx.Render("login", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "login",
			},
			"Warn": "an error on reCAPTCHA occurred, please retry",
		})
	}
	if !s {
		return ctx.Render("login", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "login",
			},
			"Warn": "reCAPTCHA not passed",
		})
	}

	u := &user.User{}
	if useEmail == "on" {
		u, err = ushio.Data.UserByEmail(email)
	} else {
		u, err = ushio.Data.UserByUsername(username)
	}

	if err != nil {
		ctx.Status(401)
		if err == sql.ErrNoRows {
			return ctx.Render("login", fiber.Map{
				"Meta": Meta{
					Config:      *ushio.Config,
					CurrentPage: "login",
				},
				"Warn": "user not found",
			})
		}
		return err
	}

	auth, err := ushio.Data.UserAuthByUID(u.UID)
	if err != nil {
		return err
	}

	if auth.Valid([]byte(password)) {
		rem := false
		if remA == "on" {
			rem = true
		}

		t := time.Now()
		s := &session.Session{
			UID:       u.UID,
			Token:     token.New(),
			UA:        string(ctx.Request().Header.UserAgent()),
			IP:        ctx.IP(),
			CreatedAt: t,
			ExpireAt:  t.Add(36 * time.Hour),
		}

		if rem {
			s.ExpireAt = t.Add(720 * time.Hour)
		}

		err := ushio.Data.InsertSession(s)
		if err != nil {
			return err
		}

		ck := &fiber.Cookie{
			Name:  "token",
			Value: s.Token,
			Path:  "/",
		}

		if rem {
			ck.Expires = s.ExpireAt
		}

		ctx.Cookie(ck)
		return ctx.Redirect("/home", 303)
	} else {
		ctx.Status(401)
		return ctx.Render("login", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "login",
			},
			"Warn": "wrong password",
		})
	}
}

func (ushio *Ushio) HandleSignUp(ctx *fiber.Ctx) error {
	ushio.Lock.RLock()
	defer ushio.Lock.RUnlock()
	tk := ctx.Query("token")
	email := ctx.Query("email")
	if len(tk) != 0 {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: email,
				Token: tk,
			},
		})
	} else {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 1",
			},
			"Step": 1,
		})
	}
}

func (ushio *Ushio) HandlePOSTSignUp(ctx *fiber.Ctx) error {
	ushio.Lock.RLock()
	defer ushio.Lock.RUnlock()
	step := ctx.Query("step")
	stepI, err := strconv.Atoi(step)
	if err != nil || (stepI != 1 && stepI != 2) {
		return fiber.NewError(http.StatusBadRequest, "Unknown step.")
	}

	switch stepI {
	case 1:
		return ushio.signUpS1PostHandler(ctx)
	default:
		return ushio.signUpS2PostHandler(ctx)
	}
}

func (ushio *Ushio) signUpS2PostHandler(ctx *fiber.Ctx) error {
	tk := ctx.FormValue("token")
	name := ctx.FormValue("name")
	username := ctx.FormValue("username")
	em := ctx.FormValue("email")
	password := ctx.FormValue("password")
	reca := ctx.FormValue("g-recaptcha-response")

	s, err := recaptcha.Verify(reca)
	if err != nil {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": "an error on reCAPTCHA occurred, please retry",
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}
	if !s {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": "reCAPTCHA not passed",
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}

	su, err := ushio.Data.SignUpByToken(tk)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Render("sign_up", fiber.Map{
				"Meta": Meta{
					Config:      *ushio.Config,
					CurrentPage: "sign up - step 2",
				},
				"Warn": template.HTML(`token not found, please use the latest token or <a href="/sign_up">re-sign up</a>`),
				"Step": 2,
				"SignUp": sign_up.SignUp{
					Email: em,
					Token: tk,
				},
			})
		}
		return err
	}

	if !su.Valid() {
		err = ushio.Data.DeleteSignUpByEmail(su.Email)
		if err != nil {
			return err
		}
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": template.HTML(`token expired, please <a href="/sign_up">re-sign up</a>`),
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}

	exists, err := ushio.Data.EmailExists(su.Email)
	if err != nil {
		return err
	}
	if exists {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": template.HTML(`user already exists, please <a href="/login">login</a> directly`),
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}

	if !validate.Username(username) {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": template.HTML(`username not valid <br />
<strong>Username MUST be:</strong>
<ul><li>length: [1,10]</li>
<li>numbers, lowercase letters, _ only</li>
<li>the first character must be lowercase letter</li></ul>`),
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}

	if !validate.Name(name) {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": template.HTML(`name not valid <br />
<strong>Name MUST be:</strong>
<ul><li>length: [1,20]</li>
<li>UTF-8 charters</li></ul>`),
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}

	exists, err = ushio.Data.UsernameExists(username)
	if err != nil {
		return err
	}
	if exists {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 2",
			},
			"Warn": "username occupied, please use another username",
			"Step": 2,
			"SignUp": sign_up.SignUp{
				Email: em,
				Token: tk,
			},
		})
	}

	nu := &user.User{
		Name:      name,
		Username:  username,
		Email:     su.Email,
		CreatedAt: time.Now(),
	}
	nu.Tidy()

	uid, err := ushio.Data.InsertUser(nu)
	if err != nil {
		return err
	}

	auth := &user.Auth{
		UID:      uid,
		Password: hash.SHA256([]byte(password)),
	}

	err = ushio.Data.InsertUserAuth(auth)
	if err != nil {
		return err
	}

	err = ushio.Data.DeleteSignUpByEmail(su.Email)
	if err != nil {
		return err
	}

	return ctx.Render("success", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "success",
		},
		"Redirect": "/login",
		"Message":  template.HTML("You have successfully signed up,<br/> please login in the next page."),
	})
}

func (ushio *Ushio) signUpS1PostHandler(ctx *fiber.Ctx) error {
	email := ctx.FormValue("email")
	reca := ctx.FormValue("g-recaptcha-response")
	s, err := recaptcha.Verify(reca)
	if err != nil {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 1",
			},
			"Warn": "an error on reCAPTCHA occurred, please retry",
			"Step": 1,
		})
	}
	if !s {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 1",
			},
			"Warn": "reCAPTCHA not passed",
			"Step": 1,
		})
	}

	if !validate.Email(email) {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 1",
			},
			"Warn": "email not valid",
			"Step": 1,
		})
	}

	exists, err := ushio.Data.EmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return ctx.Render("sign_up", fiber.Map{
			"Meta": Meta{
				Config:      *ushio.Config,
				CurrentPage: "sign up - step 1",
			},
			"Warn": "user already exists, please use login",
			"Step": 1,
		})
	}

	exists, err = ushio.Data.SignUpExists(email)
	if err != nil {
		return err
	}

	if exists {
		err := ushio.Data.DeleteSignUpByEmail(email)
		if err != nil {
			return err
		}
	}

	su := sign_up.New(email, 24*time.Hour)
	err = ushio.Data.InsertSignUp(su)
	if err != nil {
		return err
	}
	err = ushio.Config.Sender.Send(email, su.Token)
	if err != nil {
		return err
	}

	return ctx.Render("success", fiber.Map{
		"Meta": Meta{
			Config:      *ushio.Config,
			CurrentPage: "success",
		},
		"Message": template.HTML(`You have successfully signed up, please check your mailbox.<br />
If you couldn't' received verify email, please check your spam mailbox or sign up again.<br />
If all above don't work, please contact webmaster.`),
	})
}
