package session

// IsSetInt to check if variable exists
func (s *Session) IsSetInt(name string) bool {
	s.varlist.RLock()
	defer s.varlist.RUnlock()
	if _, ok := s.varlist.Int[name]; ok {
		return true
	} else {
		return false
	}
}

// GetInt returns stored variable value or default
func (s *Session) GetInt(name string, def int) int {
	s.varlist.RLock()
	defer s.varlist.RUnlock()
	if v, ok := s.varlist.Int[name]; ok {
		return v
	} else {
		return def
	}
}

// SetInt to set variable value
func (s *Session) SetInt(name string, value int) {
	isset := s.IsSetInt(name)
	s.varlist.Lock()
	defer s.varlist.Unlock()
	s.varlist.Int[name] = value
	if isset || value != 0 {
		s.changed = true
	}
}

// DelInt to remove variable
func (s *Session) DelInt(name string) {
	if s.IsSetInt(name) {
		s.varlist.Lock()
		defer s.varlist.Unlock()
		delete(s.varlist.Int, name)
		s.changed = true
	}
}
