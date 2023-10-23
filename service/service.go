package service

import (
	"encoding/csv"
	"errors"
	"mime/multipart"
	"reflect"
	"strings"

	"exportimportcsv/helper/crypt"
	"exportimportcsv/model"
	"exportimportcsv/repository"
	"exportimportcsv/utility"

)

type Service interface {
	ImportCSV(fileCSV *multipart.FileHeader) (data []model.User, err error)
	ExportCSV() (data []model.ExportUser, err error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) ImportCSV(fileCSV *multipart.FileHeader) (data []model.User, err error) {
	file, err := fileCSV.Open()
	if err != nil {
		return
	}
	defer file.Close()

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		return
	}

	headerFile := []string{"Number", "Email", "Password", "Role Name"}

	if !reflect.DeepEqual(records[0], headerFile) {
		return nil, errors.New("File headers do not match")
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		guid := utility.GenerateGoogleUUID()
		email := strings.TrimSpace(record[1])
		password, err := crypt.GenerateHashPassword(record[2])
		if err != nil {
			return []model.User{}, err
		}

		roleName := strings.TrimSpace(record[3])

		user, err := s.repository.CreateUser(guid, email, password, roleName)
		if err != nil {
			return []model.User{}, err
		}

		data = append(data, user)
	}

	return
}

func (s *service) ExportCSV() (data []model.ExportUser, err error) {
	data, err = s.repository.GetUser()
	if err != nil {
		return
	}
	
	return
}
