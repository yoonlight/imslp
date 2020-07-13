package imslpsql

import (
	"database/sql"
	errcheck "imslp/ErrorCheck"

	_ "github.com/mattn/go-sqlite3"
)

// IMSLPSql :
func IMSLPSql() {
	db, err := sql.Open("sqlite3", "file:imslp.db")

	errcheck.CheckError(err, "db start error")
	// db.SetMaxOpenConns(1)
	defer db.Close()

	// var imslpInfo IMSLPInfo
	// imslpInfo.title, imslpInfo.compose, imslpInfo.style, imslpInfo.instrument = crawler.IMSLP(saintSaens)

	// instrument := imslpInfo.instrument
	// ins := parseInstr(instrument)
	// fmt.Println(ins)
	// fmt.Println("title:", imslpInfo.title, "compose:", imslpInfo.compose, "style:", imslpInfo.style)
	// // db.Exec("CREATE TABLE `imslp` (`id` int(11) INTEGER PRIMARY KEY AUTOINCREMENT, `title` varchar(100) NOT NULL,`compose` varchar(100) NOT NULL,`style` varchar(100) NOT NULL, `instrument` varchar(100) NOT NULL,PRIMARY KEY (`id`))")
	// db.Exec("CREATE TABLE `imslp` (`title` varchar(100) NOT NULL,`compose` varchar(100) NOT NULL,`style` varchar(100) NOT NULL, `instrument` varchar(100) NOT NULL)")
	// db.Exec("CREATE TABLE `woodwind` (`flute` varchar(100) NOT NULL,`oboe` varchar(100) NOT NULL,`english horn` varchar(100) NOT NULL, `clarinet` varchar(100) NOT NULL, `bass clarinet` varchar(100) NOT NULL, `bassoon` varchar(100) NOT NULL, `contra bassoon` varchar(100) NOT NULL)")
	// db.Exec("CREATE TABLE `brasswind` (`horn` varchar(100) NOT NULL,`trumpet` varchar(100) NOT NULL,`trombone` varchar(100) NOT NULL, `tuba` varchar(100) NOT NULL)")
	// db.Exec("CREATE TABLE `percussion` (`triangle` varchar(100) NOT NULL,`cymbals` varchar(100) NOT NULL,`bass drum` varchar(100) NOT NULL, `timpani` varchar(100) NOT NULL)")
	// insertInfo, es := db.Prepare("INSERT INTO imslp VALUES (?,?,?,?)")
	// insertWood, es := db.Prepare("INSERT INTO woodwind VALUES (?,?,?,?,?,?,?)")
	// insertBrass, es := db.Prepare("INSERT INTO brasswind VALUES (?,?,?,?)")
	// insertPerc, es := db.Prepare("INSERT INTO percussion VALUES (?,?,?,?)")
	// errcheck.CheckError(es, "prepare err")
	// _, errins := insertInfo.Exec(imslpInfo.title, imslpInfo.compose, imslpInfo.style, ins.WoodWind.flute)
	// _, errins = insertWood.Exec(ins.WoodWind.flute, ins.WoodWind.oboe, ins.WoodWind.englishHorn, ins.WoodWind.clarinet, ins.WoodWind.bassClarinet, ins.WoodWind.bassoon, ins.WoodWind.contraBassoon)
	// _, errins = insertBrass.Exec(ins.BrassWind.horn, ins.BrassWind.trumpet, ins.BrassWind.trombone, ins.BrassWind.tuba)
	// _, errins = insertPerc.Exec(ins.Percussion.triangle, ins.Percussion.cymbals, ins.Percussion.bassDrum, ins.Percussion.timpani)
	// errcheck.CheckError(errins, "insert data error")
	// defer insertInfo.Close()
	// defer insertWood.Close()
	// defer insertBrass.Close()
	// defer insertPerc.Close()
}
