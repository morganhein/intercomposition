package rpm

import "fmt"

type GCCFix struct {
	*RPM
}

func (p *GCCFix) HasThis() bool {
	fmt.Printf("Builder : %v\n", p.BaseBuilder.HasThis())
	//access to struct data
	get := p.BaseBuilder.Getter
	_, err := get.Get("ok")
	if err != nil {
		fmt.Println(err)
	}

	//or struct functions
	err = p.BaseBuilder.DoSomet()
	if err != nil {
		fmt.Println(err)
	}

	//or call super() functions
	p.RPM.HasThis()
	//and then perform our own logic
	fmt.Println("custom logic")

	return false
}
