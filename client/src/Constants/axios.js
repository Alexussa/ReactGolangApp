import axios from 'axios'

export const config = {
    port: 8000,
    hostname: 'localhost'
}

export default axios.create(config)