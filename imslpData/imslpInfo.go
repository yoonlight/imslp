package imdata

// IMSLPInfo parse data
type IMSLPInfo struct {
	Title   string
	Compose string
	Style   string
	Instr   string
}

// Instr :using json
type Instr struct {
	Strings string `json:"strings"`
	Wood    Wood   `json:"wood"`
	Brass   Brass  `json:"brass"`
	Perc    Perc   `json:"perc"`
}

// Wood save 이거 사용안함.
type Wood struct {
	Piccolo       string `json:"piccolo"`
	Flute         string `json:"flute"`
	Oboe          string `json:"oboe"`
	EnglishHorn   string `json:"english horn"`
	Clarinet      string `json:"clarinet"`
	BassClarinet  string `json:"bass clarinet"`
	Bassoon       string `json:"bassoon"`
	ContraBassoon string `json:"contra bassoon"`
}

// Brass save 이거 사용안함.
type Brass struct {
	Horn     string `json:"horn"`
	Trumpet  string `json:"trumpet"`
	Trombone string `json:"trombone"`
	Tuba     string `json:"tuba"`
}

// Perc save 이거 사용안함.
type Perc struct {
	Triangle   string `json:"triangle"`
	Cymbals    string `json:"cymbal"`
	BassDrum   string `json:"bass drum"`
	SnareDrum  string `json:"snare drum"`
	Organ      string `json:"organ"`
	Piano      string `json:"piano"`
	Timpani    string `json:"timpani"`
	Harp       string `json:"harp"`
	Tambourine string `json:"tambourine"`
	Bells      string `json:"bells"`
	Cannon     string `json:"cannon"`
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
