package generate

import "github.com/lucasjones/reggen"

func GenerateNumber(pattern string) (string, error) {
	g, err := reggen.NewGenerator(pattern)
	if err != nil {
		return "", err
	}
	return g.Generate(1), nil
}
