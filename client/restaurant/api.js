const http = require('../common/http');

const createRestaurantApi = (baseUrl) => {
    const client = http.createHttpClient(baseUrl);
    return {
        getMenu: () => client.get('/menu'),
        createOrder: ({menuItemIds, tableId}) => client.post('/orders', {menuItemIds, tableId})
    }
};

module.exports = {createRestaurantApi};