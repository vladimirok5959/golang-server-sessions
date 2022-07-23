package session

// IsSetBool to check if variable exists
func (s *Session) IsSetBool(name string) bool {
	s.varlist.RLock()
	defer s.varlist.RUnlock()
	if _, ok := s.varlist.Bool[name]; ok {
		return true
	} else {
		return false
	}
}

// GetBool returns stored variable value or default
func (s *Session) GetBool(name string, def bool) bool {
	s.varlist.RLock()
	defer s.varlist.RUnlock()
	if v, ok := s.varlist.Bool[name]; ok {
		return v
	} else {
		return def
	}
}

// SetBool to set variable value
func (s *Session) SetBool(name string, value bool) {
	isset := s.IsSetBool(name)
	s.varlist.Lock()
	defer s.varlist.Unlock()
	s.varlist.Bool[name] = value
	if isset || value {
		s.changed = true
	}
}

// DelBool to remove variable
func (s *Session) DelBool(name string) {
	if s.IsSetBool(name) {
		s.varlist.Lock()
		defer s.varlist.Unlock()
		delete(s.varlist.Bool, name)
		s.changed = true
	}
}
