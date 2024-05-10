package statsig

type statsigMetadata struct {
	SDKType    string `json:"sdkType"`
	SDKVersion string `json:"sdkVersion"`
}

func getStatsigMetadata() statsigMetadata {
	return statsigMetadata{
		SDKType:    "terraform-provider",
		SDKVersion: "0.2.0",
	}
}
