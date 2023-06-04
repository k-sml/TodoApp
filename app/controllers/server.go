package controllers

import (
	"fmt"
	"gostudy/application/config"
	"gostudy/application/app/models"
	"html/template"
	"net/http"
	"strconv"
	"regexp"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string){
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	//template.Must関数はテンプレートの解析中にエラーが発生した場合にパニックを発生させるラッパー関数(ラップ（包む)関数)
	//files...の...は可変長引数の意味
	//可変長引数として渡すことでリストやスライスの中身の各要素を渡せる
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+)$")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := validPath.FindStringSubmatch(r.URL.Path)
		if q == nil {
			http.NotFound(w, r)
			return
		}
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fn(w, r, qi)
	}
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))
	return http.ListenAndServe(":" + config.Config.Port, nil)   //ホスト部分を省略するとデフォルトでlocalhostがつく
	// 第二引数をnilにするとデフォルトのマルチプレクサが使用される
	// 登録されたハンドラー関数をURLパスに基づいて呼び出す（何も登録していなかったらpage not found 404 errorが表示される)
}
