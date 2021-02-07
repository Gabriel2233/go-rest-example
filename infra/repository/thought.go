package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/mural-app/server/model"
)

type ThoughtRepositoryDb struct {
	Db *gorm.DB
}

func (r *ThoughtRepositoryDb) Add(t *model.Thought) error {
	err := r.Db.Create(t).Error

	if err != nil {
		return err
	}

	return nil
}

func (r ThoughtRepositoryDb) FindAll() []*model.Thought {
	var thoughts []*model.Thought

	r.Db.Find(&thoughts)

	return thoughts
}

func (r ThoughtRepositoryDb) FindByTag(tag string) []*model.Thought {
	var filteredThoughts []*model.Thought

	r.Db.Where("tag = ?", tag).Find(&filteredThoughts)

	return filteredThoughts
}
