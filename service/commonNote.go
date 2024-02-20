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
	repo repository.CacheRepository
}

func (r *commonNoteService) Insert(header, body string) error {
	if err := r.repo.Set(header, []byte(body)); err != nil {
		return err
	}
	return nil
}

func (r *commonNoteService) GetAllData() ([]model.UserInput, error) {
	var resultData []model.UserInput
	Userdata, err := r.repo.GetAll()

	if err != nil {
		return []model.UserInput{}, fmt.Errorf("error in service %s", err)
	}

	for i, data := range Userdata {
		userdata, err := r.repo.Get(data)

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
	if err := r.repo.Delete(header); err != nil {
		return err
	}
	return nil
}

func NewCommonNoteService(repo repository.CacheRepository) CommonNoteService {
	return &commonNoteService{repo: repo}
}
