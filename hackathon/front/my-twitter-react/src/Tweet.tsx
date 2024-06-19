import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface Tweet {
  id: number;
  content: string;
}

const Tweet: React.FC = () => {
  const [tweets, setTweets] = useState<Tweet[]>([]);
  const [tweetContent, setTweetContent] = useState<string>('');

  useEffect(() => {
    axios.get<Tweet[]>('https://your-backend-url/tweets')
      .then(response => {
        setTweets(response.data);
      })
      .catch(error => {
        console.error(error);
      });
  }, []);

  const handleTweet = () => {
    axios.post<Tweet>('https://your-backend-url/tweets', { content: tweetContent })
      .then(response => {
        setTweets([...tweets, response.data]);
        setTweetContent('');
      })
      .catch(error => {
        console.error(error);
      });
  };

  return (
    <div>
      <input type="text" value={tweetContent} onChange={(e) => setTweetContent(e.target.value)} placeholder="What's happening?" />
      <button onClick={handleTweet}>Tweet</button>
      <div>
        {tweets.map((tweet) => (
          <div key={tweet.id}>{tweet.content}</div>
        ))}
      </div>
    </div>
  );
};

export default Tweet;

