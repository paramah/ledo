package helper

type DockerServiceEnvironment struct {
	EnvironmentName  string `yaml:"environment_name,omitempty" json:"environmentName"`
	EnvironmentValue string `yaml:"environment_value,omitempty" json:"environmentValue"`
}

type DockerProjectAdditionalServiceCfg struct {
	DockerServiceMode        string                     `yaml:"docker_service_mode,omitempty" json:"dockerServiceMode"`
	DockerServiceType        string                     `yaml:"docker_service_type,omitempty" json:"dockerServiceType"`
	DockerServiceImage       string                     `yaml:"docker_service_image,omitempty" json:"dockerServiceImage"`
	DockerServiceTag         string                     `yaml:"docker_service_tag,omitempty" json:"dockerServiceTag"`
	DockerServiceEnvironment []DockerServiceEnvironment `yaml:"docker_service_environment,omitempty" json:"dockerServiceEnvironment"`
}

type DockerComposeModeCfg struct {
	DockerComposeName     string                              `yaml:"docker_compose_name,omitempty" json:"dockerComposeName"`
	DockerComposeServices []DockerProjectAdditionalServiceCfg `yaml:"docker_compose_services,omitempty" json:"dockerComposeServices"`
}

type DockerProjectCfg struct {
	HasDockerRegistry  bool                   `yaml:"has_docker_registry,omitempty" json:"hasDockerRegistry"`
	DockerBaseImage    string                 `yaml:"docker_base_image,omitempty" json:"dockerBaseImage"`
	DockerBaseTag      string                 `yaml:"docker_base_tag,omitempty" json:"dockerBaseTag"`
	DockerComposeModes []DockerComposeModeCfg `yaml:"docker_compose_modes,omitempty" json:"dockerComposeModes"`
}

//var PredefinedAdditionalService = []DockerProjectAdditionalServiceCfg{
//	{
//		DockerServiceMode:  "base",
//		DockerServiceType:  "database",
//		DockerServiceImage: "mariadb",
//		DockerServiceTag:   "",
//		DockerServiceEnvironment: []DockerServiceEnvironment{
//			{
//				EnvironmentName:  "DB_NAME",
//				EnvironmentValue: "dev",
//			},
//			},
//		},
//	{
//		DockerServiceMode:  "dev",
//		DockerServiceType:  "tools",
//		DockerServiceImage: "dbgate/dbgate",
//		DockerServiceTag:   "",
//		},
//	{
//		DockerServiceMode:  "dev",
//		DockerServiceType:  "tools",
//		DockerServiceImage: "mailhog/mailhog",
//		DockerServiceTag:   "",
//		},
//	{
//		DockerServiceMode:  "dev",
//		DockerServiceType:  "tools",
//		DockerServiceImage: "minio/minio",
//		DockerServiceTag:   "",
//		},
//	{
//		DockerServiceMode:  "base",
//		DockerServiceType:  "security",
//		DockerServiceImage: "authelia/authelia",
//		DockerServiceTag:   "",
//		},
//	{
//		DockerServiceMode:  "base",
//		DockerServiceType:  "infrastructure",
//		DockerServiceImage: "traefik",
//		DockerServiceTag:   "",
//		},
//	{
//		DockerServiceMode:  "base",
//		DockerServiceType:  "infrastructure",
//		DockerServiceImage: "portainer",
//		DockerServiceTag:   "",
//		},
//}
