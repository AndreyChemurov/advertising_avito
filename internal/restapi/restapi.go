package restapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

// Create - метод создания нового объявления
// Аргументы:
//	name: название объявления
//	description: описание объявления
//	links: список ссылок на фотографии (первая переденная будет главной)
//	price: цена за товар в объявлении
// Возвращаемые значения:
//	id: идентификатор созданного объявления
//	status_code: код результата (200 в случае успеха)
func Create(w http.ResponseWriter, r *http.Request) {
	var (
		adv      CreateRequest
		response []byte
	)

	// Проверить валидность JSON'а
	if err := json.NewDecoder(r.Body).Decode(&adv); err != nil {
		response = errorType(http.StatusBadRequest, fmt.Sprintf("%v", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	// Проврить валидность параметров
	var validate *validator.Validate = validator.New()

	err := validate.Struct(&adv)
	if err != nil {
		response = errorType(http.StatusBadRequest, fmt.Sprintf("%v", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	// Service logic

	// Статус 200 ОК
	w.WriteHeader(http.StatusOK)
	w.Write(responseOK)
}

// GetOne - метод получения конкретного объявления по его ID
// Аргументы:
//	id: уникальный идентификатор объявления.
//	fields: опциональное поле. Если оно указано [fields: true],
//		то возвращаются так же дополнительные поля: описание и все ссылки на фото.
//		Если поле fields не указано или значение false, то вышеуказанные поля не возвращаются.
// Возвращаемые значения:
//	name: название объявления
//	price: цена за объявление
//	mainlink: ссылка на главное фото
//	[description]: описание объявления
//	[alllinks]: ссылки на все фото
//	status_code: код результата (200 в случае успеха)
func GetOne(w http.ResponseWriter, r *http.Request) {
	var (
		adv      GetOneRequest
		response []byte
	)

	// Проверить валидность JSON'а
	if err := json.NewDecoder(r.Body).Decode(&adv); err != nil {
		response = errorType(http.StatusBadRequest, fmt.Sprintf("%v", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	// Проврить валидность параметров
	var validate *validator.Validate = validator.New()

	err := validate.Struct(&adv)
	if err != nil {
		response = errorType(http.StatusBadRequest, fmt.Sprintf("%v", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	// Service logic

	// Статус 200 ОК
	w.WriteHeader(http.StatusOK)
	w.Write(responseOK)
}

// GetAll - метод возвращает все объявления
// Аргументы:
//	page: пагинация. int значение, которое указывает на начало пагинации.
//		В качестве значение выступает ID объявления. Т.е., если page = 3,
//		то будут выведены объявления с ID от 3 до 12 включительно.
//	sort: сортировка. Варианты сортировки:
//		- price_asc (по возрастанию цены)
//		- price_desc (по убыванию цены)
//		- date_asc (по дате добавления)
//		- date_desc (по дате добавления в обратном порядке)
// Возвращаемые значения:
//	advertisements: список из
//		- link (ссылка на главное фото)
//		- price (цена за объявление)
//		- name (название объявления)
//	status_code: код результата (200 в случае успеха)
func GetAll(w http.ResponseWriter, r *http.Request) {
	var (
		adv      GetAllRequest
		response []byte
	)

	// Проверить валидность JSON'а
	if err := json.NewDecoder(r.Body).Decode(&adv); err != nil {
		response = errorType(http.StatusBadRequest, fmt.Sprintf("%v", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	// Проврить валидность параметров
	var validate *validator.Validate = validator.New()

	err := validate.Struct(&adv)
	if err != nil {
		response = errorType(http.StatusBadRequest, fmt.Sprintf("%v", err))

		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)

		return
	}

	// Service logic

	// Статус 200 ОК
	w.WriteHeader(http.StatusOK)
	w.Write(responseOK)
}

// NotFound вызывается, если путь не существует
func NotFound(w http.ResponseWriter, r *http.Request) {
	response := errorType(http.StatusNotFound, "not found")

	w.WriteHeader(http.StatusNotFound)
	w.Write(response)
}
