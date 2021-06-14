package restapi

// CreateRequest - данные, которые запрашиваютя при создании объявления
type CreateRequest struct {
	Name        string   `json:"name" validate:"required,max=200"`
	Description string   `json:"description" validate:"required,max=1000"`
	Links       []string `json:"links" validate:"required,min=1,max=3,dive,required"`
	Price       float64  `json:"price" validate:"required,number"`
}

// GetAllRequest - данные, которые запрашиваются при выборке одного объявления
type GetAllRequest struct {
	Page int    `json:"page"`
	Sort string `json:"sort"`
}

// GetOneRequest - данные, которые запрашиваются при выборке всех объявлений
type GetOneRequest struct {
	ID     int  `json:"id"`
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
	Description string
	AllLinks    []string
}

// GetAllResponse - данные, котоорые возвращаются после выборки всех объявлений
type GetAllResponse struct {
	Advertisements []Advertisement
}

// Advertisement - данные, которые содержатся в каждом объявлении
type Advertisement struct {
}

// Статическая переменная ответа 404 - не найдено
var responseNotFound404 map[string]string = map[string]string{
	"status_code":    "404",
	"status_message": "Not Found",
}
