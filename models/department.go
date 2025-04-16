package models

func GetDepartmentNameByID(id string) string {
    departments := map[string]string{
        "1": "QA",
        "2": "Kriyatec",
    }

    if name, exists := departments[id]; exists {
        return name
    }
    return "Unknown"
}
