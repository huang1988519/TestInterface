package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "os"
)

type Page struct {
  Title string
  Body []byte
}
func (p *Page) save() error {
  filename := p.Title + ".json"
  return ioutil.WriteFile(filename, p.Body, 0600)
}
func loadPage(title string) *Page {
  filename := title + ".json"
  body, _ := ioutil.ReadFile(filename)
  return  &Page{Title:title, Body: body}
}


func handler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path[1:]
  switch path {
  case "view":
    handleView(w, r)
  case "upload":
    handleUpload(w, r)
  default:
    http.NotFound(w, r)
    fmt.Fprintf(w, "<h1>找不到%s界面，请联系 管理员 ***</h1>",path)
  }
}
func handleUpload(w http.ResponseWriter, r *http.Request) {
  values := r.URL.Query()
  fmt.Println("%s",values)
  for k, v := range values {
    fmt.Fprintf(w,"<p> %s =  %s </p>",k,v)
  }
}
func handleView(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  p := loadPage(title)

  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>",p.Title, p.Body )
}
func fileIsExistAtPath(path string) (bool, error) {
  _, err := os.Stat(path)
  if err == nil {
    return true ,nil
  }
  if os.IsNotExist(err) {
    createErr := os.MkdirAll(path, 0777)
    if createErr == nil {
      fmt.Println("创建文件夹成功")
    }
    return false ,nil
  }
  return true , err
}
func main() {
  rootPath := "files"
  fileIsExistAtPath(rootPath)
  http.HandleFunc("/",handler)
  http.ListenAndServe(":8080", nil)
}
