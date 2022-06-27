package session

func (s *Session) IsSetInt(name string) bool {
	if _, ok := s.varlist.Int[name]; ok {
		return true
	} else {
		return false
	}
}

func (s *Session) GetInt(name string, def int) int {
	if v, ok := s.varlist.Int[name]; ok {
		return v
	} else {
		return def
	}
}

func (s *Session) SetInt(name string, value int) {
	isset := s.IsSetInt(name)
	s.varlist.Int[name] = value
	if isset || value != 0 {
		s.changed = true
	}
}

func (s *Session) DelInt(name string) {
	if s.IsSetInt(name) {
		delete(s.varlist.Int, name)
		s.changed = true
	}
}
