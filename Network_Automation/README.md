# Go scripts for network automation
- I'm currently working through examples in a book by Tom McAllen
- I've tried to deviate a little from the scripts book to test my comprehension


Resources used:
- Go Programming for Network Operations: A Golang Network Automation Handbook by Tom McAllen
- Introducing Go
- An Introduction to Programming in Go
- Go bootcamp
- Effective Go


usingTemplates/router_template_script.go 
- can be used to create router configs from go template files. 
- variables for the template can be found in the host_var.yaml file

```
Todds-MBP-3:useful_scripts toddriemenschneider$ go run router_template_script.go 

hostname cmh1-rtr1
ip name-server 8.8.4.4

router 65001
router-id 1.1.1.1
neighbor 2.2.2.2 remote-as 65002
neighbor 3.3.3.3 remote-as 65003


ntp source-interface Loopback0
ntp server 11.11.11.11 prefer
ntp server 22.22.22.22
```
