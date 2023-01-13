package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func main() {
    port := "1323"

    args :=  os.Args
    if len(args) > 1 {
        port = args[1]
    }

    exe, err:=  os.Executable()
    if err != nil {
        fmt.Println(err)
        return
    }
    dir := filepath.Dir(exe)
    path :=  filepath.Join(dir, "res.json")
    _, err =  os.Stat(path)
    if err != nil {
        fmt.Println("response file not found :", path)
        return
    } else {
        fmt.Println(path)
    }

    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err)
        return
    }

	e := echo.New()
	e.POST("/", func(c echo.Context) error {
        return c.Blob(http.StatusOK, echo.MIMEApplicationJSONCharsetUTF8, bytes)
	})
	e.Logger.Fatal(e.Start(":" + port))
}
