package vnc_rec

type Recorder interface {
	StartSession(interface{}) error
	Consume(interface{}) error
	HandleSegment(interface{}) error
}
