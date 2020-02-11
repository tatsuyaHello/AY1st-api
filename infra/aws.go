package infra

import (
	"sync"
	"time"

	"AY1st/util"

	"github.com/aws/aws-sdk-go/aws/session"
)

type awsSessionCache struct {
	awsSession *session.Session
	mut        sync.Mutex
	expiresAt  time.Time
}

var awsSessionCacheData awsSessionCache

// GetAWSSession はAWSセッションを取得します
// https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
func GetAWSSession() *session.Session {
	logger := util.GetLogger()
	awsSessionCacheData.mut.Lock()
	defer awsSessionCacheData.mut.Unlock()

	if awsSessionCacheData.awsSession == nil {
		awsSessionCacheData.awsSession, awsSessionCacheData.expiresAt = newAwsSession()
	}

	// session の有効期限が1分前なら再取得
	if awsSessionCacheData.expiresAt.Add(-1 * time.Minute).Before(time.Now().UTC()) {
		awsSessionCacheData.awsSession, awsSessionCacheData.expiresAt = newAwsSession()
		logger.Debugf("AWS Session retrieved")
	} else {
		logger.Debugf("use cached AWS Session")
	}

	return awsSessionCacheData.awsSession
}

func newAwsSession() (s *session.Session, expiresAt time.Time) {
	s = session.Must(session.NewSession())

	expiresAt, err := s.Config.Credentials.ExpiresAt()
	if err != nil {
		// credentialsから有効期限が取得出来ない場合は、15分後とする
		expiresAt = util.GetTimeNow().Add(15 * time.Minute)
	}
	if err == nil {
		util.GetLogger().Debugf("AWS Session ExpiresAt:%v", expiresAt)
	}
	return s, expiresAt
}
