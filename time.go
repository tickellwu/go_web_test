package main

import (
    "fmt"
    "strconv"
    "net/http"
    "text/template"
)



func myWeb(w http.ResponseWriter, r *http.Request) {
    fmt.Println("URL:")
    var sCompany string
    for k, v := range r.URL.Query() {
        if k == "company" {
            sCompany = v[0]
        }
        fmt.Println("key:", k, ", value:", v[0])
    }

    var mCompanyData map[string]string;
    mCompanyData = make(map[string]string)
    mCompanyData["huya"] = "2022-01-28T18:30"
    mCompanyData["bigo"] = "2022-01-29T18:30"
    mCompanyData["other"] = "2022-01-30T18:00"

    var sOffDate string
    sOffDate, ok := mCompanyData[sCompany]
    if (!ok)    {
        sOffDate = "2022-01-30T18:00"
    }
    fmt.Println(sOffDate)
    /*
    fmt.Println("PostForm:")
    for k, v := range r.PostForm {
        fmt.Println("key:", k, ", value:", v[0])
    }*/


    t, _ := template.ParseFiles("./templates/template.xml")

    data := map[string]string{
        "off_work_date": sOffDate,
    }

    t.Execute(w, data)
}

func main() {
    var port int = 80
    var sPort string = strconv.Itoa(port)
    http.HandleFunc("/", myWeb)

    http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

    fmt.Println("服务器即将开启，访问地址 http://localhost:" + sPort)

    err := http.ListenAndServe(":" + sPort, nil)
    if err != nil {
        fmt.Println("服务器开启错误: ", err)
    }
}
