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
	UserID    uuid.UUID  `json:"userId" binding:"required"`
	BlogID    uuid.UUID  `json:"blogId" binding:"required"`
	Interests []Interest `json:"interests"`
}
