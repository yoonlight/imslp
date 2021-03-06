package imparse

import (
	"regexp"
	"strconv"
)

// ParseInstr : 악기 별로 parsing 해줌
func ParseInstr(instrument string) (m map[string]string) {
	m = make(map[string]string)
	reg := `(?m)(?s)flutes|oboes|English horn|clarinets|bass clarinet|bassoons|contrabassoon|horns|trumpets|trombones|tuba|timpani|triangle|cymbals|bass drum|organ|piano|harp|piccolo|english horn|side drum|tambourine|cannon|bells`
	reg2 := `(?m)(?s)(\d.)(flutes|oboes|English horn|clarinets|bass clarinet|bassoons|contrabassoon|horns|trumpets|trombones|tuba|timpani|triangle|cymbals|bass drum|organ|piano|harp|piccolo|english horn|side drum|tambourine|cannon|bells)`
	var re = regexp.MustCompile(reg)
	for _, match := range re.FindAllStringSubmatch(instrument, -1) {
		num := strconv.Itoa(1)
		name := match[0]
		m = caseParse(num, name, m)

	}
	var re2 = regexp.MustCompile(reg2)
	for _, match := range re2.FindAllStringSubmatch(instrument, -1) {
		num := match[1]
		name := match[2]
		m = caseParse(num, name, m)
	}
	return m
}

func caseParse(num string, name string, m map[string]string) (b map[string]string) {
	switch name {
	case "piccolo":
		m["piccolo"] = num
	case "flutes":
		m["flute"] = num
	case "oboes":
		m["oboe"] = num
	case "English horn", "english horn":
		m["English horn"] = num
	case "clarinets":
		m["clarinet"] = num
	case "bass clarinet":
		m["bass clarinet"] = num
	case "bassoons":
		m["bassoons"] = num
	case "contrabassoon":
		m["contra bassoon"] = num
	case "horns":
		m["horns"] = num
	case "trumpets":
		m["trumpets"] = num
	case "trombones":
		m["trombones"] = num
	case "tuba":
		m["tuba"] = num
	case "triangle":
		m["triangle"] = num
	case "cymbals":
		m["cymbals"] = num
	case "bass drum":
		m["bass drum"] = num
	case "snare drum", "side drum":
		m["snare drum"] = num
	case "organ":
		m["organ"] = num
	case "piano":
		m["piano"] = num
	case "timpani":
		m["timpani"] = num
	case "harp":
		m["harp"] = num
	case "tambourine":
		m["tambourine"] = num
	case "bells":
		m["bells"] = num
	case "cannon":
		m["cannon"] = num
	}
	b = m
	return

}

func woodParse(num string, name string, m map[string]string) (b map[string]string) {
	switch name {
	case "piccolo":
		m["piccolo"] = num
	case "flutes":
		m["flute"] = num
	case "oboes":
		m["oboe"] = num
	case "English horn", "english horn":
		m["English horn"] = num
	case "clarinets":
		m["clarinet"] = num
	case "bass clarinet":
		m["bass clarinet"] = num
	case "bassoons":
		m["bassoons"] = num
	case "contrabassoon":
		m["contra bassoon"] = num
	}
	b = m
	return
}

func brassParse(num string, name string, m map[string]string) (b map[string]string) {
	switch name {
	case "horns":
		m["horns"] = num
	case "trumpets":
		m["trumpets"] = num
	case "trombones":
		m["trombones"] = num
	case "tuba":
		m["tuba"] = num
	}
	b = m
	return

}

func percParse(num string, name string, m map[string]string) (b map[string]string) {
	switch name {
	case "triangle":
		m["triangle"] = num
	case "cymbals":
		m["cymbals"] = num
	case "bass drum":
		m["bass drum"] = num
	case "snare drum", "side drum":
		m["snare drum"] = num
	case "organ":
		m["organ"] = num
	case "piano":
		m["piano"] = num
	case "timpani":
		m["timpani"] = num
	case "harp":
		m["harp"] = num
	case "tambourine":
		m["tambourine"] = num
	case "bells":
		m["bells"] = num
	case "cannon":
		m["cannon"] = num
	}
	b = m
	return
}
