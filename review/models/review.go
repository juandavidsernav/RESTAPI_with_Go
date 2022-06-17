package models

import (
	"errors"
	"time"
)

const MaxLengthInComments = 400

type Review struct {
	Id      int64
	Stars   int    //1 - 5
	Comment string //max 400 chars
	Date    time.Time
}

type CreateReviewCMD struct {
	Stars   int    `json:"stars"`
	Comment string `json:"comment"`
}

func (cmd *CreateReviewCMD) validate() error {
	if cmd.Stars < 1 || cmd.Stars > 5 {
		return errors.New("starts must be between 1 to 5")
	}
	if len(cmd.Comment) > MaxLengthInComments {
		return errors.New("comment must be less than 400 chars")
	}
	return nil
}
