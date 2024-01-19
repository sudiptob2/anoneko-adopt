package models

type Mysfit struct {
	MysfitId        string `json:"mysfitId"`
	Name            string `json:"name"`
	Species         string `json:"species"`
	Age             int    `json:"age"`
	Description     string `json:"description"`
	GoodEvil        string `json:"goodevil"`
	LawChaos        string `json:"lawchaos"`
	ThumbImageUri   string `json:"thumbImageUri"`
	ProfileImageUri string `json:"profileImageUri"`
}
