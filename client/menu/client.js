const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        getMenu: () => client.get('/menu'),
        createOrder: (dishIds, tableId) => client.post('/orders', { dishIds, tableId })
    }

};

module.exports = { Client };