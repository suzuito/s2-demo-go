package setting

type Env struct {
	AllowedOrigins   string `envconfig:"ALLOWED_ORIGINS"`
	AllowedMethods   string `envconfig:"ALLOWED_METHODS"`
	AllowedHeaders   string `envconfig:"ALLOWED_HEADERS"`
	GCPBucketArticle string `envconfig:"GCP_BUCKET_ARTICLE"`
	GCPBucketServer  string `envconfig:"GCP_BUCKET_SERVER"`
	OriginServer     string `envconfig:"ORIGIN_SERVER"`
	EnvironmentName  string `default:"dev" envconfig:"ENV"`
}
