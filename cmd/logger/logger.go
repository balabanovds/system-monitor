package logger

import "go.uber.org/zap"

func New(level string, production bool) (*zap.Logger, error) {
	var cfg zap.Config
	cfg.DisableStacktrace = true
	cfg.DisableCaller = true

	if production {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	al := zap.NewAtomicLevel()
	err := al.UnmarshalText([]byte(level))
	if err != nil {
		return nil, err
	}

	cfg.Level.SetLevel(al.Level())

	cfg.OutputPaths = []string{"stderr"}

	l, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	//zap.ReplaceGlobals(l)

	return l, nil
}
