package aig

type SettingOptions struct {
	Otel *bool
}

func Settings() *SettingOptions {
	return new(SettingOptions)
}

func (s *SettingOptions) EnableOtel() *SettingOptions {
	enable := true
	s.Otel = &enable
	return s
}
