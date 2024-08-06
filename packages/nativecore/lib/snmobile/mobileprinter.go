package snmobile

import "log"

type MobilePrinter interface {
	WriteToMobilePrinter(input string)
}

type internalWriter struct {
	p MobilePrinter
}

func (w internalWriter) Write(p []byte) (n int, err error) {
	w.p.WriteToMobilePrinter(string(p))
	return len(p), nil
}

func (s *SolidNativeMobile) SetPrinter(p MobilePrinter) {
	i := internalWriter{p}

	// For some reason the stderr prints twice. Use the Stout
	// TODO: Make an interface to get a printer and make a writer for it.

	log.SetOutput(i)
}
