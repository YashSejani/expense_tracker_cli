package main

import(
	"os"
	"encoding/json"
)

func (em *ExpenseManager) Load() {
	data, err := os.ReadFile(em.Filename)
	if err != nil {
		em.Storage.NextId = 1
	} else {
		json.Unmarshal(data, &em.Storage)
	}

	if em.Storage.NextId == 0 {
        em.Storage.NextId = 1
    }
}

func (em *ExpenseManager) Save() {
	jsonData, _ := json.MarshalIndent(em.Storage, "", "    ")
	os.WriteFile(em.Filename, jsonData, 0644)
}