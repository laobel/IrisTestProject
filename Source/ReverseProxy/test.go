/*package main

import (
   "github.com/kataras/iris"
   "math/rand"
   "net/http"
   "net/http/httputil"
   "net/url"
)

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
   director := func(req *http.Request) {
      target := targets[rand.Int()%len(targets)]
      req.URL.Scheme = target.Scheme
      req.URL.Host = target.Host
      req.URL.Path = target.Path
   }
   return &httputil.ReverseProxy{Director: director}
}

func main() {
   app := iris.New()

   app.Get("/", func(ctx iris.Context) {
      ctx.Writef("Hello from the server")
   })

   app.Get("/mypath", func(ctx iris.Context) {
      ctx.Writef("Hello from %s", ctx.Path())
   })

   // call .Build before use the 'app' as a http.Handler on a custom http.Server
   app.Build()

   proxy := NewMultipleHostsReverseProxy([]*url.URL{
      {
         Scheme: "http",
         Host:   "182.247.253.66:6080",
         Path: "/arcgis/rest/services/TdtYn/tdtYnImg100cm84_2017/MapServer/tile/15/14012/25411",
      },
      {
         Scheme: "http",
         Host:   "182.247.253.80:6080",
         Path: "/arcgis/rest/services/TDTYN/tdtYnImgLable84/MapServer/tile/15/14012/25411",
      },
   })
   //http.ListenAndServe(":9090", proxy)
   //log.Fatal(http.ListenAndServe(":9090", proxy))

   // create our custom server and assign the Handler/Router
   srv := &http.Server{Handler: app, Addr: ":9090"} // you have to set Handler:app and Addr, see "iris-way" which does this automatically.
   // http://localhost:8080/
   // http://localhost:8080/mypath
   println("Start a server listening on http://localhost:8080")
   srv.ListenAndServe() // same as app.Run(iris.Addr(":8080"))
}*/

package main

import (
   "github.com/gorilla/mux"
   "github.com/kataras/iris"
   "gopkg.in/couchbase/gocb.v1"
   "math/rand"
   "net/http"
   "net/http/httputil"
   "net/url"
)

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
   director := func(req *http.Request) {
      target := targets[rand.Int()%len(targets)]
      req.URL.Scheme = target.Scheme
      req.URL.Host = target.Host
      req.URL.Path = target.Path
   }
   return &httputil.ReverseProxy{Director: director}
}

type Terrain struct {
   X     int    `json:"x"`
   Y     int    `json:"y"`
   Z     int    `json:"z"`
   Value []byte `json:"value"`
}

func main() {
   app := iris.New()

   cluster, _ := gocb.Connect("couchbase://192.168.1.10")
   cluster.Authenticate(gocb.PasswordAuthenticator{
      Username: "Administrator",
      Password: "123369",
   })
   bucket, _ := cluster.OpenBucket("TestBucket", "")

   bucket.Manager("", "").CreatePrimaryIndex("", true, false)


   proxy := NewMultipleHostsReverseProxy([]*url.URL{
      {
         Scheme: "http",
         Host:   "182.247.253.66:6080",
         Path: "/arcgis/rest/services/TdtYn/tdtYnImg100cm84_2017/MapServer/tile/15/14012/25411",
      },
      {
         Scheme: "http",
         Host:   "182.247.253.80:6080",
         Path: "/arcgis/rest/services/TDTYN/tdtYnImgLable84/MapServer/tile/15/14012/25411",
      },
   })

   app.Get("/", func(ctx iris.Context) {
      ctx.Writef("Hello from the server")
   })

   app.Get("/test", func(ctx iris.Context) {
      proxy.ServeHTTP(ctx.ResponseWriter(),ctx.Request())
   })
   app.Get("/terrain", func(ctx iris.Context) {
      var terrain Terrain
      bucket.Get("010", &terrain)

      w:=ctx.ResponseWriter()
      r:=ctx.Request()

      vars := mux.Vars(r)

      headers := w.Header()
      headers.Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
      headers.Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
      headers.Set("Content-Type", "application/octet-stream")
      headers.Set("Content-Encoding", "gzip")
      headers.Set("Content-Disposition", "attachment;filename="+vars["y"]+".terrain")
      w.Write(terrain.Value)
   })

   app.Get("/mypath", func(ctx iris.Context) {
      ctx.Writef("Hello from %s", ctx.Path())
   })


   // Any custom fields here. Handler and ErrorLog are setted to the server automatically
   srv := &http.Server{Addr: ":9090",Handler:app}

   // http://localhost:8080/
   // http://localhost:8080/mypath
   app.Run(iris.Server(srv)) // same as app.Run(iris.Addr(":8080"))

   // More:
   // see "multi" if you need to use more than one server at the same app.
   //
   // for a custom listener use: iris.Listener(net.Listener) or
   // iris.TLS(cert,key) or iris.AutoTLS(), see "custom-listener" example for those.
}


