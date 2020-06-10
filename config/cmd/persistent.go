package cmd

// Persistent is a collection of global/persistent flags
var Persistent PersistentFlags

// PersistentFlags represents the persistent flags
type PersistentFlags struct {
	Path        string
	Endpoints   []string
	Concurrency int
}
