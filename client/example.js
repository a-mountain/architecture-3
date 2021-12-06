const api = require('./restaurant/api');

const restaurantApi = api.createRestaurantApi('http://localhost:8080');

// Scenario 1: Display restaurant list.
restaurantApi.getMenu()
    .then((menu) => {
        console.log('=== Scenario 1 ===');
        console.log('Menu:');
        menu.forEach((dish) => {
            const { price, name } = dish;
            console.log({ name, price })
        });
    })
    .catch((e) => {
        console.log(`Problem with menu loading: ${e.message}`);
    });

// Scenario 2: Create new order.
restaurantApi.createOrder({ menuItemIds: [0, 2, 3], tableId: 3 })
    .then((order) => {
        console.log('=== Scenario 2 ===');
        console.log('Create order response:', order);
    })
    .catch((e) => {
        console.log(`Problem creating a new order: ${e.message}`);
    });
    