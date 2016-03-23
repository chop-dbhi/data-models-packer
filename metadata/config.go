package metadata

// Config holds the configuration values for metadata creation. The fields correspond exactly with the command line flags.
type Config struct {
	DataDirPath  string
	DataVersion  string
	Etl          string
	Model        string
	ModelVersion string
	Service      string
	Site         string
}
