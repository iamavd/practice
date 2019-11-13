package bot

type Speaker interface {
	SayHello()
	CurrentTime()
	CurrentDate()
	Today()
	Bye()
	misUnderstanding()
	initLang() []string
	Introduce()
}

func Listner(s Speaker, reply string) {
	cmnds := s.initLang()

	switch reply {
	case cmnds[0], "1":
		s.SayHello()
	case cmnds[1], "2":
		s.CurrentTime()
	case cmnds[2], "3":
		s.CurrentDate()
	case cmnds[3], "4":
		s.Today()
	case cmnds[4], "5":
		s.Bye()
	default:
		s.misUnderstanding()
	}

}
