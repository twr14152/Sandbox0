package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"text/template"
)

type Config struct {
	Base Base `base`
	Bgp  Bgp  `bgp`
	Ntp  Ntp  `ntp`
}
type Base struct {
	Hostname string `hostname`
	Dns_ip   string `dns_ip`
}
type Bgp struct {
	As        string      `as`
	Id        string      `id`
	Neighbors []Neighbors `neighbors`
}
type Neighbors struct {
	Ip string `ip`
	As string `as`
}
type Ntp struct {
	Source    string `source`
	Primary   string `primary`
	Secondary string `secondary`
}

func main() {
	config := Config{}
	data, _ := ioutil.ReadFile("host_var.yaml")
	_ = yaml.Unmarshal(data, &config)
	tmpl := make(map[string]*template.Template)
	files, _ := template.ParseFiles("base.tmpl",
		"bgp.tmpl", "ntp.tmpl")
	tmpl["base.tmpl"] = template.Must(files, nil)
	_ = tmpl["base.tmpl"].ExecuteTemplate(
		os.Stdout, "base", config)
}
