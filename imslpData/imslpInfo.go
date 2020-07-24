package imdata

// IMSLPInfo parse data 이거 사용안함.
type IMSLPInfo struct {
	Title   string
	Compose string
	Style   string
	Instr   string
}

// Instrument type is parsing instrument and save 이거 사용안함.
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
	tambourine    string
	cannon        string
	bell          string
}

// Instruments array 이거 사용안함.
type Instruments []Instrument

// Wood save 이거 사용안함.
type Wood struct {
	piccolo       string
	flute         string
	oboe          string
	englishHorn   string
	clarinet      string
	bassClarinet  string
	bassoon       string
	contraBassoon string
}

// Brass save 이거 사용안함.
type Brass struct {
	horn     string
	trumpet  string
	trombone string
	tuba     string
}

// Perc save 이거 사용안함.
type Perc struct {
	triangle  string
	cymbals   string
	bassDrum  string
	snareDrum string
	organ     string
	piano     string
	timpani   string
	harp      string
}

// ImportInfo : 이거 사용안함.
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
