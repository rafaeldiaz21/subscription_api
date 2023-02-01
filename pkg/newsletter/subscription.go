package newsletter

import (
	"github.com/google/uuid"
)

type Interest string

var (
	InterestTech     Interest = "tech"
	InterestPolitics Interest = "politics"
	InterestSports   Interest = "sports"
)

type Subscription struct {
	UserID    uuid.UUID  `json:"userID" binding:"required"`
	BlogID    uuid.UUID  `json:"blogID" binding:"required"`
	Interests []Interest `json:"interests"`
}
