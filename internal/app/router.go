package app

func (s *Server) Route (){
	s.app.Post("/post", s.CreatePost)
	s.app.Get("/post/:id", s.GetPostById)
	s.app.Get("/posts", s.GetAllPosts)
	s.app.Patch("/post/:id", s.PatchPost)
	s.app.Delete("/post/:id", s.DeletePost)
}