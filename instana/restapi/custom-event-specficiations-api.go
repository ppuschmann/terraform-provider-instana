package restapi

import "errors"

const (
	//EventSpecificationBasePath path to Event Specification settings of Instana RESTful API
	EventSpecificationBasePath = EventSettingsBasePath + "/event-specifications"
	//CustomEventSpecificationResourcePath path to Custom Event Specification settings resource of Instana RESTful API
	CustomEventSpecificationResourcePath = EventSpecificationBasePath + "/custom"
)

//Severity representation of the severity in both worlds Instana API and Terraform Provider
type Severity struct {
	apiRepresentation       int
	terraformRepresentation string
}

//GetAPIRepresentation returns the integer representation of the Instana API
func (s Severity) GetAPIRepresentation() int { return s.apiRepresentation }

//GetTerraformRepresentation returns the string representation of the Terraform Provider
func (s Severity) GetTerraformRepresentation() string { return s.terraformRepresentation }

//SeverityCritical representation of the critical severity
var SeverityCritical = Severity{apiRepresentation: 10, terraformRepresentation: "critical"}

//SeverityWarning representation of the warning severity
var SeverityWarning = Severity{apiRepresentation: 5, terraformRepresentation: "warning"}

//RuleType custom type representing the type of the custom event specification rule
type RuleType string

const (
	//SystemRuleType const for RuleType of System
	SystemRuleType = "system"
	//ThresholdRuleType const for RuleType of Threshold
	ThresholdRuleType = "threshold"
	//EntityVerificationRuleType const for RuleType of Entity Verification
	EntityVerificationRuleType = "entity_verification"
)

//AggregationType custom type representing an aggregation of a custom event specification rule
type AggregationType string

//AggregationTypes custom type representing a slice of AggregationType
type AggregationTypes []AggregationType

//ToStringSlice Returns the string representations fo the aggregations
func (types AggregationTypes) ToStringSlice() []string {
	result := make([]string, len(types))
	for i, v := range types {
		result[i] = string(v)
	}
	return result
}

const (
	//AggregationSum const for a sum aggregation
	AggregationSum = AggregationType("sum")
	//AggregationAvg const for a avg aggregation
	AggregationAvg = AggregationType("avg")
	//AggregationMin const for a min aggregation
	AggregationMin = AggregationType("min")
	//AggregationMax const for a max aggregation
	AggregationMax = AggregationType("max")
)

//SupportedAggregationTypes slice of supported aggregation types
var SupportedAggregationTypes = AggregationTypes{AggregationSum, AggregationAvg, AggregationMin, AggregationMax}

//IsSupportedAggregationType check if the provided aggregation type is supported
func IsSupportedAggregationType(aggregation AggregationType) bool {
	for _, v := range SupportedAggregationTypes {
		if v == aggregation {
			return true
		}
	}
	return false
}

//ConditionOperatorType custom type representing a condition operator of a custom event specification rule
type ConditionOperatorType string

//ConditionOperatorTypes custom type representing a slice of ConditionOperatorType
type ConditionOperatorTypes []ConditionOperatorType

//ToStringSlice Returns the string representations fo the condition operators
func (types ConditionOperatorTypes) ToStringSlice() []string {
	result := make([]string, len(types))
	for i, v := range types {
		result[i] = string(v)
	}
	return result
}

const (
	//ConditionOperatorEquals const for a equals (==) condition operator
	ConditionOperatorEquals = ConditionOperatorType("==")
	//ConditionOperatorNotEqual const for a not equal (!=) condition operator
	ConditionOperatorNotEqual = ConditionOperatorType("!=")
	//ConditionOperatorLessThan const for a less than (<) condition operator
	ConditionOperatorLessThan = ConditionOperatorType("<")
	//ConditionOperatorLessThanOrEqual const for a less than or equal (<=) condition operator
	ConditionOperatorLessThanOrEqual = ConditionOperatorType("<=")
	//ConditionOperatorGreaterThan const for a greater than (>) condition operator
	ConditionOperatorGreaterThan = ConditionOperatorType(">")
	//ConditionOperatorGreaterThanOrEqual const for a greater than or equal (<=) condition operator
	ConditionOperatorGreaterThanOrEqual = ConditionOperatorType(">=")
)

