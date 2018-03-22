package gaehttp

import (
	"fmt"
	"net/http"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type User struct {
	Name string
}

// パラメータで受け取った id のユーザーを Datastore から検索して "hello user.name" と返却するハンドラー.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	id := r.FormValue("id")

	key := datastore.NewKey(ctx, "User", id, 0, nil)
	u := new(User)
	if err := datastore.Get(ctx, key, u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "hello %s", u.Name)
}
