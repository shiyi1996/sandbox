/**
 * Created by shiyi on 2017/12/21.
 * Email: shiyi@fightcoder.com
 */

package judge

import (
	"fmt"
	"testing"
)

func TestGetCaseList(t *testing.T) {
	//caseList := getCaseList("//Users/shiyi/project/fightcoder-sandbox/case")
	caseList := getCaseList(getCurrentPath() + "/case")
	fmt.Println(caseList)
}
