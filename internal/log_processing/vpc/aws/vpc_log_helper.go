package aws

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/merbinr/catcher/internal/log_processing/models"
	"github.com/merbinr/catcher/internal/web"
)

func AwsVpcLogFlowLogParsing(each_log_records web.AwsVpcLogRecordsData) (models.VpcNormalizedProcessedData, error) {
	base64_log_string := each_log_records.Data
	decodedBytes, err := base64.StdEncoding.DecodeString(base64_log_string)
	if err != nil {
		return models.VpcNormalizedProcessedData{}, fmt.Errorf("error on decoding AWS VPC log, log_data: %s",
			base64_log_string)
	}
	// string of decodedBytes will be looks like {"message" : "xx xx xxx xx xx xx"}
	var messageBody AwsVpcLogRecordMessageData
	err = json.Unmarshal(decodedBytes, &messageBody)
	if err != nil {
		return models.VpcNormalizedProcessedData{}, fmt.Errorf("error on decoding AWS VPC log, log_data: %s",
			base64_log_string)
	}
	flow_log_string := messageBody.Message

	VpcLogData, err := parseFlowLog(flow_log_string)
	if err != nil {
		return models.VpcNormalizedProcessedData{}, fmt.Errorf("unable to parse AWS vpc log message, log_data: %s, error: %s",
			base64_log_string,
			err)
	}
	fmt.Printf("%+v\n", VpcLogData)
	unique_string, err := createUniqueString(VpcLogData)
	if err != nil {
		return models.VpcNormalizedProcessedData{}, fmt.Errorf("unable to create unique string for log_data: %s, error: %s",
			base64_log_string,
			err)
	}

	var aws_vpc_log_processed_data models.VpcNormalizedProcessedData
	aws_vpc_log_processed_data.UniqueStr = unique_string
	aws_vpc_log_processed_data.Data = VpcLogData
	println(unique_string)
	return aws_vpc_log_processed_data, nil

}

func parseFlowLog(log string) (models.VpcNormalizedData, error) {
	fields := strings.Fields(log)
	if len(fields) != 14 {
		return models.VpcNormalizedData{},
			fmt.Errorf("unexpected number of fields in the AWS VPC log entry")
	}

	version, err := convertToInt(fields[0])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert version field to int type")
	}

	destinationPort, err := convertToInt(fields[5])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert destinationPort field to int type")
	}

	sourcePort, err := convertToInt(fields[6])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert sourcePort field to int type")
	}

	protocol, err := convertToInt(fields[7])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert protocol field to int type")
	}

	packets, err := convertToInt(fields[8])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert packets field to int type")
	}

	bytes, err := convertToInt(fields[9])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert bytes field to int type")
	}

	startTime, err := convertToInt(fields[10])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert startTime field to int type")
	}

	endTime, err := convertToInt(fields[11])
	if err != nil {
		return models.VpcNormalizedData{},
			fmt.Errorf("unable to convert endTime field to int type")
	}

	flowLog := models.VpcNormalizedData{
		Version:         version,
		AccountID:       fields[1],
		InterfaceID:     fields[2],
		SourceIP:        fields[3],
		DestinationIP:   fields[4],
		DestinationPort: destinationPort,
		SourcePort:      sourcePort,
		Protocol:        protocol,
		Packets:         packets,
		Bytes:           bytes,
		StartTime:       startTime,
		EndTime:         endTime,
		Action:          fields[12],
		LogStatus:       fields[13],
	}

	return flowLog, nil
}

func convertToInt(value string) (int, error) {
	output, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("unexpected error when converting to int value %s, error: %s", value, err)
	}
	return output, nil
}

func createUniqueString(flowLog models.VpcNormalizedData) (string, error) {
	unique_string_fields := os.Getenv("AWS_VPC_LOGS_UNIQUE_STRING_FIELDS")
	if unique_string_fields == "" {
		unique_string_fields = "AccountID,InterfaceID,SourceIP,SourcePort,DestinationPort"
	}
	fields := strings.Split(unique_string_fields, ",")

	val := reflect.ValueOf(flowLog)
	typ := reflect.TypeOf(flowLog)

	unique_string := ""

	for _, field := range fields {
		field = strings.TrimSpace(field)

		// Check field exist
		_, found := typ.FieldByName(field)
		if !found {
			return "", fmt.Errorf("field '%s' does not exist in the struct", field)
		}

		// Fetch value using field name
		value := val.FieldByName(field)
		unique_string = unique_string + strings.TrimSpace(value.String())
	}

	DEFAULT_UNIQUE_STRING := "awsvpclogs_"
	unique_string = DEFAULT_UNIQUE_STRING + unique_string
	return unique_string, nil

}
