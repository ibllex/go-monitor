import axios from "axios";

if (!window.axios) {
    const loc = window.location.toString()

    window.axios = axios.create({
        baseURL: loc.substring(0, loc.length - 1) + '-api'  + '/v1/'
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
