// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a route calculator resource in your AWS account. You can send requests
// to a route calculator resource to estimate travel time, distance, and get
// directions. A route calculator sources traffic and road network data from your
// chosen data provider.
func (c *Client) CreateRouteCalculator(ctx context.Context, params *CreateRouteCalculatorInput, optFns ...func(*Options)) (*CreateRouteCalculatorOutput, error) {
	if params == nil {
		params = &CreateRouteCalculatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRouteCalculator", params, optFns, c.addOperationCreateRouteCalculatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRouteCalculatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRouteCalculatorInput struct {

	// The name of the route calculator resource. Requirements:
	//
	// * Can use alphanumeric
	// characters (A–Z, a–z, 0–9) , hyphens (-), periods (.), and underscores (_).
	//
	// *
	// Must be a unique Route calculator resource name.
	//
	// * No spaces allowed. For
	// example, ExampleRouteCalculator.
	//
	// This member is required.
	CalculatorName *string

	// Specifies the data provider of traffic and road network data. This field is
	// case-sensitive. Enter the valid values as shown. For example, entering HERE
	// returns an error. Route calculators that use Esri as a data source only
	// calculate routes that are shorter than 400 km. Valid values include:
	//
	// * Esri –
	// For additional information about Esri
	// (https://docs.aws.amazon.com/location/latest/developerguide/esri.html)'s
	// coverage in your region of interest, see Esri details on street networks and
	// traffic coverage
	// (https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm).
	//
	// *
	// Here – For additional information about HERE Technologies
	// (https://docs.aws.amazon.com/location/latest/developerguide/HERE.html)' coverage
	// in your region of interest, see HERE car routing coverage
	// (https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html)
	// and HERE truck routing coverage
	// (https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html).
	//
	// For
	// additional information , see Data providers
	// (https://docs.aws.amazon.com/location/latest/developerguide/what-is-data-provider.html)
	// on the Amazon Location Service Developer Guide.
	//
	// This member is required.
	DataSource *string

	// Specifies the pricing plan for your route calculator resource. For additional
	// details and restrictions on each pricing plan option, see Amazon Location
	// Service pricing (https://aws.amazon.com/location/pricing/).
	//
	// This member is required.
	PricingPlan types.PricingPlan

	// The optional description for the route calculator resource.
	Description *string

	// Applies one or more tags to the route calculator resource. A tag is a key-value
	// pair helps manage, identify, search, and filter your resources by labelling
	// them.
	//
	// * For example: { "tag1" : "value1", "tag2" : "value2"}
	//
	// Format: "key" :
	// "value" Restrictions:
	//
	// * Maximum 50 tags per resource
	//
	// * Each resource tag must
	// be unique with a maximum of one value.
	//
	// * Maximum key length: 128 Unicode
	// characters in UTF-8
	//
	// * Maximum value length: 256 Unicode characters in UTF-8
	//
	// *
	// Can use alphanumeric characters (A–Z, a–z, 0–9), and the following characters: +
	// - = . _ : / @.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateRouteCalculatorOutput struct {

	// The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN
	// when you specify a resource across all AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:route-calculator/ExampleCalculator
	//
	// This member is required.
	CalculatorArn *string

	// The name of the route calculator resource.
	//
	// * For example,
	// ExampleRouteCalculator.
	//
	// This member is required.
	CalculatorName *string

	// The timestamp when the route calculator resource was created in ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// * For example, 2020–07-2T12:15:20.000Z+01:00
	//
	// This member is required.
	CreateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRouteCalculatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRouteCalculator{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateRouteCalculatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRouteCalculator(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateRouteCalculator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "CreateRouteCalculator",
	}
}
