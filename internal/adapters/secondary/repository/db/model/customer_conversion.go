package model

import "hexagonal-go-grpc/internal/adapters/service/entity"

func CustomerEntityToModel(m entity.Customer) Customer {
	return Customer{
		Name:      m.Name,
		Email:     m.Email,
		Phone:     m.Phone,
		Gender:    m.Gender,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func CustomerCreateEntityToModel(m entity.CustomerCreateRequest) Customer {
	return Customer{
		Name:     m.Name,
		Email:    m.Email,
		Phone:    m.Phone,
		Gender:   m.Gender,
		IsActive: m.IsActive,
	}
}

func CustomerUpdateEntityToModel(m entity.CustomerUpdateRequest) Customer {
	return Customer{
		Name:     m.Name,
		Email:    m.Email,
		Phone:    m.Phone,
		Gender:   m.Gender,
		IsActive: m.IsActive,
	}
}

func (m Customer) ToEntity() entity.Customer {
	return entity.Customer{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Phone:     m.Phone,
		Gender:    m.Gender,
		IsActive:  m.IsActive,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func CustomerEntityListToModelList(m []entity.Customer) []Customer {
	var result []Customer
	for _, v := range m {
		result = append(result, CustomerEntityToModel(v))
	}
	return result
}

func CustomerModelListToEntityList(m []Customer) []entity.Customer {
	var result []entity.Customer
	for _, v := range m {
		result = append(result, v.ToEntity())
	}
	return result
}
