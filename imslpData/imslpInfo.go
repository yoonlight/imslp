package imslpdata

// IMSLPInfo parse data
type IMSLPInfo struct {
	title      string
	compose    string
	style      string
	instrument string
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
}
