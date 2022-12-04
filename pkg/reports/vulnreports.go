package reports

import (
	"context"
	"log"

	"github.com/mkm29/vulncron/api/v1alpha1"
	"github.com/mkm29/vulncron/pkg/utils"
	"k8s.io/client-go/rest"
)

func GetVulnerabilityReportList(client *rest.RESTClient) (error, v1alpha1.VulnerabilityReportList) {
	var vrl v1alpha1.VulnerabilityReportList
	err := client.Get().
		Resource("vulnerabilityreports").
		Do(context.Background()).
		Into(&vrl)
	if err != nil {
		return err, vrl
	}
	log.Printf("got %d vulnerability reports", len(vrl.Items))
	return nil, vrl
}

func GetReportSummaries(vrl v1alpha1.VulnerabilityReportList) (Summary, []VulnerabilitySummary) {
	// create summaries
	summary := Summary{}
	summaries := []VulnerabilitySummary{}
	for _, vr := range vrl.Items {
		// get VulnerabilitySummary
		// get Namespace from meta
		vs := VulnerabilitySummary{
			Namespace:  vr.ObjectMeta.Namespace,
			Repository: utils.GetRepositoryName(vr.Report.Artifact.Repository),
			Tag:        vr.Report.Artifact.Tag,
			Critical:   vr.Report.Summary.CriticalCount,
			High:       vr.Report.Summary.HighCount,
			Medium:     vr.Report.Summary.MediumCount,
			Low:        vr.Report.Summary.LowCount,
		}
		summary.TotalCritical += vr.Report.Summary.CriticalCount
		summary.TotalHigh += vr.Report.Summary.HighCount
		summary.TotalMedium += vr.Report.Summary.MediumCount
		summary.TotalLow += vr.Report.Summary.LowCount
		summary.TotalUnknown += vr.Report.Summary.UnknownCount
		summary.TotalNone += vr.Report.Summary.NoneCount
		summaries = append(summaries, vs)
	}
	return summary, summaries
}
