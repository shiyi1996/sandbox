package judge

import "testing"

func TestNotify(t *testing.T) {
	var judger Judger
	judger.notify(Result{})
}
