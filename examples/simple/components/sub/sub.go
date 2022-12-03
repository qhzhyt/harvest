package sub0

import "net"

// @component
type (
	// @Sub
	Sub struct{}
	// @add
	add struct{}
	// @type
	number int64
	// @type
	handlerFunc func(sub *Sub) bool
	// @type
	netIP net.IP
	// @type
	starSub *Sub
	// @type
	starIP *net.IP
	// @type
	subSub Sub
	// @type
	starStarSubSub **subSub
)

// @test
func test() {

}

func (t *add) test() {

}
