package imparse

import (
	imdata "imslp/imslpData"
	"regexp"
	"strconv"
)

// ParseInstr2 : 악기 별로 parsing 해줌
func ParseInstr2(instrument string) (m imdata.Instr) {
	reg := `(?m)(?s)flutes|oboes|English horn|clarinets|bass clarinet|bassoons|contrabassoon|horns|trumpets|trombones|tuba|timpani|triangle|cymbals|bass drum|organ|piano|harp|piccolo|english horn|side drum|tambourine|cannon|bells|trumpet|flute`
	reg2 := `(?m)(?s)(\d.)(flutes|oboes|English horn|clarinets|bass clarinet|bassoons|contrabassoon|horns|trumpets|trombones|tuba|timpani|triangle|cymbals|bass drum|organ|piano|harp|piccolo|english horn|side drum|tambourine|cannon|bells|trumpet|flute)`
	var re = regexp.MustCompile(reg)
	for _, match := range re.FindAllStringSubmatch(instrument, -1) {
		num := strconv.Itoa(1)
		name := match[0]
		m = caseParse2(num, name, m)
	}
	var re2 = regexp.MustCompile(reg2)
	for _, match := range re2.FindAllStringSubmatch(instrument, -1) {
		num := match[1]
		name := match[2]
		m = caseParse2(num, name, m)
	}
	return m
}

func caseParse2(num string, name string, m imdata.Instr) imdata.Instr {
	switch name {
	case "piccolo":
		m.Wood.Piccolo = num
	case "flutes", "flute":
		m.Wood.Flute = num
	case "oboes":
		m.Wood.Oboe = num
	case "English horn", "english horn":
		m.Wood.EnglishHorn = num
	case "clarinets":
		m.Wood.Clarinet = num
	case "bass clarinet":
		m.Wood.BassClarinet = num
	case "bassoons":
		m.Wood.Bassoon = num
	case "contrabassoon":
		m.Wood.ContraBassoon = num
	case "horns":
		m.Brass.Horn = num
	case "trumpets", "trumpet":
		m.Brass.Trumpet = num
	case "trombones":
		m.Brass.Trombone = num
	case "tuba":
		m.Brass.Tuba = num
	case "triangle":
		m.Perc.Triangle = num
	case "cymbals":
		m.Perc.Cymbals = num
	case "bass drum":
		m.Perc.BassDrum = num
	case "snare drum", "side drum":
		m.Perc.SnareDrum = num
	case "organ":
		m.Perc.Organ = num
	case "piano":
		m.Perc.Piano = num
	case "timpani":
		m.Perc.Timpani = num
	case "harp":
		m.Perc.Harp = num
	case "tambourine":
		m.Perc.Tambourine = num
	case "bells":
		m.Perc.Bells = num
	case "cannon":
		m.Perc.Cannon = num
	}
	return m

}
