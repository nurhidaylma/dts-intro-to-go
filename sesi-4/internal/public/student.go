package public

type StudentResponse struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Motivasi  string `json:"motivasi"`
}

type UpdateStudentRequest struct {
	ID        int    `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Motivasi  string `json:"motivasi"`
}

type ListStudentRequest struct {
	Search string `qs:"search"`
	IDs    []int  `json:"ids"`
}
