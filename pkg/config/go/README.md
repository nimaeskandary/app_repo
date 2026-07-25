# Config

`NewJsonConfigLoaderModule` parses JSON configuration, validates struct tags, and provides a
`ConfigLoader[T]` through Fx. The application passes its complete slice of custom JSON unmarshalers
when it builds the dependency graph.

```go
type AppConfig struct {
	Name string `json:"Name" validate:"required"`
}

app := fx.New(
	config.NewJsonConfigLoaderModule[AppConfig](
		[]byte(`{"Name":"gordle"}`),
		nil,
	),
	fx.Invoke(func(loader config_core.ConfigLoader[AppConfig]) {
		appConfig := loader.GetConfig()
		// Use appConfig.
	}),
)
```

## Custom unmarshalers

Custom unmarshalers are supplied in application-defined order. For example, a `SecretString`
unmarshaler can treat its JSON value as the name of an environment variable:

```go
secretStringUnmarshaler := json.UnmarshalFunc(
	func(data []byte, value *config_core.SecretString) error {
		var environmentVariable string
		if err := json.Unmarshal(data, &environmentVariable); err != nil {
			return err
		}

		secret, found := os.LookupEnv(environmentVariable)
		if !found {
			return fmt.Errorf("environment variable %q is not set", environmentVariable)
		}

		*value = config_core.SecretString(secret)
		return nil
	},
)

app := fx.New(
	config.NewJsonConfigLoaderModule[AppConfig](
		configBytes,
		[]*json.Unmarshalers{
			secretStringUnmarshaler,
		},
	),
)
```
