package reports

type VulnerabilitySummary struct {
	Namespace  string
	Name       string
	Repository string
	Tag        string
	Critical   int
	High       int
	Medium     int
	Low        int
	Unknown    int
	None       int
}

type Summary struct {
	TotalCritical int
	TotalHigh     int
	TotalMedium   int
	TotalLow      int
	TotalUnknown  int
	TotalNone     int
}
