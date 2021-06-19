package types

// CreateRequest - данные, которые запрашиваютя при создании объявления
type CreateRequest struct {
	Name        string   `json:"name" validate:"required,max=200"`
	Description string   `json:"description" validate:"required,max=1000"`
	Links       []string `json:"links" validate:"required,min=1,max=3,dive,required,url"`
	Price       float64  `json:"price" validate:"required,min=0"`
}

// GetAllRequest - данные, которые запрашиваются при выборке одного объявления
type GetAllRequest struct {
	Page int    `json:"page" validate:"required,min=1"`
	Sort string `json:"sort" validate:"required"`
}

// GetOneRequest - данные, которые запрашиваются при выборке всех объявлений
type GetOneRequest struct {
	ID     int  `json:"id" validate:"required,min=1"`
	Fields bool `json:"fields"`
}

// CreateResponse - данные, котоорые возвращаются после создания объявления
type CreateResponse struct {
	ID int `json:"id"`
}

// GetOneResponse - данные, котоорые возвращаются после выборки одного объявления
type GetOneResponse struct {
	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	MainLink    string   `json:"main_link"`
	Description string   `json:"description,omitempty"`
	AllLinks    []string `json:"all_links,omitempty"`
}

// GetAllResponse - данные, котоорые возвращаются после выборки всех объявлений
type GetAllResponse struct {
	Advertisements []Advertisement `json:"advertisements"`
}

// Advertisement - данные, которые содержатся в каждом объявлении
type Advertisement struct {
	Name     string  `json:"name"`
	MainLink string  `json:"main_link"`
	Price    float64 `json:"price"`
}

// SortingOptions - мапа соответствий сортировок restapi-sql для GetAll
var SortingOptions = map[string]string{
	"price_desc": "price DESC",
	"price_asc":  "price",
	"date_desc":  "created_at DESC",
	"date_asc":   "created_at",
}

// GetAllForService - преобразованные данные для бизнес-логики.
// Sort имеет прямое соответствие с SortingOptions
type GetAllForService struct {
	Page int
	Sort string
}
