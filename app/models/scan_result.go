package models

type ScanResult struct {
	ID            uint
	ScanHistoryID uint
	ServerName    string
	Family        string
	Release       string
}
