package eyc

type EnvVar struct {
	ID               int    `json:"id"`
	Application      string `json:"application"`
	Application_id   int    `json:"application_id"`
	Application_name string `json:"application_name"`
	Environment      string `json:"environment"`
	Environment_id   int    `json:"environment_id`
	Environment_name string `json:"environment_name"`
	Name             string `json:"name"`
	Value            string `json:"value"`
}

type EnvVarParam struct {
	Environment_variable EnvVarNameValue `json:"environment_variable`
	Application_id       int             `json:"application_id`
	Environment_id       int             `json:"environment_id`
}

type EnvVarNameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type EnvVars struct {
	Environment_variables []EnvVar `json:"environment_variables"`
}
