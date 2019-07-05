package config

var options *Options

// Current exports available config options for rest of the app to access
func Current() *Options {
	return options
}
