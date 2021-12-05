const channels = require('./menu/client');

const client = channels.Client('http://localhost:8080');

// Scenario 1: Display menu list.
client.getMenu()
    .then((menu) => {
        console.log('=== Scenario 1 ===');
        console.log('Menu:');
        list.forEach((dish) => {
            const { price, name } = dish;
            console.log({ name, price })
        });
    })
    .catch((e) => {
        console.log(`Problem with menu loading: ${e.message}`);
    });

// Scenario 2: Create new order.
client.createOrder({ dishIds: [0, 2, 3], tableId: 3 })
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create order response:', resp);
    })
    .catch((e) => {
        console.log(`Problem creating a new order: ${e.message}`);
    });
    