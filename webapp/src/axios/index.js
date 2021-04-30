import axios from 'axios'
import store from './../store'

let api = axios.create({
    baseURL: 'http://localhost:8080/api/v1',
    headers: {
        'Authorization': `Bearer ${ store.getters.token }`
    }
})

api.interceptors.request.use((config) => {
    const token = store.getters.token
    console.log("ABCD TOKEN", token)
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }

    return config
})

export {
    api
}