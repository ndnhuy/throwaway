import { Grid, Stack, TextField, Typography } from "@mui/material";
import { Box } from "@mui/system";
import axios from "axios";
import { useEffect, useState } from "react";

type Message = {
  fromUser: string;
  toUser: string;
  content: string;
}
export const axiosClient = axios.create({
  baseURL: 'http://localhost:8888',
  headers: {
    'Content-Type': 'application/json'
  }
});
export default function ChatHistory() {
  const [history, setHistory] = useState<Message[]>([]);
  const [fetchCount, setFetchCount] = useState(0);
  const [message, setMessage] = useState('');
  const [currentUser, setCurrentUser] = useState('');
  const [toUser, setToUser] = useState('');

  useEffect(() => {
    setMessage('');
    axiosClient.get('/messages')
      .then(res => {
        setHistory(res.data);
      })
  }, [currentUser, fetchCount]);

  useEffect(() => {
    if (history && history.length > 0) {
      if (!currentUser || currentUser == '') {
        setCurrentUser(history[0].fromUser);
        setToUser(history[0].toUser);
      }
    }
  }, [history])

  const handleEnter = (e) => {
    if (e?.key === 'Enter') {
      axiosClient
        .post('/send', {
          "fromUser": currentUser,
          "toUser": '',
          "content": message,
          "createdAt": "2023-01-24"
        })
        .then(res => {
          setFetchCount(fetchCount+1);
        })
    }
  }

  return (
    <Box sx={{ p: 2.5 }}>
      <Grid container spacing={3}>
        <Grid item xs={12} sm={6}>
          <Grid item xs={12}>
            <label htmlFor='current-user'>current user: </label>
            <select id= 'current-user' name='current-user' value={currentUser} onChange={e => setCurrentUser(e.target.value)}>
              {[...new Set(history.map(h => h.fromUser))].map((user, index) => (
                <option key={index} value={user}>{user}</option>
              ))}
            </select>
          </Grid><Grid item xs={12}>
            <label htmlFor='to-user'>to user: </label>
            <select id= 'to-user' name='to-user' value={toUser} onChange={e => setToUser(e.target.value)}>
              {[...new Set(history.map(h => h.toUser))].map((user, index) => (
                <option key={index} value={user}>{user}</option>
              ))}
            </select>
          </Grid>
          <Grid container>
            {history.map((msg, index) => (
              <Grid item xs={12} key={index}>
                {msg.fromUser === currentUser ? (
                  <Stack direction="row" justifyContent="flex-end" alignItems="flex-start">
                    <Typography>{`${msg.fromUser}: ${msg.content}`}</Typography>
                  </Stack>
                ) : msg.fromUser === toUser ? (
                  <Stack direction="row" justifyContent="flex-start" alignItems="flex-start">
                    <Typography>{`${msg.fromUser}: ${msg.content}`}</Typography>
                  </Stack>
                ) : (<></>)}
              </Grid>
            ))}
            <Grid
              item
              xs={12}
              sx={{ mt: 3, borderTop: `1px solid` }}
            >
              <TextField
                fullWidth
                multiline
                rows={4}
                placeholder="Your Message..."
                value={message}
                onChange={(e) => setMessage(e.target.value.length <= 1 ? e.target.value.trim() : e.target.value)}
                onKeyPress={handleEnter}
                variant="standard"
                sx={{
                  pr: 2
                }}
              />
            </Grid>
          </Grid>
        </Grid>
      </Grid>
    </Box>
  )
}
