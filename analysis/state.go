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

func (s *State) OpenDocument(uri string, content string) {
	s.logger.Printf("Opened file with content %s", content)
	s.Documents[uri] = content
}

func (s *State) UpdateDocument(uri string, content string) {
	s.logger.Printf("Update file with content %s", content)
	s.Documents[uri] = content
}
