package repository

type petData struct {
	Age     int32
	PetName string
}

type Repository struct {
	Data map[string]petData
}

func NewRepository() Repository {

	data := make(map[string]petData)

	data["Piotrek"] = petData{
		Age:     17,
		PetName: "Radek",
	}

	return Repository{
		Data: data,
	}
}
