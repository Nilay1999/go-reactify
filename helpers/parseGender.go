package helpers

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

func ParseGender(genderStr string) (Gender, error) {
	switch genderStr {
	case "male":
		return Male, nil
	case "female":
		return Female, nil
	default:
		return "", nil
	}
}
