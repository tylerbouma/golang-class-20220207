package main

import "fmt"

type Virtmach struct {
	ip            string
	hostname      string
	diskgb        int
	ram           int
	status        string
	resourceGroup string
	tags          []Tags
}

type Tags struct {
	tagName string
	tagVal  string
}

type ResourceGroup struct {
	name string
	vms  []Virtmach
}

func (v *Virtmach) dealloc() {
	// deallocate a VM if it is running
	if v.status != "off" {
		fmt.Println("deallocating vm:", v.hostname)
		v.status = "off"
	} else {
		fmt.Println(v.hostname, "is already deallocated")
	}
}

func (v Virtmach) sysinfo() {
	// print VM info
	fmt.Print("\n--------------\n")
	fmt.Println("hostname:", v.hostname)
	fmt.Println("ip:", v.ip)
	fmt.Println("disk space:", v.diskgb, "GB")
	fmt.Println("VM status:", v.status)
	fmt.Println("resource group:", v.resourceGroup)
	for i, c := range v.tags {
		fmt.Printf("tag%v name: %v\n", i, c.tagName)
		fmt.Printf("tag%v value: %v\n", i, c.tagVal)
	}
	fmt.Printf("--------------\n\n")
}

func (v *Virtmach) associate(rg *ResourceGroup) {
	// Associate VM with a group
	v.resourceGroup = rg.name
	rg.vms = append(rg.vms, *v)
}

func main() {
	rg := ResourceGroup{name: "empire", vms: []Virtmach{}}
	v := Virtmach{ip: "10.10.10.0",
		hostname:      "deathstar",
		diskgb:        100,
		ram:           8,
		resourceGroup: "none",
		status:        "running",
		tags:          []Tags{{"priority", "critical"}, {"org", "enterprise"}},
	}

	// mess around with the vm and view changes
	v.sysinfo()
	v.dealloc()
	v.sysinfo()
	v.associate(&rg)
	v.sysinfo()

	// take a look at the resource group we created
	fmt.Println(rg)

}
