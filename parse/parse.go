package parse

func parseInstr(instrument string) imslpInfo.Instrument {
	var category string
	// reg := `(?m)(?s)flutes|oboes|English horn|clarinets|bass clarinet|bassoons|contrabassoon|horns|trumpets|trombones|tuba|timpani|triangle|cymbals|bass drum|organ|piano|strings `
	reg2 := `(?m)(?s)(\d.)(flutes|oboes|English horn|clarinets|bass clarinet|bassoons|contrabassoon|horns|trumpets|trombones|tuba|timpani|triangle|cymbals|bass drum|organ|piano|strings)`
	var re = regexp.MustCompile(reg2)
	for _, match := range re.FindAllStringSubmatch(instrument, -1) {
		num := match[1]
		name := match[2]
		switch name {
		case "flutes":
			category.WoodWind.flute = num
		case "oboes":
			category.WoodWind.oboe = num
		case "English horn":
			category.WoodWind.englishHorn = num
		case "clarinets":
			category.WoodWind.clarinet = num
		case "bass clarinet":
			category.WoodWind.bassClarinet = num
		case "bassoons":
			category.WoodWind.contraBassoon = num
		case "horns":
			category.BrassWind.horn = num
		case "trumpets":
			category.BrassWind.trumpet = num
		case "trombones":
			category.BrassWind.trombone = num
		case "tuba":
			category.BrassWind.tuba = num

		}
	}
	return category
}
