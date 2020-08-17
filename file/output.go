package file

import (
	"os"
	"os/exec"
)

func PutGoLang(name ,content string)error {
	f, err := os.Open(name)
	if err != nil {
		f,err=os.Create(name)
	}
	defer f.Close()
	f.WriteString(content)
	cmd:= exec.Command("gofmt", "-w", f.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err = cmd.Run(); err != nil {
		return err
	}
	return  nil
}

func PutJson(name ,content string) error{
	f, err := os.Open(name)
	if err != nil {
		f,err=os.Create(name)
	}
	defer f.Close()
	f.WriteString(content)
	return nil
}

func PutIdl(name string,subStructs []string,content string)error{
	f,err:=os.Open(name)
	if err!=nil{
		f,err=os.Create(name)
	}
	defer f.Close()
	c:="namespace php json.to.idl \n"
	for _,s:=range subStructs{
		c+=s
	}
	f.WriteString(c+content)
	//fmt.Println(c,content)
	return err
}
