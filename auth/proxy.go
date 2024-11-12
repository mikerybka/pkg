package auth

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/mikerybka/pkg/english"
	"github.com/mikerybka/pkg/util"
)

type Proxy struct {
	Data       *Data
	Twilio     *util.TwilioClient
	BackendURL string
}

func (s *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/auth/send-login-code":
		s.SendLoginCode(w, r)
	case "/auth/login":
		s.Login(w, r)
	case "/auth/logout":
		s.Logout(w, r)
	default:
		session := s.Session(r)
		if session != nil {
			// Proxy the request to the backend adding a `UserID` header
			url, err := url.Parse(s.BackendURL)
			if err != nil {
				panic(err)
			}
			proxy := httputil.NewSingleHostReverseProxy(url)
			proxy.Rewrite = func(r *httputil.ProxyRequest) {
				r.Out.Header.Add("UserID", session.UserID)
			}
			proxy.ServeHTTP(w, r)
		} else {
			// Redirect to the login page
			target := r.URL.Path
			if r.URL.RawQuery != "" {
				target += "?" + r.URL.RawQuery
			}
			params := url.Values{}
			params.Add("target", target)
			redirectURL := url.URL{
				Path:     "/auth/login",
				RawQuery: params.Encode(),
			}
			http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
		}
	}
}

func (s *Proxy) SendLoginCode(w http.ResponseWriter, r *http.Request) {
	form := &util.Form{
		Name: english.NewName("Send Login Code"),
		Fields: []util.Field{
			{
				Name: english.NewName("Phone"),
				Type: &util.Type{
					IsScalar: true,
					Kind:     "phone",
				},
			},
		},
		Handle: func(w http.ResponseWriter, r *http.Request) {
			// Parse
			req := &SendLoginCodeRequest{}
			if util.ContentType(r, "application/json") {
				err := json.NewDecoder(r.Body).Decode(req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
			} else {
				req.Phone = util.ParseJSON[util.PhoneNumber](r.FormValue("phone"))
			}

			// Validate
			err := req.Phone.Validate()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Execute
			res := s.sendLoginCode(req)

			// Respond
			if util.Accept(r, "application/json") {
				json.NewEncoder(w).Encode(res)
			} else {
				nextURL := fmt.Sprintf("/auth/login?userID=%s", res.UserID)
				http.Redirect(w, r, nextURL, http.StatusSeeOther)
			}
		},
	}
	form.ServeHTTP(w, r)
}

func (s *Proxy) sendLoginCode(req *SendLoginCodeRequest) *SendLoginCodeResponse {
	u, err := s.Data.UserByPhone(req.Phone)
	if err.Error() == "not found" {
		u, err = s.createUser(req.Phone)
	}
	if err != nil {
		panic(err)
	}
	err = u.SendLoginCode(s.Twilio)
	if err != nil {
		panic(err)
	}
	err = s.Data.SaveUser(u)
	if err != nil {
		panic(err)
	}
	return &SendLoginCodeResponse{
		UserID: u.ID,
	}
}

func (s *Proxy) Login(w http.ResponseWriter, r *http.Request) {
	form := &util.Form{
		Name: english.NewName("Login"),
		Fields: []util.Field{
			{
				Name: english.NewName("Code"),
				Type: LoginCodeType,
			},
		},
		Handle: func(w http.ResponseWriter, r *http.Request) {
			req := &LoginRequest{}
			if util.ContentType(r, "application/json") {
				err := json.NewDecoder(r.Body).Decode(&req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			} else {
				panic("not implemented")
				// TODO
			}

			res := s.login(req)

			if util.Accept(r, "application/json") {
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic("programmer error")
				}
			} else {
				http.SetCookie(w, &http.Cookie{
					Name:  "UserID",
					Value: res.UserID,
				})
				http.SetCookie(w, &http.Cookie{
					Name:  "Token",
					Value: res.Token,
				})
				http.Redirect(w, r, "/"+res.UserID, http.StatusSeeOther)
			}
		},
	}
	form.ServeHTTP(w, r)
}

func (s *Proxy) login(req *LoginRequest) *LoginResponse {
	u := s.Data.User(req.UserID)
	if u == nil {
		return &LoginResponse{
			Error: "no such user",
		}
	}

	token, err := u.Login(req.Code)
	if err != nil {
		return &LoginResponse{
			Error: err.Error(),
		}
	}
	err = s.Data.SaveUser(u)
	if err != nil {
		return &LoginResponse{
			Error: err.Error(),
		}
	}

	return &LoginResponse{
		Ok:     true,
		UserID: u.ID,
		Token:  token,
	}
}

func (s *Proxy) Logout(w http.ResponseWriter, r *http.Request) {
	form := &util.Form{
		Name: english.NewName("Logout"),
		Handle: func(w http.ResponseWriter, r *http.Request) {
			req := &LogoutRequest{}
			if util.ContentType(r, "application/json") {
				err := json.NewDecoder(r.Body).Decode(&req)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			} else {
				req.Session = s.Session(r)
			}

			res := s.logout(req)

			if util.Accept(r, "application/json") {
				err := json.NewEncoder(w).Encode(res)
				if err != nil {
					panic("programmer error")
				}
			} else {
				http.SetCookie(w, &http.Cookie{
					Name:  "UserID",
					Value: "",
				})
				http.SetCookie(w, &http.Cookie{
					Name:  "Token",
					Value: "",
				})
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		},
	}
	form.ServeHTTP(w, r)
}

func (s *Proxy) logout(req *LogoutRequest) *LogoutResponse {
	u := s.Data.User(req.Session.UserID)
	u.Logout(req.Session.Token)
	return &LogoutResponse{
		Ok: true,
	}
}

func (s *Proxy) Session(r *http.Request) *Session {
	c, err := r.Cookie("UserID")
	if err != nil {
		return nil
	}
	userID := c.Value
	u := s.Data.User(userID)

	c, err = r.Cookie("Token")
	if err != nil {
		return nil
	}
	token := c.Value

	if !u.ValidSession(token) {
		return nil
	}

	return &Session{
		UserID: userID,
		Token:  token,
	}
}

func (s *Proxy) canRead(userID, orgID string) bool {
	return s.Data.Org(orgID).CanRead(userID)
}

func (s *Proxy) canWrite(userID, orgID string) bool {
	return s.Data.Org(orgID).CanWrite(userID)
}

func (s *Proxy) Authorized(r *http.Request) bool {
	pathParts := util.ParsePath(r.URL.Path)
	if len(pathParts) == 0 {
		return true
	}
	orgID := pathParts[0]
	session := s.Session(r)
	if r.Method == http.MethodGet || r.Method == http.MethodHead {
		return s.canRead(session.UserID, orgID)
	} else {
		return s.canWrite(session.UserID, orgID)
	}
}

func (s *Proxy) createUser(phone util.PhoneNumber) (*User, error) {
	u := &User{
		ID:    util.RandomID(),
		Phone: phone,
	}
	err := s.Data.SaveUser(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
