# twitterStreaming

twitterStreaming uses [anaconda][anaconda] library to fetch tweets from twitter's 
streaming api. Anaconda is an open source(MIT license) library in golang for accessing
Twitter api. I chose this set of tools(golang & anaconda) because golang is reasonably fast
and has nice support for concurrency. Anaconda is a nice library that also supports the twitter streaming
api. 

Objective of this project is to setup a system fetch tweets for *sentiment analysis*.
As sentimet analysis is the focus, many attributes of a [tweet object][tweet_obj], that do not add 
value to sentiment analysis have been ignored.

[anaconda]: https://github.com/ChimeraCoder/anaconda
[tweet_obj]: https://dev.twitter.com/overview/api/tweets
