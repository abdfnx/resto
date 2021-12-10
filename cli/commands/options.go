package commands

type Auth struct {
	TokenAuth string
	BasicAuthUsername string
	BasicAuthPassword string
	Type string
}

type Method struct {
	AuthType *Auth
	JustShowBody bool
	JustShowHeaders bool
	SaveFile string
	ContentType string
	OpenEditor bool
	Body string
	IsBodyStdin bool
}

type Options struct {
	Method *Method
	URL string
}
