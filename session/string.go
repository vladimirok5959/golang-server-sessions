package session

// IsSetString to check if variable exists
func (s *Session) IsSetString(name string) bool {
	if _, ok := s.varlist.String[name]; ok {
		return true
	} else {
		return false
	}
}

// GetString returns stored variable value or default
func (s *Session) GetString(name string, def string) string {
	if v, ok := s.varlist.String[name]; ok {
		return v
	} else {
		return def
	}
}

// SetString to set variable value
func (s *Session) SetString(name string, value string) {
	isset := s.IsSetString(name)
	s.varlist.String[name] = value
	if isset || value != "" {
		s.changed = true
	}
}

// DelString to remove variable
func (s *Session) DelString(name string) {
	if s.IsSetString(name) {
		delete(s.varlist.String, name)
		s.changed = true
	}
}
