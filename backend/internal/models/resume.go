package models

type ResumeRequest struct {
	Name       string `json:"name"`
	Role       string `json:"role"`
	Skills     string `json:"skills"`
	Experience string `json:"experience"`
}

type ResumeResponse struct {
	ResumeText string `json:"resume_text"`
}
