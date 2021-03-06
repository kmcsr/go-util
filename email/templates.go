
package kpnmmail

import (
	bytes "bytes"
	// texttemp "text/template"
	htmltemp "html/template"
	ioutil "io/ioutil"

	filex "github.com/kmcsr/go-util/file"
)

var (
	// texttp *texttemp.Template
	htmltp *htmltemp.Template = htmltemp.New("template_html")
)

func LoadHtmlFiles(paths ...string)(err error){
	_, err = htmltp.ParseFiles(paths...)
	return err
}

func ExeHtmlTemp(path string, _value... interface{})(text string, err error){
	var value interface{} = nil
	if len(_value) > 0 {
		value = _value[0]
	}

	buf := bytes.NewBuffer([]byte{})
	err = htmltp.ExecuteTemplate(buf, path, value)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func LoadTemplateDir(dirpath string){
	{ // load mail template files
		var templateFiles []string = make([]string, 0)
		basePath := filex.AbsPath(dirpath)
		var findFunc func(path string)
		findFunc = func(path string){
			finfos, err := ioutil.ReadDir(path)
			if err != nil {
				panic(err)
			}
			for _, info := range finfos {
				fpath := filex.JoinPath(path, info.Name())
				if info.IsDir() {
					findFunc(fpath)
				}else{
					templateFiles = append(templateFiles, fpath)
				}
			}
		}
		findFunc(basePath)

		if len(templateFiles) > 0 {
			LoadHtmlFiles(templateFiles...)
		}
	}
}

func init(){
	htmltp.Funcs(htmltemp.FuncMap{
		"odd": func(num int)(bool){ return num % 2 == 0 },
	})
}

