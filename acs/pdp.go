package acs

import "fmt"

type Data struct {
	AuthnFlow      string
	SceneAALMap    map[string]int32
	ResourceAALMap map[string]int32
}

type Policy func(d *Data) bool

func PolicyA(d *Data) bool {
	return true
}

func PolicyB(d *Data) bool {
	return false
}

func PolicyC(d *Data) bool {
	if len(d.AuthnFlow) > 0 {
		return true
	}
	return false
}

func execute() {
	d := &Data{
		AuthnFlow:      "SMS",
		SceneAALMap:    nil,
		ResourceAALMap: nil,
	}
	policySet := []Policy{PolicyA, PolicyB, PolicyC}
	for _, p := range policySet {
		fmt.Println(p(d))
	}
}
