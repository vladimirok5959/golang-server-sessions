package session

func (s *Session) IsSetString(name string) bool {
	if _, ok := s.v.String[name]; ok {
		return true
	} else {
		return false
	}
}

func (s *Session) GetString(name string, def string) string {
	if v, ok := s.v.String[name]; ok {
		return v
	} else {
		return def
	}
}

func (s *Session) SetString(name string, value string) {
	isset := s.IsSetString(name)
	s.v.String[name] = value
	if isset || value != "" {
		s.c = true
	}
}
