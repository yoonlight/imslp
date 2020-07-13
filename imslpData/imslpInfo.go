package imslpdata

// IMSLPInfo parse data
type IMSLPInfo struct {
	title      string
	compose    string
	style      string
	instrument Instrument
}

// Instrument type is parsing instrument and save
type Instrument struct {
	strings    string
	WoodWind   WoodWind
	BrassWind  BrassWind
	Percussion Percussion
}

// WoodWind save
type WoodWind struct {
	piccolo       string
	flute         string
	oboe          string
	englishHorn   string
	clarinet      string
	bassClarinet  string
	bassoon       string
	contraBassoon string
}

// BrassWind save
type BrassWind struct {
	horn     string
	trumpet  string
	trombone string
	tuba     string
}

// Percussion save
type Percussion struct {
	triangle  string
	cymbals   string
	bassDrum  string
	snareDrum string
	organ     string
	piano     string
	timpani   string
	harp      string
}

// ImportInfo :
func ImportInfo(m map[string]string) Instrument {
	var (
		WoodWind   WoodWind
		BrassWind  BrassWind
		Percussion Percussion
	)
	WoodWind.piccolo = m["piccolo"]
	WoodWind.flute = m["flute"]
	WoodWind.oboe = m["oboe"]
	WoodWind.englishHorn = m["English horn"]
	WoodWind.clarinet = m["clarinet"]
	WoodWind.bassClarinet = m["bass clarinet"]
	WoodWind.bassoon = m["bassoons"]
	WoodWind.contraBassoon = m["contra bassoon"]
	BrassWind.horn = m["horns"]
	BrassWind.trumpet = m["trumpets"]
	BrassWind.trombone = m["trombones"]
	BrassWind.tuba = m["tuba"]
	Percussion.triangle = m["triangle"]
	Percussion.cymbals = m["cymbals"]
	Percussion.bassDrum = m["bass drum"]
	Percussion.snareDrum = m["snare drum"]
	Percussion.organ = m["organ"]
	Percussion.piano = m["piano"]
	Percussion.timpani = m["timpani"]
	Percussion.harp = m["harp"]

	return Instrument{
		strings:    "1",
		WoodWind:   WoodWind,
		BrassWind:  BrassWind,
		Percussion: Percussion,
	}

}
