package bot

type Speaker interface {
	SayHello()
	currentTime()
	currentDate()
	today()
	bye()
	misUnderstanding()
	initLang() []string
}

func Listner(s Speaker, reply string) {
	cmnds := s.initLang()

	switch reply {
	case cmnds[0], "1":
		s.SayHello()
	case cmnds[1], "2":
		s.currentTime()
	case cmnds[2], "3":
		s.currentDate()
	case cmnds[3], "4":
		s.today()
	case cmnds[4], "5":
		s.bye()
	default:
		s.misUnderstanding()
	}

}