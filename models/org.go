package models

func GetOrganizationNameByID(id string) string {
    organizations := map[string]string{
        "1": "Magnolia Community Health",
        "2": "Kriyatec",
    }

    if name, exists := organizations[id]; exists {
        return name
    }
    return "Unknown"
}
