package test
import
(
	"strings"
	"testing"
)

func TestImportConverter(t *testing.T) {
		eachLine := "import MidAdBanner from '../components/home/MidAdBanner';"
		replacedImport := strings.ReplaceAll(eachLine,"import","const")
		replacedImport = strings.ReplaceAll(replacedImport,"from ","= dynamic(() => import(")
		replacedImport = strings.ReplaceAll(replacedImport, ";","")
		replacedImport = replacedImport + "));"	

		if replacedImport != "const MidAdBanner = dynamic(() => import('../components/home/MidAdBanner'));" {
			t.Errorf("Output is incorrect, got: %s, want: %s.", replacedImport, "const MidAdBanner = dynamic(() => import('../components/home/MidAdBanner'));\nconst BottomAdBanner = dynamic(() => import('../components/home/BottomAdBanner'));")
		}
}