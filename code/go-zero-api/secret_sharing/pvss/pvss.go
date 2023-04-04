package pvss

import (
	"github.com/Secured-Finance/kyber"
	"github.com/Secured-Finance/kyber/group/edwards25519"
	"github.com/Secured-Finance/kyber/share"
	"github.com/Secured-Finance/kyber/share/pvss"
)

func Init() (suite *edwards25519.SuiteEd25519, G, H kyber.Point) {

	suite = edwards25519.NewBlakeSHA256Ed25519()
	G = suite.Point().Base()
	H = suite.Point().Pick(suite.RandomStream())
	return suite, G, H
}

func CreateKeyPair(suite *edwards25519.SuiteEd25519, n int) (x []kyber.Scalar, X []kyber.Point) {

	x = make([]kyber.Scalar, n)
	X = make([]kyber.Point, n)
	for i := 0; i < n; i++ {
		x[i] = suite.Scalar().Pick(suite.RandomStream())
		X[i] = suite.Point().Mul(x[i], nil)
	}
	return x, X
}

func CreateSecret(suite *edwards25519.SuiteEd25519) (secret kyber.Scalar) {

	secret = suite.Scalar().Pick(suite.RandomStream())
	return secret
}

func CreateEnShare(suite pvss.Suite, H kyber.Point, X []kyber.Point, secret kyber.Scalar, t int) (shares []*pvss.PubVerShare, commit *share.PubPoly, err error) {

	return pvss.EncShares(suite, H, X, secret, t)
}

func VerifyEncShare(suite pvss.Suite, H kyber.Point, X kyber.Point, commit *share.PubPoly, encShare *pvss.PubVerShare) (verified bool) {

	sH := commit.Eval(encShare.S.I).V
	verified = nil == pvss.VerifyEncShare(suite, H, X, sH, encShare)
	return verified
}

func DecShare(suite pvss.Suite, H kyber.Point, X kyber.Point, x kyber.Scalar, encShare *pvss.PubVerShare, commit *share.PubPoly) (*pvss.PubVerShare, error) {

	sH := commit.Eval(encShare.S.I).V
	return pvss.DecShare(suite, H, X, sH, x, encShare)
}

func RecoverSecret(suite pvss.Suite, G kyber.Point, X []kyber.Point, encShares []*pvss.PubVerShare, decShares []*pvss.PubVerShare, t int, n int) (kyber.Point, error) {

	return pvss.RecoverSecret(suite, G, X, encShares, decShares, t, n)
}

type KyberPoint = kyber.Point
type PubVerShare = pvss.PubVerShare
