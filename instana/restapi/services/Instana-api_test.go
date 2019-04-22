package services_test

import (
	"testing"

	. "github.com/gessnerfl/terraform-provider-instana/instana/restapi/services"
)

func TestShouldReturnResourcesFromInstanaAPI(t *testing.T) {
	api := NewInstanaAPI("api-token", "endpoint")

	t.Run("Should return CustomEventSpecificationResource instance", func(t *testing.T) {
		customEventSpecificationResource := api.CustomEventSpecifications()
		if customEventSpecificationResource == nil {
			t.Fatal("Expected instance of CustomEventSpecificationResource to be returned")
		}
	})
	t.Run("Should return RuleResource instance", func(t *testing.T) {
		ruleResource := api.Rules()
		if ruleResource == nil {
			t.Fatal("Expected instance of RuleResource to be returned")
		}
	})
	t.Run("Should return RuleBindingResource instance", func(t *testing.T) {
		ruleBindingResource := api.RuleBindings()
		if ruleBindingResource == nil {
			t.Fatal("Expected instance of RuleBindingResource to be returned")
		}
	})
	t.Run("Should return UserRoleResource instance", func(t *testing.T) {
		userRoleResource := api.UserRoles()
		if userRoleResource == nil {
			t.Fatal("Expected instance of UserRoleResource to be returned")
		}
	})
	t.Run("Should return ApplicationConfigResource instance", func(t *testing.T) {
		applicationConfigResource := api.ApplicationConfigs()
		if applicationConfigResource == nil {
			t.Fatal("Expected instance of ApplicationConfigResource to be returned")
		}
	})
}
