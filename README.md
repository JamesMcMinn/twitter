# Twitter Utilities
A set of tools and utilities for dealing with Tweets and Twitter content.

## Requirements
The following packages are required:
 - github.com/mrjones/oauth
 - github.com/araddon/httpstream

## Getting started
To fill a channel with data from the Twitter firehose:

    firehose chan twitter.Tweet = make(chan twitter.Tweet, 1000)
    go twitter.FillStream(firehose, "consumer key", "consumer secret", "oauth token", "oauth secret")