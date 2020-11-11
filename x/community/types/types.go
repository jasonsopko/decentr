package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gofrs/uuid"
)

type Category uint8

const (
	InvalidCategory Category = iota
	WorldNewsCategory
	TravelAndTourismCategory
	ScienceAndTechnologyCategory
	StrangeWorldCategory
	HealthAndCultureCategory
	FitnessAndExerciseCategory
)

type Post struct {
	UUID          uuid.UUID      `json:"uuid"`
	Owner         sdk.AccAddress `json:"owner"`
	Title         string         `json:"title"`
	PreviewImage  string         `json:"previewImage"`
	Category      Category       `json:"category"`
	Text          string         `json:"text"`
	LikesCount    uint32         `json:"likesCount"`
	DislikesCount uint32         `json:"dislikesCount"`
	CreatedAt     int64          `json:"createdAt"`
}