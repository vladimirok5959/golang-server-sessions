package session

// IsSetInt64 to check if variable exists
func (s *Session) IsSetInt64(name string) bool {
	if _, ok := s.varlist.Int64[name]; ok {
		return true
	} else {
		return false
	}
}

// GetInt64 returns stored variable value or default
func (s *Session) GetInt64(name string, def int64) int64 {
	if v, ok := s.varlist.Int64[name]; ok {
		return v
	} else {
		return def
	}
}

// SetInt64 to set variable value
func (s *Session) SetInt64(name string, value int64) {
	isset := s.IsSetInt64(name)
	s.varlist.Int64[name] = value
	if isset || value != 0 {
		s.changed = true
	}
}

// DelInt64 to remove variable
func (s *Session) DelInt64(name string) {
	if s.IsSetInt64(name) {
		delete(s.varlist.Int64, name)
		s.changed = true
	}
}
