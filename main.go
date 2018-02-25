package main

import (
	"fmt"	
	"flag" 
	"os"
	"path"
	"io/ioutil" 
	"regexp"
	"strings"
)
 

func main(){ 
	var (
		cmdBuild, cmdCreate bool
		cpos string
	)
	flag.BoolVar(&cmdBuild, "b", false, "build contracts")
	flag.BoolVar(&cmdCreate, "c", false, "create a new project")
	flag.StringVar(&cpos, "cpos", "./contract", "contract path, using './contract' when default")
	
	flag.Parse() 
	fmt.Println(cpos)
	if cmdBuild {
		fmt.Printf("build contract for proj %s\n", cpos)
		buildContracts(cpos);
	} else if cmdCreate {
		fmt.Printf("build contract for proj %s\n", cpos)
		buildContracts(cpos);
	} 
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func isExist(path string) bool {
	_, err := os.Stat(path)
    if err != nil {   
        return os.IsExist(err)  
    }  
    return true  
}

func isDir(path string) bool {
	s, err := os.Stat(path)  
    if err != nil {  
        return false  
    }  
    return s.IsDir() 
}

func buildContracts(cpos string) {
	pthSrc := fmt.Sprintf("%s/src/main.sol", cpos)
	code := readContracts(pthSrc, 0)
	regpragma := regexp.MustCompile(`pragma.+;`) 
	regcontinuesn := regexp.MustCompile(`[\n]+`) 
	imports := regpragma.FindAllString(code, -1)
	 
	if len(imports) <= 0 {
		fmt.Println("must have solidity progma .")
		return
	}

	code = imports[0] + regpragma.ReplaceAllString(code, "") 
	code = regcontinuesn.ReplaceAllString(code, "\n")

	fmt.Println("\n  ======\nCode Generated \n  ======\n" + code)

	pthOut := fmt.Sprintf("%s/out/compiled.sol", cpos)
	ioutil.WriteFile(pthOut, []byte(code), 0644)
}

func readContracts(pth string, dep int) string {
	dat, err := ioutil.ReadFile(pth)
    check(err)
	
	src := string(dat)
	//fmt.Print(src)  

	reg := regexp.MustCompile(`[\s]*import[\s]+\"([^\s]+)\";[\s]*`)
	imports := reg.FindAllString(src, -1)

	for _, porigin := range imports{
		fmt.Printf(".%s.", porigin)
		realpath :=  reg.ReplaceAllString(porigin, "${1}")
		if realpath[0] == '.' {
			realpath = path.Join(path.Dir(pth), realpath)
		} 
		if isExist(realpath) {
			if isDir(realpath) {
				realpath = realpath+"/_.sol" 
			} 
			fmt.Printf("load file : %s \n", realpath)
			src = strings.Replace(src, porigin, "\n// " + strings.Repeat("-- ", dep) + "import from " + realpath + " ======" + readContracts(realpath, dep + 1) + "\n", -1)
		} else {
			fmt.Printf("%s not find.", realpath)
		} 
	}
	return src
}