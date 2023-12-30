package preference

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
)

type PreferenceRepo interface {
	Create(ctx context.Context, userId string) (UserPreference, error)
	FindByUserId(ctx context.Context, userId string) (UserPreference, error)
}

type PreferenceRepoImpl struct{ b.BaseRepo }
