package franken

type FrankenColor string

const (
	FrankenColorPrimary   FrankenColor = "primary"
	FrankenColorSecondary FrankenColor = "secondary"
	FrankenColorDanger    FrankenColor = "danger"
	FrankenColorLink      FrankenColor = "link"
)

type FrankenConditionalClasses map[string]bool
