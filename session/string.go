package session

func (s *Session) IsSetString(name string) bool {
	if _, ok := s.varlist.String[name]; ok {
		return true
	} else {
		return false
	}
}

func (s *Session) GetString(name string, def string) string {
	if v, ok := s.varlist.String[name]; ok {
		return v
	} else {
		return def
	}
}

func (s *Session) SetString(name string, value string) {
	isset := s.IsSetString(name)
	s.varlist.String[name] = value
	if isset || value != "" {
		s.changed = true
	}
}
