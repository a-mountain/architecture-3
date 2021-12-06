const request = require('request');

const createHttpClient = (baseUrl) => {
  return {
    get: (path) => {
      return new Promise((resolve, reject) => {
        request(`${baseUrl}${path}`, { json: true }, (err, res, body) => {
          if (err) {
            reject(err);
            return;
          }
          resolve(body);
        });
      });
    },
    post: (path, data) => {
      return new Promise((resolve, reject) => {
        request(
          `${baseUrl}${path}`,
          { json: true, method: 'POST', body: data },
          (err, res, body) => {
            if (err) {
              reject(err);
              return;
            }
            resolve(body);
          }
        );
      });
    },
  };
};

module.exports = { createHttpClient };
