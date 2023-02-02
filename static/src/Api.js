import axios from "axios";

if (!window.axios) {
    window.axios = axios.create({
        baseURL: 'http://localhost:8080/monitor/api/v1/'
    });
}

const Api = {
    system: async () => {
        const data = await window.axios.get('/system');

        return data.data;
    },
    history: async () => {
        const data = await window.axios.get('/history');

        return data.data;
    }
};

export default Api;
