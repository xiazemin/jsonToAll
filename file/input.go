package file
import (
	"io/ioutil"
	"os"
)

func GetJson(name string)[]byte{
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return content
}
