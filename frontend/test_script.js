const axios = require('axios');

axios.get('http://127.0.0.1:3000/api/v1/board/1')
    .then(function (response) {
        // handle success
        console.log(response.data);
    });