//SupportedConditionOperatorTypes slice of supported aggregation types
var SupportedConditionOperatorTypes = ConditionOperatorTypes{ConditionOperatorEquals, ConditionOperatorNotEqual, ConditionOperatorLessThan, ConditionOperatorLessThanOrEqual, ConditionOperatorGreaterThan, ConditionOperatorGreaterThanOrEqual}

//IsSupportedConditionOperatorType check if the provided condition operator type is supported
func IsSupportedConditionOperatorType(operator ConditionOperatorType) bool {
	for _, v := range SupportedConditionOperatorTypes {
		if v == operator {
			return true
		}
	}
	return false
}

//MatchingOperatorType custom type representing a matching operator of a custom event specification rule
type MatchingOperatorType string

//MatchingOperatorTypes custom type representing a slice of MatchingOperatorType
type MatchingOperatorTypes []MatchingOperatorType

//ToStringSlice Returns the string representations fo the matching operators
func (types MatchingOperatorTypes) ToStringSlice() []string {
	result := make([]string, len(types))
	for i, v := range types {
		result[i] = string(v)
	}
	return result
}

const (
	//MatchingOperatorIs const for IS condition operator
	MatchingOperatorIs = MatchingOperatorType("is")
	//MatchingOperatorContains const for CONTAINS condition operator
	MatchingOperatorContains = MatchingOperatorType("contains")
	//MatchingOperatorStartsWith const for STARTS_WITH condition operator
	MatchingOperatorStartsWith = MatchingOperatorType("starts_with")
	//MatchingOperatorEndsWith const for ENDS_WITH condition operator
	MatchingOperatorEndsWith = MatchingOperatorType("ends_with")
	//MatchingOperatorNone const for NONE condition operator
	MatchingOperatorNone = MatchingOperatorType("none")
)

//SupportedMatchingOperatorTypes slice of supported matching operatorTypes types
var SupportedMatchingOperatorTypes = MatchingOperatorTypes{MatchingOperatorIs, MatchingOperatorContains, MatchingOperatorStartsWith, MatchingOperatorEndsWith, MatchingOperatorNone}

//IsSupportedMatchingOperatorType check if the provided matching operator type is supported
func IsSupportedMatchingOperatorType(operator MatchingOperatorType) bool {
	for _, v := range SupportedMatchingOperatorTypes {
		if v == operator {
			return true
		}
	}
	return false
}

//NewSystemRuleSpecification creates a new instance of System Rule
func NewSystemRuleSpecification(systemRuleID string, severity int) RuleSpecification {
	return RuleSpecification{
		DType:        SystemRuleType,
		SystemRuleID: &systemRuleID,
		Severity:     severity,
	}
}

//NewEntityVerificationRuleSpecification creates a new instance of Entity Verification Rule
func NewEntityVerificationRuleSpecification(matchingEntityLabel string, matchingEntityType string, matchingOperator MatchingOperatorType, offlineDuration int, severity int) RuleSpecification {
	return RuleSpecification{
		DType:               EntityVerificationRuleType,
		MatchingEntityLabel: &matchingEntityLabel,
		MatchingEntityType:  &matchingEntityType,
		MatchingOperator:    &matchingOperator,
		OfflineDuration:     &offlineDuration,
		Severity:            severity,
	}
}

//RuleSpecification representation of a rule specification for a CustomEventSpecification
type RuleSpecification struct {
	//Common Fields
	DType    RuleType `json:"ruleType"`
	Severity int      `json:"severity"`

	//System Rule fields
	SystemRuleID *string `json:"systemRuleId"`

	//Threshold Rule fields
	MetricName        *string                `json:"metricName"`
	Rollup            *int                   `json:"rollup"`
	Window            *int                   `json:"window"`
	Aggregation       *AggregationType       `json:"aggregation"`
	ConditionOperator *ConditionOperatorType `json:"conditionOperator"`
	ConditionValue    *float64               `json:"conditionValue"`

	//Entity Verification Rule
	MatchingEntityType  *string               `json:"matchingEntityType"`
	MatchingOperator    *MatchingOperatorType `json:"matchingOperator"`
	MatchingEntityLabel *string               `json:"matchingEntityLabel"`
	OfflineDuration     *int                  `json:"offlineDuration"`
}

