package image

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Achno/gowall/config"
	cpkg "github.com/Achno/gowall/internal/backends/color"
)

type Theme struct {
	Name   string
	Colors []color.Color
}

var themes = map[string]Theme{
	"caelus": Caelus,
	"autumn": Autumn,
	"cyberpunk": Cyberpunk,
	"dusky": Dusky,
	"everforest": Everforest,
	"evergreen": Evergreen,
	"gruvbox": Gruvbox,
	"pink-crimson": PinkCrimson,
	"catppuccin-mocha-mauve": CatppuccinMochaMauve,
	"garnet": Garnet,
	"kanagawa-kasumi": KanagawaKasumi,
	"kemuri-koke": KemuriKoke,
	"lilac-amoled": LilacAmoled,
	"matecito": Matecito,
	"mizuki-akiyama": MizukiAkiyama,
	"murata": Murata,
	"noctalia-legacy": NoctaliaLegacy,
	"osaka-jade": OsakaJade,
	"shien": Shien,
	"shinonome": Shinonome,
	"tokyo-night-moon": TokyoNightMoon,
	"nord": Nord,
	"rose-pine": RosePine,
}







func LoadCustomThemes() {

	for _, tw := range config.GowallConfig.Themes {
		valid := true
		if tw.Name == "" || len(tw.Colors) == 0 {
			// skip invalid color
			continue
		}

		theme := Theme{
			Name:   tw.Name,
			Colors: make([]color.Color, len(tw.Colors)),
		}

		for i, hexColor := range tw.Colors {
			col, err := cpkg.HexToRGBA(hexColor)
			if err != nil {
				log.Printf("invalid color %s in theme %s: %v", hexColor, tw.Name, err)
				valid = false
				break
			}
			theme.Colors[i] = col
		}

		if valid && !themeExists(theme.Name) {

			themes[strings.ToLower(theme.Name)] = theme
		}
	}
}

func ListThemes() []string {
	allThemes := make([]string, 0, len(themes))
	for theme := range themes {
		allThemes = append(allThemes, theme)
	}
	return allThemes
}

func SelectTheme(theme string) (Theme, error) {
	selectedTheme, exists := themes[strings.ToLower(theme)]

	if !exists {
		return Theme{}, errors.New("unknown theme")
	}

	return selectedTheme, nil
}

func themeExists(theme string) bool {

	_, exists := themes[theme]

	return exists
}

// returns themeName that was inserted to the theme map
func LoadThemeFromJson(jsonTheme string) (string, error) {
	reader, err := os.Open(jsonTheme)
	if err != nil {
		return "", fmt.Errorf("while opening json file: %w", err)
	}
	defer reader.Close()
	data, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("while reading the json file")
	}
	var tm struct {
		Name   string   `json:"name"`
		Colors []string `json:"colors"`
	}

	if err := json.Unmarshal(data, &tm); err != nil {
		return "", fmt.Errorf("while parsing json theme file, ensure your .json is written correctly")
	}
	if len(tm.Name) <= 0 || len(tm.Colors) < 1 {
		return "", fmt.Errorf("json file does not contain a name or colors field(s)")
	}
	clrs, err := cpkg.HexToRGBASlice(tm.Colors)
	if err != nil {
		return "", err
	}
	themes[strings.ToLower(tm.Name)] = Theme{
		Name:   tm.Name,
		Colors: clrs,
	}

	return tm.Name, nil
}

// returns the colors of the theme in hex code format
func GetThemeColors(theme string) ([]string, error) {
	var colors []string

	selectedTheme, err := SelectTheme(theme)

	if err != nil {
		return nil, err
	}

	for _, clr := range selectedTheme.Colors {
		rgba, ok := clr.(color.RGBA)

		if !ok {
			return nil, fmt.Errorf("color is not of type color.RGBA")
		}
		hexCode := cpkg.RGBtoHex(rgba)
		colors = append(colors, hexCode)
	}

	return colors, nil

}

