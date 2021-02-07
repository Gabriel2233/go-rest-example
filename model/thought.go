package model

import (
	"log"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Thought struct {
	ID          string    `json:"id" gorm:"type:uuid;primaryKey" valid:"uuid"`
	Title       string    `json:"title" gorm:"varchar(20); not null" valid:"notnull"`
	Description string    `json:"description" gorm:"varchar(255); not null" valid:"notnull"`
	Tag         string    `json:"tag" gorm:"varchar(80)" valid:"-"`
	CreatedAt   time.Time `json:"created_at"`
}

type ThoughtRepositoryInterface interface {
	Add(t *Thought) error
	FindAll() ([]*Thought, error)
	FindByTagName(tag string) ([]*Thought, error)
}

func (t *Thought) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewThought() *Thought {
	thought := &Thought{
		ID: uuid.NewV4().String(),
	}

	err := thought.isValid()

	if err != nil {
		log.Fatal("Invalid Thought")
		return nil
	}

	return thought
}
