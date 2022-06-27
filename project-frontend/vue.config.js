const fs = require('fs')
module.exports = {
    devServer: {
        https:true,
        https: {
          cert: fs.readFileSync('ssl/server.crt'),
          key: fs.readFileSync('ssl/server.key'),
        },
    }
}