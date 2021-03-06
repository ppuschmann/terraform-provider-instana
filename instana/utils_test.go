package instana_test

import (
	"testing"

	. "github.com/gessnerfl/terraform-provider-instana/instana"
	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"
	"github.com/stretchr/testify/assert"
)

func TestRandomID(t *testing.T) {
	id := RandomID()

	assert.NotEqual(t, 0, len(id))
}

func TestReadStringArrayParameterFromResourceWhenParameterIsProvided(t *testing.T) {
	ruleIds := []interface{}{"test1", "test2"}
	data := make(map[string]interface{})
	data[CustomEventSpecificationDownstreamIntegrationIds] = ruleIds
	resourceData := NewTestHelper(t).CreateResourceDataForResourceHandle(NewCustomEventSpecificationWithSystemRuleResourceHandle(), data)
	result := ReadStringArrayParameterFromResource(resourceData, CustomEventSpecificationDownstreamIntegrationIds)

	assert.NotNil(t, result)
	assert.Equal(t, []string{"test1", "test2"}, result)
}

func TestReadStringArrayParameterFromResourceWhenParameterIsMissing(t *testing.T) {
	resourceData := NewTestHelper(t).CreateEmptyResourceDataForResourceHandle(NewCustomEventSpecificationWithSystemRuleResourceHandle())
	result := ReadStringArrayParameterFromResource(resourceData, CustomEventSpecificationDownstreamIntegrationIds)

	assert.Nil(t, result)
}

func TestShouldReturnStringRepresentationOfSeverityWarning(t *testing.T) {
	testShouldReturnStringRepresentationOfSeverity(restapi.SeverityWarning, t)
}

func TestShouldReturnStringRepresentationOfSeverityCritical(t *testing.T) {
	testShouldReturnStringRepresentationOfSeverity(restapi.SeverityCritical, t)
}

func testShouldReturnStringRepresentationOfSeverity(severity restapi.Severity, t *testing.T) {
	result, err := ConvertSeverityFromInstanaAPIToTerraformRepresentation(severity.GetAPIRepresentation())

	assert.Nil(t, err)
	assert.Equal(t, severity.GetTerraformRepresentation(), result)
}

func TestShouldFailToConvertStringRepresentationForSeverityWhenIntValueIsNotValid(t *testing.T) {
	result, err := ConvertSeverityFromInstanaAPIToTerraformRepresentation(1)

	assert.NotNil(t, err)
	assert.Equal(t, "INVALID", result)
}

func TestShouldReturnIntRepresentationOfSeverityWarning(t *testing.T) {
	testShouldReturnIntRepresentationOfSeverity(restapi.SeverityWarning, t)
}

func TestShouldReturnIntRepresentationOfSeverityCritical(t *testing.T) {
	testShouldReturnIntRepresentationOfSeverity(restapi.SeverityCritical, t)
}

func testShouldReturnIntRepresentationOfSeverity(severity restapi.Severity, t *testing.T) {
	result, err := ConvertSeverityFromTerraformToInstanaAPIRepresentation(severity.GetTerraformRepresentation())

	assert.Nil(t, err)
	assert.Equal(t, severity.GetAPIRepresentation(), result)
}

func TestShouldFailToConvertIntRepresentationForSeverityWhenStringValueIsNotValid(t *testing.T) {
	result, err := ConvertSeverityFromTerraformToInstanaAPIRepresentation("foo")

	assert.NotNil(t, err)
	assert.Equal(t, -1, result)
}
