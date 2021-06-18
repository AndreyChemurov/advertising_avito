package types

// CreateRequest - данные, которые запрашиваютя при создании объявления
type CreateRequest struct {
	Name        string   `json:"name" validate:"required,max=200"`
	Description string   `json:"description" validate:"required,max=1000"`
	Links       []string `json:"links" validate:"required,min=1,max=3,dive,required"`
	Price       float64  `json:"price" validate:"required"`
}

// GetAllRequest - данные, которые запрашиваются при выборке одного объявления
type GetAllRequest struct {
	Page int    `json:"page" validate:"required,min=1"`
	Sort string `json:"sort" validate:"required"`
}

// GetOneRequest - данные, которые запрашиваются при выборке всех объявлений
type GetOneRequest struct {
	ID     int  `json:"id" validate:"required"`
	Fields bool `json:"fields"`
}

// CreateResponse - данные, котоорые возвращаются после создания объявления
type CreateResponse struct {
	ID int
}

// GetOneResponse - данные, котоорые возвращаются после выборки одного объявления
type GetOneResponse struct {
	Name        string
	Price       float64
	MainLink    string
	Description string   `json:",omitempty"`
	AllLinks    []string `json:",omitempty"`
}

// GetAllResponse - данные, котоорые возвращаются после выборки всех объявлений
type GetAllResponse struct {
	Advertisements []Advertisement
}

// Advertisement - данные, которые содержатся в каждом объявлении
type Advertisement struct {
}
