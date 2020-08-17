package listener

import (
	"strings"
)

func stripQuotes(s string) string {
	if s == "" || !strings.Contains(s, "\"") {
		return s
	}
	return s[1 : len(s)-1]
}

func stripArr(s string) string {
	if s == "" || s[0]!='['  || s[len(s)-1]!=']'{
		return s
	}
	return s[1 : len(s)-1]
}

func stripNewLine(s string) string {
	if s == "" || s[0]!='\n' {
		return s
	}
	return s[1 :]
}

func captical(s string) string {
	if s == "" {
		return s
	}
	arr := []rune(s)
	//fmt.Println(s,arr[0],'A','a')
	if arr[0] >= 'A' && arr[0] <= 'Z' {
		return s
	}
	if arr[0] >= 'a' && arr[0] <= 'z' {
		arr[0] = arr[0] - 'a' + 'A'
		return string(arr)
	}
	return "K" + s
}

type AutoGenerated []struct {
	description string
	logs        struct {
		level string
		dir   string
	}
	host string
	//admin[]"parrt"
	//aliases[]
	Spaces struct {
	}
	//keys[]1
	//null_ default
	//bool_ default
	structs []struct {
		a string
		b string
	}
}

type AutoGenerated1 []struct {
	description string
	logs        struct {
		level string
		dir   string
	}
	host  string
	admin []string
	//aliases[]
	Spaces struct {
	}
	keys    []float64
	null_   string
	bool_   string
	structs []struct {
		a string
		b string
	}
}

type AutoGenerated3 []struct {
	description string
	logs        struct {
		level string
		dir   string
	}
	host    string
	admin   []string
	aliases []interface{}
	Spaces  struct {
	}
	keys    []float64
	null_   string
	bool_   string
	structs []struct {
		a string
		b string
	}
}

type AutoGenerated4 []struct {
	Description string
	Logs        struct {
		Level string
		Dir   string
	}
	Host    string
	Admin   []string
	Aliases []interface{}
	Spaces  struct {
	}
	Keys    []float64
	Null_   string
	Bool_   string
	Structs []struct {
		A string
		B string
	}
}

type AutoGenerated5 []struct {
	Description string `json:"description"`
	Logs        struct {
		Level string `json:"level"`
		Dir   string `json:"dir"`
	} `json:"logs"`
	Host    string        `json:"host"`
	Admin   []string      `json:"admin"`
	Aliases []interface{} `json:"aliases"`
	Spaces  struct {
	} `json:"Spaces"`
	Keys    []float64 `json:"keys"`
	Null_   string    `json:"null_"`
	Bool_   string    `json:"bool_"`
	Structs []struct {
		A string `json:"a"`
		B string `json:"b"`
	} `json:"structs"`
}

type AutoGenerated6 []struct {
	Description string `json:"description"`
	Logs        struct {
		Level string `json:"level"`
		Dir   string `json:"dir"`
	} `json:"logs"`
	Host    string        `json:"host"`
	Admin   []string      `json:"admin"`
	Aliases []interface{} `json:"aliases"`
	Spaces  struct {
	} `json:"Spaces"`
	Keys    []float64   `json:"keys"`
	Null_   interface{} `json:"null_"`
	Bool_   bool        `json:"bool_"`
	Structs []struct {
		A string `json:"a"`
		B string `json:"b"`
	} `json:"structs"`
}
