package server

import (
	st "../data_struct"
	s "../service"
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	config     *st.Config
	urlService *s.UrlColorService
}

type Data struct {
	UrlImgData []st.UrlImage
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.UrlcolorHttp)
	return mux
}

func (s *Server) Run() {

	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}
	httpServer.ListenAndServe()
}

func (s *Server) RunUrlColor() {
	s.urlService.CreateTable()
	s.urlService.InsertUrl(s.urlService.MakeUrlColor())
}

func (s *Server) UrlcolorHttp(w http.ResponseWriter, r *http.Request) {

	urlColors := s.urlService.FindAll()

	data := Data{
		UrlImgData: urlColors,
	}
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, data)
	log.Print("Server is listening...")
}

func NewServer(config *st.Config, service *s.UrlColorService) *Server {
	return &Server{
		config:     config,
		urlService: service,
	}
}
