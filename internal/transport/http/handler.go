package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/eozdeveci/CommentRestApi/internal/comment"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods(http.MethodGet)
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods(http.MethodPost)
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods(http.MethodPut)
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods(http.MethodDelete)

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)

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
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Error retrieving all comments")
	}

	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode json body")
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from id")
	}

	var newComment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		fmt.Fprintf(w, "Failed to decode json body")
	}

	comment, err := h.Service.UpdateComment(uint(commentID), newComment)
	if err != nil {
		fmt.Fprintf(w, "Failed to update comment")
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse UINT from id")
	}

	if err := h.Service.DeleteComment(uint(commentID)); err != nil {
		fmt.Fprintf(w, "Failed to delete comment")
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}
}
