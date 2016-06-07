package awsreg

import "os"

func RegionFromEnvOrDefault(defaultReg string) *string {
	envReg := os.Getenv("AWS_REGION")
	switch envReg {
	case "":
		return &defaultReg
	default:
		return &envReg
	}
}
