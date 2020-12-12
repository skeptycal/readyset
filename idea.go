package readyset

import (
	"errors"

	log "github.com/sirupsen/logrus"

	"golang.org/x/text/language"
	_ "golang.org/x/text/language"
)

var (
	Lang language.Tag
)

type readySettings struct {
	// - GNU compatible flags
	//
	// - Import flags
	//
	// - Create flags (something like modeling with clay ...)
	//
	// - String based flags specification
	//
	gnuflags string

	// ansi colors
	//
	// - strings.Buffer
	//
	// - automatic replacement of "color tags"
	//
	// - what is the best template tag??
	//
	// - convert to / from all common color formats
	//
	colors string

	// strings
	//
	// - parsers
	//
	// - markdown
	//
	// - pretty formatters
	//
	// - translators (python to go would be fun)
	//
	strings string

	// language
	//
	// todo - use golang.org/x/text ??
	// quote from source
	//
	// text is a repository of text-related packages related to internationalization
	// (i18n) and localization (l10n), such as character encodings, text
	// transformations, and locale-specific text handling.
	//
	// There is a 30 minute video, recorded on 2017-11-30, on the "State of
	// golang.org/x/text" at https://www.youtube.com/watch?v=uYrDrMEGu58
	//
	// reference: https://phrase.com/blog/posts/internationalization-i18n-go/
	//
	language string

	// data is configuration information for the data package.
	//
	// mysql database connection information
	//
	// - parser
	//
	// - inferStructure
	//
	// - cleanDataSet
	//
	// - translate
	//
	// - io.Reader implementations
	//
	// - io.Writer implementations
	//
	data string

	// math is the configuration information for the math package
	//
	// - data analysis
	//
	// - statistics
	//
	// - boolean algebra
	//
	// - linear algebra
	//
	// - vector analysis
	//
	// - predictions
	//
	// - solver

	// AI is the configuration information for the AI package
	//
	// - pattern recognition
	//
	// - prediction engine
	//
	// - identifier
	//
	// - ignore engine (maximize ignoring to reduce scope and increase precision)
	//
	ai string

	// regex is the configuration information, constants and patterns
	//
	// - collection of standard compiled regex patterns
	//
	// - simple regex builder (like lego bricks)
	//
	regex string

	// io - stdout, stderr, stdin functionality
	//
	// - colorized output
	//
	// - tee output
	//
	// - multiwriter
	//
	// - jsonwriter
	//
	// - toml, yaml, CSV
	//
	// - SQL writer
	//
	// - ghost logger
	//
	io string

	// logger is the configuration information for the logger package
	//
	// using mainly logrus
	//
	// - config file
	//
	// - io.writer
	//
	// - io.reader
	//
	// - regexLogReader("xxxxxx")
	//
	// - logFind("xxxxxx")
	//
	// - duplicate to CSV, JSON, SQL, etc.
	//
	logger string

	// shell command wrapper configuration
	//
	shell string

	// repl is a shell REPL written in Go
	//
	repl string
}

func NewReadySettings() (*readySettings, error) {
	s := &readySettings{}
	if s == nil {
		return nil, errors.New("readySettings could not be initialized")
	}
	return s, nil
}

func init() {
	RSG, err := NewReadySettings()
	if err != nil {
		log.Info(err)
	}
	Lang, err = language.Parse("en-US")
	if err != nil {
		log.Info(err)
	}
	log.Info("readySettings are initialized: %v", RSG)
}
