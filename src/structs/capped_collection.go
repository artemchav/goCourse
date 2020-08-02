package structs

// В этой коллекции мы будем хранить ограниченное количество элементов
// (текущая строка и то количество строк до нее, которое указано при вызове
// программы во флаге -A )
//
// Как только слайс заполнен, мы переходим к первому элементу слайса и
// перезаписываем значения сверху вних с каждой добавленной строкой.
//
// таким образом мы добиваемся того, что мы храним текущую строку и
// указанное количество строк до нее, что позволяет выводить строки
// "до" и "после" при поиске строки за один проход по файлу

type CappedStringCollection struct {
	currentIndex int
	limit        int
	storage      []string
}

func NewCappedStringCollection(cap int) *CappedStringCollection {
	return &CappedStringCollection{
		limit:   cap,
		storage: make([]string, 0, cap),
	}
}
func (c *CappedStringCollection) increaseCurrentIndex() {
	c.currentIndex++
	if c.currentIndex == c.limit {
		c.currentIndex = 0
	}
}

func (c *CappedStringCollection) AddString(s string) []string {
	if len(c.storage) > 0 {
		c.increaseCurrentIndex()
	}

	if len(c.storage) < c.limit {
		c.storage = append(c.storage, s)
	} else {
		c.storage[c.currentIndex] = s
	}

	return c.GetSavedStrings()
}

// Последняя добавленная строка находится по адресу c.currentIndex,
// то есть "где-то" ) при заполненном слайсе самые "старые" записи находятся под
// текущей позицией, и "молодеют" вниз. Далее они молодеют от начала слайса до текущей
// строки включительно.
func (c *CappedStringCollection) GetSavedStrings() []string {

	result := make([]string, 0, c.limit)
	if len(c.storage) > 0 {
		if c.currentIndex+1 < c.limit && len(c.storage) == c.limit {
			result = append(result, c.storage[c.currentIndex+1:]...)
		}

		result = append(result, c.storage[:c.currentIndex+1]...)
	}
	return result
}
