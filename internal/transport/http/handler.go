package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/eozdeveci/CommentRestApi/internal/comment"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment/", h.PostComment).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am alive!")
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from id")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprintf(w, "Error retrieving comment by id")
	}
	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {

	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Error retrieving all comments")
	}
	fmt.Fprintf(w, "%+v", comments)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from id")
	}

	comment, err := h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to delete comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}
