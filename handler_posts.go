package main

import (
	"fmt"
	"net/http"

	"github.com/ivcp/blog-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerGetPostByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not get posts: %s", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
