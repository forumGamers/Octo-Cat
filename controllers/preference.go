package controllers

import (
	"context"
	"strings"

	h "github.com/forumGamers/Octo-Cat/helpers"
	"github.com/forumGamers/Octo-Cat/pkg/preference"
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewPreferenceController(
	w web.ResponseWriter,
	r web.RequestReader,
	repo preference.PreferenceRepo,
	validate *validator.Validate,
) PreferenceController {
	return &PreferenceControllerImpl{w, r, repo, validate}
}

func (pc *PreferenceControllerImpl) CreateData(c *gin.Context) {
	//pakai producer nanti
	userId := h.GetUser(c).UUID
	if _, err := pc.Repo.FindByUserId(context.Background(), userId); err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "data not found") || err != mongo.ErrNoDocuments {
			pc.AbortHttp(c, err)
			return
		}
	}

	data, err := pc.Repo.Create(context.Background(), userId)
	if err != nil {
		pc.AbortHttp(c, err)
		return
	}

	pc.Write201Response(c, "Success", data)
}
