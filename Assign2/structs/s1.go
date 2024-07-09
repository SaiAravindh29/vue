package structs

type Personal struct {
	FatherName string `json:"father"`
	MotherName string `json:"mother"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
}

type Address struct {
	StudentAddress string `json:"studentaddress"`
	Pincode        string `json:"pincode"`
	Landmark       string `json:"landmark"`
}

type School10th struct {
	SchoolName string `json:"schoolname"`
	Place      string `json:"place"`
	Pincode    string `json:"pincode"`
	Address    string `json:"address"`
	Type       string `json:"type"`
}
type School12th struct {
	SchoolName string `json:"schoolname"`
	Place      string `json:"place"`
	Pincode    string `json:"pincode"`
	Address    string `json:"address"`
	Type       string `json:"type"`
}

type School struct {
	S10 *School10th `json:"s10"`
	S12 *School12th `json:"s12"`
}

type Mark10th struct {
	English int `json:"english"`
	Tamil   int `json:"tamil"`
	Maths   int `json:"maths"`
	Science int `json:"science"`
	Social  int `json:"social"`
}

type Mark12th struct {
	English   int `json:"english"`
	Tamil     int `json:"tamil"`
	Maths     int `json:"maths"`
	Biology   int `json:"biology"`
	Chemistry int `json:"chemistry"`
}

type Marks struct {
	M10 *Mark10th `json:"m10"`
	M12 *Mark12th `json:"m12"`
}
