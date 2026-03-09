package model

// Finding represents a single issue detected by an analyzer.
type Finding struct {
	AnalyzerName   string
	Severity       Severity
	ServiceName    string   // empty if finding applies to multiple services
	ServiceNames   []string // populated for multi-service findings (e.g., cycles)
	Title          string
	Description    string
	Recommendation string
}
