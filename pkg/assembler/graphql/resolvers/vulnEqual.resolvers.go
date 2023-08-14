package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// IngestVulnEqual is the resolver for the ingestVulnEqual field.
func (r *mutationResolver) IngestVulnEqual(ctx context.Context, vulnerability model.VulnerabilityInputSpec, otherVulnerability model.VulnerabilityInputSpec, vulnEqual model.VulnEqualInputSpec) (*model.VulnEqual, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	return r.Backend.IngestVulnEqual(ctx,
		model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityID)},
		model.VulnerabilityInputSpec{Type: strings.ToLower(otherVulnerability.Type), VulnerabilityID: strings.ToLower(otherVulnerability.VulnerabilityID)},
		vulnEqual)
}

// VulnEqual is the resolver for the vulnEqual field.
func (r *queryResolver) VulnEqual(ctx context.Context, vulnEqualSpec model.VulnEqualSpec) ([]*model.VulnEqual, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	if len(vulnEqualSpec.Vulnerabilities) > 0 {
		var lowercaseVulnFilterList []*model.VulnerabilitySpec
		for _, v := range vulnEqualSpec.Vulnerabilities {
			lowercaseVulnFilter := model.VulnerabilitySpec{
				Type:            toLower(v.Type),
				VulnerabilityID: toLower(v.VulnerabilityID),
			}
			lowercaseVulnFilterList = append(lowercaseVulnFilterList, &lowercaseVulnFilter)
		}

		lowercaseVulnEqualFilter := model.VulnEqualSpec{
			ID:              vulnEqualSpec.ID,
			Vulnerabilities: lowercaseVulnFilterList,
			Justification:   vulnEqualSpec.Justification,
			Origin:          vulnEqualSpec.Origin,
			Collector:       vulnEqualSpec.Collector,
		}
		return r.Backend.VulnEqual(ctx, &lowercaseVulnEqualFilter)
	} else {
		return r.Backend.VulnEqual(ctx, &vulnEqualSpec)
	}
}
