package main

import (
	"os"
       "fmt"
       "io/ioutil"
//        "reflect"
	"encoding/json"
        "simplesurveygo/dao"
	//"sync"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

type MovieDataField struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Director string  `json:"director"`
	Cast string	`json:"cast"`
	Genre string    `json:"genre"`
	Notes string    `json:"notes"`

}

/*
func getData(startIndex,endIndex int,ldata []MovieDataField)[]MovieDataField{
	return ldata[startIndex:endIndex]

}
*/


func  f(lmoviedata []MovieDataField,name string,done chan bool){

        fmt.Println("in go return")
	session := dao.MgoSession.Clone()
	clctn := session.DB("movies").C("movietable")
	defer session.Close()
	for _,data1:= range lmoviedata{	
	p := &MovieDataField{

		Title  : data1.Title,
		Year   : data1.Year,
		Director : data1.Director,
		Cast :  data1.Cast,
		Genre : data1.Genre,
		Notes : data1.Notes,

	}
	err1 := clctn.Insert(p)
	check(err1)
       }
}

func main() {


   done :=make (chan bool)
   if(len(os.Args) > 1){
	   dat, err := ioutil.ReadFile(os.Args[1])
           check(err)
           fmt.Println(os.Args[0:1])
	   var lData []MovieDataField
	   err1 := json.Unmarshal(dat,&lData)
	   check(err1)

	   toNoThread := len(lData)

	   p := toNoThread / 4
            fmt.Println(p) 
            fmt.Println(lData[0].Year) 

	    go f(lData[0:p],"1",done)
	   go f(lData[p:p*2],"2",done)
	    go f(lData[p*2:p*3],"3",done)
	    go f(lData[p*3:toNoThread],"4",done)
	    <-done
  }
//           return 

    fmt.Println("please specify the movie data file")
}


