import { useState } from 'react'
import axios, { config } from './Constants/axios'

import Button from '@mui/material/Button';

function App() {
    const [connection, setConnection] = useState('Not Connected')

    const connectionInit = () => {
        console.log('test')
        axios.get('/connection', config).then(response => {
            const { data } = response;

            if(data.status) {
                setConnection('Connection Available')
            } else {
                setConnection('Connection Unavailable')
            }
        })
    }

    return (
        <div className="App">
          <header className="App-header">
              <h1>Test app with React and Golang</h1>
            <p>
                Connection Status: {connection}
            </p>
              <Button variant="contained" onClick={connectionInit}>Click to connect to server</Button>
          </header>
        </div>
    );
}

export default App;
