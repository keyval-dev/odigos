package destinations

type Destination struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Type        string `yaml:"type"`
	DisplayName string `yaml:"displayName"`
	Category    string `yaml:"category"`
}

type Spec struct {
	Image   string `yaml:"image"`
	Signals struct {
		Traces struct {
			Supported bool `yaml:"supported"`
		}
		Metrics struct {
			Supported bool `yaml:"supported"`
		}
		Logs struct {
			Supported bool `yaml:"supported"`
		}
	}
	Fields []Field `yaml:"fields"`
}

type Field struct {
	Name           string                 `yaml:"name"`
	DisplayName    string                 `yaml:"displayName"`
	VideoURL       string                 `yaml:"videoUrl"`
	ComponentType  string                 `yaml:"componentType"`
	ComponentProps map[string]interface{} `yaml:"componentProps"`
}
