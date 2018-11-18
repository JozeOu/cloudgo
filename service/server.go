package service

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
		Extensions: []string{".html", ".jpg"},
	})
	// negroni 包中的初始化服务
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
	mx.PathPrefix("/show").Handler(http.StripPrefix("/show/", http.FileServer(http.Dir("assets/"))))

}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		// 获取请求参数
		vars := mux.Vars(req)
		id := vars["id"]
		// 使用 render 包中的 JSON 函数进行解析
		formatter.JSON(w, http.StatusOK, struct{ ShowHello string }{"Hello " + id})
	}
}
