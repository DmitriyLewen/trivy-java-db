package types

const (
	// types of files
	JarType = "jar"
	AarType = "aar"
)

type Index struct {
	GroupID    string
	ArtifactID string
	Version    string
	Sha1       string
	Type       string
}
