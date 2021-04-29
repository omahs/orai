package types

// provider module event types
const (
	// TODO: Create your event types
	// EventType<Action>    		= "action"
	// Event type for set ai oracle
	EventTypeRequestWithData        = "ai_oracle_data"
	EventTypeReportWithData         = "report_data"
	EventTypeTestCaseReportWithData = "testcase_report_data"
	EventTypeSetAIRequest           = "set_ai_oracle"
	EventTypeSetReport              = "set_report"
	EventTypeSetTestCaseReport      = "set_testcase_report"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	// AttributeKeyAddress  		= "address"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
	// Attribute for request
	AttributeRequestID        = "request_id"
	AttributeRequestValidator = "request_validator"
	AttributeRequestCreator   = "request_creator"
	// Attributes for oracle script
	AttributeContract              = "contract"
	AttributeRequestValidatorCount = "request_validator_count"
	AttributeRequestInput          = "request_input"
	AttributeRequestExpectedOutput = "expected_request_output"
	AttributeRequestDSources       = "data_sources"
	AttributeRequestTCases         = "test_cases"
	AttributeBaseReport            = "base_report"
	AttributeReport                = "report"
	AttributeTestCaseReport        = "testcase_report"
	AttributeKeyValidator          = "validator"
)
