package reports

import (
	"context"
	"log"

	"github.com/mkm29/vulncron/api/v1alpha1"
	"github.com/mkm29/vulncron/pkg/utils"
	"k8s.io/client-go/rest"
)

type vulnSummary map[string]Summary

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

func GetReportSummaries(vrl v1alpha1.VulnerabilityReportList) (map[string]Summary, []VulnerabilitySummary) {
	// create summaries
	// summary := Summary{}
	vss := make(map[string]Summary)
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
		if _, ok := vss[vr.ObjectMeta.Namespace]; !ok {
			// Namespace does not exist in vss so add it
			vss[vr.ObjectMeta.Namespace] = Summary{}
		}
		summary := vss[vr.ObjectMeta.Namespace]
		summary.TotalCritical += vr.Report.Summary.CriticalCount
		summary.TotalHigh += vr.Report.Summary.HighCount
		summary.TotalMedium += vr.Report.Summary.MediumCount
		summary.TotalLow += vr.Report.Summary.LowCount
		summary.TotalUnknown += vr.Report.Summary.UnknownCount
		summary.TotalNone += vr.Report.Summary.NoneCount
		summaries = append(summaries, vs)
	}
	return vss, summaries
}

func VssToHtml(vs []VulnerabilitySummary, cluster string) string {
	var html string = ohtml
	// replace ${title}
	html = strings.Replace(html, "${title}", fmt.Sprintf("Trivy Report: %s", cluster), 1)
	rhtml := ""
	// iterate over vs
	for _, s := range vs {
		rhtml = fmt.Sprintf(`<tr>
  <td>%s</td>
  <td>%s</td>
  <td>%s</td>
  <td>%d</td>
  <td>%d</td>
  <td>%d</td>
  <td>%d</td>
</tr>
`, s.Namespace,
			s.Repository,
			s.Tag,
			s.Critical,
			s.High,
			s.Medium,
			s.Low,
		)
		// append row html to html
		html = html + rhtml
	}

	return fmt.Sprintf("%s%s", html, chtml)
}
