package analysis

import "log"

type State struct {
	logger    *log.Logger
	Documents map[string]string
}

func NewState(logger *log.Logger) State {
	return State{
		logger:    logger,
		Documents: make(map[string]string),
	}
}

func (s *State) HasDocument(uri string) bool {
	_, ok := s.Documents[uri]
	return ok
}

func (s *State) OpenDocument(uri string, content string) {
	s.Documents[uri] = content
}

func (s *State) UpdateDocument(uri string, content string) {
	s.Documents[uri] = content
}
