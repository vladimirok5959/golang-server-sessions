package session

func (s *Session) IsSetInt64(name string) bool {
	if _, ok := s.varlist.Int64[name]; ok {
		return true
	} else {
		return false
	}
}

func (s *Session) GetInt64(name string, def int64) int64 {
	if v, ok := s.varlist.Int64[name]; ok {
		return v
	} else {
		return def
	}
}

func (s *Session) SetInt64(name string, value int64) {
	isset := s.IsSetInt64(name)
	s.varlist.Int64[name] = value
	if isset || value != 0 {
		s.changed = true
	}
}
