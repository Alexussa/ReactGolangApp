import axios from 'axios'

export const config = {
    port: 3030,
    hostname: 'localhost'
}

export default axios.create(config)