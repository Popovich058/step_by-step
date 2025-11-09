package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("data parsing error '%s': %v", data, err)
			continue 
	}

	// Формируем и выводим информацию
	info, err := dp.ActionInfo()
	if err != nil {
		log.Printf("error generating information for the data '%s': %v", data, err)
		continue 
	}

		// Выводим результат
	log.Println(info)
	}
}
