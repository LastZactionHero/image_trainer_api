package main

// Classification option for Image
type Classification struct {
	ID     int64
	Name   string `json:"name"`
	Hotkey string `json:"hotkey"`
}

// Valid does the classification have all necessary data
func (c Classification) Valid() bool {
	return (len(c.Name) > 0 && len(c.Hotkey) == 1)
}

// FindClassificationByName find a classification matching the name
func FindClassificationByName(name string) *Classification {
	var classification Classification
	db.Where("name = ?", name).First(&classification)
	if classification.ID > 0 {
		return &classification
	}
	return nil
}

// FindClassificationByHotkey find a classification matching the hotkey
func FindClassificationByHotkey(hotkey string) *Classification {
	var classification Classification
	db.Where("hotkey = ?", hotkey).First(&classification)
	if classification.ID > 0 {
		return &classification
	}
	return nil
}