var (
	Caelus = Theme{
		Name: "Caelus",
		Colors: []color.Color{
			color.RGBA{R: 239, G: 147, B: 77, A: 255},
			color.RGBA{R: 15, G: 15, B: 15, A: 255},
			color.RGBA{R: 126, G: 201, B: 126, A: 255},
			color.RGBA{R: 126, G: 201, B: 163, A: 255},
			color.RGBA{R: 241, G: 110, B: 101, A: 255},
			color.RGBA{R: 244, G: 222, B: 205, A: 255},
			color.RGBA{R: 30, G: 30, B: 30, A: 255},
			color.RGBA{R: 191, G: 175, B: 158, A: 255},
			color.RGBA{R: 53, G: 53, B: 53, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 41, G: 41, B: 41, A: 255},
			color.RGBA{R: 239, G: 191, B: 113, A: 255},
			color.RGBA{R: 113, G: 180, B: 214, A: 255},
			color.RGBA{R: 226, G: 141, B: 198, A: 255},
		},
	}

	Autumn = Theme{
		Name: "Autumn",
		Colors: []color.Color{
			color.RGBA{R: 232, G: 163, B: 61, A: 255},
			color.RGBA{R: 61, G: 42, B: 12, A: 255},
			color.RGBA{R: 74, G: 52, B: 24, A: 255},
			color.RGBA{R: 251, G: 217, B: 160, A: 255},
			color.RGBA{R: 181, G: 74, B: 50, A: 255},
			color.RGBA{R: 245, G: 239, B: 230, A: 255},
			color.RGBA{R: 61, G: 32, B: 24, A: 255},
			color.RGBA{R: 240, G: 201, B: 188, A: 255},
			color.RGBA{R: 127, G: 160, B: 184, A: 255},
			color.RGBA{R: 18, G: 36, B: 46, A: 255},
			color.RGBA{R: 36, G: 51, B: 61, A: 255},
			color.RGBA{R: 207, G: 225, B: 235, A: 255},
			color.RGBA{R: 226, G: 96, B: 74, A: 255},
			color.RGBA{R: 61, G: 15, B: 8, A: 255},
			color.RGBA{R: 74, G: 28, B: 18, A: 255},
			color.RGBA{R: 245, G: 207, B: 194, A: 255},
			color.RGBA{R: 21, G: 19, B: 15, A: 255},
			color.RGBA{R: 243, G: 233, B: 216, A: 255},
			color.RGBA{R: 31, G: 26, B: 19, A: 255},
			color.RGBA{R: 201, G: 185, B: 154, A: 255},
			color.RGBA{R: 122, G: 104, B: 73, A: 255},
			color.RGBA{R: 61, G: 51, B: 35, A: 255},
			color.RGBA{R: 42, G: 35, B: 24, A: 255},
			color.RGBA{R: 138, G: 90, B: 29, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
		},
	}

	Cyberpunk = Theme{
		Name: "Cyberpunk",
		Colors: []color.Color{
			color.RGBA{R: 196, G: 168, B: 46, A: 255},
			color.RGBA{R: 14, G: 16, B: 21, A: 255},
			color.RGBA{R: 209, G: 67, B: 88, A: 255},
			color.RGBA{R: 0, G: 166, B: 108, A: 255},
			color.RGBA{R: 179, G: 45, B: 45, A: 255},
			color.RGBA{R: 10, G: 13, B: 20, A: 255},
			color.RGBA{R: 92, G: 138, B: 196, A: 255},
			color.RGBA{R: 17, G: 21, B: 29, A: 255},
			color.RGBA{R: 155, G: 107, B: 193, A: 255},
			color.RGBA{R: 69, G: 160, B: 214, A: 255},
			color.RGBA{R: 9, G: 13, B: 19, A: 255},
			color.RGBA{R: 216, G: 224, B: 255, A: 255},
			color.RGBA{R: 12, G: 14, B: 20, A: 255},
			color.RGBA{R: 230, G: 69, B: 114, A: 255},
			color.RGBA{R: 137, G: 211, B: 106, A: 255},
			color.RGBA{R: 215, G: 162, B: 58, A: 255},
			color.RGBA{R: 79, G: 143, B: 255, A: 255},
			color.RGBA{R: 157, G: 109, B: 255, A: 255},
			color.RGBA{R: 67, G: 201, B: 255, A: 255},
			color.RGBA{R: 183, G: 196, B: 242, A: 255},
			color.RGBA{R: 43, G: 49, B: 74, A: 255},
		},
	}

	Dusky = Theme{
		Name: "Dusky",
		Colors: []color.Color{
			color.RGBA{R: 216, G: 93, B: 123, A: 255},
			color.RGBA{R: 19, G: 17, B: 26, A: 255},
			color.RGBA{R: 155, G: 116, B: 184, A: 255},
			color.RGBA{R: 98, G: 156, B: 184, A: 255},
			color.RGBA{R: 224, G: 90, B: 90, A: 255},
			color.RGBA{R: 226, G: 222, B: 234, A: 255},
			color.RGBA{R: 31, G: 27, B: 41, A: 255},
			color.RGBA{R: 156, G: 150, B: 170, A: 255},
			color.RGBA{R: 64, G: 57, B: 79, A: 255},
			color.RGBA{R: 12, G: 10, B: 18, A: 255},
			color.RGBA{R: 45, G: 39, B: 58, A: 255},
			color.RGBA{R: 216, G: 157, B: 106, A: 255},
			color.RGBA{R: 91, G: 143, B: 172, A: 255},
			color.RGBA{R: 155, G: 109, B: 184, A: 255},
		},
	}

	Everforest = Theme{
		Name: "Everforest",
		Colors: []color.Color{
			color.RGBA{R: 167, G: 192, B: 128, A: 255},
			color.RGBA{R: 39, G: 46, B: 51, A: 255},
			color.RGBA{R: 127, G: 187, B: 179, A: 255},
			color.RGBA{R: 214, G: 153, B: 182, A: 255},
			color.RGBA{R: 230, G: 126, B: 128, A: 255},
			color.RGBA{R: 211, G: 198, B: 170, A: 255},
			color.RGBA{R: 55, G: 65, B: 69, A: 255},
			color.RGBA{R: 157, G: 169, B: 160, A: 255},
			color.RGBA{R: 73, G: 81, B: 86, A: 255},
			color.RGBA{R: 31, G: 38, B: 42, A: 255},
			color.RGBA{R: 230, G: 152, B: 117, A: 255},
			color.RGBA{R: 219, G: 188, B: 127, A: 255},
			color.RGBA{R: 131, G: 192, B: 146, A: 255},
		},
	}

	Evergreen = Theme{
		Name: "Evergreen",
		Colors: []color.Color{
			color.RGBA{R: 74, G: 154, B: 104, A: 255},
			color.RGBA{R: 8, G: 13, B: 10, A: 255},
			color.RGBA{R: 106, G: 174, B: 82, A: 255},
			color.RGBA{R: 82, G: 152, B: 122, A: 255},
			color.RGBA{R: 200, G: 122, B: 92, A: 255},
			color.RGBA{R: 16, G: 25, B: 19, A: 255},
			color.RGBA{R: 161, G: 175, B: 156, A: 255},
			color.RGBA{R: 40, G: 48, B: 43, A: 255},
			color.RGBA{R: 121, G: 131, B: 117, A: 255},
			color.RGBA{R: 74, G: 104, B: 74, A: 255},
			color.RGBA{R: 196, G: 166, B: 78, A: 255},
			color.RGBA{R: 138, G: 120, B: 86, A: 255},
			color.RGBA{R: 216, G: 146, B: 104, A: 255},
			color.RGBA{R: 130, G: 196, B: 90, A: 255},
			color.RGBA{R: 212, G: 186, B: 96, A: 255},
			color.RGBA{R: 90, G: 174, B: 122, A: 255},
			color.RGBA{R: 160, G: 142, B: 102, A: 255},
			color.RGBA{R: 72, G: 170, B: 130, A: 255},
			color.RGBA{R: 185, G: 195, B: 181, A: 255},
		},
	}

	Gruvbox = Theme{
		Name: "Gruvbox",
		Colors: []color.Color{
			color.RGBA{R: 215, G: 153, B: 33, A: 255},
			color.RGBA{R: 40, G: 40, B: 40, A: 255},
			color.RGBA{R: 142, G: 192, B: 124, A: 255},
			color.RGBA{R: 69, G: 133, B: 136, A: 255},
			color.RGBA{R: 204, G: 36, B: 29, A: 255},
			color.RGBA{R: 251, G: 241, B: 199, A: 255},
			color.RGBA{R: 235, G: 219, B: 178, A: 255},
			color.RGBA{R: 60, G: 56, B: 54, A: 255},
			color.RGBA{R: 168, G: 153, B: 132, A: 255},
			color.RGBA{R: 80, G: 73, B: 69, A: 255},
			color.RGBA{R: 29, G: 32, B: 33, A: 255},
			color.RGBA{R: 104, G: 157, B: 106, A: 255},
			color.RGBA{R: 211, G: 134, B: 155, A: 255},
			color.RGBA{R: 131, G: 165, B: 152, A: 255},
		},
	}

	PinkCrimson = Theme{
		Name: "Pink-Crimson",
		Colors: []color.Color{
			color.RGBA{R: 224, G: 32, B: 48, A: 255},
			color.RGBA{R: 232, G: 174, B: 182, A: 255},
			color.RGBA{R: 240, G: 139, B: 176, A: 255},
			color.RGBA{R: 26, G: 18, B: 20, A: 255},
			color.RGBA{R: 232, G: 106, B: 154, A: 255},
			color.RGBA{R: 44, G: 31, B: 34, A: 255},
			color.RGBA{R: 138, G: 110, B: 120, A: 255},
			color.RGBA{R: 184, G: 192, B: 200, A: 255},
			color.RGBA{R: 58, G: 36, B: 40, A: 255},
			color.RGBA{R: 95, G: 191, B: 74, A: 255},
			color.RGBA{R: 74, G: 58, B: 62, A: 255},
			color.RGBA{R: 255, G: 74, B: 88, A: 255},
			color.RGBA{R: 122, G: 217, B: 100, A: 255},
			color.RGBA{R: 248, G: 192, B: 212, A: 255},
			color.RGBA{R: 255, G: 194, B: 204, A: 255},
		},
	}

	CatppuccinMochaMauve = Theme{
		Name: "Catppuccin-Mocha-Mauve",
		Colors: []color.Color{
			color.RGBA{R: 203, G: 166, B: 247, A: 255},
			color.RGBA{R: 30, G: 30, B: 46, A: 255},
			color.RGBA{R: 180, G: 190, B: 254, A: 255},
			color.RGBA{R: 137, G: 220, B: 235, A: 255},
			color.RGBA{R: 243, G: 139, B: 168, A: 255},
			color.RGBA{R: 205, G: 214, B: 244, A: 255},
			color.RGBA{R: 49, G: 50, B: 68, A: 255},
			color.RGBA{R: 166, G: 173, B: 200, A: 255},
			color.RGBA{R: 108, G: 112, B: 134, A: 255},
			color.RGBA{R: 17, G: 17, B: 27, A: 255},
			color.RGBA{R: 69, G: 71, B: 90, A: 255},
			color.RGBA{R: 166, G: 227, B: 161, A: 255},
			color.RGBA{R: 249, G: 226, B: 175, A: 255},
			color.RGBA{R: 137, G: 180, B: 250, A: 255},
			color.RGBA{R: 245, G: 194, B: 231, A: 255},
			color.RGBA{R: 148, G: 226, B: 213, A: 255},
			color.RGBA{R: 186, G: 194, B: 222, A: 255},
			color.RGBA{R: 88, G: 91, B: 112, A: 255},
			color.RGBA{R: 245, G: 224, B: 220, A: 255},
		},
	}

	Garnet = Theme{
		Name: "Garnet",
		Colors: []color.Color{
			color.RGBA{R: 153, G: 0, B: 0, A: 255},
			color.RGBA{R: 245, G: 238, B: 238, A: 255},
			color.RGBA{R: 200, G: 150, B: 42, A: 255},
			color.RGBA{R: 26, G: 20, B: 20, A: 255},
			color.RGBA{R: 122, G: 0, B: 0, A: 255},
			color.RGBA{R: 240, G: 232, B: 232, A: 255},
			color.RGBA{R: 214, G: 69, B: 69, A: 255},
			color.RGBA{R: 18, G: 18, B: 18, A: 255},
			color.RGBA{R: 232, G: 224, B: 224, A: 255},
			color.RGBA{R: 24, G: 22, B: 22, A: 255},
			color.RGBA{R: 196, G: 188, B: 188, A: 255},
			color.RGBA{R: 42, G: 36, B: 36, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 179, G: 0, B: 0, A: 255},
			color.RGBA{R: 90, G: 158, B: 90, A: 255},
			color.RGBA{R: 201, G: 162, B: 39, A: 255},
			color.RGBA{R: 74, G: 127, B: 181, A: 255},
			color.RGBA{R: 168, G: 90, B: 122, A: 255},
			color.RGBA{R: 74, G: 163, B: 163, A: 255},
			color.RGBA{R: 58, G: 58, B: 58, A: 255},
			color.RGBA{R: 127, G: 199, B: 127, A: 255},
			color.RGBA{R: 230, G: 193, B: 77, A: 255},
			color.RGBA{R: 111, G: 163, B: 212, A: 255},
			color.RGBA{R: 201, G: 127, B: 160, A: 255},
			color.RGBA{R: 111, G: 202, B: 202, A: 255},
			color.RGBA{R: 245, G: 239, B: 239, A: 255},
		},
	}

	KanagawaKasumi = Theme{
		Name: "Kanagawa-Kasumi",
		Colors: []color.Color{
			color.RGBA{R: 217, G: 167, B: 139, A: 255},
			color.RGBA{R: 26, G: 32, B: 38, A: 255},
			color.RGBA{R: 119, G: 148, B: 166, A: 255},
			color.RGBA{R: 139, G: 163, B: 122, A: 255},
			color.RGBA{R: 217, G: 108, B: 108, A: 255},
			color.RGBA{R: 217, G: 209, B: 186, A: 255},
			color.RGBA{R: 32, G: 39, B: 46, A: 255},
			color.RGBA{R: 49, G: 60, B: 71, A: 255},
			color.RGBA{R: 10, G: 12, B: 16, A: 255},
			color.RGBA{R: 184, G: 122, B: 142, A: 255},
			color.RGBA{R: 119, G: 163, B: 160, A: 255},
			color.RGBA{R: 230, G: 126, B: 126, A: 255},
			color.RGBA{R: 159, G: 186, B: 140, A: 255},
			color.RGBA{R: 242, G: 190, B: 160, A: 255},
			color.RGBA{R: 138, G: 178, B: 201, A: 255},
			color.RGBA{R: 209, G: 149, B: 168, A: 255},
			color.RGBA{R: 140, G: 190, B: 186, A: 255},
			color.RGBA{R: 235, G: 228, B: 204, A: 255},
		},
	}

	KemuriKoke = Theme{
		Name: "Kemuri-Koke",
		Colors: []color.Color{
			color.RGBA{R: 139, G: 163, B: 126, A: 255},
			color.RGBA{R: 36, G: 33, B: 32, A: 255},
			color.RGBA{R: 166, G: 150, B: 128, A: 255},
			color.RGBA{R: 203, G: 145, B: 104, A: 255},
			color.RGBA{R: 198, G: 104, B: 93, A: 255},
			color.RGBA{R: 229, G: 222, B: 201, A: 255},
			color.RGBA{R: 54, G: 49, B: 47, A: 255},
			color.RGBA{R: 102, G: 92, B: 88, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 103, G: 128, B: 138, A: 255},
			color.RGBA{R: 162, G: 120, B: 134, A: 255},
			color.RGBA{R: 118, G: 159, B: 147, A: 255},
			color.RGBA{R: 219, G: 129, B: 119, A: 255},
			color.RGBA{R: 156, G: 183, B: 147, A: 255},
			color.RGBA{R: 226, G: 170, B: 130, A: 255},
			color.RGBA{R: 128, G: 155, B: 166, A: 255},
			color.RGBA{R: 188, G: 145, B: 159, A: 255},
			color.RGBA{R: 146, G: 185, B: 173, A: 255},
			color.RGBA{R: 244, G: 238, B: 220, A: 255},
		},
	}

	LilacAmoled = Theme{
		Name: "Lilac-Amoled",
		Colors: []color.Color{
			color.RGBA{R: 181, G: 143, B: 255, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 199, G: 154, B: 255, A: 255},
			color.RGBA{R: 216, G: 180, B: 255, A: 255},
			color.RGBA{R: 255, G: 111, B: 155, A: 255},
			color.RGBA{R: 232, G: 216, B: 255, A: 255},
			color.RGBA{R: 17, G: 13, B: 26, A: 255},
			color.RGBA{R: 76, G: 58, B: 112, A: 255},
			color.RGBA{R: 168, G: 230, B: 207, A: 255},
			color.RGBA{R: 224, G: 193, B: 255, A: 255},
			color.RGBA{R: 255, G: 140, B: 179, A: 255},
			color.RGBA{R: 184, G: 240, B: 216, A: 255},
			color.RGBA{R: 230, G: 209, B: 255, A: 255},
			color.RGBA{R: 201, G: 168, B: 255, A: 255},
			color.RGBA{R: 212, G: 184, B: 255, A: 255},
			color.RGBA{R: 240, G: 224, B: 255, A: 255},
			color.RGBA{R: 245, G: 240, B: 255, A: 255},
		},
	}

	Matecito = Theme{
		Name: "Matecito",
		Colors: []color.Color{
			color.RGBA{R: 141, G: 163, B: 131, A: 255},
			color.RGBA{R: 30, G: 35, B: 32, A: 255},
			color.RGBA{R: 203, G: 155, B: 124, A: 255},
			color.RGBA{R: 127, G: 145, B: 147, A: 255},
			color.RGBA{R: 194, G: 121, B: 121, A: 255},
			color.RGBA{R: 211, G: 205, B: 195, A: 255},
			color.RGBA{R: 39, G: 46, B: 42, A: 255},
			color.RGBA{R: 157, G: 164, B: 154, A: 255},
			color.RGBA{R: 74, G: 83, B: 77, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 114, G: 138, B: 150, A: 255},
			color.RGBA{R: 164, G: 143, B: 161, A: 255},
			color.RGBA{R: 122, G: 153, B: 141, A: 255},
			color.RGBA{R: 71, G: 82, B: 75, A: 255},
			color.RGBA{R: 212, G: 143, B: 143, A: 255},
			color.RGBA{R: 161, G: 181, B: 152, A: 255},
			color.RGBA{R: 219, G: 176, B: 149, A: 255},
			color.RGBA{R: 141, G: 163, B: 176, A: 255},
			color.RGBA{R: 186, G: 169, B: 184, A: 255},
			color.RGBA{R: 147, G: 173, B: 26, A: 255},
			color.RGBA{R: 228, G: 223, B: 213, A: 255},
		},
	}

	MizukiAkiyama = Theme{
		Name: "Mizuki-Akiyama",
		Colors: []color.Color{
			color.RGBA{R: 230, G: 166, B: 200, A: 255},
			color.RGBA{R: 43, G: 20, B: 34, A: 255},
			color.RGBA{R: 175, G: 162, B: 216, A: 255},
			color.RGBA{R: 29, G: 24, B: 48, A: 255},
			color.RGBA{R: 127, G: 182, B: 214, A: 255},
			color.RGBA{R: 7, G: 31, B: 45, A: 255},
			color.RGBA{R: 240, G: 138, B: 155, A: 255},
			color.RGBA{R: 45, G: 11, B: 20, A: 255},
			color.RGBA{R: 16, G: 17, B: 27, A: 255},
			color.RGBA{R: 238, G: 231, B: 240, A: 255},
			color.RGBA{R: 29, G: 26, B: 42, A: 255},
			color.RGBA{R: 200, G: 185, B: 202, A: 255},
			color.RGBA{R: 77, G: 70, B: 92, A: 255},
			color.RGBA{R: 5, G: 6, B: 10, A: 255},
			color.RGBA{R: 59, G: 48, B: 78, A: 255},
			color.RGBA{R: 246, G: 234, B: 243, A: 255},
			color.RGBA{R: 143, G: 208, B: 184, A: 255},
			color.RGBA{R: 217, G: 184, B: 116, A: 255},
			color.RGBA{R: 216, G: 162, B: 203, A: 255},
			color.RGBA{R: 124, G: 206, B: 217, A: 255},
			color.RGBA{R: 255, G: 156, B: 175, A: 255},
			color.RGBA{R: 167, G: 222, B: 201, A: 255},
			color.RGBA{R: 239, G: 208, B: 141, A: 255},
			color.RGBA{R: 157, G: 206, B: 231, A: 255},
			color.RGBA{R: 230, G: 182, B: 213, A: 255},
			color.RGBA{R: 154, G: 221, B: 228, A: 255},
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
		},
	}

	Murata = Theme{
		Name: "Murata",
		Colors: []color.Color{
			color.RGBA{R: 219, G: 109, B: 109, A: 255},
			color.RGBA{R: 26, G: 27, B: 28, A: 255},
			color.RGBA{R: 209, G: 179, B: 148, A: 255},
			color.RGBA{R: 143, G: 174, B: 177, A: 255},
			color.RGBA{R: 205, G: 197, B: 189, A: 255},
			color.RGBA{R: 37, G: 38, B: 39, A: 255},
			color.RGBA{R: 166, G: 157, B: 150, A: 255},
			color.RGBA{R: 90, G: 91, B: 92, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 145, G: 173, B: 131, A: 255},
			color.RGBA{R: 184, G: 146, B: 177, A: 255},
			color.RGBA{R: 143, G: 184, B: 167, A: 255},
			color.RGBA{R: 62, G: 64, B: 66, A: 255},
			color.RGBA{R: 229, G: 142, B: 142, A: 255},
			color.RGBA{R: 176, G: 204, B: 157, A: 255},
			color.RGBA{R: 230, G: 204, B: 177, A: 255},
			color.RGBA{R: 172, G: 198, B: 201, A: 255},
			color.RGBA{R: 209, G: 177, B: 204, A: 255},
			color.RGBA{R: 172, G: 201, B: 188, A: 255},
			color.RGBA{R: 229, G: 223, B: 217, A: 255},
		},
	}

	NoctaliaLegacy = Theme{
		Name: "Noctalia-Legacy",
		Colors: []color.Color{
			color.RGBA{R: 199, G: 161, B: 216, A: 255},
			color.RGBA{R: 26, G: 21, B: 31, A: 255},
			color.RGBA{R: 169, G: 132, B: 196, A: 255},
			color.RGBA{R: 243, G: 237, B: 247, A: 255},
			color.RGBA{R: 224, G: 183, B: 201, A: 255},
			color.RGBA{R: 32, G: 22, B: 31, A: 255},
			color.RGBA{R: 233, G: 137, B: 157, A: 255},
			color.RGBA{R: 30, G: 20, B: 24, A: 255},
			color.RGBA{R: 28, G: 24, B: 34, A: 255},
			color.RGBA{R: 233, G: 228, B: 240, A: 255},
			color.RGBA{R: 38, G: 33, B: 48, A: 255},
			color.RGBA{R: 167, G: 154, B: 176, A: 255},
			color.RGBA{R: 62, G: 54, B: 78, A: 255},
			color.RGBA{R: 18, G: 15, B: 24, A: 255},
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
		},
	}

	OsakaJade = Theme{
		Name: "Osaka-Jade",
		Colors: []color.Color{
			color.RGBA{R: 30, G: 145, B: 119, A: 255},
			color.RGBA{R: 184, G: 200, B: 196, A: 255},
			color.RGBA{R: 22, G: 122, B: 99, A: 255},
			color.RGBA{R: 38, G: 165, B: 137, A: 255},
			color.RGBA{R: 147, G: 54, B: 54, A: 255},
			color.RGBA{R: 8, G: 21, B: 18, A: 255},
			color.RGBA{R: 166, G: 181, B: 177, A: 255},
			color.RGBA{R: 15, G: 37, B: 31, A: 255},
			color.RGBA{R: 153, G: 168, B: 164, A: 255},
			color.RGBA{R: 27, G: 99, B: 82, A: 255},
			color.RGBA{R: 4, G: 10, B: 9, A: 255},
			color.RGBA{R: 20, G: 27, B: 30, A: 255},
			color.RGBA{R: 218, G: 218, B: 218, A: 255},
			color.RGBA{R: 35, G: 42, B: 45, A: 255},
			color.RGBA{R: 229, G: 116, B: 116, A: 255},
			color.RGBA{R: 140, G: 207, B: 126, A: 255},
			color.RGBA{R: 229, G: 199, B: 107, A: 255},
			color.RGBA{R: 103, G: 176, B: 232, A: 255},
			color.RGBA{R: 196, G: 127, B: 213, A: 255},
			color.RGBA{R: 108, G: 191, B: 191, A: 255},
			color.RGBA{R: 179, G: 185, B: 184, A: 255},
			color.RGBA{R: 70, G: 78, B: 80, A: 255},
			color.RGBA{R: 239, G: 126, B: 126, A: 255},
			color.RGBA{R: 150, G: 217, B: 136, A: 255},
			color.RGBA{R: 244, G: 214, B: 122, A: 255},
			color.RGBA{R: 113, G: 186, B: 242, A: 255},
			color.RGBA{R: 206, G: 137, B: 223, A: 255},
			color.RGBA{R: 103, G: 203, B: 231, A: 255},
			color.RGBA{R: 189, G: 195, B: 194, A: 255},
		},
	}

	Shien = Theme{
		Name: "Shien",
		Colors: []color.Color{
			color.RGBA{R: 155, G: 139, B: 193, A: 255},
			color.RGBA{R: 21, G: 19, B: 27, A: 255},
			color.RGBA{R: 120, G: 102, B: 163, A: 255},
			color.RGBA{R: 93, G: 80, B: 124, A: 255},
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
			color.RGBA{R: 184, G: 122, B: 135, A: 255},
			color.RGBA{R: 191, G: 179, B: 219, A: 255},
			color.RGBA{R: 36, G: 32, B: 44, A: 255},
			color.RGBA{R: 70, G: 61, B: 92, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 122, G: 168, B: 159, A: 255},
			color.RGBA{R: 196, G: 183, B: 148, A: 255},
			color.RGBA{R: 207, G: 147, B: 159, A: 255},
			color.RGBA{R: 145, G: 191, B: 182, A: 255},
			color.RGBA{R: 222, G: 212, B: 182, A: 255},
			color.RGBA{R: 239, G: 238, B: 241, A: 255},
		},
	}

	Shinonome = Theme{
		Name: "Shinonome",
		Colors: []color.Color{
			color.RGBA{R: 213, G: 172, B: 169, A: 255},
			color.RGBA{R: 26, G: 29, B: 32, A: 255},
			color.RGBA{R: 179, G: 141, B: 151, A: 255},
			color.RGBA{R: 197, G: 186, B: 175, A: 255},
			color.RGBA{R: 235, G: 207, B: 178, A: 255},
			color.RGBA{R: 45, G: 51, B: 57, A: 255},
			color.RGBA{R: 66, G: 75, B: 84, A: 255},
			color.RGBA{R: 0, G: 0, B: 0, A: 255},
			color.RGBA{R: 151, G: 167, B: 181, A: 255},
			color.RGBA{R: 201, G: 162, B: 171, A: 255},
			color.RGBA{R: 226, G: 195, B: 193, A: 255},
			color.RGBA{R: 243, G: 223, B: 202, A: 255},
			color.RGBA{R: 176, G: 194, B: 212, A: 255},
			color.RGBA{R: 220, G: 213, B: 206, A: 255},
			color.RGBA{R: 253, G: 251, B: 247, A: 255},
		},
	}

	TokyoNightMoon = Theme{
		Name: "Tokyo-Night-Moon",
		Colors: []color.Color{
			color.RGBA{R: 122, G: 136, B: 207, A: 255},
			color.RGBA{R: 31, G: 35, B: 53, A: 255},
			color.RGBA{R: 215, G: 114, B: 159, A: 255},
			color.RGBA{R: 156, G: 213, B: 138, A: 255},
			color.RGBA{R: 247, G: 118, B: 142, A: 255},
			color.RGBA{R: 169, G: 177, B: 214, A: 255},
			color.RGBA{R: 44, G: 49, B: 74, A: 255},
			color.RGBA{R: 192, G: 202, B: 245, A: 255},
			color.RGBA{R: 75, G: 81, B: 122, A: 255},
			color.RGBA{R: 24, G: 27, B: 42, A: 255},
			color.RGBA{R: 230, G: 195, B: 132, A: 255},
			color.RGBA{R: 123, G: 176, B: 192, A: 255},
			color.RGBA{R: 130, G: 137, B: 166, A: 255},
			color.RGBA{R: 84, G: 92, B: 126, A: 255},
		},
	}

	Nord = Theme{
		Name: "Nord",
		Colors: []color.Color{
			color.RGBA{R: 180, G: 142, B: 173, A: 255},
			color.RGBA{R: 46, G: 52, B: 64, A: 255},
			color.RGBA{R: 163, G: 190, B: 140, A: 255},
			color.RGBA{R: 235, G: 203, B: 139, A: 255},
			color.RGBA{R: 191, G: 97, B: 106, A: 255},
			color.RGBA{R: 236, G: 239, B: 244, A: 255},
			color.RGBA{R: 59, G: 66, B: 82, A: 255},
			color.RGBA{R: 229, G: 233, B: 240, A: 255},
			color.RGBA{R: 67, G: 76, B: 94, A: 255},
			color.RGBA{R: 76, G: 86, B: 106, A: 255},
			color.RGBA{R: 216, G: 222, B: 233, A: 255},
			color.RGBA{R: 198, G: 208, B: 245, A: 255},
			color.RGBA{R: 94, G: 129, B: 172, A: 255},
			color.RGBA{R: 244, G: 184, B: 228, A: 255},
			color.RGBA{R: 143, G: 188, B: 187, A: 255},
		},
	}

	RosePine = Theme{
		Name: "Rose-Pine",
		Colors: []color.Color{
			color.RGBA{R: 234, G: 154, B: 151, A: 255},
			color.RGBA{R: 35, G: 33, B: 54, A: 255},
			color.RGBA{R: 156, G: 207, B: 216, A: 255},
			color.RGBA{R: 62, G: 143, B: 176, A: 255},
			color.RGBA{R: 224, G: 222, B: 244, A: 255},
			color.RGBA{R: 235, G: 111, B: 146, A: 255},
			color.RGBA{R: 57, G: 53, B: 82, A: 255},
			color.RGBA{R: 144, G: 140, B: 170, A: 255},
			color.RGBA{R: 68, G: 65, B: 90, A: 255},
			color.RGBA{R: 86, G: 82, B: 110, A: 255},
			color.RGBA{R: 246, G: 193, B: 119, A: 255},
			color.RGBA{R: 196, G: 167, B: 231, A: 255},
			color.RGBA{R: 110, G: 106, B: 134, A: 255},
		},
	}
)
