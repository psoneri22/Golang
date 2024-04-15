package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Get Method"))
		case http.MethodPost:
			w.Write([]byte("Post Method"))
		default:
			w.WriteHeader(http.StatusNotAcceptable)
		}
	}
	// else {
	// 	// matches := uc.userIDPattern.FindAllStringSubmatch(r.URL.Path)
	// 	// if len(matches) == 0 {
	// 	// 	w.WriteHeader(http.StatusNotFound)
	// 	// }
	// 	// id, err := strconv.Atoi(matches[1])
	// 	// if err != nil {
	// 	// 	w.WriteHeader(http.StatusNotFound)
	// 	// }
	// 	// switch r.Method {
	// 	// case http.MethodGet:
	// 	// 	uc.get(id, w)
	// 	// case http.MethodPut:
	// 	// 	uc.put(id, w, r)
	// 	// case http.MethodDelete:
	// 	// 	uc.delete(id, w, r)
	// 	// default:
	// 	// 	w.WriteHeader(http.StatusNotFound)
	// 	// }
	// }
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^users/(\d+)/?`),
	}
}
