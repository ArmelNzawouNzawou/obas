package io

import (
	"errors"
	"obas/api"
	domain "obas/domain/location"
)

const locationTypeUrl = api.BASE_URL + "/location"

type LocationType domain.LocationType

func GetLocationTypes() ([]domain.LocationType, error) {
	entites := []domain.LocationType{}
	resp, _ := api.Rest().Get(locationTypeUrl + "/type/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetLocationType(id string) (domain.LocationType, error) {
	entity := domain.LocationType{}
	resp, _ := api.Rest().Get(locationTypeUrl + "/type/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func CreateLocationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationTypeUrl + "/type/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateLocationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationTypeUrl + "/type/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteLocationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationTypeUrl + "/type/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
