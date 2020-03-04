const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');
const path = require('path');


const app = express();

//Middlewares
app.use(bodyParser.json());
app.use(cors());

//Rotas
// app.use('/api/', authRoutes);

//const server = "192.168.0.105";
const server = "localhost";

app.listen(3000, server, _ => {
    console.log(`Server is listening at http://${server}:3000`);
});