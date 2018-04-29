package cmd

import (
    "fmt"
    "os"
    "text/template"
    "github.com/im-auld/waypoint/waypoint/state"
)

const (
    master  = `Versions:{{block "list" .}}{{"\n"}}{{range .}}{{println "- Version: " .SemVer}}{{println "  Date: " .Timestamp}}{{end}}{{end}}`
)

func listAll(versions state.Versions)  {
    masterTmpl, err := template.New("master").Parse(master)
    if err != nil {
        fmt.Errorf(err.Error())
    }
    if err := masterTmpl.Execute(os.Stdout, versions); err != nil {
        fmt.Errorf(err.Error())
    }
}
