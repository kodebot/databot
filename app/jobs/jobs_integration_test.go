// build integration

package jobs

import "testing"

func TestIntegrationLoadArticlesFromFeedsJob(t *testing.T) {
	// this retrieves data from internet and loads into the database
	// use this as a test harness for new feed
	job := LoadArticlesFromFeedsJob{
		FeedInfoPath: "../../conf/feed/ready/"}
	job.Run()
}
