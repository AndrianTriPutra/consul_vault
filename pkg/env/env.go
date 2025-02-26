package env

import "time"

type Setting struct {
	Host  string
	Token string
	Path  string
}

type (
	Config struct {
		Internal Internal `json:"internal"`
		Interval Interval `json:"interval"`
		Database Database `json:"database"`
		Rabbit   Rabbit   `json:"rabbit"`
	}

	Internal struct {
		Logger string `json:"logger"`
		Core   uint   `json:"core"`
	}
	Interval struct {
		Environment time.Duration `json:"environment"`
		Schedulle   time.Duration `json:"schedulle"`
	}
	Database struct {
		Host string `json:"host"`
		Port uint   `json:"port"`
		Name string `json:"name"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}
	Rabbit struct {
		Host string   `json:"host"`
		Tag  string   `json:"tag"`
		Que  string   `json:"que"`
		Key  []string `json:"key"`
	}
)
