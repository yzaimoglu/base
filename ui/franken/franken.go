package franken

import (
	"fmt"

	"github.com/a-h/templ"
)

func NoComponent() templ.Component {
	return nil
}

type FrankenColor string

type TailwindColorType string
type TailwindColorWeight string
type TailwindColor struct {
	Type   TailwindColorType
	Weight TailwindColorWeight
}

const (
	TailwindColorRed    TailwindColorType = "red"
	TailwindColorSlate  TailwindColorType = "slate"
	TailwindColorGreen  TailwindColorType = "green"
	TailwindColorPurple TailwindColorType = "purple"
	TailwindColorIndigo TailwindColorType = "indigo"
	TailwindColorBlue   TailwindColorType = "blue"

	TailwindColor50  TailwindColorWeight = "50"
	TailwindColor100 TailwindColorWeight = "100"
	TailwindColor200 TailwindColorWeight = "200"
	TailwindColor300 TailwindColorWeight = "300"
	TailwindColor400 TailwindColorWeight = "400"
	TailwindColor500 TailwindColorWeight = "500"
	TailwindColor600 TailwindColorWeight = "600"
	TailwindColor700 TailwindColorWeight = "700"
	TailwindColor800 TailwindColorWeight = "800"
	TailwindColor900 TailwindColorWeight = "900"
)

func NewTailwindColor(colorType TailwindColorType, colorWeight TailwindColorWeight) TailwindColor {
	return TailwindColor{colorType, colorWeight}
}

func (t TailwindColor) String() string {
	if !t.Isset() {
		return ""
	}
	return fmt.Sprintf("%s-%s", t.Type, t.Weight)
}

func (t TailwindColor) Isset() bool {
	if string(t.Type) != "" && string(t.Weight) != "" {
		return true
	}
	return false
}

func (f FrankenColor) Isset() bool {
	if string(f) != "" {
		return true
	}
	return false
}

const (
	FrankenColorDefault   FrankenColor = "default"
	FrankenColorGhost     FrankenColor = "ghost"
	FrankenColorPrimary   FrankenColor = "primary"
	FrankenColorSecondary FrankenColor = "secondary"
	FrankenColorDanger    FrankenColor = "danger"
	FrankenColorText      FrankenColor = "text"
	FrankenColorLink      FrankenColor = "link"
)

type FrankenConditionalClasses map[string]bool