//Validate Rule interface implementation for SystemRule
func (r *RuleSpecification) Validate() error {
	if len(r.DType) == 0 {
		return errors.New("type of rule is missing")
	}
	if r.DType == SystemRuleType {
		return r.validateSystemRule()
	} else if r.DType == ThresholdRuleType {
		return r.validateThresholdRule()
	} else if r.DType == EntityVerificationRuleType {
		return r.validateEntityVerificationRule()
	}
	return errors.New("Unsupported rule type " + string(r.DType))
}

func (r *RuleSpecification) validateSystemRule() error {
	if r.SystemRuleID == nil || len(*r.SystemRuleID) == 0 {
		return errors.New("id of system rule is missing")
	}
	return nil
}

func (r *RuleSpecification) validateThresholdRule() error {
	if r.MetricName == nil || len(*r.MetricName) == 0 {
		return errors.New("metric name of threshold rule is missing")
	}
	if (r.Window == nil && r.Rollup == nil) || (r.Window != nil && r.Rollup != nil && *r.Window == 0 && *r.Rollup == 0) {
		return errors.New("either rollup or window and condition must be defined")
	}

	if r.Window != nil && (r.Aggregation == nil || !IsSupportedAggregationType(*r.Aggregation)) {
		return errors.New("aggregation type of threshold rule is mission or not valid")
	}

	if r.ConditionOperator == nil || !IsSupportedConditionOperatorType(*r.ConditionOperator) {
		return errors.New("condition operator of threshold rule is missing or not valid")
	}

	return nil
}

func (r *RuleSpecification) validateEntityVerificationRule() error {
	if r.MatchingEntityLabel == nil || len(*r.MatchingEntityLabel) == 0 {
		return errors.New("matching entity label of entity verification rule is missing")
	}
	if r.MatchingEntityType == nil || len(*r.MatchingEntityType) == 0 {
		return errors.New("matching entity type of entity verification rule is missing")
	}
	if r.MatchingOperator == nil || !IsSupportedMatchingOperatorType(*r.MatchingOperator) {
		return errors.New("matching operator of entity verification rule is missing or not valid")
	}
	if r.OfflineDuration == nil {
		return errors.New("offline duration of entity verification rule is missing")
	}
	return nil
}

//EventSpecificationDownstream definition of downstream reporting for the event specification
type EventSpecificationDownstream struct {
	IntegrationIds                []string `json:"integrationIds"`
	BroadcastToAllAlertingConfigs bool     `json:"broadcastToAllAlertingConfigs"`
}

//Validate validates the consitency of an EventSpecificationDownstream
func (d EventSpecificationDownstream) Validate() error {
	if len(d.IntegrationIds) == 0 {
		return errors.New("At least one integration id must be defined for a downstream specification")
	}
	return nil
}

//CustomEventSpecification is the representation of a custom event specification in Instana
type CustomEventSpecification struct {
	ID             string                        `json:"id"`
	Name           string                        `json:"name"`
	EntityType     string                        `json:"entityType"`
	Query          *string                       `json:"query"`
	Triggering     bool                          `json:"triggering"`
	Description    *string                       `json:"description"`
	ExpirationTime *int                          `json:"expirationTime"`
	Enabled        bool                          `json:"enabled"`
	Rules          []RuleSpecification           `json:"rules"`
	Downstream     *EventSpecificationDownstream `json:"downstream"`
}

//GetID implemention of the interface InstanaDataObject
func (spec CustomEventSpecification) GetID() string {
	return spec.ID
}

//Validate implementation of the interface InstanaDataObject to verify if data object is correct
func (spec CustomEventSpecification) Validate() error {
	if len(spec.ID) == 0 {
		return errors.New("ID is missing")
	}
	if len(spec.Name) == 0 {
		return errors.New("name is missing")
	}
	if len(spec.EntityType) == 0 {
		return errors.New("entity type is missing")
	}
	if len(spec.Rules) != 1 {
		return errors.New("exactly one rule must be defined")
	}
	for _, r := range spec.Rules {
		if err := r.Validate(); err != nil {
			return err
		}
	}

	if spec.Downstream != nil {
		if err := spec.Downstream.Validate(); err != nil {
			return err
		}
	}
	return nil
}
