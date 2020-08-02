package structs

type Match struct {
	IsDone     bool
	afterCount int
	matched    string
	before     []string
	after      []string
}

func (m *Match) AddAfter(s string) {
	if len(m.after) < m.afterCount {
		m.after = append(m.after, s)
		if len(m.after) == m.afterCount {
			m.IsDone = true
		}
	}
}

func (m *Match) Result() []string {
	result := m.before
	result = append(result, m.matched)
	result = append(result, m.after...)
	return result
}

func NewMatchesStorage(afterCount int) *Matches {
	m := &Matches{
		afterCount: afterCount,
		m:          make([]*Match, 0),
	}
	return m
}

type Matches struct {
	afterCount int
	m          []*Match
}

func (m *Matches) newMatch(s string, before []string) {
	m.m = append(m.m, &Match{
		afterCount: m.afterCount,
		matched:    s,
		before:     before,
		after:      make([]string, 0, m.afterCount),
	})
}

func (m *Matches) Result() []*Match {

	result := make([]*Match, 0)
	for _, match := range m.m {
		if match.IsDone {
			result = append(result, match)
		}
	}
	m.cleanDone()
	return result
}

func (m *Matches) cleanDone() {
	lastDoneIndex := -1
	// т.к. мы идем по порядку и не может быть ситуации, когда текущая задача выполнена,
	// а предыдущая - нет, мы можем смело удалять все выполненные задачи до текущего индекса
	for i, match := range m.m {
		if match.IsDone {
			lastDoneIndex = i
		}
	}
	if lastDoneIndex >= 0 {
		if lastDoneIndex == len(m.m)-1 {
			m.m = make([]*Match, 0)
		} else {
			m.m = m.m[lastDoneIndex+1:]
		}
	}
}

func (m *Matches) Process(s string, isMatched bool, before []string) []*Match {
	for _, match := range m.m {
		match.AddAfter(s)
	}
	if isMatched {
		m.newMatch(s, before)
	}

	return m.Result()
}

func (m *Matches) Finish() []*Match {
	for _, match := range m.m {
		match.IsDone = true
	}
	return m.Result()
}
