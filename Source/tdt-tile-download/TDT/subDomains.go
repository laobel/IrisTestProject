package TDT

import "math/rand"

var  subdomains=[] string{"t0","t1","t2","t3","t4","t5","t6","t7"}

func Subdomains() []string {
	return subdomains
}

func RandomSubdomain() string{
	return  subdomains[rand.Intn(len(subdomains))]
}