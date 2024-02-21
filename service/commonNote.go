package service

import (
	"fmt"
	"list/model"
	"list/repository"
)

type CommonNoteService interface {
	Insert(header, body string) error
	GetAllData() ([]model.UserInput, error)
	Delete(header string) error
}

type commonNoteService struct {
	cache repository.CacheRepository
	repo  repository.NoteRepository
}

func (r *commonNoteService) Insert(header, body string) error {
	//database
	if err := r.repo.Insert(model.UserInput{
		Header: header,
		Body:   body,
	}); err != nil {
		return err
	}

	//cache
	if err := r.cache.Set(header, []byte(body)); err != nil {
		return err
	}
	return nil
}

func (r *commonNoteService) GetAllData() ([]model.UserInput, error) {
	var resultData []model.UserInput
	Userdata, err := r.cache.GetAll()

	if err != nil {
		return []model.UserInput{}, fmt.Errorf("error in service %s", err)
	}

	for i, data := range Userdata {
		userdata, err := r.cache.Get(data)

		if userdata == nil {
			res, err := r.repo.GetByHeader(Userdata[i])

			if err != nil {
				return []model.UserInput{}, fmt.Errorf("error in service %s", err)
			}
			userdata = []byte(res.Body)
		}

		if err != nil {
			return []model.UserInput{}, fmt.Errorf("error in service %s", err)
		}
		data := model.UserInput{
			Header: Userdata[i],
			Body:   string(userdata),
		}
		resultData = append(resultData, data)
	}
	return resultData, nil
}

func (r *commonNoteService) Delete(header string) error {
	if err := r.cache.Delete(header); err != nil {
		return err
	}
	if err := r.repo.Delete(header); err != nil {
		return err
	}
	return nil
}

func NewCommonNoteService(cache repository.CacheRepository, repo repository.NoteRepository) CommonNoteService {
	return &commonNoteService{
		cache: cache,
		repo:  repo,
	}
}
