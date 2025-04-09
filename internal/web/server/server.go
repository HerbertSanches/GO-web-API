package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nomedoseuprojeto/internal/domain"
	"nomedoseuprojeto/internal/service"
)

type Server struct {
	AuthorService *service.AuthorService
	port 		  string
}

func NewServer(authorSevice *service.AuthorService, port string) *Server{
	return &Server{
		AuthorService: authorSevice,
		port: port,
	}
}

func (s *Server) ConfigRoutes(){
	http.HandleFunc("/authors", s.handleCreateAuthor)
}

func (s *Server) handleCreateAuthor(w http.ResponseWriter, r *http.Request){
	var authorInput domain.AuthorInput
	if r.Method != http.MethodPost{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&authorInput); err != nil{
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := s.AuthorService.CreateAuthor(authorInput.Name)
	if err != nil{
		http.Error(w, "Erro ao criar autor", http.StatusInternalServerError)
		return	
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Autor criado com sucesso")
}

func (s *Server) Start() error {
	fmt.Println("Servidor rodando na porta", s.port)
	return http.ListenAndServe(":"+s.port, nil)
}