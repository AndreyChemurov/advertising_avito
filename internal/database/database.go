package database

type database interface {
	Create(string, string, []string, float64) (int, error)
	GetOne(int, bool) (string, float64, string, string, []string)
	GetAll(int, string)
}
