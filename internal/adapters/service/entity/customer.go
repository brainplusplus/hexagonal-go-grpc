package entity

type Customer struct {
	ID        int32
	Name      string
	Email     string
	Phone     string
	Gender    string
	IsActive  bool
	CreatedAt int64
	UpdatedAt int64
}

type CustomerIdRequest GenericIdRequest[int32]

type CustomerListRequest GenericListRequest[int32]

type CustomerEmailOrPhoneRequest struct {
	Email string
	Phone string
}

type CustomerCreateRequest struct {
	Name     string
	Email    string
	Phone    string
	Gender   string
	IsActive bool
}
type CustomerUpdateRequest struct {
	ID       int32
	Name     string
	Email    string
	Phone    string
	Gender   string
	IsActive bool
}

type CustomerRegisterRequest struct {
	Name   string
	Email  string
	Phone  string
	Gender string
}
