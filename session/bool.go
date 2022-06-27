package session

func (s *Session) IsSetBool(name string) bool {
	if _, ok := s.varlist.Bool[name]; ok {
		return true
	} else {
		return false
	}
}

func (s *Session) GetBool(name string, def bool) bool {
	if v, ok := s.varlist.Bool[name]; ok {
		return v
	} else {
		return def
	}
}

func (s *Session) SetBool(name string, value bool) {
	isset := s.IsSetBool(name)
	s.varlist.Bool[name] = value
	if isset || value {
		s.changed = true
	}
}

func (s *Session) DelBool(name string) {
	if s.IsSetBool(name) {
		delete(s.varlist.Bool, name)
		s.changed = true
	}
}
