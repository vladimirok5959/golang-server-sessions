package session

func (s *Session) IsSetBool(name string) bool {
	if _, ok := s.v.Bool[name]; ok {
		return true
	} else {
		return false
	}
}

func (s *Session) GetBool(name string, def bool) bool {
	if v, ok := s.v.Bool[name]; ok {
		return v
	} else {
		return def
	}
}

func (s *Session) SetBool(name string, value bool) {
	isset := s.IsSetBool(name)
	s.v.Bool[name] = value
	if isset || value {
		s.c = true
	}
}
