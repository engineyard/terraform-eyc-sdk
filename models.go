package eyc

type EnvVar struct {
	id               int    `json:"id"`
	application      string `json:"application"`
	application_id   int    `json:"application_id"`
	application_name string `json:"application_name"`
	environment      string `json:"environment"`
	environment_id   int    `json:"environment_id`
	environment_name string `json:"environment_name"`
	name             string `json:"name"`
	value            string `json:"value"`
}

type EnvVarParam struct {
	environment_variable EnvVarNameValue `json:"environment_variable`
	application_id       int             `json:"application_id`
	environment_id       int             `json:"environment_id`
}

type EnvVarNameValue struct {
	name  string `json:"name"`
	value string `json:"value"`
}

type EnvVars struct {
	environment_variables []EnvVar `json:"environment_variables"`
}
