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
	strings       string
	piccolo       string
	flute         string
	oboe          string
	englishHorn   string
	clarinet      string
	bassClarinet  string
	bassoon       string
	contraBassoon string
	horn          string
	trumpet       string
	trombone      string
	tuba          string
	triangle      string
	cymbals       string
	bassDrum      string
	snareDrum     string
	organ         string
	piano         string
	timpani       string
	harp          string
}

// Instruments array
type Instruments []Instrument

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
func ImportInfo(m map[string]string) (Instrument Instrument) {

	Instrument.strings = "1"
	Instrument.piccolo = m["piccolo"]
	Instrument.flute = m["flute"]
	Instrument.oboe = m["oboe"]
	Instrument.englishHorn = m["English horn"]
	Instrument.clarinet = m["clarinet"]
	Instrument.bassClarinet = m["bass clarinet"]
	Instrument.bassoon = m["bassoons"]
	Instrument.contraBassoon = m["contra bassoon"]
	Instrument.horn = m["horns"]
	Instrument.trumpet = m["trumpets"]
	Instrument.trombone = m["trombones"]
	Instrument.tuba = m["tuba"]
	Instrument.triangle = m["triangle"]
	Instrument.cymbals = m["cymbals"]
	Instrument.bassDrum = m["bass drum"]
	Instrument.snareDrum = m["snare drum"]
	Instrument.organ = m["organ"]
	Instrument.piano = m["piano"]
	Instrument.timpani = m["timpani"]
	Instrument.harp = m["harp"]

	return
}
